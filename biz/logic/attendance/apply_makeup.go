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

// ApplyMakeupLogic 申请补签业务逻辑
func ApplyMakeupLogic(workerID int64, req *attendance.ApplyMakeupReq) (*attendance.ApplyMakeupResp, error) {
	// 解析补签日期和时间
	makeupDate, err := time.Parse("2006-01-02 15:04:05", req.MakeupTime)
	if err != nil {
		// 尝试只解析日期
		makeupDate, err = time.Parse("2006-01-02", req.MakeupDate)
		if err != nil {
			return &attendance.ApplyMakeupResp{
				Base: &common.BaseResp{
					Code:      400,
					Message:   "日期格式错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
	}

	// 检查该日期是否已有考勤记录
	existingRecord, err := mysql.GetTodayAttendanceRecord(nil, workerID, req.JobID)
	if err != nil {
		utils.Errorf("检查考勤记录失败: %v", err)
		return &attendance.ApplyMakeupResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if existingRecord != nil {
		return &attendance.ApplyMakeupResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "该日期已有考勤记录",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 创建补签考勤记录
	record := &models.AttendanceRecord{
		JobID:     req.JobID,
		WorkerID:  workerID,
		CheckIn:   &makeupDate,
		Status:    "normal",
		WorkHours: decimal.Zero,
	}

	err = mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.CreateAttendanceRecord(tx, record)
	})

	if err != nil {
		utils.Errorf("申请补签失败: %v", err)
		return &attendance.ApplyMakeupResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "申请补签失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &attendance.ApplyMakeupResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "申请补签成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		MakeupID: record.ID,
	}, nil
}
