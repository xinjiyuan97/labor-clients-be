package system

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/system"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// SubmitFeedbackLogic 提交反馈业务逻辑
func SubmitFeedbackLogic(req *system.SubmitFeedbackReq) (*system.SubmitFeedbackResp, error) {
	// 创建反馈
	feedback := &models.Feedback{
		UserID:  0, // 需要从JWT token中获取
		Type:    req.FeedbackType,
		Content: req.Content,
		Contact: req.ContactInfo,
		Status:  "pending",
	}

	err := mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.CreateFeedback(tx, feedback)
	})

	if err != nil {
		utils.Errorf("提交反馈失败: %v", err)
		return &system.SubmitFeedbackResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "提交反馈失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &system.SubmitFeedbackResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "提交反馈成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		FeedbackID: feedback.ID,
	}, nil
}
