package models

// User 用户基础信息表
type User struct {
	BaseModel
	Username     string `json:"username" gorm:"column:username;type:varchar(50);not null;comment:用户名"`
	Phone        string `json:"phone" gorm:"column:phone;type:varchar(20);uniqueIndex;not null;comment:手机号"`
	Avatar       string `json:"avatar" gorm:"column:avatar;type:varchar(255);comment:头像URL"`
	PasswordHash string `json:"password_hash" gorm:"column:password_hash;type:varchar(255);not null;comment:密码哈希"`
	Role         string `json:"role" gorm:"column:role;type:enum('worker','employer','admin');not null;index;comment:用户角色"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
