package services

import (
	"context"

	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository"
	"github.com/samims/merchant-api/logger"
)

type UserService interface {
	SignUp(context.Context, models.User) (models.User, error)
	GetAll(context.Context) ([]models.User, error)
}

type userService struct {
	userRepo repository.UserRepo
}

func (svc *userService) SignUp(ctx context.Context, user models.User) (models.User, error) {
	logger.Log.Info("Signup service is being called!!!!")

	svc.userRepo.Save(ctx, &user)
	return user, nil
}

func (svc *userService) GetAll(ctx context.Context) ([]models.User, error) {
	logger.Log.Info("User GetAll service is being called!!")
	users, err := svc.userRepo.GetAll(ctx)
	if err != nil {
		logger.Log.WithError(err).Error("GetAll_userService")
		return nil, err
	}
	return users, nil

}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
