package message

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/message"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetMessageDetailLogic 获取消息详情业务逻辑
func GetMessageDetailLogic(req *message.GetMessageDetailReq) (*message.GetMessageDetailResp, error) {
	// 获取消息详情
	messageModel, err := mysql.GetMessageByID(nil, req.MessageID)
	if err != nil {
		utils.Errorf("获取消息详情失败: %v", err)
		return &message.GetMessageDetailResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if messageModel == nil {
		return &message.GetMessageDetailResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "消息不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建消息信息
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

	return &message.GetMessageDetailResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取消息详情成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Message: messageInfo,
	}, nil
}
