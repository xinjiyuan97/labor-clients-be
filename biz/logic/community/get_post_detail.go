package community

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/community"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetPostDetailLogic 获取帖子详情业务逻辑
func GetPostDetailLogic(ctx context.Context, req *community.GetPostDetailReq) (*community.GetPostDetailResp, error) {
	// 获取帖子详情
	post, err := mysql.GetCommunityPostByID(nil, req.PostID)
	if err != nil {
		utils.Errorf("获取帖子详情失败: %v", err)
		return &community.GetPostDetailResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if post == nil {
		return &community.GetPostDetailResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "帖子不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 增加浏览数
	_ = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.IncrementPostViewCount(tx, req.PostID)
	})

	// 构建帖子信息
	postInfo := &common.CommunityPostInfo{
		PostID:    post.ID,
		Title:     post.Title,
		Content:   post.Content,
		AuthorID:  post.AuthorID,
		PostType:  post.PostType,
		ViewCount: int32(post.ViewCount),
		LikeCount: int32(post.LikeCount),
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
	}

	return &community.GetPostDetailResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取帖子详情成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Post: postInfo,
	}, nil
}
