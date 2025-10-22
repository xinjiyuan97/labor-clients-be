package admin

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// ResetAdminPasswordLogic 重置管理员密码业务逻辑
func ResetAdminPasswordLogic(req *admin.ResetAdminPasswordReq) (*admin.ResetAdminPasswordResp, error) {
	// 获取管理员信息
	adminUser, err := mysql.GetAdminByID(nil, req.AdminID)
	if err != nil {
		utils.Errorf("获取管理员信息失败: %v", err)
		return &admin.ResetAdminPasswordResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if adminUser == nil {
		return &admin.ResetAdminPasswordResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "管理员不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 加密新密码
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		utils.Errorf("密码加密失败: %v", err)
		return &admin.ResetAdminPasswordResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 更新密码
	err = mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.UpdateUserPassword(tx, req.AdminID, hashedPassword)
	})

	if err != nil {
		utils.Errorf("重置管理员密码失败: %v", err)
		return &admin.ResetAdminPasswordResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "重置密码失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &admin.ResetAdminPasswordResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "重置密码成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
