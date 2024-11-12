package usecases

import (
	"context"
	"errors"
	constants "go-jwt-auth/internal"
	"go-jwt-auth/internal/domain"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetIsUserExists(ctx context.Context, email string, username string) (bool, error)
	Create(ctx context.Context, dto *domain.RegisterDTO) (*domain.UserModel, error)
	GetByEmail(ctx context.Context, email string) (*domain.UserModel, error)
}

type authUsecase struct {
	repo      UserRepository
	jwtSecret []byte
	jwtExp    time.Duration
}

func (u *authUsecase) Register(ctx context.Context, dto *domain.RegisterDTO) (*domain.JWTResponse, error) {
	isUserExits, err := u.repo.GetIsUserExists(ctx, dto.Email, dto.Username)
	if err != nil {
		return nil, err
	}

	if isUserExits {
		return nil, errors.New(constants.ErrUserAlreadyExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	dto.Password = string(hashedPassword)

	user, err := u.repo.Create(ctx, dto)
	if err != nil {
		return nil, err
	}

	token, err := u.SignToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.JWTResponse{
		AccessToken: token,
	}, nil
}

func (u *authUsecase) Login(ctx context.Context, dto *domain.LoginDTO) (*domain.JWTResponse, error) {
	user, err := u.repo.GetByEmail(ctx, dto.Email)
	if err != nil {
		return nil, errors.New(constants.ErrInvalidCredentials)
	}

	if err := u.VerifyPassword(dto.Password, user.Password); err != nil {
		return nil, errors.New(constants.ErrInvalidCredentials)
	}

	token, err := u.SignToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.JWTResponse{
		AccessToken: token,
	}, nil
}

func (u *authUsecase) SignToken(userId uint64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "auth-app",
		"sub": strconv.Itoa(int(userId)),
		"exp": time.Now().Add(u.jwtExp).Unix(),
		"iat": time.Now().Unix(),
	}).SignedString(u.jwtSecret)
}

func (u *authUsecase) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return u.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New(constants.ErrTokenIvalid)
	}
	return token, nil
}

func (u *authUsecase) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *authUsecase) VerifyPassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func NewAuthUsecase(repo UserRepository, jwtSecret []byte) *authUsecase {
	return &authUsecase{
		repo:      repo,
		jwtSecret: jwtSecret,
		jwtExp:    time.Minute * 15,
	}
}
