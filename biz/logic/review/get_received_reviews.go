package review

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/review"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetReceivedReviewsLogic 获取收到的评价列表业务逻辑
func GetReceivedReviewsLogic(userID int64, req *review.GetReceivedReviewsReq) (*review.GetReceivedReviewsResp, error) {
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

	reviewType := "" // 不按类型过滤

	// 获取收到的评价列表
	reviews, err := mysql.GetReceivedReviews(nil, userID, reviewType, offset, limit)
	if err != nil {
		utils.Errorf("获取收到的评价列表失败: %v", err)
		return &review.GetReceivedReviewsResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取总数
	total, err := mysql.CountReceivedReviews(nil, userID, reviewType)
	if err != nil {
		utils.Errorf("获取收到的评价总数失败: %v", err)
		return &review.GetReceivedReviewsResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建评价信息
	var reviewInfos []*common.ReviewInfo
	for _, reviewModel := range reviews {
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
		reviewInfos = append(reviewInfos, reviewInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &review.GetReceivedReviewsResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取收到的评价列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Reviews:  reviewInfos,
	}, nil
}
