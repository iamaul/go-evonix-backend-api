package auth

import (
	"context"

	"github.com/iamaul/go-evonix-backend-api/internal/domain"
)

type Repository interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByEmailOrUsername(ctx context.Context, username string, email string) (bool, error)
}
