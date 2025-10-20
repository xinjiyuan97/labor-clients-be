package user

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/user"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UnfavoriteJobLogic 取消收藏工作业务逻辑
func UnfavoriteJobLogic(req *user.UnfavoriteJobReq, userID int64) (*user.UnfavoriteJobResp, error) {
	// 检查是否已经收藏
	exists, err := mysql.CheckUserFavoriteJob(nil, userID, req.JobID)
	if err != nil {
		utils.Errorf("检查收藏状态失败: %v", err)
		return &user.UnfavoriteJobResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if !exists {
		return &user.UnfavoriteJobResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "未收藏该工作",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 使用事务删除收藏记录
	err = mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.DeleteUserFavoriteJob(tx, userID, req.JobID)
	})

	if err != nil {
		utils.Errorf("取消收藏失败: %v", err)
		return &user.UnfavoriteJobResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &user.UnfavoriteJobResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "取消收藏成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
