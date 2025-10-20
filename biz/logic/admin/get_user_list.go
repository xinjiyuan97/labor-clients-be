package admin

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetUserListLogic 获取用户列表业务逻辑
func GetUserListLogic(req *admin.GetUserListReq) (*admin.GetUserListResp, error) {
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

	// 获取用户列表
	var users []*models.User
	var total int64
	var err error

	if req.Role != "" {
		// 根据角色获取用户
		users, err = mysql.GetUsersByRole(nil, req.Role, offset, limit)
		if err != nil {
			utils.Errorf("根据角色获取用户列表失败: %v", err)
			return &admin.GetUserListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountUsersByRole(nil, req.Role)
	} else {
		// 获取所有用户
		users, err = mysql.GetUsers(nil, offset, limit)
		if err != nil {
			utils.Errorf("获取用户列表失败: %v", err)
			return &admin.GetUserListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountUsers(nil)
	}

	if err != nil {
		utils.Errorf("获取用户总数失败: %v", err)
		return &admin.GetUserListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建用户信息
	var userInfos []*admin.BrandUserInfo
	for _, user := range users {
		userInfo := &admin.BrandUserInfo{
			UserID:   user.ID,
			Username: user.Username,
		}
		userInfos = append(userInfos, userInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &admin.GetUserListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取用户列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageInfo: pageResp,
		Users:    userInfos,
	}, nil
}
