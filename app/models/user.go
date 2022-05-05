package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/samims/merchant-api/utils"
)

// User represents a user of the system assigned to a merchant
type User struct {
	BaseModel
	Email        string    `json:"email" orm:"column(email); unique" validate:"required"`
	FirstName    string    `json:"first_name" orm:"column(first_name)" validate:"required"`
	LastName     string    `json:"last_name" orm:"column(last_name)" validate:"required"`
	PasswordHash string    `json:"password" orm:"column(password_hash)" validate:"required"`
	Merchant     *Merchant `json:"merchant" orm:"rel(fk);null;on_delete(cascade)"`
}

// TableName returns table name which will be created on db
func (u *User) TableName() string {
	return "users"
}

func (u *User) GeneratePasswordHash() error {
	hash, err := utils.GenerateBCryptHash(u.PasswordHash)
	if err != nil {
		return err
	}
	u.PasswordHash = hash
	return nil
}

func (u *User) ValidatePasswordHash(password string) error {
	return utils.ValidateBCryptHash(password, u.PasswordHash)
}

func init() {
	orm.RegisterModel(new(User))
}

// PublicUser represents a user of the system without sensitive information
type PublicUser struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Serialize serializes user to PublicUser
func (u *User) Serialize() PublicUser {
	return PublicUser{
		ID:        u.Id,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

type UserQuery struct {
	User
	Pagination *Pagination
}
