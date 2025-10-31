package review

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/review"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CreateReviewLogic 创建评价业务逻辑
func CreateReviewLogic(ctx context.Context, workerID int64, req *review.CreateReviewReq) (*review.CreateReviewResp, error) {
	// 检查是否已经评价过
	existingReview, err := mysql.GetReviewByJobAndUsers(nil, req.JobID, req.EmployerID, workerID, req.ReviewType)
	if err != nil {
		utils.Errorf("检查评价是否存在失败: %v", err)
		return &review.CreateReviewResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if existingReview != nil {
		return &review.CreateReviewResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "已经评价过该工作了",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 创建评价
	reviewModel := &models.Review{
		JobID:      req.JobID,
		EmployerID: req.EmployerID,
		WorkerID:   workerID,
		Rating:     int8(req.Rating),
		Content:    req.Content,
		ReviewType: req.ReviewType,
	}

	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.CreateReview(tx, reviewModel)
	})

	if err != nil {
		utils.Errorf("创建评价失败: %v", err)
		return &review.CreateReviewResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "创建评价失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &review.CreateReviewResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "创建评价成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		ReviewID: reviewModel.ID,
	}, nil
}
