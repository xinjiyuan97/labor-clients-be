package models

import "time"

// Feedback 用户反馈表
type Feedback struct {
	BaseModel
	UserID    int64      `json:"user_id" gorm:"column:user_id;type:bigint;not null;index;comment:用户ID"`
	Type      string     `json:"type" gorm:"column:type;type:varchar(20);not null;comment:反馈类型"`
	Content   string     `json:"content" gorm:"column:content;type:text;not null;comment:反馈内容"`
	Contact   string     `json:"contact" gorm:"column:contact;type:varchar(100);comment:联系方式"`
	Status    string     `json:"status" gorm:"column:status;type:enum('pending','processing','resolved','closed');default:pending;index;comment:处理状态"`
	Reply     string     `json:"reply" gorm:"column:reply;type:text;comment:回复内容"`
	ReplyTime *time.Time `json:"reply_time" gorm:"column:reply_time;type:datetime;comment:回复时间"`
}

// TableName 指定表名
func (Feedback) TableName() string {
	return "feedbacks"
}

// Notice 系统通知表
type Notice struct {
	BaseModel
	Title     string     `json:"title" gorm:"column:title;type:varchar(200);not null;comment:通知标题"`
	Content   string     `json:"content" gorm:"column:content;type:text;not null;comment:通知内容"`
	Type      string     `json:"type" gorm:"column:type;type:varchar(20);not null;comment:通知类型"`
	Priority  string     `json:"priority" gorm:"column:priority;type:enum('low','normal','high','urgent');default:normal;comment:优先级"`
	Status    string     `json:"status" gorm:"column:status;type:enum('draft','published','archived');default:draft;index;comment:状态"`
	StartTime *time.Time `json:"start_time" gorm:"column:start_time;type:datetime;comment:开始时间"`
	EndTime   *time.Time `json:"end_time" gorm:"column:end_time;type:datetime;comment:结束时间"`
}

// TableName 指定表名
func (Notice) TableName() string {
	return "notices"
}

// VersionInfo 版本信息
type VersionInfo struct {
	Version     string    `json:"version" gorm:"column:version;type:varchar(20);not null;comment:版本号"`
	BuildNumber string    `json:"build_number" gorm:"column:build_number;type:varchar(20);not null;comment:构建号"`
	MinVersion  string    `json:"min_version" gorm:"column:min_version;type:varchar(20);not null;comment:最低支持版本"`
	ForceUpdate bool      `json:"force_update" gorm:"column:force_update;type:boolean;default:false;comment:是否强制更新"`
	UpdateURL   string    `json:"update_url" gorm:"column:update_url;type:varchar(255);comment:更新下载地址"`
	UpdateNote  string    `json:"update_note" gorm:"column:update_note;type:text;comment:更新说明"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null;comment:创建时间"`
}

// TableName 指定表名
func (VersionInfo) TableName() string {
	return "version_info"
}
