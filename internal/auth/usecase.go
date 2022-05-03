package auth

import (
	"context"

	"github.com/iamaul/go-evonix-backend-api/internal/domain"
)

type Usecase interface {
	Register(ctx context.Context, user *domain.User) (*domain.UserWithToken, error)
}
