package models

import "time"

// JobApplication 岗位申请表
type JobApplication struct {
	BaseModel
	JobID          int64      `json:"job_id" gorm:"column:job_id;type:bigint;not null;uniqueIndex:uk_job_worker;comment:岗位ID"`
	WorkerID       int64      `json:"worker_id" gorm:"column:worker_id;type:bigint;not null;uniqueIndex:uk_job_worker;index;comment:零工ID"`
	Status         string     `json:"status" gorm:"column:status;type:enum('applied','confirmed','rejected','cancelled','completed');default:applied;index;comment:申请状态"`
	AppliedAt      time.Time  `json:"applied_at" gorm:"column:applied_at;type:timestamp;default:CURRENT_TIMESTAMP;index;comment:申请时间"`
	ConfirmedAt    *time.Time `json:"confirmed_at" gorm:"column:confirmed_at;type:timestamp;comment:确认时间"`
	CancelReason   string     `json:"cancel_reason" gorm:"column:cancel_reason;type:text;comment:取消原因"`
	WorkerRating   int8       `json:"worker_rating" gorm:"column:worker_rating;type:tinyint;comment:零工评分"`
	EmployerRating int8       `json:"employer_rating" gorm:"column:employer_rating;type:tinyint;comment:雇主评分"`
	Review         string     `json:"review" gorm:"column:review;type:text;comment:评价内容"`
}

// TableName 指定表名
func (JobApplication) TableName() string {
	return "job_applications"
}
