package models

import "time"

// Employer 雇主信息表
type Employer struct {
	BaseModel
	UserID          int64      `json:"user_id" gorm:"column:user_id;type:bigint;not null;index;comment:关联用户ID"`
	BrandID         int64      `json:"brand_id" gorm:"column:brand_id;type:bigint;not null;index;comment:所属品牌ID"`
	CompanyName     string     `json:"company_name" gorm:"column:company_name;type:varchar(100);comment:公司名称"`
	ContactPerson   string     `json:"contact_person" gorm:"column:contact_person;type:varchar(50);comment:联系人姓名"`
	ContactPhone    string     `json:"contact_phone" gorm:"column:contact_phone;type:varchar(20);comment:联系人手机"`
	BusinessLicense string     `json:"business_license" gorm:"column:business_license;type:varchar(100);comment:营业执照号"`
	AuthStatus      string     `json:"auth_status" gorm:"column:auth_status;type:enum('pending','approved','rejected');default:pending;index;comment:认证状态"`
	AuthTime        *time.Time `json:"auth_time" gorm:"column:auth_time;type:datetime;comment:认证时间"`
}

// TableName 指定表名
func (Employer) TableName() string {
	return "employers"
}
