package models

// UserRole 用户角色关联表
type UserRole struct {
	BaseModel
	UserID   int64  `json:"user_id" gorm:"column:user_id;type:bigint;not null;index;comment:用户ID"`
	RoleType string `json:"role_type" gorm:"column:role_type;type:enum('brand_admin','store_admin');not null;index;comment:角色类型"`
	BrandID  *int64 `json:"brand_id" gorm:"column:brand_id;type:bigint;index;comment:关联品牌ID"`
	StoreID  *int64 `json:"store_id" gorm:"column:store_id;type:bigint;index;comment:关联门店ID"`
	Status   string `json:"status" gorm:"column:status;type:enum('active','disabled');not null;default:'active';index;comment:角色状态"`
}

// TableName 指定表名
func (UserRole) TableName() string {
	return "user_roles"
}
