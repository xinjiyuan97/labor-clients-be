package community

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/community"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// DeletePostLogic 删除帖子业务逻辑
func DeletePostLogic(ctx context.Context, req *community.DeletePostReq) (*community.DeletePostResp, error) {
	err := mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.DeleteCommunityPost(tx, req.PostID)
	})

	if err != nil {
		utils.Errorf("删除帖子失败: %v", err)
		return &community.DeletePostResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "删除帖子失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &community.DeletePostResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "删除帖子成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
