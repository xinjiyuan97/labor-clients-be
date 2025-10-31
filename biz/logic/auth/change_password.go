package auth

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// ChangePasswordLogic 修改密码业务逻辑
func ChangePasswordLogic(ctx context.Context, userID int64, req *auth.ChangePasswordReq) (*auth.ChangePasswordResp, error) {
	// 获取用户信息
	user, err := mysql.GetUserByID(nil, userID)
	if err != nil {
		utils.Errorf("获取用户信息失败: %v", err)
		return &auth.ChangePasswordResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if user == nil {
		return &auth.ChangePasswordResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "用户不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 验证旧密码
	if err := utils.CheckPassword(user.PasswordHash, req.OldPassword); err != nil {
		return &auth.ChangePasswordResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "旧密码错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 加密新密码
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		utils.Errorf("密码加密失败: %v", err)
		return &auth.ChangePasswordResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 更新密码
	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.UpdateUserPassword(tx, userID, hashedPassword)
	})

	if err != nil {
		utils.Errorf("修改密码失败: %v", err)
		return &auth.ChangePasswordResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "修改密码失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &auth.ChangePasswordResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "修改密码成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
