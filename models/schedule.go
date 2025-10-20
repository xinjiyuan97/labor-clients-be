package models

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Schedule 个人日程表
type Schedule struct {
	BaseModel
	WorkerID        int64     `json:"worker_id" gorm:"column:worker_id;type:bigint;not null;index;comment:零工ID"`
	JobID           *int64    `json:"job_id" gorm:"column:job_id;type:bigint;index;comment:关联岗位ID"`
	Title           string    `json:"title" gorm:"column:title;type:varchar(100);not null;comment:日程标题"`
	StartTime       time.Time `json:"start_time" gorm:"column:start_time;type:datetime;not null;index;comment:开始时间"`
	EndTime         time.Time `json:"end_time" gorm:"column:end_time;type:datetime;not null;comment:结束时间"`
	Location        string    `json:"location" gorm:"column:location;type:varchar(255);comment:地点"`
	Notes           string    `json:"notes" gorm:"column:notes;type:text;comment:备注"`
	Status          string    `json:"status" gorm:"column:status;type:enum('pending','in_progress','completed','cancelled');default:pending;index;comment:状态"`
	ReminderMinutes int       `json:"reminder_minutes" gorm:"column:reminder_minutes;type:int;default:15;comment:提前提醒分钟数"`
}

// TableName 指定表名
func (Schedule) TableName() string {
	return "schedules"
}

func (s *Schedule) ToThrift() *common.ScheduleInfo {
	return &common.ScheduleInfo{
		ScheduleID:      s.ID,
		WorkerID:        s.WorkerID,
		JobID:           utils.GetOrDefault(s.JobID, 0),
		Title:           s.Title,
		StartTime:       s.StartTime.Format(time.RFC3339),
		EndTime:         s.EndTime.Format(time.RFC3339),
		Location:        s.Location,
		Notes:           s.Notes,
		Status:          s.Status,
		ReminderMinutes: int32(s.ReminderMinutes),
	}
}
