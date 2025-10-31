package admin

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/config"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetBrandListLogic 获取品牌列表业务逻辑
func GetBrandListLogic(req *admin.GetBrandListReq) (*admin.GetBrandListResp, error) {
	// 验证分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	// 计算偏移量
	offset := (req.Page - 1) * req.Limit

	// 查询品牌列表
	brands, err := mysql.GetBrandsForAdmin(nil, int(offset), int(req.Limit))
	if err != nil {
		utils.Errorf("查询品牌列表失败: %v", err)
		return &admin.GetBrandListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "查询品牌列表失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 统计总数
	total, err := mysql.CountBrandsForAdmin(nil)
	if err != nil {
		utils.Errorf("统计品牌总数失败: %v", err)
		return &admin.GetBrandListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "统计品牌总数失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取上传服务用于签名URL
	uploadService, err := utils.GetUploadService(&config.GetGlobalConfig().OSS)
	if err != nil {
		utils.Warnf("获取上传服务失败，将使用原始URL: %v", err)
	}

	// 转换为响应对象
	brandList := make([]*admin.BrandDetail, 0, len(brands))
	for _, brand := range brands {
		brandDetail := brand.ToThriftBrand()

		// 如果logo不为空且上传服务可用，则生成签名URL
		if brandDetail.Logo != "" && uploadService != nil {
			signedURL, signErr := uploadService.GetSignedURL(brandDetail.Logo, 3600) // 1小时有效期
			if signErr != nil {
				utils.Warnf("为品牌 %d 的logo生成签名URL失败: %v, 使用原始URL", brand.ID, signErr)
			} else {
				brandDetail.Logo = signedURL
			}
		}

		brandList = append(brandList, brandDetail)
	}

	utils.Infof("获取品牌列表成功, page: %d, limit: %d, total: %d", req.Page, req.Limit, total)
	return &admin.GetBrandListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取品牌列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageInfo: &common.PageResp{
			Total: int32(total),
			Page:  req.Page,
			Limit: req.Limit,
		},
		Brands: brandList,
	}, nil
}
