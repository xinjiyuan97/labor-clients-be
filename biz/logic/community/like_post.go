package community

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/community"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// LikePostLogic 点赞帖子业务逻辑
func LikePostLogic(ctx context.Context, req *community.LikePostReq) (*community.LikePostResp, error) {
	userID := int64(0) // 需要从JWT token中获取

	// 检查是否已点赞
	liked, err := mysql.CheckPostLike(nil, req.PostID, userID)
	if err != nil {
		utils.Errorf("检查点赞状态失败: %v", err)
		return &community.LikePostResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if liked {
		return &community.LikePostResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "已经点赞过了",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 创建点赞记录并增加点赞数
	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		// 创建点赞记录
		like := &models.PostLike{
			PostID: req.PostID,
			UserID: userID,
		}
		if err := mysql.CreatePostLike(tx, like); err != nil {
			return err
		}

		// 增加点赞数
		return mysql.IncrementPostLikeCount(tx, req.PostID)
	})

	if err != nil {
		utils.Errorf("点赞失败: %v", err)
		return &community.LikePostResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "点赞失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &community.LikePostResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "点赞成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
