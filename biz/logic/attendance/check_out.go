package attendance

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/attendance"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CheckOutLogic 签退业务逻辑
func CheckOutLogic(ctx context.Context, workerID int64, req *attendance.CheckOutReq) (*attendance.CheckOutResp, error) {
	// 获取今日打卡记录
	record, err := mysql.GetTodayAttendanceRecord(nil, workerID, req.JobID)
	if err != nil {
		utils.Errorf("获取今日打卡记录失败: %v", err)
		return &attendance.CheckOutResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if record == nil {
		return &attendance.CheckOutResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "今日未打卡",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if record.CheckOut != nil {
		return &attendance.CheckOutResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "已经签退",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 更新签退信息
	now := time.Now()
	record.CheckOut = &now
	record.CheckOutLocation = req.CheckOutLocation

	// 计算工作时长
	if record.CheckIn != nil {
		record.WorkHours = mysql.CalculateWorkHours(*record.CheckIn, now)
	}

	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.UpdateAttendanceRecord(tx, record)
	})

	if err != nil {
		utils.Errorf("签退失败: %v", err)
		return &attendance.CheckOutResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "签退失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	workHours, _ := record.WorkHours.Float64()

	return &attendance.CheckOutResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "签退成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		RecordID:     int32(record.ID),
		CheckOutTime: now.Format(time.RFC3339),
		WorkHours:    workHours,
	}, nil
}
