package review

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/review"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UpdateReviewLogic 更新评价业务逻辑
func UpdateReviewLogic(req *review.UpdateReviewReq) (*review.UpdateReviewResp, error) {
	// 获取现有评价
	reviewModel, err := mysql.GetReviewByID(nil, req.ReviewID)
	if err != nil {
		utils.Errorf("获取评价失败: %v", err)
		return &review.UpdateReviewResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if reviewModel == nil {
		return &review.UpdateReviewResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "评价不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 更新字段
	if req.Rating > 0 {
		reviewModel.Rating = int8(req.Rating)
	}
	if req.Content != "" {
		reviewModel.Content = req.Content
	}

	err = mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.UpdateReview(tx, reviewModel)
	})

	if err != nil {
		utils.Errorf("更新评价失败: %v", err)
		return &review.UpdateReviewResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "更新评价失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &review.UpdateReviewResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "更新评价成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
