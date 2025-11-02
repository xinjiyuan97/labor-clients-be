package auth

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// WeChatLoginWithSMSCodeLogic 微信手机号验证码登录业务逻辑
func WeChatLoginWithSMSCodeLogic(ctx context.Context, c *app.RequestContext, req *auth.WeChatLoginWithSMSCodeReq) (*auth.LoginResp, error) {
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

	// 3. 获取微信用户信息（从请求头）
	openid := string(c.GetHeader("X-WX-OPENID"))
	unionid := string(c.GetHeader("X-WX-UNIONID"))
	appid := string(c.GetHeader("X-WX-APPID"))

	// 4. 查询或创建用户
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

	// 如果用户不存在，自动创建新用户
	if user == nil {
		isNewUser = true

		// 生成一个随机密码
		defaultPassword := "default_password_123456" // TODO: 考虑使用随机密码
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
		}).Info("微信登录自动创建新用户")
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

	// 5. 如果提供了微信信息，自动绑定微信账号
	if openid != "" && appid != "" {
		// 检查是否已有该微信的绑定
		existingBinding, err := mysql.GetThirdPartyBindingByPlatformAndOpenID(ctx, "wechat", openid)
		if err != nil {
			utils.Errorf("查询微信绑定失败: %v", err)
			// 不阻止登录，只是绑定失败
		} else if existingBinding == nil {
			// 没有绑定，创建新绑定
			now := time.Now()
			binding := &models.ThirdPartyBinding{
				UserID:      user.ID,
				Platform:    "wechat",
				OpenID:      openid,
				UnionID:     unionid,
				AppID:       appid,
				Status:      "active",
				LastLoginAt: &now,
			}

			if err := mysql.CreateThirdPartyBinding(ctx, binding); err != nil {
				utils.Errorf("创建微信绑定失败: %v", err)
				// 不阻止登录，只是绑定失败
			} else {
				utils.Infof("微信绑定成功, UserID: %d, OpenID: %s", user.ID, openid)
			}
		} else if existingBinding.UserID != user.ID {
			// 已有绑定但是不同的用户，记录警告但不阻止登录
			utils.LogWithFields(map[string]interface{}{
				"existing_user_id": existingBinding.UserID,
				"current_user_id":  user.ID,
				"openid":           openid,
			}).Warn("微信账号已绑定到其他用户，跳过绑定")
		} else {
			// 已绑定到当前用户，更新登录时间
			now := time.Now()
			updateData := map[string]interface{}{
				"unionid":       unionid,
				"last_login_at": now,
			}
			if err := mysql.UpdateThirdPartyBinding(ctx, existingBinding.ID, updateData); err != nil {
				utils.Warnf("更新微信绑定失败: %v", err)
			}
		}
	}

	// 6. 生成token
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
		"user_id":    user.ID,
		"phone":      req.Phone,
		"role":       user.Role,
		"is_new_user": isNewUser,
		"has_wechat": openid != "",
	}).Info("微信手机号验证码登录成功")

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

