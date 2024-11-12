package usecases_test

import (
	"context"
	"go-jwt-auth/internal/domain"
	"go-jwt-auth/internal/usecases"
	"testing"
)

func TestUsersUsecase(t *testing.T) {
	repo := &userRepoUsersMock{}
	usecase := usecases.NewUsersUsecase(repo)

	t.Run("Get all users", func(t *testing.T) {
		users, err := usecase.GetAllUsers(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if len(users) != 1 {
			t.Error("expected 1 user, got", len(users))
		}
	})

	t.Run("Update username", func(t *testing.T) {
		user, err := usecase.UpdateUsername(context.Background(), 1, "test")
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != 1 {
			t.Error("expected id 1, got", user.ID)
		}
		if user.Username != "test" {
			t.Error("expected username test, got", user.Username)
		}
	})
}

type userRepoUsersMock struct{}

func (u *userRepoUsersMock) GetAllUsers(ctx context.Context) ([]domain.UserModel, error) {
	return []domain.UserModel{{}}, nil
}

func (u *userRepoUsersMock) UpdateUsername(ctx context.Context, id uint64, username string) (*domain.UserModel, error) {
	return &domain.UserModel{ID: id, Username: username}, nil
}

func (u *userRepoUsersMock) GetIsUserExists(ctx context.Context, email string, username string) (bool, error) {
	return false, nil
}
