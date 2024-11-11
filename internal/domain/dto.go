package domain

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

type UpdateUsernameDTO struct {
	Username string `json:"username" validate:"required,min=3"`
}
