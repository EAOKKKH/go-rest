package repository

import (
	"context"
	"go-rest/internal/models"

	"github.com/jmoiron/sqlx"
)

type IAuthRepository interface {
	SaveUser(ctx context.Context, user *models.User) (*models.User, error)
	FindUserByLogin(ctx context.Context, login string) (*models.User, error)
}

type authRepository struct {
	db *sqlx.DB
}

// FindUserByLogin implements IAuthRepository.
func (r *authRepository) FindUserByLogin(ctx context.Context, login string) (*models.User, error) {
	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, `SELECT login, password FROM users WHERE login = $1`,
		login).StructScan(u); err != nil {
		return nil, err
	}
	return u, nil
}

// SaveUser implements IAuthRepository.
func (r *authRepository) SaveUser(ctx context.Context, user *models.User) (*models.User, error) {
	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING login, password`,
		user.Login, user.Password).StructScan(u); err != nil {
		return nil, err
	}
	return u, nil
}

func NewUserRepository(db *sqlx.DB) IAuthRepository {
	return &authRepository{db: db}
}
