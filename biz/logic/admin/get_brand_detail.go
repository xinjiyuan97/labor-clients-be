package admin

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
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
