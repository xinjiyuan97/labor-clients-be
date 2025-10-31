package auth

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// RegisterLogic 用户注册业务逻辑
func RegisterLogic(ctx context.Context, req *auth.RegisterReq) (*auth.RegisterResp, error) {
	// 验证角色
	if req.Role != "worker" && req.Role != "employer" {
		return &auth.RegisterResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "角色只能是worker或employer",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 检查手机号是否已存在
	exists, err := mysql.CheckUserExistsByPhone(nil, req.Phone)
	if err != nil {
		utils.Errorf("检查手机号是否存在失败: %v", err)
		return &auth.RegisterResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if exists {
		return &auth.RegisterResp{
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
		return &auth.RegisterResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 创建用户
	user := &models.User{
		Username:     req.Username,
		Phone:        req.Phone,
		PasswordHash: hashedPassword,
		Role:         req.Role,
	}

	// 使用事务创建用户
	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.CreateUser(ctx, user)
	})

	if err != nil {
		utils.Errorf("创建用户失败: %v", err)
		return &auth.RegisterResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.Errorf("生成token失败: %v", err)
		return &auth.RegisterResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &auth.RegisterResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "注册成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(utils.TokenExpire).Format(time.RFC3339),
	}, nil
}
