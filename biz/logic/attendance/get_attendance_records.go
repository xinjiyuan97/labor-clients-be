package attendance

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/attendance"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetAttendanceRecordsLogic 获取考勤记录列表业务逻辑
func GetAttendanceRecordsLogic(req *attendance.GetAttendanceRecordsReq) (*attendance.GetAttendanceRecordsResp, error) {
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

	// 获取用户ID，这里需要从JWT token中获取
	workerID := int64(0) // 需要从JWT token中获取

	// 获取考勤记录列表
	records, err := mysql.GetAttendanceRecords(nil, workerID, &req.JobID, req.StartDate, req.EndDate, offset, limit)
	if err != nil {
		utils.Errorf("获取考勤记录列表失败: %v", err)
		return &attendance.GetAttendanceRecordsResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取总数
	total, err := mysql.CountAttendanceRecords(nil, workerID, &req.JobID, req.StartDate, req.EndDate)
	if err != nil {
		utils.Errorf("获取考勤记录总数失败: %v", err)
		return &attendance.GetAttendanceRecordsResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建考勤记录信息
	var recordInfos []*common.AttendanceRecordInfo
	for _, record := range records {
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

		recordInfos = append(recordInfos, recordInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &attendance.GetAttendanceRecordsResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取考勤记录列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp: pageResp,
		Records:  recordInfos,
	}, nil
}
