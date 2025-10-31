package admin

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
)

// UpdateBrandAdminStatus 更新品牌/门店管理员状态
func UpdateBrandAdminStatus(ctx context.Context, req *admin.UpdateBrandAdminStatusReq) (*admin.UpdateBrandAdminStatusResp, error) {
	// 验证角色是否存在
	role, err := mysql.GetUserRoleByID(ctx, req.RoleID)
	if err != nil || role == nil {
		return &admin.UpdateBrandAdminStatusResp{
			Base: &common.BaseResp{
				Code:    404,
				Message: "角色不存在",
			},
		}, nil
	}

	// 验证状态值
	if req.Status != "active" && req.Status != "disabled" {
		return &admin.UpdateBrandAdminStatusResp{
			Base: &common.BaseResp{
				Code:    400,
				Message: "无效的状态值",
			},
		}, nil
	}

	// 更新状态
	updates := map[string]interface{}{
		"status": req.Status,
	}

	if err := mysql.UpdateUserRole(ctx, req.RoleID, updates); err != nil {
		return &admin.UpdateBrandAdminStatusResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "更新状态失败",
			},
		}, err
	}

	return &admin.UpdateBrandAdminStatusResp{
		Base: &common.BaseResp{
			Code:    200,
			Message: "success",
		},
	}, nil
}
