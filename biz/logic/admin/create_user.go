package admin

import (
	"context"
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/errno"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CreateUserLogic 创建用户业务逻辑
func CreateUserLogic(ctx context.Context, req *admin.CreateUserReq) (*admin.CreateUserResp, error) {
	// 检查手机号是否已存在
	existingUser, err := mysql.GetUserByPhone(nil, req.Phone)
	if err != nil {
		utils.Errorf("检查手机号失败: %v", err)
		return &admin.CreateUserResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if existingUser != nil {
		return &admin.CreateUserResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "手机号已存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.Errorf("密码加密失败: %v", err)
		return &admin.CreateUserResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 创建用户
	user := &models.User{
		Phone:        req.Phone,
		Username:     req.Phone, // 使用手机号作为用户名
		PasswordHash: hashedPassword,
		Role:         req.Role,
	}

	err = mysql.CreateUser(ctx, user)

	if err != nil {
		utils.Errorf("创建用户失败: %v", err)
		return &admin.CreateUserResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "创建用户失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, errno.NewError(500, "创建用户失败")
	}

	return &admin.CreateUserResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "创建用户成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		UserID: user.ID,
	}, nil
}
