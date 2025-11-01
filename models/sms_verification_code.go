package models

// SMSVerificationCode 短信验证码表
type SMSVerificationCode struct {
	BaseModel
	Phone     string  `json:"phone" gorm:"column:phone;type:varchar(20);not null;index;comment:手机号"`
	Code      string  `json:"code" gorm:"column:code;type:varchar(10);not null;comment:验证码"`
	Status    string  `json:"status" gorm:"column:status;type:enum('unused','used','expired');not null;default:'unused';index;comment:状态"`
	ExpiresAt string  `json:"expires_at" gorm:"column:expires_at;type:datetime;not null;index;comment:过期时间"`
	UsedAt    *string `json:"used_at" gorm:"column:used_at;type:datetime;comment:使用时间"`
}

// TableName 指定表名
func (SMSVerificationCode) TableName() string {
	return "sms_verification_codes"
}

