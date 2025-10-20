package schedule

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/schedule"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UpdateScheduleLogic 更新日程业务逻辑
func UpdateScheduleLogic(req *schedule.UpdateScheduleReq) (*schedule.UpdateScheduleResp, error) {
	// 检查日程是否存在
	existingSchedule, err := mysql.GetScheduleByID(nil, req.ScheduleID)
	if err != nil {
		utils.Errorf("获取日程详情失败: %v", err)
		return &schedule.UpdateScheduleResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if existingSchedule == nil {
		return &schedule.UpdateScheduleResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "日程不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 解析时间
	var startTime, endTime time.Time
	if req.StartTime != "" {
		startTime, err = time.Parse(time.RFC3339, req.StartTime)
		if err != nil {
			utils.Errorf("解析开始时间失败: %v", err)
			return &schedule.UpdateScheduleResp{
				Base: &common.BaseResp{
					Code:      400,
					Message:   "开始时间格式错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
	} else {
		startTime = existingSchedule.StartTime
	}

	if req.EndTime != "" {
		endTime, err = time.Parse(time.RFC3339, req.EndTime)
		if err != nil {
			utils.Errorf("解析结束时间失败: %v", err)
			return &schedule.UpdateScheduleResp{
				Base: &common.BaseResp{
					Code:      400,
					Message:   "结束时间格式错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
	} else {
		endTime = existingSchedule.EndTime
	}

	// 验证时间
	if endTime.Before(startTime) {
		return &schedule.UpdateScheduleResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "结束时间不能早于开始时间",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 更新日程
	scheduleModel := &models.Schedule{
		WorkerID:        existingSchedule.WorkerID,
		Title:           req.Title,
		StartTime:       startTime,
		EndTime:         endTime,
		Location:        req.Location,
		Notes:           req.Notes,
		Status:          existingSchedule.Status,
		ReminderMinutes: int(req.ReminderMinutes),
	}

	err = mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.UpdateSchedule(tx, scheduleModel)
	})

	if err != nil {
		utils.Errorf("更新日程失败: %v", err)
		return &schedule.UpdateScheduleResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "更新日程失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &schedule.UpdateScheduleResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "更新日程成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
