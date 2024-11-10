package domain

import "time"

type UserEntity struct {
	ID        uint64    `db:"user_id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

func (u *UserEntity) ToModel() *UserModel {
	return &UserModel{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Password:  string(u.Password),
		CreatedAt: u.CreatedAt,
	}
}
