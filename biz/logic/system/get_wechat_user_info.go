package system

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/system"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetWeChatUserInfoLogic 获取微信用户信息业务逻辑
func GetWeChatUserInfoLogic(ctx context.Context, c *app.RequestContext) (*system.GetWeChatUserInfoResp, error) {
	// 获取微信云托管容器内的用户信息
	// 微信云托管会在请求头中自动注入以下字段：
	// X-WX-OPENID: 微信小程序用户的openid
	// X-WX-UNIONID: 微信用户的unionid（可选，需要用户在微信开放平台绑定）
	// X-WX-APPID: 微信小程序appid
	// X-WX-ENV: 环境ID
	// X-WX-CLOUDBASE-ACCESS-TOKEN: 访问令牌

	openid := string(c.GetHeader("X-WX-OPENID"))
	unionid := string(c.GetHeader("X-WX-UNIONID"))
	appid := string(c.GetHeader("X-WX-APPID"))
	env := string(c.GetHeader("X-WX-ENV"))
	cloudbaseAccessToken := string(c.GetHeader("X-WX-CLOUDBASE-ACCESS-TOKEN"))

	// 记录所有微信用户信息到日志
	utils.LogWithFields(map[string]interface{}{
		"openid":                      openid,
		"unionid":                     unionid,
		"appid":                       appid,
		"env":                         env,
		"cloudbase_access_token":      cloudbaseAccessToken,
		"has_openid":                  openid != "",
		"has_unionid":                 unionid != "",
		"has_appid":                   appid != "",
		"has_env":                     env != "",
		"has_cloudbase_access_token":  cloudbaseAccessToken != "",
	}).Info("获取微信用户信息")

	// 如果没有获取到openid，说明不是从微信云托管调用
	if openid == "" {
		utils.Warn("未获取到微信用户openid，可能是非微信云托管环境调用")
	}

	return &system.GetWeChatUserInfoResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取微信用户信息成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Openid:                openid,
		Unionid:               unionid,
		Appid:                 appid,
		Env:                   env,
		CloudbaseAccessToken:  cloudbaseAccessToken,
	}, nil
}

