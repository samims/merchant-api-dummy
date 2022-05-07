package services

import (
	"context"
	"errors"
	"math"

	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
)

type MerchantService interface {
	Create(context.Context, models.Merchant) (models.PublicMerchant, error)
	Get(context.Context, int64) (models.PublicMerchant, error)
	GetTeamMembers(context.Context, int64, *int64, *int64) (models.TeamMemberResponse, error)
	AddTeamMember(context.Context, int64, int64) (map[string]interface{}, error)
	RemoveTeamMember(context.Context, int64, int64) (map[string]interface{}, error)
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

	relatedUsers, _, err := svc.userRepo.GetAll(ctx, userQuery)
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
	merchantQ := models.Merchant{
		BaseModel: models.BaseModel{Id: id},
	}
	merchant, err := svc.merchantRepo.FindOne(ctx, merchantQ)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.PublicMerchant{}, err
	}

	userQuery := models.UserQuery{
		User: models.User{
			Merchant: merchant,
		},
	}

	relatedUsers, _, err := svc.userRepo.GetAll(ctx, userQuery)

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

func (svc *merchantService) GetTeamMembers(ctx context.Context, id int64, page, pageSize *int64) (models.TeamMemberResponse, error) {
	groupError := "GetTeamMembers_merchantService"
	res := models.TeamMemberResponse{
		Members: []models.PublicUser{},
	}
	if page != nil && *page < 1 {
		return res, errors.New(constants.PaginationError)
	}
	merchantQ := models.Merchant{
		BaseModel: models.BaseModel{Id: id},
	}
	merchant, err := svc.merchantRepo.FindOne(ctx, merchantQ)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
	}

	userQuery := models.UserQuery{
		User: models.User{
			Merchant: merchant,
		},
		Pagination: &models.Pagination{
			Page: page,
			Size: pageSize,
		},
	}

	relatedUsers, totalRecords, err := svc.userRepo.GetAll(ctx, userQuery)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
	}

	for _, reladedUser := range relatedUsers {
		serializedUser := reladedUser.Serialize()
		res.Members = append(res.Members, serializedUser)
	}

	res.TotalRecord = totalRecords
	if page != nil && pageSize != nil {
		res.CurrentPage = *page
		totalPage := int64(math.Ceil(float64(totalRecords) / float64(*pageSize)))
		res.TotalPage = totalPage
	}

	return res, nil
}

func (svc *merchantService) AddTeamMember(ctx context.Context, merchantId int64, userId int64) (map[string]interface{}, error) {
	groupError := "AddTeamMember_merchantService"
	res := map[string]interface{}{
		"success": false,
	}
	merchantQ := models.Merchant{
		BaseModel: models.BaseModel{Id: merchantId},
	}
	merchant, err := svc.merchantRepo.FindOne(ctx, merchantQ)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
	}
	userQ := models.User{
		BaseModel: models.BaseModel{Id: userId},
	}
	user, err := svc.userRepo.FindOne(ctx, userQ)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
	}

	if user.Merchant != nil {
		return res, errors.New(constants.UserAlreadyPartOfAMerchant)
	}

	user.Merchant = merchant
	err = svc.userRepo.Update(ctx, user, []string{"merchant_id"})
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
	}
	res["success"] = true
	return res, nil
}

func (svc *merchantService) RemoveTeamMember(ctx context.Context, merchantId int64, userId int64) (map[string]interface{}, error) {
	groupError := "RemoveTeamMember_merchantService"
	res := map[string]interface{}{
		"success": false,
	}
	merchantQ := models.Merchant{
		BaseModel: models.BaseModel{Id: merchantId},
	}
	merchant, err := svc.merchantRepo.FindOne(ctx, merchantQ)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
	}
	userQ := models.User{
		BaseModel: models.BaseModel{Id: userId},
		Merchant:  merchant,
	}
	user, err := svc.userRepo.FindOne(ctx, userQ)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
	}

	if user.Merchant == nil {
		return res, errors.New(constants.UserNotPartOfAnyMerchant)
	}

	if user.Merchant.Id != merchantId {
		return res, errors.New(constants.UserNotPartOfCurrentMerchant)
	}

	user.Merchant = nil
	err = svc.userRepo.Update(ctx, user, []string{"merchant_id"})
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
	}
	res["success"] = true
	return res, nil
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
