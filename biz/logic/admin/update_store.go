package admin

import (
	"context"

	"github.com/shopspring/decimal"
	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
)

// UpdateStore 更新门店信息
func UpdateStore(ctx context.Context, req *admin.UpdateStoreReq) (*admin.UpdateStoreResp, error) {
	// 检查门店是否存在
	exists, err := mysql.CheckStoreExists(ctx, req.StoreID)
	if err != nil || !exists {
		return &admin.UpdateStoreResp{
			Base: &common.BaseResp{
				Code:    404,
				Message: "门店不存在",
			},
		}, err
	}

	// 构建更新字段
	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Address != "" {
		updates["address"] = req.Address
	}
	if req.Latitude != "" {
		lat, err := decimal.NewFromString(req.Latitude)
		if err == nil {
			updates["latitude"] = lat
		}
	}
	if req.Longitude != "" {
		lng, err := decimal.NewFromString(req.Longitude)
		if err == nil {
			updates["longitude"] = lng
		}
	}
	if req.ContactPhone != "" {
		updates["contact_phone"] = req.ContactPhone
	}
	if req.ContactPerson != "" {
		updates["contact_person"] = req.ContactPerson
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	// 更新门店
	if err := mysql.UpdateStore(ctx, req.StoreID, updates); err != nil {
		return &admin.UpdateStoreResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "更新门店失败",
			},
		}, err
	}

	return &admin.UpdateStoreResp{
		Base: &common.BaseResp{
			Code:    200,
			Message: "更新门店成功",
		},
	}, nil
}
