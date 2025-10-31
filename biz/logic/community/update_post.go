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

// UpdatePostLogic 更新帖子业务逻辑
func UpdatePostLogic(ctx context.Context, req *community.UpdatePostReq) (*community.UpdatePostResp, error) {
	// 获取现有帖子
	post, err := mysql.GetCommunityPostByID(nil, req.PostID)
	if err != nil {
		utils.Errorf("获取帖子失败: %v", err)
		return &community.UpdatePostResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if post == nil {
		return &community.UpdatePostResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "帖子不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 更新字段
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}
	if req.PostType != "" {
		post.PostType = req.PostType
	}

	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.UpdateCommunityPost(tx, post)
	})

	if err != nil {
		utils.Errorf("更新帖子失败: %v", err)
		return &community.UpdatePostResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "更新帖子失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &community.UpdatePostResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "更新帖子成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
