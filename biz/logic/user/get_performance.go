package user

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/user"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetPerformanceLogic 获取绩效统计业务逻辑
func GetPerformanceLogic(req *user.GetPerformanceReq, userID int64) (*user.GetPerformanceResp, error) {
	// 获取用户绩效统计
	totalApplications, completedJobs, successRate, averageRating, err := mysql.GetUserPerformanceStats(nil, userID)
	if err != nil {
		utils.Errorf("获取用户绩效统计失败: %v", err)
		return &user.GetPerformanceResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &user.GetPerformanceResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取绩效统计成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		CompletedJobs:     int32(completedJobs),
		TotalApplications: int32(totalApplications),
		SuccessRate:       successRate,
		AverageRating:     averageRating,
	}, nil
}
