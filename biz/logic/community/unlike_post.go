package community

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/community"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UnlikePostLogic 取消点赞业务逻辑
func UnlikePostLogic(req *community.UnlikePostReq) (*community.UnlikePostResp, error) {
	userID := int64(0) // 需要从JWT token中获取

	// 检查是否已点赞
	liked, err := mysql.CheckPostLike(nil, req.PostID, userID)
	if err != nil {
		utils.Errorf("检查点赞状态失败: %v", err)
		return &community.UnlikePostResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if !liked {
		return &community.UnlikePostResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "还未点赞",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 删除点赞记录并减少点赞数
	err = mysql.Transaction(func(tx *gorm.DB) error {
		// 删除点赞记录
		if err := mysql.DeletePostLike(tx, req.PostID, userID); err != nil {
			return err
		}

		// 减少点赞数
		return mysql.DecrementPostLikeCount(tx, req.PostID)
	})

	if err != nil {
		utils.Errorf("取消点赞失败: %v", err)
		return &community.UnlikePostResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "取消点赞失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &community.UnlikePostResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "取消点赞成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
