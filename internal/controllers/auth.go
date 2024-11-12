package controllers

import (
	"context"
	"errors"
	constants "go-jwt-auth/internal"
	"go-jwt-auth/internal/domain"
	"go-jwt-auth/pkg/utils"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthUsecase interface {
	Register(ctx context.Context, dto *domain.RegisterDTO) (*domain.JWTResponse, error)
	Login(ctx context.Context, dto *domain.LoginDTO) (*domain.JWTResponse, error)
}

type authController struct {
	useCase  AuthUsecase
	validate *validator.Validate
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	dto := new(domain.LoginDTO)

	if err := utils.ParseJSON(r, dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := c.validate.Struct(dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	jwtResponse, err := c.useCase.Login(r.Context(), dto)
	if err != nil {
		if err.Error() == constants.ErrInvalidCredentials {
			utils.WriteError(w, http.StatusUnauthorized, errors.New("invalid credentials"))
			return
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusOK, jwtResponse)
}

func (c *authController) Register(w http.ResponseWriter, r *http.Request) {
	dto := new(domain.RegisterDTO)

	if err := utils.ParseJSON(r, dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := c.validate.Struct(dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	jwtResponse, err := c.useCase.Register(r.Context(), dto)
	if err != nil {
		if err.Error() == constants.ErrUserAlreadyExists {
			utils.WriteError(w, http.StatusConflict, errors.New("user already exists"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusCreated, jwtResponse)
}

func NewAuthController(useCase AuthUsecase, validate *validator.Validate) *authController {
	return &authController{
		useCase:  useCase,
		validate: validate,
	}
}

func (c *authController) RegisterRoutes(router *http.ServeMux) {
	auth := http.NewServeMux()
	auth.HandleFunc("POST /login", c.Login)
	auth.HandleFunc("POST /register", c.Register)

	router.Handle("/auth/", http.StripPrefix("/auth", auth))
	slog.Info("auth router registered")
}
