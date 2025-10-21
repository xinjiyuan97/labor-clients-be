package community

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/community"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetPostListLogic 获取帖子列表业务逻辑
func GetPostListLogic(req *community.GetPostListReq) (*community.GetPostListResp, error) {
	// 设置默认分页参数
	page := 1
	limit := 10
	if req.PageReq != nil {
		if req.PageReq.Page > 0 {
			page = int(req.PageReq.Page)
		}
		if req.PageReq.Limit > 0 {
			limit = int(req.PageReq.Limit)
		}
	}

	offset := (page - 1) * limit

	// 获取帖子列表
	posts, err := mysql.GetCommunityPosts(nil, req.PostType, offset, limit)
	if err != nil {
		utils.Errorf("获取帖子列表失败: %v", err)
		return &community.GetPostListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取总数
	total, err := mysql.CountCommunityPosts(nil, req.PostType)
	if err != nil {
		utils.Errorf("获取帖子总数失败: %v", err)
		return &community.GetPostListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建帖子信息
	var postInfos []*common.CommunityPostInfo
	for _, post := range posts {
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
		postInfos = append(postInfos, postInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &community.GetPostListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取帖子列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Posts:    postInfos,
	}, nil
}
