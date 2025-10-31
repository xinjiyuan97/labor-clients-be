package admin

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
)

// DeleteBrandAdmin 删除品牌/门店管理员
func DeleteBrandAdmin(ctx context.Context, req *admin.DeleteBrandAdminReq) (*admin.DeleteBrandAdminResp, error) {
	// 验证角色是否存在
	role, err := mysql.GetUserRoleByID(ctx, req.RoleID)
	if err != nil || role == nil {
		return &admin.DeleteBrandAdminResp{
			Base: &common.BaseResp{
				Code:    404,
				Message: "角色不存在",
			},
		}, nil
	}

	// 删除角色
	if err := mysql.DeleteUserRole(ctx, req.RoleID); err != nil {
		return &admin.DeleteBrandAdminResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "删除管理员失败",
			},
		}, err
	}

	return &admin.DeleteBrandAdminResp{
		Base: &common.BaseResp{
			Code:    200,
			Message: "Success",
		},
	}, nil
}
