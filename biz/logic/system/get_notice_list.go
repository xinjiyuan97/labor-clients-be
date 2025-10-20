package system

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/system"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetNoticeListLogic 获取通知列表业务逻辑
func GetNoticeListLogic(req *system.GetNoticeListReq) (*system.GetNoticeListResp, error) {
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

	// 获取通知列表
	notices, err := mysql.GetNotices(nil, offset, limit)
	if err != nil {
		utils.Errorf("获取通知列表失败: %v", err)
		return &system.GetNoticeListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取总数
	total, err := mysql.CountNotices(nil)
	if err != nil {
		utils.Errorf("获取通知总数失败: %v", err)
		return &system.GetNoticeListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建通知信息
	var noticeInfos []*common.CommunityPostInfo
	for _, notice := range notices {
		noticeInfo := &common.CommunityPostInfo{
			PostID:    notice.ID,
			Title:     notice.Title,
			Content:   notice.Content,
			CreatedAt: notice.CreatedAt.Format(time.RFC3339),
		}
		noticeInfos = append(noticeInfos, noticeInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &system.GetNoticeListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取通知列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Notices:  noticeInfos,
	}, nil
}
