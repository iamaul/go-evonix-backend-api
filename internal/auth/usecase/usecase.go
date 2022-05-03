package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/iamaul/go-evonix-backend-api/internal/auth"
	"github.com/iamaul/go-evonix-backend-api/internal/domain"
	"github.com/iamaul/go-evonix-backend-api/pkg/errors"
	"github.com/iamaul/go-evonix-backend-api/pkg/hash"
	"github.com/iamaul/go-evonix-backend-api/pkg/jwt"
	"github.com/iamaul/go-evonix-backend-api/pkg/logger"
)

type authUsecase struct {
	authRepo auth.Repository
	hasher   hash.PasswordHasher
	tokenJwt jwt.TokenManager
	logger   logger.Logger
}

func NewAuthUsecase(authRepo auth.Repository, hasher hash.PasswordHasher, tokenJwt jwt.TokenManager, logger logger.Logger) auth.Usecase {
	return &authUsecase{authRepo: authRepo, hasher: hasher, tokenJwt: tokenJwt, logger: logger}
}

func (u *authUsecase) Register(ctx context.Context, user *domain.User) (*domain.UserWithToken, error) {
	userExist, err := u.authRepo.FindByEmailOrUsername(ctx, user.Name, user.Email)
	if !userExist && err != nil {
		return nil, errors.NewRestErrorWithMessage(http.StatusBadRequest, errors.UsernameOrEmailAlreadyExists, nil)
	}

	passwordHash, err := u.hasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = passwordHash
	user.RegisteredDate = int32(time.Now().Unix())

	createdUser, err := u.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	userID := fmt.Sprintf("%v", createdUser.ID)
	accessToken, err := u.tokenJwt.NewJWT(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := u.tokenJwt.NewRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	return &domain.UserWithToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
