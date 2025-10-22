package models

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
)

// User 用户基础信息表
type User struct {
	BaseModel
	Username     string `json:"username" gorm:"column:username;type:varchar(50);not null;comment:用户名"`
	Phone        string `json:"phone" gorm:"column:phone;type:varchar(20);uniqueIndex;not null;comment:手机号"`
	Avatar       string `json:"avatar" gorm:"column:avatar;type:varchar(255);comment:头像URL"`
	PasswordHash string `json:"password_hash" gorm:"column:password_hash;type:varchar(255);not null;comment:密码哈希"`
	Role         string `json:"role" gorm:"column:role;type:enum('worker','employer','admin');not null;index;comment:用户角色"`
	Status       string `json:"status" gorm:"column:status;type:enum('active','disabled','pending');not null;default:'active';index;comment:账号状态"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

func (u *User) ToThriftUser() *common.UserInfo {
	return &common.UserInfo{
		UserID:   u.ID,
		Username: u.Username,
		Phone:    u.Phone,
		Avatar:   u.Avatar,
		Role:     u.Role,
	}
}

func (u *User) ToThriftAdmin() *admin.AdminInfo {
	return &admin.AdminInfo{
		AdminID:  u.ID,
		Username: u.Username,
		RealName: u.Username,
		Role:     u.Role,
		// Permissions: u.Permissions,
		CreatedAt:     u.CreatedAt.Format(time.RFC3339),
		AccountStatus: u.Status,
	}
}
