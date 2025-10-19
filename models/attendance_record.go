package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// AttendanceRecord 考勤记录表
type AttendanceRecord struct {
	BaseModel
	JobID            int64           `json:"job_id" gorm:"column:job_id;type:bigint;not null;index:idx_job_worker;comment:岗位ID"`
	WorkerID         int64           `json:"worker_id" gorm:"column:worker_id;type:bigint;not null;index:idx_job_worker;comment:零工ID"`
	CheckIn          *time.Time      `json:"check_in" gorm:"column:check_in;type:datetime;index;comment:打卡时间"`
	CheckOut         *time.Time      `json:"check_out" gorm:"column:check_out;type:datetime;comment:签退时间"`
	WorkHours        decimal.Decimal `json:"work_hours" gorm:"column:work_hours;type:decimal(4,2);comment:工作时长"`
	CheckInLocation  string          `json:"check_in_location" gorm:"column:check_in_location;type:varchar(255);comment:打卡位置"`
	CheckOutLocation string          `json:"check_out_location" gorm:"column:check_out_location;type:varchar(255);comment:签退位置"`
	Status           string          `json:"status" gorm:"column:status;type:enum('normal','late','early_leave','absent','leave');default:normal;comment:考勤状态"`
}

// TableName 指定表名
func (AttendanceRecord) TableName() string {
	return "attendance_records"
}
