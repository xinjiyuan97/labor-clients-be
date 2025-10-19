package auth

import (
	"strings"
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// LogoutLogic 用户登出业务逻辑
func LogoutLogic(req *auth.LogoutReq) (*auth.LogoutResp, error) {
	// 从Authorization header中提取token
	token := strings.TrimPrefix(req.Token, "Bearer ")

	// 验证token
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return &auth.LogoutResp{
			Base: &common.BaseResp{
				Code:      401,
				Message:   "无效的token",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 这里可以添加token黑名单逻辑，将token加入Redis黑名单
	// 由于当前没有Redis配置，暂时只返回成功
	utils.Infof("用户登出成功, UserID: %d", claims.UserID)

	return &auth.LogoutResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "登出成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
