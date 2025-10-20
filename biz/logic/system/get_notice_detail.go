package system

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/system"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetNoticeDetailLogic 获取通知详情业务逻辑
func GetNoticeDetailLogic(req *system.GetNoticeDetailReq) (*system.GetNoticeDetailResp, error) {
	// 获取通知详情
	notice, err := mysql.GetNoticeByID(nil, req.NoticeID)
	if err != nil {
		utils.Errorf("获取通知详情失败: %v", err)
		return &system.GetNoticeDetailResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if notice == nil {
		return &system.GetNoticeDetailResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "通知不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建通知信息
	noticeInfo := &common.CommunityPostInfo{
		PostID:    notice.ID,
		Title:     notice.Title,
		Content:   notice.Content,
		CreatedAt: notice.CreatedAt.Format(time.RFC3339),
	}

	return &system.GetNoticeDetailResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取通知详情成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Notice: noticeInfo,
	}, nil
}
