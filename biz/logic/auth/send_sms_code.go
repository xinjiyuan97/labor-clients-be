package auth

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/redis"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// SendSMSCodeLogic 发送短信验证码业务逻辑
func SendSMSCodeLogic(ctx context.Context, req *auth.SendSMSCodeReq) (*auth.SendSMSCodeResp, error) {
	// 生成6位随机验证码
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// 将验证码存储到Redis，有效期5分钟
	key := fmt.Sprintf("sms_code:%s", req.Phone)
	err := redis.Set(key, code, 5*time.Minute)
	if err != nil {
		utils.Errorf("存储短信验证码到Redis失败: %v", err)
		return &auth.SendSMSCodeResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "发送验证码失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// TODO: 集成真实短信服务提供商API发送短信
	// 这里仅打印到日志，生产环境需要替换为真实的短信服务
	utils.LogWithFields(map[string]interface{}{
		"phone": req.Phone,
		"code":  code,
		"key":   key,
	}).Info("发送短信验证码")

	// 返回验证码（仅开发环境，生产环境应该注释掉）
	return &auth.SendSMSCodeResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "验证码发送成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Code:       code, // TODO: 生产环境删除此字段
		ExpiresIn: 300,   // 5分钟有效期
	}, nil
}

