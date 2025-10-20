package user

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/user"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetIncomeLogic 获取收入统计业务逻辑
func GetIncomeLogic(req *user.GetIncomeReq, userID int64) (*user.GetIncomeResp, error) {
	// 根据时间段计算开始和结束日期
	var startDate, endDate string
	now := time.Now()

	switch req.Period {
	case "week":
		// 本周
		weekday := int(now.Weekday())
		if weekday == 0 {
			weekday = 7 // 周日
		}
		startDate = now.AddDate(0, 0, -weekday+1).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "month":
		// 本月
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "year":
		// 本年
		startDate = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	case "custom":
		// 自定义时间段
		if req.Year > 0 && req.Month > 0 {
			startDate = time.Date(int(req.Year), time.Month(req.Month), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
			endDate = time.Date(int(req.Year), time.Month(req.Month+1), 0, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		}
	default:
		// 默认本月
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		endDate = now.Format("2006-01-02")
	}

	// 获取收入统计
	totalIncome, pendingIncome, paidIncome, err := mysql.GetUserIncomeStats(nil, userID, startDate, endDate)
	if err != nil {
		utils.Errorf("获取收入统计失败: %v", err)
		return &user.GetIncomeResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取最近的支付记录
	payments, err := mysql.GetUserPayments(nil, userID, 0, 10)
	if err != nil {
		utils.Errorf("获取支付记录失败: %v", err)
		return &user.GetIncomeResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建支付信息
	var paymentInfos []*common.PaymentInfo
	for _, payment := range payments {
		paymentInfo := &common.PaymentInfo{
			PaymentID:     payment.ID,
			JobID:         payment.JobID,
			WorkerID:      payment.WorkerID,
			EmployerID:    payment.EmployerID,
			Amount:        payment.Amount.InexactFloat64(),
			PaymentMethod: payment.PaymentMethod,
			Status:        payment.Status,
			PlatformFee:   payment.PlatformFee.InexactFloat64(),
		}
		if payment.PaidAt != nil {
			paymentInfo.PaidAt = payment.PaidAt.Format(time.RFC3339)
		}
		paymentInfos = append(paymentInfos, paymentInfo)
	}

	return &user.GetIncomeResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取收入统计成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		TotalIncome:   totalIncome,
		PendingIncome: pendingIncome,
		PaidIncome:    paidIncome,
		Payments:      paymentInfos,
	}, nil
}
