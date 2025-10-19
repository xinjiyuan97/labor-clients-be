package models

// Message 消息表
type Message struct {
	BaseModel
	FromUser    int64  `json:"from_user" gorm:"column:from_user;type:bigint;not null;index;comment:发送用户ID"`
	ToUser      int64  `json:"to_user" gorm:"column:to_user;type:bigint;not null;index;comment:接收用户ID"`
	MessageType string `json:"message_type" gorm:"column:message_type;type:varchar(20);not null;comment:消息类型"`
	Content     string `json:"content" gorm:"column:content;type:text;not null;comment:消息内容"`
	MsgCategory string `json:"msg_category" gorm:"column:msg_category;type:enum('system','chat','community');default:chat;index;comment:消息分类"`
	IsRead      bool   `json:"is_read" gorm:"column:is_read;type:boolean;default:false;comment:是否已读"`
}

// TableName 指定表名
func (Message) TableName() string {
	return "messages"
}
