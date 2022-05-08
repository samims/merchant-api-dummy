package services

import (
	"context"
	"errors"
	"time"

	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository"
	"github.com/samims/merchant-api/config"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
	"github.com/samims/merchant-api/utils"
)

type UserService interface {
	SignUp(context.Context, models.User) (models.PublicUser, error)
	SignIn(context.Context, models.LoginModel) (models.SignInResponse, error)
	GetAll(context.Context) ([]models.PublicUser, error)
	Update(context.Context, int64, models.User) (models.PublicUser, error)
}

type userService struct {
	userRepo repository.UserRepo
	cfg      config.Configuration
}

// SignUp is a service that creates a new user
func (svc *userService) SignUp(ctx context.Context, user models.User) (models.PublicUser, error) {
	groupError := "SignUp_userService"

	err := user.GeneratePasswordHash()
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return user.Serialize(), errors.New(constants.ErrorEmptyString)
	}

	err = svc.userRepo.Save(ctx, &user)
	// to remove password hash from response
	publicUser := user.Serialize()
	return publicUser, err
}

// SignIn is a service that signs in a user and returns a token if the user is valid and the password is correct
func (svc *userService) SignIn(ctx context.Context, loginModel models.LoginModel) (models.SignInResponse, error) {
	groupError := "SignIn_userService"

	userQ := models.User{
		Email: loginModel.Email,
	}
	user, err := svc.userRepo.FindOne(ctx, userQ)
	logger.Log.Error(user)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.SignInResponse{}, errors.New(constants.ErrorInvalidCredentials)
	}

	err = user.ValidatePasswordHash(loginModel.Password)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.SignInResponse{}, errors.New(constants.ErrorInvalidCredentials)
	}

	token, err := utils.GenerateJWT(user.Id, user.Email, svc.cfg.AppConfig().GetSecretKey())
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.SignInResponse{}, errors.New(constants.ErrorInvalidCredentials)
	}

	return models.SignInResponse{
		Token: token,
		ID:    user.Id,
		Email: user.Email,
		Exp:   time.Now().Add(time.Minute * constants.JWT_EXPIRATION_DELTA_IN_MINUTES).Unix(),
	}, nil
}

// GetAll is a service that returns all users
func (svc *userService) GetAll(ctx context.Context) ([]models.PublicUser, error) {
	users, _, err := svc.userRepo.GetAll(ctx, models.UserQuery{})
	if err != nil {
		logger.Log.WithError(err).Error("GetAll_userService")
		return nil, err
	}
	publicUsers := make([]models.PublicUser, len(users))
	for i, user := range users {
		publicUsers[i] = user.Serialize()
	}
	return publicUsers, nil

}

func (svc *userService) Update(ctx context.Context, id int64, doc models.User) (models.PublicUser, error) {
	grouptError := "Updateuser_repository"

	userQ := models.User{
		BaseModel: models.BaseModel{Id: id},
	}
	user, err := svc.userRepo.FindOne(ctx, userQ)
	if err != nil {
		logger.Log.WithError(err).Error(grouptError)
		return models.PublicUser{}, err
	}

	user.FirstName = utils.CheckAndSetString(user.FirstName, doc.FirstName)
	user.LastName = utils.CheckAndSetString(user.LastName, doc.LastName)
	user.Email = utils.CheckAndSetString(user.Email, doc.Email)

	err = svc.userRepo.Update(ctx, user, []string{"first_name", "last_name", "email"})
	if err != nil {
		logger.Log.WithError(err).Error(grouptError)
		return models.PublicUser{}, err
	}
	return user.Serialize(), nil

}

func NewUserService(userRepo repository.UserRepo, cfg config.Configuration) UserService {
	return &userService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}
