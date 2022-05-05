package services

import (
	"context"

	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository"
	"github.com/samims/merchant-api/logger"
)

type MerchantService interface {
	Create(context.Context, models.Merchant) (models.PublicMerchant, error)
	Get(context.Context, int64) (models.PublicMerchant, error)
}

type merchantService struct {
	merchantRepo repository.MerchantRepo
	userRepo     repository.UserRepo
}

func (svc *merchantService) Create(ctx context.Context, merchant models.Merchant) (models.PublicMerchant, error) {
	groupError := "Create_merchantService"

	err := merchant.AssignCode()
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return merchant.Serialize(nil), err
	}

	err = svc.merchantRepo.Save(ctx, &merchant)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return merchant.Serialize(nil), err
	}

	/*
		merchant := Merchant{
			id 1
			url 'abc.com'
			teams []
		}
	*/
	userQuery := models.UserQuery{
		User: models.User{
			Merchant: &merchant,
		},
	}

	relatedUsers, err := svc.userRepo.GetAll(ctx, userQuery)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return merchant.Serialize(nil), err
	}
	var relatedUserPublic []*models.PublicUser
	for _, reladedUser := range relatedUsers {
		serializedUser := reladedUser.Serialize()
		relatedUserPublic = append(relatedUserPublic, &serializedUser)
	}

	return merchant.Serialize(relatedUserPublic), err
}

func (svc *merchantService) Get(ctx context.Context, id int64) (models.PublicMerchant, error) {
	groupError := "Get_merchantService"
	merchant := models.Merchant{
		BaseModel: models.BaseModel{Id: id},
	}
	resp, err := svc.merchantRepo.FindOne(ctx, merchant)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.PublicMerchant{}, err
	}

	userQuery := models.UserQuery{
		User: models.User{
			Merchant: resp,
		},
	}

	relatedUsers, err := svc.userRepo.GetAll(ctx, userQuery)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return merchant.Serialize(nil), err
	}
	var relatedUserPublic []*models.PublicUser
	for _, reladedUser := range relatedUsers {
		serializedUser := reladedUser.Serialize()
		relatedUserPublic = append(relatedUserPublic, &serializedUser)
	}

	return resp.Serialize(relatedUserPublic), err

}

func NewMerchantService(
	merchantRepo repository.MerchantRepo,
	userRepo repository.UserRepo,
) MerchantService {
	return &merchantService{
		merchantRepo: merchantRepo,
		userRepo:     userRepo,
	}
}
