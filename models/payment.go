package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// Payment 支付记录表
type Payment struct {
	BaseModel
	JobID         int64           `json:"job_id" gorm:"column:job_id;type:bigint;not null;comment:岗位ID"`
	WorkerID      int64           `json:"worker_id" gorm:"column:worker_id;type:bigint;not null;index;comment:零工ID"`
	EmployerID    int64           `json:"employer_id" gorm:"column:employer_id;type:bigint;not null;index;comment:雇主ID"`
	Amount        decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(10,2);not null;comment:支付金额"`
	PaymentMethod string          `json:"payment_method" gorm:"column:payment_method;type:varchar(20);not null;comment:支付方式"`
	Status        string          `json:"status" gorm:"column:status;type:enum('pending','processing','completed','failed');default:pending;index;comment:支付状态"`
	PaidAt        *time.Time      `json:"paid_at" gorm:"column:paid_at;type:datetime;comment:支付时间"`
	PlatformFee   decimal.Decimal `json:"platform_fee" gorm:"column:platform_fee;type:decimal(10,2);default:0;comment:平台费用"`
}

// TableName 指定表名
func (Payment) TableName() string {
	return "payments"
}
