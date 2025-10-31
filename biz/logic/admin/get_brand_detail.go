package admin

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/config"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetBrandDetailLogic 获取品牌详情业务逻辑
func GetBrandDetailLogic(brandID int64) (*admin.GetBrandDetailResp, error) {
	// 验证品牌ID
	if brandID <= 0 {
		utils.Warn("品牌ID无效")
		return &admin.GetBrandDetailResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "品牌ID无效",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 查询品牌信息
	brand, err := mysql.GetBrandByID(nil, brandID)
	if err != nil {
		utils.Errorf("查询品牌信息失败: %v", err)
		return &admin.GetBrandDetailResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "查询品牌信息失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if brand == nil {
		utils.Warnf("品牌不存在, brand_id: %d", brandID)
		return &admin.GetBrandDetailResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "品牌不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 转换为响应对象
	brandDetail := brand.ToThriftBrand()

	// Signed logo and cert urls
	uploadService, err := utils.GetUploadService(&config.GetGlobalConfig().OSS)
	if err != nil {
		utils.Errorf("获取上传服务失败: %v", err)
		return &admin.GetBrandDetailResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "获取上传服务失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}
	if brandDetail.Logo != "" {
		brandDetail.Logo, err = uploadService.GetSignedURL(brandDetail.Logo, 3600)
		if err != nil {
			utils.Errorf("获取签名URL失败: %v", err)
		}
	}
	if brandDetail.BusinessLicense != "" {
		brandDetail.BusinessLicense, err = uploadService.GetSignedURL(brandDetail.BusinessLicense, 3600)
		if err != nil {
			utils.Errorf("获取签名URL失败: %v", err)
		}
	}

	// 查询联系人信息
	if brand.ContactUserID != nil && *brand.ContactUserID > 0 {
		contactUser, err := mysql.GetUserByID(mysql.DB, *brand.ContactUserID)
		if err != nil {
			utils.Warnf("查询联系人信息失败: %v, user_id: %d", err, *brand.ContactUserID)
		} else if contactUser != nil {
			brandDetail.ContactPerson = contactUser.Username
			brandDetail.ContactPhone = contactUser.Phone
			// 如果用户表有email字段，可以在这里设置
			// brandDetail.ContactEmail = contactUser.Email
		}
	}

	utils.Infof("获取品牌详情成功, brand_id: %d", brandID)
	return &admin.GetBrandDetailResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取品牌详情成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		BrandInfo: brandDetail,
	}, nil
}
