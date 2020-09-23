package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/iamaul/evonix-backend-api/app/models"
	"github.com/iamaul/evonix-backend-api/app/user"
)

type userUsecase struct {
	userRepo       user.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(ur user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepo:       ur,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) GetByID(c context.Context, id int64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	res, err := uu.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uu *userUsecase) GetByName(c context.Context, name string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	res, err := uu.userRepo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uu *userUsecase) GetByEmail(c context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	res, err := uu.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uu *userUsecase) Store(c context.Context, um *models.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	userExist, _ := uu.GetByName(ctx, um.Name)
	if userExist != nil {
		return errors.New("That username is already in use")
	}

	emailExist, _ := uu.GetByEmail(ctx, um.Email)
	if emailExist != nil {
		return errors.New("The email address that you've entered is already exist")
	}

	err := uu.userRepo.Store(ctx, um)
	if err != nil {
		return err
	}

	return nil
}
