package models

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
)

// Store 门店信息表
type Store struct {
	BaseModel
	BrandID       int64           `json:"brand_id" gorm:"column:brand_id;type:bigint;not null;index;comment:所属品牌ID"`
	Name          string          `json:"name" gorm:"column:name;type:varchar(100);not null;comment:门店名称"`
	Address       string          `json:"address" gorm:"column:address;type:varchar(255);not null;comment:门店地址"`
	Latitude      decimal.Decimal `json:"latitude" gorm:"column:latitude;type:decimal(10,8);comment:纬度"`
	Longitude     decimal.Decimal `json:"longitude" gorm:"column:longitude;type:decimal(11,8);comment:经度"`
	ContactPhone  string          `json:"contact_phone" gorm:"column:contact_phone;type:varchar(20);comment:联系电话"`
	ContactPerson string          `json:"contact_person" gorm:"column:contact_person;type:varchar(50);comment:联系人"`
	Description   string          `json:"description" gorm:"column:description;type:text;comment:门店描述"`
	Status        string          `json:"status" gorm:"column:status;type:enum('active','disabled');not null;default:'active';index;comment:门店状态"`
}

// TableName 指定表名
func (Store) TableName() string {
	return "stores"
}

// ToThriftStore 转换为Thrift门店信息
func (s *Store) ToThriftStore() *admin.StoreDetail {
	return &admin.StoreDetail{
		StoreID:       s.ID,
		BrandID:       s.BrandID,
		Name:          s.Name,
		Address:       s.Address,
		Latitude:      s.Latitude.String(),
		Longitude:     s.Longitude.String(),
		ContactPhone:  s.ContactPhone,
		ContactPerson: s.ContactPerson,
		Description:   s.Description,
		Status:        s.Status,
		CreatedAt:     s.CreatedAt.Format(time.DateTime),
		UpdatedAt:     s.UpdatedAt.Format(time.DateTime),
	}
}
