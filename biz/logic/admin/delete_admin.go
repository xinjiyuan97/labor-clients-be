package admin

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// DeleteAdminLogic 删除管理员业务逻辑
func DeleteAdminLogic(ctx context.Context, req *admin.DeleteAdminReq) (*admin.DeleteAdminResp, error) {
	// 获取管理员信息
	adminUser, err := mysql.GetAdminByID(nil, req.AdminID)
	if err != nil {
		utils.Errorf("获取管理员信息失败: %v", err)
		return &admin.DeleteAdminResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if adminUser == nil {
		return &admin.DeleteAdminResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "管理员不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 软删除管理员
	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.DeleteUser(tx, req.AdminID)
	})

	if err != nil {
		utils.Errorf("删除管理员失败: %v", err)
		return &admin.DeleteAdminResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "删除管理员失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &admin.DeleteAdminResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "删除管理员成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
