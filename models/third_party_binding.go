package models

import (
	"time"
)

// ThirdPartyBinding 第三方账号绑定表
type ThirdPartyBinding struct {
	BaseModel
	UserID      int64      `json:"user_id" gorm:"column:user_id;type:bigint;not null;index:idx_user_id;comment:用户ID"`
	Platform    string     `json:"platform" gorm:"column:platform;type:varchar(50);not null;index;comment:第三方平台"`
	OpenID      string     `json:"openid" gorm:"column:openid;type:varchar(255);not null;uniqueIndex:idx_platform_openid;comment:第三方平台OpenID"`
	UnionID     string     `json:"unionid" gorm:"column:unionid;type:varchar(255);index:idx_unionid;comment:第三方平台UnionID"`
	AppID       string     `json:"appid" gorm:"column:appid;type:varchar(100);not null;comment:应用AppID"`
	Nickname    string     `json:"nickname" gorm:"column:nickname;type:varchar(100);comment:第三方平台昵称"`
	Avatar      string     `json:"avatar" gorm:"column:avatar;type:varchar(500);comment:第三方平台头像"`
	Status      string     `json:"status" gorm:"column:status;type:enum('active','disabled');not null;default:'active';comment:绑定状态"`
	LastLoginAt *time.Time `json:"last_login_at" gorm:"column:last_login_at;type:datetime;comment:最后登录时间"`
}

// TableName 指定表名
func (ThirdPartyBinding) TableName() string {
	return "third_party_bindings"
}
