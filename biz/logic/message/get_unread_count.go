package message

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/message"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetUnreadCountLogic 获取未读消息数量业务逻辑
func GetUnreadCountLogic(req *message.GetUnreadCountReq) (*message.GetUnreadCountResp, error) {
	// 获取未读消息数量 - 这里需要从JWT token中获取UserID
	userID := int64(0) // 需要从JWT token中获取
	count, err := mysql.GetUnreadCount(nil, userID, req.MsgCategory)
	if err != nil {
		utils.Errorf("获取未读消息数量失败: %v", err)
		return &message.GetUnreadCountResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &message.GetUnreadCountResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取未读消息数量成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		UnreadCount: int32(count),
	}, nil
}
