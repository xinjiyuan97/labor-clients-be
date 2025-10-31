package admin

import (
	"context"

	"github.com/shopspring/decimal"
	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
)

// CreateStore 创建门店
func CreateStore(ctx context.Context, req *admin.CreateStoreReq) (*admin.CreateStoreResp, error) {
	// 创建门店对象
	store := &models.Store{
		BrandID:       req.BrandID,
		Name:          req.Name,
		Address:       req.Address,
		ContactPhone:  req.ContactPhone,
		ContactPerson: req.ContactPerson,
		Description:   req.Description,
		Status:        "active",
	}

	// 处理经纬度
	if req.Latitude != "" {
		lat, err := decimal.NewFromString(req.Latitude)
		if err == nil {
			store.Latitude = lat
		}
	}
	if req.Longitude != "" {
		lng, err := decimal.NewFromString(req.Longitude)
		if err == nil {
			store.Longitude = lng
		}
	}

	// 创建门店
	if err := mysql.CreateStore(ctx, store); err != nil {
		return &admin.CreateStoreResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "创建门店失败",
			},
		}, err
	}

	return &admin.CreateStoreResp{
		Base: &common.BaseResp{
			Code:    200,
			Message: "创建门店成功",
		},
		StoreID: store.ID,
	}, nil
}
