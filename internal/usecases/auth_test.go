package usecases_test

import (
	"context"
	"go-jwt-auth/internal/domain"
	"go-jwt-auth/internal/usecases"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestAuthUsecase(t *testing.T) {
	repo := &userRepoAuthMock{}
	usecase := usecases.NewAuthUsecase(repo, []byte("secret"))
	t.Run("Sign and verify token", func(t *testing.T) {
		token, err := usecase.SignToken(1)
		if err != nil {
			t.Fatal(err)
		}

		result, err := usecase.VerifyToken(token)
		if err != nil {
			t.Fatal(err)
		}
		if !result.Valid {
			t.Error("token is not valid")
		}

		userId, err := result.Claims.GetSubject()
		if err != nil {
			t.Fatal(err)
		}
		if userId != "1" {
			t.Error("user id is not 1")
		}
	})

	t.Run("Register", func(t *testing.T) {
		response, err := usecase.Register(context.Background(), &domain.RegisterDTO{
			Email:    "a@b.com",
			Username: "test",
			Password: "test",
		})
		if err != nil {
			t.Fatal(err)
		}

		verifiedToken, err := usecase.VerifyToken(response.AccessToken)
		if err != nil {
			t.Fatal(err)
		}
		if !verifiedToken.Valid {
			t.Error("token is not valid")
		}
	})

	t.Run("Login", func(t *testing.T) {
		response, err := usecase.Login(context.Background(), &domain.LoginDTO{
			Email:    "a@b.com",
			Password: "test",
		})
		if err != nil {
			t.Fatal(err)
		}

		verifiedToken, err := usecase.VerifyToken(response.AccessToken)
		if err != nil {
			t.Fatal(err)
		}
		if !verifiedToken.Valid {
			t.Error("token is not valid")
		}
	})
}

type userRepoAuthMock struct{}

func (u *userRepoAuthMock) Create(ctx context.Context, dto *domain.RegisterDTO) (*domain.UserModel, error) {
	return &domain.UserModel{}, nil
}

func (u *userRepoAuthMock) GetIsUserExists(ctx context.Context, email string, username string) (bool, error) {
	return false, nil
}

func (u *userRepoAuthMock) GetByEmail(ctx context.Context, email string) (*domain.UserModel, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	return &domain.UserModel{Password: string(hashedPassword), Email: email}, nil
}
