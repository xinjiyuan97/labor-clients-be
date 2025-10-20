package user

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/user"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetIncomeDetailLogic 获取收入详情业务逻辑
func GetIncomeDetailLogic(req *user.GetIncomeDetailReq, userID int64) (*user.GetIncomeDetailResp, error) {
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

	// 获取支付记录
	var payments []*models.Payment
	var total int64
	var err error

	if req.StartDate != "" || req.EndDate != "" {
		// 根据时间段获取
		payments, err = mysql.GetUserPaymentsByPeriod(nil, userID, req.StartDate, req.EndDate, offset, limit)
		if err != nil {
			utils.Errorf("根据时间段获取支付记录失败: %v", err)
			return &user.GetIncomeDetailResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountUserPaymentsByPeriod(nil, userID, req.StartDate, req.EndDate)
	} else if req.Status != "" {
		// 根据状态获取
		payments, err = mysql.GetUserPaymentsByStatus(nil, userID, req.Status, offset, limit)
		if err != nil {
			utils.Errorf("根据状态获取支付记录失败: %v", err)
			return &user.GetIncomeDetailResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountUserPaymentsByStatus(nil, userID, req.Status)
	} else {
		// 获取所有记录
		payments, err = mysql.GetUserPayments(nil, userID, offset, limit)
		if err != nil {
			utils.Errorf("获取支付记录失败: %v", err)
			return &user.GetIncomeDetailResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountUserPayments(nil, userID)
	}

	if err != nil {
		utils.Errorf("获取支付记录总数失败: %v", err)
		return &user.GetIncomeDetailResp{
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

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &user.GetIncomeDetailResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取收入详情成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Payments: paymentInfos,
	}, nil
}
