package admin

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetAdminListLogic 获取管理员列表业务逻辑
func GetAdminListLogic(req *admin.GetAdminListReq) (*admin.GetAdminListResp, error) {
	// 设置默认分页参数
	page := 1
	limit := 10
	if req.Page > 0 {
		page = int(req.Page)
	}
	if req.Limit > 0 {
		limit = int(req.Limit)
	}

	offset := (page - 1) * limit

	// 获取管理员列表
	admins, err := mysql.GetAdmins(nil, offset, limit)
	if err != nil {
		utils.Errorf("获取管理员列表失败: %v", err)
		return &admin.GetAdminListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取总数
	total, err := mysql.CountAdmins(nil)
	if err != nil {
		utils.Errorf("获取管理员总数失败: %v", err)
		return &admin.GetAdminListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建管理员信息
	var adminInfos []*admin.AdminInfo
	for _, adminUser := range admins {
		adminInfo := adminUser.ToThriftAdmin()
		adminInfos = append(adminInfos, adminInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &admin.GetAdminListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取管理员列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageInfo: pageResp,
		Admins:   adminInfos,
	}, nil
}
