package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type AuthUsecase interface{}

type authController struct {
	useCase  AuthUsecase
	validate *validator.Validate
}

func NewAuthController(useCase AuthUsecase, validate *validator.Validate) *authController {
	return &authController{
		useCase:  useCase,
		validate: validate,
	}
}

func (c *authController) RegisterRoutes(router *chi.Mux) {
	router.Route("/auth", func(r chi.Router) {
		router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("login"))
		})

		router.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("register"))
		})
	})
}
