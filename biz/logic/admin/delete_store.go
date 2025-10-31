package admin

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
)

// DeleteStore 删除门店
func DeleteStore(ctx context.Context, req *admin.DeleteStoreReq) (*admin.DeleteStoreResp, error) {
	// 检查门店是否存在
	exists, err := mysql.CheckStoreExists(ctx, req.StoreID)
	if err != nil || !exists {
		return &admin.DeleteStoreResp{
			Base: &common.BaseResp{
				Code:    404,
				Message: "门店不存在",
			},
		}, err
	}

	// 删除门店（软删除）
	if err := mysql.DeleteStore(ctx, req.StoreID); err != nil {
		return &admin.DeleteStoreResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "删除门店失败",
			},
		}, err
	}

	return &admin.DeleteStoreResp{
		Base: &common.BaseResp{
			Code:    200,
			Message: "删除门店成功",
		},
	}, nil
}
