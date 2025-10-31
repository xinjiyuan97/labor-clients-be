package job

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/job"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetBrandListLogic 获取品牌列表业务逻辑
func GetBrandListLogic(req *job.GetBrandListReq) (*job.GetBrandListResp, error) {
	// 设置默认分页参数
	page := 1
	limit := 10
	if req.PageReq != nil {
		if req.PageReq.Page > 0 {
			page = int(req.PageReq.Page)
		}
		if req.PageReq.Limit > 0 {
			limit = int(req.PageReq.Limit)
		}
	}

	offset := (page - 1) * limit

	// 获取品牌列表
	var brands []*models.Brand
	var total int64
	var err error

	if req.Name != "" && req.AuthStatus != "" {
		// 根据名称和认证状态获取
		brands, err = mysql.GetBrandsByAuthStatus(nil, req.AuthStatus, offset, limit)
		if err != nil {
			utils.Errorf("根据认证状态获取品牌列表失败: %v", err)
			return &job.GetBrandListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountBrandsByAuthStatus(nil, req.AuthStatus)
	} else if req.Name != "" {
		// 根据名称搜索
		brands, err = mysql.GetBrandsByName(nil, req.Name, offset, limit)
		if err != nil {
			utils.Errorf("根据名称搜索品牌失败: %v", err)
			return &job.GetBrandListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountBrandsByName(nil, req.Name)
	} else if req.AuthStatus != "" {
		// 根据认证状态获取
		brands, err = mysql.GetBrandsByAuthStatus(nil, req.AuthStatus, offset, limit)
		if err != nil {
			utils.Errorf("根据认证状态获取品牌列表失败: %v", err)
			return &job.GetBrandListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountBrandsByAuthStatus(nil, req.AuthStatus)
	} else {
		// 获取所有已认证品牌
		brands, err = mysql.GetBrands(nil, offset, limit)
		if err != nil {
			utils.Errorf("获取品牌列表失败: %v", err)
			return &job.GetBrandListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountBrands(nil)
	}

	if err != nil {
		utils.Errorf("获取品牌总数失败: %v", err)
		return &job.GetBrandListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建品牌信息
	var brandInfos []*common.BrandInfo
	for _, brand := range brands {
		brandInfo := brand.ToBrandInfo()
		brandInfos = append(brandInfos, brandInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &job.GetBrandListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取品牌列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Brands:   brandInfos,
	}, nil
}
