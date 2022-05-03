package repository

import (
	"context"

	"github.com/iamaul/go-evonix-backend-api/internal/auth"
	"github.com/iamaul/go-evonix-backend-api/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) auth.Repository {
	return &authRepo{db: db}
}

func (r *authRepo) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	newUser := &domain.User{}
	if err := r.db.QueryRowxContext(ctx, createUserQuery, &user.Name, &user.Password, &user.Email, &user.RegisteredDate, &user.RegisterIP).StructScan(newUser); err != nil {
		return nil, errors.Wrap(err, "[auth/repository] Register.StructScan")
	}

	return newUser, nil
}

func (r *authRepo) FindByEmailOrUsername(ctx context.Context, username string, email string) (bool, error) {
	foundUser := &domain.User{}
	if err := r.db.QueryRowxContext(ctx, getUserByEmailOrUsername, username, email).StructScan(foundUser); err != nil {
		return false, errors.Wrap(err, "[auth/repository] FindByEmailOrUsername.StructScan")
	}

	return true, nil
}
