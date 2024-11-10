package domain

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
