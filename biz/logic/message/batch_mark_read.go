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

// BatchMarkReadLogic 批量标记已读业务逻辑
func BatchMarkReadLogic(ctx context.Context, req *message.BatchMarkReadReq) (*message.BatchMarkReadResp, error) {
	err := mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.BatchMarkMessagesRead(tx, req.MessageIds)
	})

	if err != nil {
		utils.Errorf("批量标记消息已读失败: %v", err)
		return &message.BatchMarkReadResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "批量标记消息已读失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &message.BatchMarkReadResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "批量标记消息已读成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
