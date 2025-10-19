package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/middleware"
)

// GetCurrentUserProfile 获取当前登录用户的信息（需要鉴权）
// @router /api/v1/auth/me [GET]
func GetCurrentUserProfile(ctx context.Context, c *app.RequestContext) {
	// 从上下文中获取用户ID
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(consts.StatusUnauthorized, &common.BaseResp{
			Code:      401,
			Message:   "未登录",
			Timestamp: middleware.GetCurrentTime(),
		})
		return
	}

	// 从上下文中获取用户角色
	userRole, _ := middleware.GetUserRoleFromContext(c)

	// 构建响应
	resp := map[string]interface{}{
		"base": &common.BaseResp{
			Code:      200,
			Message:   "获取用户信息成功",
			Timestamp: middleware.GetCurrentTime(),
		},
		"user_id":   userID,
		"user_role": userRole,
	}

	c.JSON(consts.StatusOK, resp)
}

// UpdateUserProfile 更新用户信息（需要worker角色）
// @router /api/v1/auth/profile [PUT]
func UpdateUserProfile(ctx context.Context, c *app.RequestContext) {
	// 从上下文中获取用户ID
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		c.JSON(consts.StatusUnauthorized, &common.BaseResp{
			Code:      401,
			Message:   "未登录",
			Timestamp: middleware.GetCurrentTime(),
		})
		return
	}

	// 从上下文中获取用户角色
	userRole, _ := middleware.GetUserRoleFromContext(c)

	// 检查角色权限
	if userRole != "worker" {
		c.JSON(consts.StatusForbidden, &common.BaseResp{
			Code:      403,
			Message:   "只有worker可以更新个人信息",
			Timestamp: middleware.GetCurrentTime(),
		})
		return
	}

	// 这里可以添加更新用户信息的逻辑
	resp := map[string]interface{}{
		"base": &common.BaseResp{
			Code:      200,
			Message:   "更新用户信息成功",
			Timestamp: middleware.GetCurrentTime(),
		},
		"user_id": userID,
	}

	c.JSON(consts.StatusOK, resp)
}
