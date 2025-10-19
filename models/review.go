package models

// Review 评价表
type Review struct {
	BaseModel
	JobID      int64  `json:"job_id" gorm:"column:job_id;type:bigint;not null;comment:岗位ID"`
	EmployerID int64  `json:"employer_id" gorm:"column:employer_id;type:bigint;not null;index;comment:雇主ID"`
	WorkerID   int64  `json:"worker_id" gorm:"column:worker_id;type:bigint;not null;index;comment:零工ID"`
	Rating     int8   `json:"rating" gorm:"column:rating;type:tinyint;not null;check:rating >= 1 AND rating <= 5;index;comment:评分(1-5)"`
	Content    string `json:"content" gorm:"column:content;type:text;comment:评价内容"`
	ReviewType string `json:"review_type" gorm:"column:review_type;type:enum('employer_to_worker','worker_to_employer');not null;comment:评价类型"`
}

// TableName 指定表名
func (Review) TableName() string {
	return "reviews"
}
