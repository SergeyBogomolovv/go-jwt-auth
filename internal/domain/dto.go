package domain

import "time"

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterDTO struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type JWTResponse struct {
	AccessToken string `json:"access_token"`
}

type JWTPayload struct {
	UserID uint64        `json:"user_id"`
	Exp    time.Duration `json:"exp"`
}

const (
	ErrUserAlreadyExists  = "user already exists"
	ErrTokenIvalid        = "token ivalid"
	ErrInvalidCredentials = "invalid credentials"
)
