package repository

import (
	"context"

	"github.com/astaxie/beego/orm"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/logger"
)

type UserRepo interface {
	Save(context.Context, *models.User) error
	GetAll(context.Context) ([]models.User, error)
}

type userRepo struct {
	db orm.Ormer
}

// Save writes data to db
// It returns error/nil but we generally need ID after save
// so we are using model pointer which will be modified and ID will
// be assigned and can be accessed from any point
func (repo *userRepo) Save(ctx context.Context, doc *models.User) error {
	groupError := "Save_userRepo"
	id, err := repo.db.Insert(doc)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	doc.Id = id

	return nil
}

// get user list
func (repo *userRepo) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	groupError := "GetUserList_userRepo"
	qs := repo.db.QueryTable(new(models.User))
	qs = qs.RelatedSel()
	_, err := qs.All(&users)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return nil, err
	}

	return users, nil
}

func NewUserRepo(db orm.Ormer) UserRepo {
	return &userRepo{
		db: db,
	}
}
