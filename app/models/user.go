package models

import "github.com/astaxie/beego/orm"

// User represents a user of the system assigned to a merchant
type User struct {
	BaseModel
	Email        string    `json:"email" orm:"column(email); unique"`
	FirstName    string    `json:"first_name" orm:"column(first_name)"`
	LastName     string    `json:"last_name" orm:"column(last_name)"`
	PasswordHash string    `json:"password" orm:"column(password_hash)"`
	Merchant     *Merchant `json:"merchant" orm:"rel(fk);null;on_delete(cascade)"`
}

// TableName returns table name which will be created on db
func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(User))
}
