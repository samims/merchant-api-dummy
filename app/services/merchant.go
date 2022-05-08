package services

import (
	"context"
	"errors"
	"math"

	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
	"github.com/samims/merchant-api/utils"
)

type MerchantService interface {
	Create(context.Context, models.Merchant) (models.PublicMerchant, error)
	Get(context.Context, int64) (models.PublicMerchant, error)
	Update(context.Context, int64, models.Merchant) (models.PublicMerchant, error)
	Delete(context.Context, int64) (map[string]interface{}, error)
	GetTeamMembers(context.Context, int64, *int64, *int64) (models.TeamMemberResponse, error)
	AddTeamMember(context.Context, int64, int64) (map[string]interface{}, error)
	RemoveTeamMember(context.Context, int64, int64) (map[string]interface{}, error)
	VerifyObjectPermission(context.Context, int64) error
}

type merchantService struct {
	merchantRepo repository.MerchantRepo
	userRepo     repository.UserRepo
}

// VerifyObjectPermission checks if the user has permission to access the merchant
func (ctlr *merchantService) VerifyObjectPermission(ctx context.Context, merchantID int64) error {
	groupError := "verifyObjectPermission_merchantService"

	userID := ctx.Value(constants.UserIDContextKey).(int64)
	userQ := models.User{
		BaseModel: models.BaseModel{Id: userID},
	}

	user, err := ctlr.userRepo.FindOne(ctx, userQ)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	// if user is not part of a merchant, we should return an error
	if user.Merchant.Id != merchantID {
		return errors.New(string(constants.PermissionDenied))
	}

	return nil
}

// Create merchant service holds the business logic for creating a merchant
func (svc *merchantService) Create(ctx context.Context, merchant models.Merchant) (models.PublicMerchant, error) {
	groupError := "Create_merchantService"

	requestUserID := ctx.Value(constants.UserIDContextKey).(int64)
	userQ := models.User{
		BaseModel: models.BaseModel{Id: requestUserID},
	}

	user, err := svc.userRepo.FindOne(ctx, userQ)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.PublicMerchant{}, err
	}

	if user.Merchant != nil {
		return models.PublicMerchant{}, errors.New(constants.UserAlreadyPartOfAMerchant)
	}

	err = merchant.AssignCode()

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.PublicMerchant{}, err
	}

	err = svc.merchantRepo.Save(ctx, &merchant)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.PublicMerchant{}, err
	}

	user.Merchant = &merchant
	err = svc.userRepo.Update(ctx, user, []string{"merchant_id"})
	if err != nil {
		logger.Log.WithError(err).Error(groupError)

		// rollback merchant creation cause user association failed
		merchantDeletionErr := svc.merchantRepo.Delete(ctx, merchant)
		if err != nil {
			logger.Log.WithError(merchantDeletionErr).Error(groupError)
			return models.PublicMerchant{}, merchantDeletionErr
		}

	}

	return merchant.Serialize(), err
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

	return merchant.Serialize(), err

}

func (svc *merchantService) Update(ctx context.Context, id int64, doc models.Merchant) (models.PublicMerchant, error) {
	groupError := "Update_merchantService"
	merchantQ := models.Merchant{
		BaseModel: models.BaseModel{Id: id},
	}

	err := svc.VerifyObjectPermission(ctx, id)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.PublicMerchant{}, err
	}

	merchant, err := svc.merchantRepo.FindOne(ctx, merchantQ)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.PublicMerchant{}, err
	}
	merchant.Name = utils.CheckAndSetString(merchant.Name, doc.Name)
	merchant.URL = utils.CheckAndSetString(merchant.URL, doc.URL)

	err = svc.merchantRepo.Update(ctx, merchant, []string{"name", "url"})
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return models.PublicMerchant{}, err
	}
	return merchant.Serialize(), err

}

// Delete merchant
func (svc *merchantService) Delete(ctx context.Context, id int64) (map[string]interface{}, error) {
	groupError := "Delete_merchantService"

	resp := map[string]interface{}{
		"success": false,
	}
	err := svc.VerifyObjectPermission(ctx, id)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return resp, err
	}

	merchantQ := models.Merchant{
		BaseModel: models.BaseModel{Id: id},
	}

	_, err = svc.merchantRepo.FindOne(ctx, merchantQ)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return resp, err
	}

	err = svc.merchantRepo.Delete(ctx, merchantQ)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return resp, err
	}
	resp["success"] = true
	return resp, nil

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

	err := svc.VerifyObjectPermission(ctx, merchantId)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
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

	err := svc.VerifyObjectPermission(ctx, merchantId)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return res, err
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
