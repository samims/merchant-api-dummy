package models

import "github.com/astaxie/beego/orm"

type Merchant struct {
	BaseModel
	Name  string  `json:"name" orm:"column(name)"`
	URL   string  `json:"url" orm:"column(url)"`
	Code  string  `json:"code" orm:"column(code); unique"`
	Users []*User `json:"users" orm:"reverse(many)"`
}

func (m *Merchant) TableName() string {
	return "merchants"
}

func init() {
	orm.RegisterModel(new(Merchant))
}
