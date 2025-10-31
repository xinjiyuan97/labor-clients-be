package admin

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
)

// GetBrandAdmins 获取品牌管理员列表
func GetBrandAdmins(ctx context.Context, req *admin.GetBrandAdminsReq) (*admin.GetBrandAdminsResp, error) {
	offset := (req.Page - 1) * req.Limit

	// 查询品牌管理员列表
	admins, total, err := mysql.GetBrandAdminsList(ctx, req.BrandID, req.RoleType, req.Status, int(offset), int(req.Limit))
	if err != nil {
		return &admin.GetBrandAdminsResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "查询品牌管理员列表失败",
			},
		}, err
	}

	// 转换为Thrift结构
	adminList := make([]*admin.BrandAdminInfo, 0, len(admins))
	for _, a := range admins {
		adminInfo := &admin.BrandAdminInfo{
			UserID:    a.UserID,
			RoleID:    a.RoleID,
			Username:  a.Username,
			Phone:     a.Phone,
			RoleType:  a.RoleType,
			Status:    a.Status,
			CreatedAt: a.CreatedAt,
		}

		if a.BrandID != nil {
			adminInfo.BrandID = *a.BrandID
			adminInfo.BrandName = a.BrandName
		}

		if a.StoreID != nil {
			adminInfo.StoreID = *a.StoreID
			adminInfo.StoreName = a.StoreName
		}

		adminList = append(adminList, adminInfo)
	}

	return &admin.GetBrandAdminsResp{
		Base: &common.BaseResp{
			Code:    200,
			Message: "获取品牌管理员列表成功",
		},
		PageInfo: &common.PageResp{
			Total: int32(total),
			Page:  req.Page,
			Limit: req.Limit,
		},
		Admins: adminList,
	}, nil
}
