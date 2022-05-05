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
	Id   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
	Code string `json:"code"`

	Members []*PublicUser `json:"members"`
}

func (m *Merchant) Serialize(members []*PublicUser) PublicMerchant {
	return PublicMerchant{
		Id:      m.Id,
		Name:    m.Name,
		URL:     m.URL,
		Code:    m.Code,
		Members: members,
	}
}

type TeamMemberResponse struct {
	CurrentPage int64        `json:"current_page"`
	TotalPage   int64        `json:"total_page"`
	TotalRecord int64        `json:"total_record"`
	Members     []PublicUser `json:"members"`
}
