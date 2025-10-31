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

// CommentPostLogic 评论帖子业务逻辑
func CommentPostLogic(ctx context.Context, userID int64, req *community.CommentPostReq) (*community.CommentPostResp, error) {
	// 创建评论
	comment := &models.PostComment{
		PostID:   req.PostID,
		UserID:   userID,
		Content:  req.Content,
		ParentID: &req.ParentID,
	}

	err := mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.CreatePostComment(tx, comment)
	})

	if err != nil {
		utils.Errorf("评论失败: %v", err)
		return &community.CommentPostResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "评论失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &community.CommentPostResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "评论成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		CommentID: comment.ID,
	}, nil
}
