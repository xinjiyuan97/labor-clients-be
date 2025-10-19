package auth

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// LoginLogic 用户登录业务逻辑
func LoginLogic(req *auth.LoginReq) (*auth.LoginResp, error) {
	// 根据手机号查询用户
	user, err := mysql.GetUserByPhone(nil, req.Phone)
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

	if user == nil {
		return &auth.LoginResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "用户不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 验证密码
	err = utils.CheckPassword(user.PasswordHash, req.Password)
	if err != nil {
		return &auth.LoginResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "密码错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 生成token
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
