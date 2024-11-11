package repositories

import (
	"context"
	"go-jwt-auth/internal/domain"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func (r *userRepository) Create(ctx context.Context, dto *domain.RegisterDTO) (*domain.UserModel, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *`

	entity := new(domain.UserEntity)
	if err := r.db.GetContext(ctx, entity, query, dto.Username, dto.Email, dto.Password); err != nil {
		return nil, err
	}

	return entity.ToModel(), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.UserModel, error) {
	query := `SELECT * FROM users WHERE email = $1`

	entity := new(domain.UserEntity)
	if err := r.db.GetContext(ctx, entity, query, email); err != nil {
		return nil, err
	}

	return entity.ToModel(), nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]domain.UserModel, error) {
	query := `SELECT * FROM users`

	entities := make([]domain.UserEntity, 0)
	if err := r.db.SelectContext(ctx, &entities, query); err != nil {
		return nil, err
	}

	users := make([]domain.UserModel, 0)
	for _, entity := range entities {
		users = append(users, *entity.ToModel())
	}

	return users, nil
}

func (r *userRepository) UpdateUsername(ctx context.Context, id uint64, username string) (*domain.UserModel, error) {
	query := `UPDATE users SET username = $1 WHERE user_id = $2 RETURNING *`

	entity := new(domain.UserEntity)
	if err := r.db.GetContext(ctx, entity, query, username, id); err != nil {
		return nil, err
	}

	return entity.ToModel(), nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.UserModel, error) {
	query := `SELECT * FROM users WHERE username = $1`

	entity := new(domain.UserEntity)
	if err := r.db.GetContext(ctx, entity, query, username); err != nil {
		return nil, err
	}

	return entity.ToModel(), nil
}

func (r *userRepository) GetIsUserExists(ctx context.Context, email string, username string) (bool, error) {
	query := `SELECT is_user_exists($1, $2)`

	var exists bool
	if err := r.db.GetContext(ctx, &exists, query, username, email); err != nil {
		return false, err
	}

	return exists, nil
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
