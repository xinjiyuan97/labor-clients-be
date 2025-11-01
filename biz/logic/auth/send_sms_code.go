package auth

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// SendSMSCodeLogic 发送短信验证码业务逻辑
func SendSMSCodeLogic(ctx context.Context, req *auth.SendSMSCodeReq) (*auth.SendSMSCodeResp, error) {
	// 1. 检查是否在1分钟内已发送过验证码（防止频繁发送）
	hasRecentCode, err := mysql.CheckRecentCodeExists(ctx, req.Phone, 1)
	if err != nil {
		utils.Errorf("检查最近验证码失败: %v", err)
		return &auth.SendSMSCodeResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if hasRecentCode {
		return &auth.SendSMSCodeResp{
			Base: &common.BaseResp{
				Code:      429,
				Message:   "发送过于频繁，请稍后再试",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 2. 生成6位随机验证码
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// 3. 将验证码存储到MySQL
	err = mysql.CreateSMSVerificationCode(ctx, req.Phone, code)
	if err != nil {
		utils.Errorf("存储短信验证码到MySQL失败: %v", err)
		return &auth.SendSMSCodeResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "发送验证码失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 4. 定时清理过期验证码（异步）
	go func() {
		if err := mysql.CleanExpiredSMSVerificationCodes(context.Background()); err != nil {
			utils.Warnf("清理过期验证码失败: %v", err)
		}
	}()

	// TODO: 集成真实短信服务提供商API发送短信
	// 这里仅打印到日志，生产环境需要替换为真实的短信服务
	utils.LogWithFields(map[string]interface{}{
		"phone": req.Phone,
		"code":  code,
	}).Info("发送短信验证码")

	// 返回验证码（仅开发环境，生产环境应该注释掉）
	return &auth.SendSMSCodeResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "验证码发送成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Code:      code, // TODO: 生产环境删除此字段
		ExpiresIn: 300,  // 5分钟有效期
	}, nil
}
