package models

import "github.com/astaxie/beego/orm"

// Merchant represents a merchant
type Merchant struct {
	BaseModel
	Name string  `json:"name" orm:"column(name)"`
	URL  string  `json:"url" orm:"column(url)"`
	Code string  `json:"code" orm:"column(code); unique"`
	Team []*User `json:"teams" orm:"reverse(many)"`
}

func (m *Merchant) TableName() string {
	return "merchants"
}

func init() {
	orm.RegisterModel(new(Merchant))
}
