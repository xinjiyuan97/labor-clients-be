package system

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/system"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetVersionLogic 获取版本信息业务逻辑
func GetVersionLogic(req *system.GetVersionReq) (*system.GetVersionResp, error) {
	// 获取版本信息
	versionInfo, err := mysql.GetVersionInfo(nil)
	if err != nil {
		utils.Errorf("获取版本信息失败: %v", err)
		return &system.GetVersionResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &system.GetVersionResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取版本信息成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Version:   versionInfo.Version,
		BuildTime: versionInfo.CreatedAt.Format(time.RFC3339),
		GitCommit: versionInfo.BuildNumber,
	}, nil
}
