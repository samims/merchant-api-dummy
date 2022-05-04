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

func NewMerchantRepo(db orm.Ormer) MerchantRepo {
	return &merchantRepo{
		db: db,
	}
}
