package models

// UserFavoriteJob 用户收藏岗位表
type UserFavoriteJob struct {
	BaseModel
	UserID int64 `json:"user_id" gorm:"column:user_id;type:bigint;not null;uniqueIndex:uk_user_job;index;comment:用户ID"`
	JobID  int64 `json:"job_id" gorm:"column:job_id;type:bigint;not null;uniqueIndex:uk_user_job;comment:岗位ID"`
}

// TableName 指定表名
func (UserFavoriteJob) TableName() string {
	return "user_favorite_jobs"
}
