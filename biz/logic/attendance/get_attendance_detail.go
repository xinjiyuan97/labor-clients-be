package attendance

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/attendance"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetAttendanceDetailLogic 获取考勤记录详情业务逻辑
func GetAttendanceDetailLogic(req *attendance.GetAttendanceDetailReq) (*attendance.GetAttendanceDetailResp, error) {
	// 获取考勤记录详情
	record, err := mysql.GetAttendanceRecordByID(nil, int64(req.RecordID))
	if err != nil {
		utils.Errorf("获取考勤记录详情失败: %v", err)
		return &attendance.GetAttendanceDetailResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if record == nil {
		return &attendance.GetAttendanceDetailResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "考勤记录不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建考勤记录信息
	workHours, _ := record.WorkHours.Float64()

	recordInfo := &common.AttendanceRecordInfo{
		RecordID:         int32(record.ID),
		JobID:            record.JobID,
		WorkerID:         record.WorkerID,
		CheckIn:          "",
		CheckOut:         "",
		WorkHours:        workHours,
		CheckInLocation:  record.CheckInLocation,
		CheckOutLocation: record.CheckOutLocation,
		Status:           record.Status,
	}

	if record.CheckIn != nil {
		recordInfo.CheckIn = record.CheckIn.Format(time.RFC3339)
	}
	if record.CheckOut != nil {
		recordInfo.CheckOut = record.CheckOut.Format(time.RFC3339)
	}

	return &attendance.GetAttendanceDetailResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取考勤记录详情成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Record: recordInfo,
	}, nil
}
