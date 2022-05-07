package repository

import (
	"context"
	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
)

type UserRepo interface {
	Save(context.Context, *models.User) error
	GetAll(context.Context, models.UserQuery) ([]models.User, int64, error)
	FindOne(context.Context, models.User) (*models.User, error)
	Update(context.Context, *models.User, []string) error
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
func (repo *userRepo) GetAll(ctx context.Context, query models.UserQuery) ([]models.User, int64, error) {
	var users []models.User
	groupError := "GetUserList_userRepo"
	qs := repo.db.QueryTable(new(models.User)).OrderBy("-created_at")

	if query.Pagination != nil && query.Pagination.Page != nil && query.Pagination.Size != nil {
		// using pointer in pagination to make sure it will be nil if not set rather than 0
		qs = qs.Offset((*query.Pagination.Page - 1) * *query.Pagination.Size).Limit(*query.Pagination.Size)
	}

	if query.Merchant != nil && query.Merchant.Id != 0 {
		qs = qs.Filter("merchant_id", query.Merchant.Id)
	}

	qs = qs.RelatedSel()

	_, err := qs.All(&users)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return nil, 0, err
	}
	// Count ignores pagination and returns total number of records without limit
	count, err := qs.Count()
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return nil, 0, err
	}

	return users, count, nil
}

// FindOne fetch one object by query
func (repo *userRepo) FindOne(ctx context.Context, user models.User) (*models.User, error) {
	groupError := "FindOne_userRepo"
	if user.Id == 0 {
		return nil, errors.New(constants.UserIDIsRequired)
	}
	qs := repo.db.QueryTable(new(models.User))
	if user.Id != 0 {
		qs = qs.Filter("id", user.Id)
	}

	err := qs.One(&user)

	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		if err == orm.ErrNoRows {
			return nil, errors.New(constants.UserNotFound)
		}
		return nil, err
	}
	return &user, nil
}

func (repo *userRepo) Update(ctx context.Context, doc *models.User, fieldsToUpdate []string) error {
	groupError := "Update_userRepo"
	_, err := repo.db.Update(doc, fieldsToUpdate...)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		return err
	}
	return nil
}

func NewUserRepo(db orm.Ormer) UserRepo {
	return &userRepo{
		db: db,
	}
}
