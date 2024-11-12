package controllers

import (
	"context"
	"errors"
	constants "go-jwt-auth/internal"
	"go-jwt-auth/internal/domain"
	"go-jwt-auth/internal/middleware"
	"go-jwt-auth/pkg/utils"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UsersUsecase interface {
	GetAllUsers(ctx context.Context) ([]domain.UserModel, error)
	UpdateUsername(ctx context.Context, id uint64, username string) (*domain.UserModel, error)
}

type usersController struct {
	useCase  UsersUsecase
	validate *validator.Validate
}

func (c *usersController) RegisterRoutes(router *http.ServeMux, mw middleware.Middleware) {
	users := http.NewServeMux()
	users.HandleFunc("GET /all", c.GetAllUsers)
	users.HandleFunc("PUT /update-username", c.UpdateUsername)

	router.Handle("/users/", http.StripPrefix("/users", mw(users)))
	slog.Info("users router registered")
}

func (c *usersController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.useCase.GetAllUsers(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, users)
}

func (c *usersController) UpdateUsername(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(constants.AuthUserId).(uint64)

	dto := new(domain.UpdateUsernameDTO)
	if err := utils.ParseJSON(r, dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := c.validate.Struct(dto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := c.useCase.UpdateUsername(r.Context(), id, dto.Username)
	if err != nil {
		if err.Error() == constants.ErrUserAlreadyExists {
			utils.WriteError(w, http.StatusConflict, errors.New("user already exists"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}

func NewUsersController(useCase UsersUsecase, validate *validator.Validate) *usersController {
	return &usersController{
		useCase:  useCase,
		validate: validate,
	}
}
