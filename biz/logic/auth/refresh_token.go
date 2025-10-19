package auth

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// RefreshTokenLogic 刷新token业务逻辑
func RefreshTokenLogic(req *auth.RefreshTokenReq) (*auth.RefreshTokenResp, error) {
	// 验证refresh token
	claims, err := utils.ValidateToken(req.RefreshToken)
	if err != nil {
		return &auth.RefreshTokenResp{
			Base: &common.BaseResp{
				Code:      401,
				Message:   "无效的refresh token",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 检查用户是否仍然存在
	user, err := mysql.GetUserByID(nil, claims.UserID)
	if err != nil {
		utils.Errorf("查询用户失败: %v", err)
		return &auth.RefreshTokenResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if user == nil {
		return &auth.RefreshTokenResp{
			Base: &common.BaseResp{
				Code:      401,
				Message:   "用户不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 生成新的token
	newToken, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.Errorf("生成新token失败: %v", err)
		return &auth.RefreshTokenResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &auth.RefreshTokenResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "token刷新成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Token:     newToken,
		ExpiresAt: time.Now().Add(utils.TokenExpire).Format(time.RFC3339),
	}, nil
}
