package message

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/message"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetMessageListLogic 获取消息列表业务逻辑
func GetMessageListLogic(req *message.GetMessageListReq) (*message.GetMessageListResp, error) {
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

	// 获取消息列表 - 这里需要从JWT token中获取UserID
	userID := int64(0) // 需要从JWT token中获取
	messages, err := mysql.GetMessages(nil, userID, req.MsgCategory, offset, limit)
	if err != nil {
		utils.Errorf("获取消息列表失败: %v", err)
		return &message.GetMessageListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取总数
	total, err := mysql.CountMessages(nil, userID, req.MsgCategory)
	if err != nil {
		utils.Errorf("获取消息总数失败: %v", err)
		return &message.GetMessageListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建消息信息
	var messageInfos []*common.MessageInfo
	for _, messageModel := range messages {
		messageInfo := &common.MessageInfo{
			MessageID:   messageModel.ID,
			FromUser:    messageModel.FromUser,
			ToUser:      messageModel.ToUser,
			MessageType: messageModel.MessageType,
			Content:     messageModel.Content,
			MsgCategory: messageModel.MsgCategory,
			IsRead:      messageModel.IsRead,
			CreatedAt:   messageModel.CreatedAt.Format(time.RFC3339),
		}
		messageInfos = append(messageInfos, messageInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &message.GetMessageListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取消息列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Messages: messageInfos,
	}, nil
}
