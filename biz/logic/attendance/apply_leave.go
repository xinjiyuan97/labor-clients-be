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

// ApplyLeaveLogic 申请请假业务逻辑
func ApplyLeaveLogic(workerID int64, req *attendance.ApplyLeaveReq) (*attendance.ApplyLeaveResp, error) {
	// 解析请假日期
	leaveDate, err := time.Parse("2006-01-02", req.LeaveDate)
	if err != nil {
		return &attendance.ApplyLeaveResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "日期格式错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 检查该日期是否已有考勤记录
	existingRecord, err := mysql.GetTodayAttendanceRecord(nil, workerID, req.JobID)
	if err != nil {
		utils.Errorf("检查考勤记录失败: %v", err)
		return &attendance.ApplyLeaveResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if existingRecord != nil {
		return &attendance.ApplyLeaveResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "该日期已有考勤记录",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 创建请假考勤记录
	record := &models.AttendanceRecord{
		JobID:     req.JobID,
		WorkerID:  workerID,
		CheckIn:   &leaveDate,
		Status:    "leave",
		WorkHours: decimal.Zero,
	}

	err = mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.CreateAttendanceRecord(tx, record)
	})

	if err != nil {
		utils.Errorf("申请请假失败: %v", err)
		return &attendance.ApplyLeaveResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "申请请假失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &attendance.ApplyLeaveResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "申请请假成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		LeaveID: record.ID,
	}, nil
}
