package community

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/community"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetCommentListLogic 获取评论列表业务逻辑
func GetCommentListLogic(req *community.GetCommentListReq) (*community.GetCommentListResp, error) {
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

	// 获取评论列表
	comments, err := mysql.GetPostComments(nil, req.PostID, offset, limit)
	if err != nil {
		utils.Errorf("获取评论列表失败: %v", err)
		return &community.GetCommentListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取总数
	total, err := mysql.CountPostComments(nil, req.PostID)
	if err != nil {
		utils.Errorf("获取评论总数失败: %v", err)
		return &community.GetCommentListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建评论信息
	var commentInfos []*common.MessageInfo
	for _, comment := range comments {
		commentInfo := &common.MessageInfo{
			MessageID:   comment.ID,
			FromUser:    comment.UserID,
			ToUser:      0, // 评论的目标是帖子，这里可以设为0或帖子作者ID
			MessageType: "comment",
			Content:     comment.Content,
			MsgCategory: "community",
			IsRead:      true,
			CreatedAt:   comment.CreatedAt.Format(time.RFC3339),
		}
		commentInfos = append(commentInfos, commentInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &community.GetCommentListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取评论列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Comments: commentInfos,
	}, nil
}
