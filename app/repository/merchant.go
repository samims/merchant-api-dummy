package repository

import (
	"context"

	"github.com/astaxie/beego/orm"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/logger"
)

type MerchantRepo interface {
	Save(context.Context, *models.Merchant) error
	GetAll(context.Context) ([]models.Merchant, error)
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

func (repo *merchantRepo) GetAll(ctx context.Context) ([]models.Merchant, error) {
	var merchants []models.Merchant
	groupError := "GetMerchantList_merchantRepo"
	qs := repo.db.QueryTable(new(models.Merchant))
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
		return nil, err
	}
	return &merchant, nil
}

func NewMerchantRepo(db orm.Ormer) MerchantRepo {
	return &merchantRepo{
		db: db,
	}
}
