package admin

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
)

// GetStoreList 获取门店列表
func GetStoreList(ctx context.Context, req *admin.GetStoreListReq) (*admin.GetStoreListResp, error) {
	offset := (req.Page - 1) * req.Limit

	var brandID *int64
	if req.BrandID > 0 {
		brandID = &req.BrandID
	}

	stores, total, err := mysql.GetStoreList(ctx, brandID, req.Status, req.Name, int(offset), int(req.Limit))
	if err != nil {
		return &admin.GetStoreListResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "获取门店列表失败",
			},
		}, err
	}

	// 转换为Thrift对象
	storeList := make([]*admin.StoreDetail, 0, len(stores))
	for _, store := range stores {
		storeList = append(storeList, store.ToThriftStore())
	}

	return &admin.GetStoreListResp{
		Base: &common.BaseResp{
			Code:    200,
			Message: "获取门店列表成功",
		},
		PageInfo: &common.PageResp{
			Total: int32(total),
			Page:  req.Page,
			Limit: req.Limit,
		},
		Stores: storeList,
	}, nil
}
