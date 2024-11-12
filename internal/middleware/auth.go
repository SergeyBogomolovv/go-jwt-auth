package middleware

import (
	"context"
	"errors"
	constants "go-jwt-auth/internal"
	"go-jwt-auth/pkg/utils"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUseCase interface {
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type authMiddleware struct {
	useCase AuthUseCase
}

func (m *authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.WriteError(w, http.StatusUnauthorized, errors.New(constants.ErrUnathorized))
			return
		}

		token, err := m.useCase.VerifyToken(tokenString[len("Bearer "):])
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, errors.New(constants.ErrUnathorized))
			return
		}
		userId, err := token.Claims.GetSubject()
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, errors.New("invalid subject"))
			return
		}

		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, errors.New("invalid user id"))
			return
		}

		ctx := context.WithValue(r.Context(), constants.AuthUserId, uint64(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewAuthMiddleware(useCase AuthUseCase) *authMiddleware {
	return &authMiddleware{
		useCase: useCase,
	}
}
