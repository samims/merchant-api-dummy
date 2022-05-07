package repository

import (
	"context"
	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
)

type MerchantRepo interface {
	Save(context.Context, *models.Merchant) error
	Update(context.Context, *models.Merchant, []string) error
	GetAll(context.Context) ([]models.Merchant, error)
	Delete(context.Context, models.Merchant) error
	FindOne(context.Context, models.Merchant) (*models.Merchant, error)
}

type merchantRepo struct {
	db orm.Ormer
}

func (repo *merchantRepo) Save(ctx context.Context, doc *models.Merchant) error {
	groupError := "Save_merchantRepo"
	id, err := repo.db.Insert(doc)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	doc.Id = id
	return nil

}

func (repo *merchantRepo) Update(ctx context.Context, doc *models.Merchant, fieldsToUpdate []string) error {
	groupError := "Update_merchantRepo"

	_, err := repo.db.Update(doc, fieldsToUpdate...)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
	}

	return err
}

// Delete merchant by id
func (repo *merchantRepo) Delete(ctx context.Context, merchant models.Merchant) error {
	groupError := "Delete_merchantRepo"
	_, err := repo.db.Delete(&merchant)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	return nil
}

func (repo *merchantRepo) GetAll(ctx context.Context) ([]models.Merchant, error) {
	var merchants []models.Merchant
	groupError := "GetMerchantList_merchantRepo"
	qs := repo.db.QueryTable(new(models.Merchant)).OrderBy("-created_at")
	_, err := qs.All(&merchants)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return nil, err
	}

	return merchants, nil
}

// FindOne fetch one object by query
func (repo *merchantRepo) FindOne(ctx context.Context, merchant models.Merchant) (*models.Merchant, error) {
	groupError := "FindOne_merchantRepo"
	qs := repo.db.QueryTable(new(models.Merchant))
	if merchant.Id != 0 {
		qs = qs.Filter("id", merchant.Id)
	}
	if len(merchant.Code) != 0 {
		qs = qs.Filter("code", merchant.Code)
	}
	err := qs.One(&merchant)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		if err == orm.ErrNoRows {
			return nil, errors.New(constants.MerchantNotFound)
		}
		return nil, err
	}
	return &merchant, nil
}

func NewMerchantRepo(db orm.Ormer) MerchantRepo {
	return &merchantRepo{
		db: db,
	}
}
