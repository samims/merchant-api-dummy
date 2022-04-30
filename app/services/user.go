package services

import (
	"context"

	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository"
	"github.com/samims/merchant-api/logger"
)

type UserService interface {
	SignUp(context.Context, models.User) (models.User, error)
}

type userService struct {
	userRepo repository.UserRepo
}

func (svc *userService) SignUp(ctx context.Context, user models.User) (models.User, error) {
	logger.Log.Info("Signup service is being called!!!!")

	svc.userRepo.Save(ctx, &user)
	return user, nil
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
