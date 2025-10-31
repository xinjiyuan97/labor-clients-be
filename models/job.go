package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// Job 岗位信息表
type Job struct {
	BaseModel
	EmployerID     int64           `json:"employer_id" gorm:"column:employer_id;type:bigint;not null;index;comment:雇主ID"`
	BrandID        int64           `json:"brand_id" gorm:"column:brand_id;type:bigint;not null;index;comment:所属品牌ID"`
	StoreID        *int64          `json:"store_id" gorm:"column:store_id;type:bigint;index;comment:所属门店ID"`
	CategoryID     int64           `json:"category_id" gorm:"column:category_id;type:bigint;not null;index;comment:分类ID"`
	Title          string          `json:"title" gorm:"column:title;type:varchar(100);not null;comment:岗位标题"`
	JobType        string          `json:"job_type" gorm:"column:job_type;type:enum('standard','rush','transfer');not null;default:standard;index;comment:岗位类型"`
	Description    string          `json:"description" gorm:"column:description;type:text;comment:岗位描述"`
	Salary         decimal.Decimal `json:"salary" gorm:"column:salary;type:decimal(10,2);not null;comment:薪资"`
	SalaryUnit     string          `json:"salary_unit" gorm:"column:salary_unit;type:varchar(20);default:天;comment:结算单位"`
	Location       string          `json:"location" gorm:"column:location;type:varchar(255);not null;comment:工作地点"`
	Latitude       decimal.Decimal `json:"latitude" gorm:"column:latitude;type:decimal(10,8);comment:纬度"`
	Longitude      decimal.Decimal `json:"longitude" gorm:"column:longitude;type:decimal(11,8);comment:经度"`
	Requirements   string          `json:"requirements" gorm:"column:requirements;type:text;comment:工作要求"`
	Benefits       string          `json:"benefits" gorm:"column:benefits;type:text;comment:福利待遇"`
	StartTime      time.Time       `json:"start_time" gorm:"column:start_time;type:datetime;not null;index;comment:开始时间"`
	EndTime        time.Time       `json:"end_time" gorm:"column:end_time;type:datetime;not null;comment:结束时间"`
	Status         string          `json:"status" gorm:"column:status;type:enum('draft','published','filled','completed','cancelled');default:draft;index;comment:岗位状态"`
	MaxApplicants  int             `json:"max_applicants" gorm:"column:max_applicants;type:int;not null;default:1;comment:该岗位最大招募人数"`
	ApplicantCount int             `json:"applicant_count" gorm:"column:applicant_count;type:int;default:0;comment:报名人数"`
}

// TableName 指定表名
func (Job) TableName() string {
	return "jobs"
}
