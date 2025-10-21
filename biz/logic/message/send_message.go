package message

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/message"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// SendMessageLogic 发送消息业务逻辑
func SendMessageLogic(req *message.SendMessageReq) (*message.SendMessageResp, error) {
	// 创建消息
	messageModel := &models.Message{
		FromUser:    0, // 需要从JWT token中获取
		ToUser:      req.ToUser,
		MessageType: req.MessageType,
		Content:     req.Content,
		MsgCategory: req.MsgCategory,
		IsRead:      false,
	}

	err := mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.CreateMessage(tx, messageModel)
	})

	if err != nil {
		utils.Errorf("发送消息失败: %v", err)
		return &message.SendMessageResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "发送消息失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &message.SendMessageResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "发送消息成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		MessageID: messageModel.ID,
	}, nil
}
