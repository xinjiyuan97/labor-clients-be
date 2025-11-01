package models

// WeChatBinding 微信账号绑定表
type WeChatBinding struct {
	BaseModel
	UserID     int64  `json:"user_id" gorm:"column:user_id;type:bigint;not null;index:idx_user_id;comment:用户ID"`
	OpenID     string `json:"openid" gorm:"column:openid;type:varchar(255);not null;uniqueIndex:idx_openid;comment:微信OpenID"`
	UnionID    string `json:"unionid" gorm:"column:unionid;type:varchar(255);index:idx_unionid;comment:微信UnionID"`
	AppID      string `json:"appid" gorm:"column:appid;type:varchar(100);not null;comment:小程序AppID"`
	Nickname   string `json:"nickname" gorm:"column:nickname;type:varchar(100);comment:微信昵称"`
	Avatar     string `json:"avatar" gorm:"column:avatar;type:varchar(500);comment:微信头像"`
	Status     string `json:"status" gorm:"column:status;type:enum('active','disabled');not null;default:'active';comment:绑定状态"`
	LastLoginAt string `json:"last_login_at" gorm:"column:last_login_at;type:datetime;comment:最后登录时间"`
}

// TableName 指定表名
func (WeChatBinding) TableName() string {
	return "wechat_bindings"
}

