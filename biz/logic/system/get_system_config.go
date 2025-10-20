package system

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/system"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetSystemConfigLogic 获取系统配置业务逻辑
func GetSystemConfigLogic(req *system.GetSystemConfigReq) (*system.GetSystemConfigResp, error) {
	// 获取系统配置
	config, err := mysql.GetSystemConfig(nil)
	if err != nil {
		utils.Errorf("获取系统配置失败: %v", err)
		return &system.GetSystemConfigResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &system.GetSystemConfigResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取系统配置成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Configs: config,
	}, nil
}
