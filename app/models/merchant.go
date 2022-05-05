package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/samims/merchant-api/logger"
	"github.com/samims/merchant-api/utils"
)

// Merchant represents a merchant
type Merchant struct {
	BaseModel
	Name string  `json:"name" orm:"column(name)" validate:"required"`
	URL  string  `json:"url" orm:"column(url)"`
	Code string  `json:"code" orm:"column(code); unique"`
	Team []*User `json:"teams" orm:"reverse(many)"`
}

func (m *Merchant) AssignCode() error {
	uuidString, err := utils.GenerateUUIDString()

	if err != nil {
		logger.Log.WithError(err).Error("GenerateCode_merchant")
		return err
	}
	m.Code = uuidString
	return nil

}

func (m *Merchant) TableName() string {
	return "merchants"
}

func init() {
	orm.RegisterModel(new(Merchant))
}

type PublicMerchant struct {
	Merchant
	Teams []*PublicUser `json:"teams"`
}

func (m *Merchant) Serialize(teams []*PublicUser) PublicMerchant {
	return PublicMerchant{
		Merchant: *m,
		Teams:    teams,
	}
}
