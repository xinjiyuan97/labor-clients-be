package admin

import (
	"context"
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/constants"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UpdateBrandLogic 更新品牌业务逻辑
func UpdateBrandLogic(ctx context.Context, req *admin.UpdateBrandReq) (*admin.UpdateBrandResp, error) {
	// 1. 验证品牌是否存在
	brand, err := mysql.GetBrandByID(mysql.DB, req.BrandID)
	if err != nil {
		utils.Errorf("查询品牌失败: %v", err)
		return &admin.UpdateBrandResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "查询品牌失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if brand == nil {
		utils.Warnf("品牌不存在, brand_id: %d", req.BrandID)
		return &admin.UpdateBrandResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "品牌不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 2. 更新品牌字段（只更新非空字段）
	if req.CompanyName != "" {
		brand.Name = req.CompanyName
	}
	if req.CompanyShortName != "" {
		brand.CompanyShortName = req.CompanyShortName
	}
	if req.Logo != "" {
		brand.Logo = req.Logo
	}
	if req.Description != "" {
		brand.Description = req.Description
	}
	if req.Website != "" {
		brand.Website = req.Website
	}
	if req.Industry != "" {
		brand.Industry = req.Industry
	}
	if req.CompanySize != "" {
		brand.CompanySize = req.CompanySize
	}
	if req.CreditCode != "" {
		brand.CreditCode = req.CreditCode
	}
	if req.CompanyAddress != "" {
		brand.CompanyAddress = req.CompanyAddress
	}
	if req.BusinessScope != "" {
		brand.BusinessScope = req.BusinessScope
	}
	if req.EstablishedDate != "" {
		// 解析日期
		if date, err := time.Parse("2006-01-02", req.EstablishedDate); err == nil {
			brand.EstablishedDate = &date
		}
	}
	if req.RegisteredCapital > 0 {
		brand.RegisteredCapital = req.RegisteredCapital
	}
	if req.ContactPosition != "" {
		brand.ContactPosition = req.ContactPosition
	}
	if req.IDCardNumber != "" {
		brand.IDCardNumber = req.IDCardNumber
	}
	if req.IDCardFront != "" {
		brand.IDCardFront = req.IDCardFront
	}
	if req.IDCardBack != "" {
		brand.IDCardBack = req.IDCardBack
	}
	if req.BusinessLicense != "" {
		brand.BusinessLicense = req.BusinessLicense
	}
	if req.TaxCertificate != "" {
		brand.TaxCertificate = req.TaxCertificate
	}
	if req.OrgCodeCertificate != "" {
		brand.OrgCodeCertificate = req.OrgCodeCertificate
	}
	if req.BankLicense != "" {
		brand.BankLicense = req.BankLicense
	}
	if req.OtherCertificates != "" {
		brand.OtherCertificates = req.OtherCertificates
	}
	if req.BankAccount != "" {
		brand.BankAccount = req.BankAccount
	}
	if req.SettlementCycle != "" {
		brand.SettlementCycle = req.SettlementCycle
	}
	if req.DepositAmount > 0 {
		brand.DepositAmount = req.DepositAmount
	}
	if req.AuthStatus != "" {
		brand.AuthStatus = constants.BrandAuthStatus(req.AuthStatus)
	}
	if req.AccountStatus != "" {
		brand.AccountStatus = constants.BrandAccountStatus(req.AccountStatus)
	}

	// 更新联系人（如果提供了新的联系电话）
	if req.ContactPhone != "" {
		// 查找或创建联系人用户
		contactUser, err := createBrandAdminForContact(ctx, req.ContactPhone, req.ContactPerson)
		if err != nil {
			utils.Warnf("更新联系人失败: %v", err)
		} else {
			brand.ContactUserID = &contactUser.ID
		}
	}

	// 3. 保存更新
	if err := mysql.UpdateBrand(mysql.DB, brand); err != nil {
		utils.Errorf("更新品牌失败: %v, brand_id: %d", err, req.BrandID)
		return &admin.UpdateBrandResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "更新品牌失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	utils.Infof("品牌更新成功, brand_id: %d", req.BrandID)
	return &admin.UpdateBrandResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "品牌更新成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
