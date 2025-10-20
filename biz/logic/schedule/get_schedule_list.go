package schedule

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/schedule"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetScheduleListLogic 获取日程列表业务逻辑
func GetScheduleListLogic(req *schedule.GetScheduleListReq) (*schedule.GetScheduleListResp, error) {
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

	// 获取日程列表
	var schedules []*models.Schedule
	var total int64
	var err error

	if req.StartDate != "" && req.EndDate != "" {
		// 根据日期范围获取
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			utils.Errorf("解析开始日期失败: %v", err)
			return &schedule.GetScheduleListResp{
				Base: &common.BaseResp{
					Code:      400,
					Message:   "开始日期格式错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			utils.Errorf("解析结束日期失败: %v", err)
			return &schedule.GetScheduleListResp{
				Base: &common.BaseResp{
					Code:      400,
					Message:   "结束日期格式错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		schedules, err = mysql.GetSchedulesByDateRange(nil, 0, startDate, endDate, offset, limit) // WorkerID需要从JWT获取
		if err != nil {
			utils.Errorf("根据日期范围获取日程列表失败: %v", err)
			return &schedule.GetScheduleListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountSchedulesByDateRange(nil, 0, startDate, endDate)
	} else if req.Status != "" {
		// 根据状态获取
		schedules, err = mysql.GetSchedulesByStatus(nil, 0, req.Status, offset, limit) // WorkerID需要从JWT获取
		if err != nil {
			utils.Errorf("根据状态获取日程列表失败: %v", err)
			return &schedule.GetScheduleListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountSchedulesByStatus(nil, 0, req.Status)
	} else {
		// 获取所有日程
		schedules, err = mysql.GetSchedules(nil, 0, offset, limit) // WorkerID需要从JWT获取
		if err != nil {
			utils.Errorf("获取日程列表失败: %v", err)
			return &schedule.GetScheduleListResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
		total, err = mysql.CountSchedules(nil, 0)
	}

	if err != nil {
		utils.Errorf("获取日程总数失败: %v", err)
		return &schedule.GetScheduleListResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建日程信息
	var scheduleInfos []*common.ScheduleInfo
	for _, scheduleModel := range schedules {
		var jobID int64
		if scheduleModel.JobID != nil {
			jobID = *scheduleModel.JobID
		}
		scheduleInfo := &common.ScheduleInfo{
			ScheduleID:      scheduleModel.ID,
			WorkerID:        scheduleModel.WorkerID,
			JobID:           jobID,
			Title:           scheduleModel.Title,
			StartTime:       scheduleModel.StartTime.Format(time.RFC3339),
			EndTime:         scheduleModel.EndTime.Format(time.RFC3339),
			Location:        scheduleModel.Location,
			Notes:           scheduleModel.Notes,
			Status:          scheduleModel.Status,
			ReminderMinutes: int32(scheduleModel.ReminderMinutes),
		}
		scheduleInfos = append(scheduleInfos, scheduleInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &schedule.GetScheduleListResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取日程列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp:  pageResp,
		Schedules: scheduleInfos,
	}, nil
}
