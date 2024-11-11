package constants

type ContextKey string

const (
	ErrUserAlreadyExists             = "user already exists"
	ErrTokenIvalid                   = "token ivalid"
	ErrInvalidCredentials            = "invalid credentials"
	ErrUnathorized                   = "unauthorized"
	UserIdKey             ContextKey = "user_id"
)
