package user

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/user"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// FavoriteJobLogic 收藏工作业务逻辑
func FavoriteJobLogic(req *user.FavoriteJobReq, userID int64) (*user.FavoriteJobResp, error) {
	// 检查是否已经收藏
	exists, err := mysql.CheckUserFavoriteJob(nil, userID, req.JobID)
	if err != nil {
		utils.Errorf("检查收藏状态失败: %v", err)
		return &user.FavoriteJobResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if exists {
		return &user.FavoriteJobResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "已经收藏过该工作",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 使用事务创建收藏记录
	var favoriteID int64
	err = mysql.Transaction(func(tx *gorm.DB) error {
		favorite := &models.UserFavoriteJob{
			UserID: userID,
			JobID:  req.JobID,
		}

		if err := mysql.CreateUserFavoriteJob(tx, favorite); err != nil {
			return err
		}

		favoriteID = favorite.ID
		return nil
	})

	if err != nil {
		utils.Errorf("收藏工作失败: %v", err)
		return &user.FavoriteJobResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &user.FavoriteJobResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "收藏成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		FavoriteID: favoriteID,
	}, nil
}
