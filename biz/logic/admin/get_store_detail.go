package admin

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
)

// GetStoreDetail 获取门店详情
func GetStoreDetail(ctx context.Context, req *admin.GetStoreDetailReq) (*admin.GetStoreDetailResp, error) {
	store, err := mysql.GetStoreByID(ctx, req.StoreID)
	if err != nil {
		return &admin.GetStoreDetailResp{
			Base: &common.BaseResp{
				Code:    404,
				Message: "门店不存在",
			},
		}, err
	}

	return &admin.GetStoreDetailResp{
		Base: &common.BaseResp{
			Code:    0,
			Message: "success",
		},
		StoreInfo: store.ToThriftStore(),
	}, nil
}
