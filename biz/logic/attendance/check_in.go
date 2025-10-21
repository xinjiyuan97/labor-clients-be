package attendance

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/attendance"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CheckInLogic 打卡签到业务逻辑
func CheckInLogic(workerID int64, req *attendance.CheckInReq) (*attendance.CheckInResp, error) {

	// 检查今天是否已经打卡
	existingRecord, err := mysql.GetTodayAttendanceRecord(nil, workerID, req.JobID)
	if err != nil {
		utils.Errorf("检查今日打卡记录失败: %v", err)
		return &attendance.CheckInResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if existingRecord != nil {
		return &attendance.CheckInResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "今日已经打卡",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 创建打卡记录
	now := time.Now()
	record := &models.AttendanceRecord{
		JobID:           req.JobID,
		WorkerID:        workerID,
		CheckIn:         &now,
		CheckInLocation: req.CheckInLocation,
		Status:          "normal",
		WorkHours:       decimal.Zero,
	}

	err = mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.CreateAttendanceRecord(tx, record)
	})

	if err != nil {
		utils.Errorf("打卡失败: %v", err)
		return &attendance.CheckInResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "打卡失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &attendance.CheckInResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "打卡成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		RecordID:    int32(record.ID),
		CheckInTime: now.Format(time.RFC3339),
	}, nil
}
