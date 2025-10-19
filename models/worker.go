package models

import "github.com/shopspring/decimal"

// Worker 零工详细信息表
type Worker struct {
	BaseModel
	UserID         int64           `json:"user_id" gorm:"column:user_id;type:bigint;not null;index;comment:用户ID"`
	RealName       string          `json:"real_name" gorm:"column:real_name;type:varchar(50);index;comment:真实姓名"`
	Gender         string          `json:"gender" gorm:"column:gender;type:enum('male','female');comment:性别"`
	Age            uint8           `json:"age" gorm:"column:age;type:tinyint unsigned;comment:年龄"`
	IDCard         string          `json:"id_card" gorm:"column:id_card;type:varchar(20);comment:身份证号"`
	HealthCert     string          `json:"health_cert" gorm:"column:health_cert;type:varchar(255);comment:健康证URL"`
	Education      string          `json:"education" gorm:"column:education;type:varchar(50);comment:学历"`
	Height         decimal.Decimal `json:"height" gorm:"column:height;type:decimal(4,1);comment:身高(cm)"`
	Introduction   string          `json:"introduction" gorm:"column:introduction;type:text;comment:个人介绍"`
	WorkExperience string          `json:"work_experience" gorm:"column:work_experience;type:text;comment:工作经历"`
	ExpectedSalary decimal.Decimal `json:"expected_salary" gorm:"column:expected_salary;type:decimal(10,2);comment:期望薪资"`
}

// TableName 指定表名
func (Worker) TableName() string {
	return "workers"
}
