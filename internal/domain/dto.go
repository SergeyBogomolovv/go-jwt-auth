package domain

import "time"

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
