package models

// JobTag 岗位标签表
type JobTag struct {
	BaseModel
	JobID   int64  `json:"job_id" gorm:"column:job_id;type:bigint;not null;index;comment:岗位ID"`
	TagName string `json:"tag_name" gorm:"column:tag_name;type:varchar(50);not null;comment:标签名称"`
	TagType string `json:"tag_type" gorm:"column:tag_type;type:varchar(20);not null;index;comment:标签类型"`
}

// TableName 指定表名
func (JobTag) TableName() string {
	return "job_tags"
}
