package auth

import (
	"context"
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// LoginWithSMSCodeLogic 短信验证码登录业务逻辑
func LoginWithSMSCodeLogic(ctx context.Context, req *auth.LoginWithSMSCodeReq) (*auth.LoginResp, error) {
	// 1. 验证短信验证码
	smsCode, err := mysql.GetSMSVerificationCode(ctx, req.Phone, req.Code)
	if err != nil {
		utils.Errorf("从MySQL获取验证码失败: %v", err)
		return &auth.LoginResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 检查验证码是否存在或已使用
	if smsCode == nil {
		utils.LogWithFields(map[string]interface{}{
			"phone":      req.Phone,
			"input_code": req.Code,
		}).Warn("验证码不存在或已使用")
		return &auth.LoginResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "验证码错误或已过期",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 2. 验证码验证成功，标记为已使用
	err = mysql.MarkSMSVerificationCodeUsed(ctx, req.Phone, req.Code)
	if err != nil {
		utils.Warnf("标记验证码已使用失败: %v", err)
	}

	// 3. 查询或创建用户
	user, err := mysql.GetUserByPhone(ctx, req.Phone)
	isNewUser := false

	if err != nil {
		utils.Errorf("查询用户失败: %v", err)
		return &auth.LoginResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 如果用户不存在，自动创建新用户（类似微信绑定逻辑）
	if user == nil {
		isNewUser = true
		
		// 生成一个随机密码
		// TODO: 考虑使用更安全的密码生成策略
		defaultPassword := "default_password_123456"
		hashedPassword, err := utils.HashPassword(defaultPassword)
		if err != nil {
			utils.Errorf("密码加密失败: %v", err)
			return &auth.LoginResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}

		// 创建新用户（默认为worker角色）
		user = &models.User{
			Phone:        req.Phone,
			Username:     req.Phone, // 使用手机号作为用户名
			PasswordHash: hashedPassword,
			Role:         "worker",
			Status:       "active",
		}

		if err := mysql.CreateUser(ctx, user); err != nil {
			utils.Errorf("创建用户失败: %v", err)
			return &auth.LoginResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "创建用户失败",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}

		utils.LogWithFields(map[string]interface{}{
			"user_id": user.ID,
			"phone":   req.Phone,
		}).Info("通过短信验证码自动创建新用户")
	}

	// 检查账号状态
	if user.Status != "active" {
		return &auth.LoginResp{
			Base: &common.BaseResp{
				Code:      403,
				Message:   "账号已被禁用",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 记录是否为新用户
	if isNewUser {
		utils.LogWithFields(map[string]interface{}{
			"user_id":    user.ID,
			"phone":      req.Phone,
			"is_new_user": true,
		}).Info("新用户通过短信验证码登录")
	}

	// 4. 生成token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.Errorf("生成token失败: %v", err)
		return &auth.LoginResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	utils.LogWithFields(map[string]interface{}{
		"user_id": user.ID,
		"phone":   req.Phone,
		"role":    user.Role,
	}).Info("短信验证码登录成功")

	return &auth.LoginResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "登录成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(utils.TokenExpire).Format(time.RFC3339),
	}, nil
}

