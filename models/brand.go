package models

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
)

// Brand 品牌信息表
type Brand struct {
	BaseModel
	Name        string `json:"name" gorm:"column:name;type:varchar(100);not null;index;comment:品牌名称"`
	Logo        string `json:"logo" gorm:"column:logo;type:varchar(255);comment:品牌Logo URL"`
	Description string `json:"description" gorm:"column:description;type:text;comment:品牌描述"`
	AuthStatus  string `json:"auth_status" gorm:"column:auth_status;type:enum('pending','approved','rejected');default:pending;comment:认证状态"`
}

// TableName 指定表名
func (Brand) TableName() string {
	return "brands"
}

// ToThriftBrand 转换为Thrift品牌信息
func (b *Brand) ToThriftBrand() *admin.BrandDetail {
	return &admin.BrandDetail{
		BrandID:       b.ID,
		CompanyName:   b.Name,
		Logo:          b.Logo,
		Description:   b.Description,
		AuthStatus:    b.AuthStatus,
		CreatedAt:     b.CreatedAt.Format(time.DateTime),
		UpdatedAt:     b.UpdatedAt.Format(time.DateTime),
		AccountStatus: "active",
		JobCount:      0,
		ActivityLevel: "normal",
	}
}
