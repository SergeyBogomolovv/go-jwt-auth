package usecases

import (
	"context"
	"errors"
	constants "go-jwt-auth/internal"
	"go-jwt-auth/internal/domain"
)

type Repository interface {
	GetIsUserExists(ctx context.Context, email string, username string) (bool, error)
	GetAllUsers(ctx context.Context) ([]domain.UserModel, error)
	UpdateUsername(ctx context.Context, id uint64, username string) (*domain.UserModel, error)
}

type usersUsecase struct {
	repo Repository
}

func (u *usersUsecase) GetAllUsers(ctx context.Context) ([]domain.UserModel, error) {
	return u.repo.GetAllUsers(ctx)
}

func (u *usersUsecase) UpdateUsername(ctx context.Context, id uint64, username string) (*domain.UserModel, error) {
	isUsernameExists, err := u.repo.GetIsUserExists(ctx, "", username)
	if err != nil || isUsernameExists {
		return nil, errors.New(constants.ErrUserAlreadyExists)
	}

	return u.repo.UpdateUsername(ctx, id, username)
}

func NewUsersUsecase(repo Repository) *usersUsecase {
	return &usersUsecase{
		repo: repo,
	}
}
