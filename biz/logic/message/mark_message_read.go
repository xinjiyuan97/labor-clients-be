package message

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/message"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// MarkMessageReadLogic 标记消息已读业务逻辑
func MarkMessageReadLogic(ctx context.Context, req *message.MarkMessageReadReq) (*message.MarkMessageReadResp, error) {
	err := mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.MarkMessageRead(tx, req.MessageID)
	})

	if err != nil {
		utils.Errorf("标记消息已读失败: %v", err)
		return &message.MarkMessageReadResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "标记消息已读失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &message.MarkMessageReadResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "标记消息已读成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
