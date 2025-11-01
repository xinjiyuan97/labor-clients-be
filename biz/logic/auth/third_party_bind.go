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

// ThirdPartyLoginBindLogic 第三方登录绑定业务逻辑
func ThirdPartyLoginBindLogic(ctx context.Context, req *auth.ThirdPartyLoginBindReq) (*auth.ThirdPartyLoginBindResp, error) {
	// 1. 验证短信验证码
	smsCode, err := mysql.GetSMSVerificationCode(ctx, req.Phone, req.Code)
	if err != nil {
		utils.Errorf("从MySQL获取验证码失败: %v", err)
		return &auth.ThirdPartyLoginBindResp{
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
			"phone": req.Phone,
		}).Warn("验证码不存在或已使用")
		return &auth.ThirdPartyLoginBindResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "验证码错误或已过期",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 2. 标记验证码为已使用
	err = mysql.MarkSMSVerificationCodeUsed(ctx, req.Phone, req.Code)
	if err != nil {
		utils.Warnf("标记验证码已使用失败: %v", err)
	}

	// 3. 检查是否已有该平台的绑定
	existingBinding, err := mysql.GetThirdPartyBindingByPlatformAndOpenID(ctx, req.Platform, req.Openid)
	if err != nil {
		utils.Errorf("查询第三方绑定失败: %v", err)
		return &auth.ThirdPartyLoginBindResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 4. 查询或创建用户
	user, err := mysql.GetUserByPhone(ctx, req.Phone)
	isNewUser := false

	if err != nil {
		utils.Errorf("查询用户失败: %v", err)
		return &auth.ThirdPartyLoginBindResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if user == nil {
		// 用户不存在，创建新用户
		isNewUser = true
		
		// 生成一个默认密码（用户可以通过手机号+验证码登录）
		defaultPassword := "default_password_123456" // TODO: 考虑使用随机密码
		hashedPassword, err := utils.HashPassword(defaultPassword)
		if err != nil {
			utils.Errorf("密码加密失败: %v", err)
			return &auth.ThirdPartyLoginBindResp{
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
			Avatar:       req.Avatar,
		}

		if err := mysql.CreateUser(ctx, user); err != nil {
			utils.Errorf("创建用户失败: %v", err)
			return &auth.ThirdPartyLoginBindResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "创建用户失败",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}

		utils.Infof("创建新用户成功, UserID: %d", user.ID)
	}

	// 5. 如果已有该平台绑定，更新绑定信息；否则创建新绑定
	if existingBinding != nil {
		// 更新绑定信息
		updateData := map[string]interface{}{
			"user_id":      user.ID,
			"unionid":      req.Unionid,
			"nickname":     req.Nickname,
			"avatar":       req.Avatar,
			"last_login_at": time.Now().Format("2006-01-02 15:04:05"),
		}
		
		if err := mysql.UpdateThirdPartyBinding(ctx, existingBinding.ID, updateData); err != nil {
			utils.Errorf("更新第三方绑定失败: %v", err)
			return &auth.ThirdPartyLoginBindResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "更新绑定失败",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
	} else {
		// 创建新的第三方绑定
		lastLoginAt := time.Now().Format("2006-01-02 15:04:05")
		binding := &models.ThirdPartyBinding{
			UserID:      user.ID,
			Platform:    req.Platform,
			OpenID:      req.Openid,
			UnionID:     req.Unionid,
			AppID:       req.Appid,
			Nickname:    req.Nickname,
			Avatar:      req.Avatar,
			Status:      "active",
			LastLoginAt: &lastLoginAt,
		}

		if err := mysql.CreateThirdPartyBinding(ctx, binding); err != nil {
			utils.Errorf("创建第三方绑定失败: %v", err)
			return &auth.ThirdPartyLoginBindResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "创建绑定失败",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
	}

	// 6. 生成token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.Errorf("生成token失败: %v", err)
		return &auth.ThirdPartyLoginBindResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	utils.LogWithFields(map[string]interface{}{
		"user_id":     user.ID,
		"phone":       req.Phone,
		"platform":    req.Platform,
		"openid":      req.Openid,
		"is_new_user": isNewUser,
	}).Info("第三方登录绑定成功")

	return &auth.ThirdPartyLoginBindResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "登录成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		IsNewUser:  isNewUser,
		UserID:     user.ID,
		Token:      token,
		ExpiresAt:  time.Now().Add(utils.TokenExpire).Format(time.RFC3339),
	}, nil
}
