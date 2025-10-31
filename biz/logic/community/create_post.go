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

// CreatePostLogic 创建帖子业务逻辑
func CreatePostLogic(ctx context.Context, userID int64, req *community.CreatePostReq) (*community.CreatePostResp, error) {
	// 创建帖子
	post := &models.CommunityPost{
		AuthorID: userID,
		Title:    req.Title,
		Content:  req.Content,
		PostType: req.PostType,
		Status:   "published",
	}

	err := mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.CreateCommunityPost(tx, post)
	})

	if err != nil {
		utils.Errorf("创建帖子失败: %v", err)
		return &community.CreatePostResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "创建帖子失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &community.CreatePostResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "创建帖子成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PostID: post.ID,
	}, nil
}
