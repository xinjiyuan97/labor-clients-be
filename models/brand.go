package models

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/constants"
)

// Brand 品牌信息表
type Brand struct {
	BaseModel

	// 基本信息
	Name             string `json:"name" gorm:"column:name;type:varchar(100);not null;index;comment:品牌名称（公司全称）"`
	CompanyShortName string `json:"company_short_name" gorm:"column:company_short_name;type:varchar(50);comment:公司简称"`
	Logo             string `json:"logo" gorm:"column:logo;type:varchar(255);comment:品牌Logo URL"`
	Description      string `json:"description" gorm:"column:description;type:text;comment:品牌描述"`
	Website          string `json:"website" gorm:"column:website;type:varchar(255);comment:公司网站"`

	// 公司信息
	Industry          string     `json:"industry" gorm:"column:industry;type:varchar(50);comment:所属行业"`
	CompanySize       string     `json:"company_size" gorm:"column:company_size;type:varchar(20);comment:公司规模"`
	CreditCode        string     `json:"credit_code" gorm:"column:credit_code;type:varchar(50);index;comment:统一社会信用代码"`
	CompanyAddress    string     `json:"company_address" gorm:"column:company_address;type:varchar(255);comment:公司地址"`
	BusinessScope     string     `json:"business_scope" gorm:"column:business_scope;type:text;comment:经营范围"`
	EstablishedDate   *time.Time `json:"established_date" gorm:"column:established_date;type:date;comment:成立日期"`
	RegisteredCapital float64    `json:"registered_capital" gorm:"column:registered_capital;type:decimal(15,2);comment:注册资本"`

	// 联系人信息（存储user_id）
	ContactUserID   *int64 `json:"contact_user_id" gorm:"column:contact_user_id;type:bigint;index;comment:联系人用户ID"`
	ContactPosition string `json:"contact_position" gorm:"column:contact_position;type:varchar(50);comment:联系人职位"`

	// 证件信息
	IDCardNumber       string `json:"id_card_number" gorm:"column:id_card_number;type:varchar(50);comment:身份证号"`
	IDCardFront        string `json:"id_card_front" gorm:"column:id_card_front;type:varchar(255);comment:身份证正面照URL"`
	IDCardBack         string `json:"id_card_back" gorm:"column:id_card_back;type:varchar(255);comment:身份证反面照URL"`
	BusinessLicense    string `json:"business_license" gorm:"column:business_license;type:varchar(255);comment:营业执照URL"`
	TaxCertificate     string `json:"tax_certificate" gorm:"column:tax_certificate;type:varchar(255);comment:税务登记证URL"`
	OrgCodeCertificate string `json:"org_code_certificate" gorm:"column:org_code_certificate;type:varchar(255);comment:组织机构代码证URL"`
	BankLicense        string `json:"bank_license" gorm:"column:bank_license;type:varchar(255);comment:开户许可证URL"`
	OtherCertificates  string `json:"other_certificates" gorm:"column:other_certificates;type:text;comment:其他证件URL（JSON数组）"`

	// 财务信息
	BankAccount     string  `json:"bank_account" gorm:"column:bank_account;type:varchar(50);comment:银行账号"`
	SettlementCycle string  `json:"settlement_cycle" gorm:"column:settlement_cycle;type:varchar(20);comment:结算周期"`
	DepositAmount   float64 `json:"deposit_amount" gorm:"column:deposit_amount;type:decimal(15,2);comment:保证金金额"`

	// 状态信息
	AuthStatus    constants.BrandAuthStatus    `json:"auth_status" gorm:"column:auth_status;type:enum('pending','approved','rejected');default:pending;index;comment:认证状态"`
	AccountStatus constants.BrandAccountStatus `json:"account_status" gorm:"column:account_status;type:enum('active','disabled','frozen');default:active;index;comment:账号状态"`
}

// TableName 指定表名
func (Brand) TableName() string {
	return "brands"
}

// ToThriftBrand 转换为Thrift品牌信息
func (b *Brand) ToThriftBrand() *admin.BrandDetail {
	detail := &admin.BrandDetail{
		BrandID:            b.ID,
		CompanyName:        b.Name,
		CompanyShortName:   b.CompanyShortName,
		Logo:               b.Logo,
		Description:        b.Description,
		Website:            b.Website,
		Industry:           b.Industry,
		CompanySize:        b.CompanySize,
		CreditCode:         b.CreditCode,
		CompanyAddress:     b.CompanyAddress,
		BusinessScope:      b.BusinessScope,
		RegisteredCapital:  b.RegisteredCapital,
		ContactPosition:    b.ContactPosition,
		IDCardNumber:       b.IDCardNumber,
		IDCardFront:        b.IDCardFront,
		IDCardBack:         b.IDCardBack,
		BusinessLicense:    b.BusinessLicense,
		TaxCertificate:     b.TaxCertificate,
		OrgCodeCertificate: b.OrgCodeCertificate,
		BankLicense:        b.BankLicense,
		OtherCertificates:  b.OtherCertificates,
		BankAccount:        b.BankAccount,
		SettlementCycle:    b.SettlementCycle,
		DepositAmount:      b.DepositAmount,
		AuthStatus:         string(b.AuthStatus),
		AccountStatus:      string(b.AccountStatus),
		CreatedAt:          b.CreatedAt.Format(time.DateTime),
		UpdatedAt:          b.UpdatedAt.Format(time.DateTime),
	}

	// 转换成立日期
	if b.EstablishedDate != nil {
		detail.EstablishedDate = b.EstablishedDate.Format("2006-01-02")
	}

	// TODO: 从联系人用户ID获取联系人信息
	// 这部分逻辑需要在业务层实现，查询user表获取联系人姓名、电话、邮箱

	return detail
}

func (b *Brand) ToBrandInfo() *common.BrandInfo {
	return &common.BrandInfo{
		BrandID:     b.ID,
		Name:        b.Name,
		Logo:        b.Logo,
		Description: b.Description,
		AuthStatus:  string(b.AuthStatus),
	}
}
