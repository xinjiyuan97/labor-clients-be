package admin

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// DisableAdminLogic 停用管理员业务逻辑
func DisableAdminLogic(req *admin.DisableAdminReq) (*admin.DisableAdminResp, error) {
	// 获取管理员信息
	adminUser, err := mysql.GetAdminByID(nil, req.AdminID)
	if err != nil {
		utils.Errorf("获取管理员信息失败: %v", err)
		return &admin.DisableAdminResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if adminUser == nil {
		return &admin.DisableAdminResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "管理员不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 更新账号状态为停用
	err = mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.UpdateUserStatus(tx, req.AdminID, "disabled")
	})

	if err != nil {
		utils.Errorf("停用管理员失败: %v", err)
		return &admin.DisableAdminResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "停用管理员失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &admin.DisableAdminResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "停用管理员成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
