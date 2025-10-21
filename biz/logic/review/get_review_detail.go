package review

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/review"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetReviewDetailLogic 获取评价详情业务逻辑
func GetReviewDetailLogic(req *review.GetReviewDetailReq) (*review.GetReviewDetailResp, error) {
	// 获取评价详情
	reviewModel, err := mysql.GetReviewByID(nil, req.ReviewID)
	if err != nil {
		utils.Errorf("获取评价详情失败: %v", err)
		return &review.GetReviewDetailResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if reviewModel == nil {
		return &review.GetReviewDetailResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "评价不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建评价信息
	reviewInfo := &common.ReviewInfo{
		ReviewID:   reviewModel.ID,
		JobID:      reviewModel.JobID,
		EmployerID: reviewModel.EmployerID,
		WorkerID:   reviewModel.WorkerID,
		Rating:     int32(reviewModel.Rating),
		Content:    reviewModel.Content,
		ReviewType: reviewModel.ReviewType,
		CreatedAt:  reviewModel.CreatedAt.Format(time.RFC3339),
	}

	return &review.GetReviewDetailResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取评价详情成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Review: reviewInfo,
	}, nil
}
