package schedule

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/schedule"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CreateScheduleLogic 创建日程业务逻辑
func CreateScheduleLogic(ctx context.Context, req *schedule.CreateScheduleReq) (*schedule.CreateScheduleResp, error) {
	// 解析时间
	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		utils.Errorf("解析开始时间失败: %v", err)
		return &schedule.CreateScheduleResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "开始时间格式错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		utils.Errorf("解析结束时间失败: %v", err)
		return &schedule.CreateScheduleResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "结束时间格式错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 验证时间
	if endTime.Before(startTime) {
		return &schedule.CreateScheduleResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "结束时间不能早于开始时间",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 创建日程
	scheduleModel := &models.Schedule{
		WorkerID:        0, // 需要从JWT token中获取
		JobID:           &req.JobID,
		Title:           req.Title,
		StartTime:       startTime,
		EndTime:         endTime,
		Location:        req.Location,
		Notes:           req.Notes,
		Status:          "pending",
		ReminderMinutes: int(req.ReminderMinutes),
	}

	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.CreateSchedule(tx, scheduleModel)
	})

	if err != nil {
		utils.Errorf("创建日程失败: %v", err)
		return &schedule.CreateScheduleResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "创建日程失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &schedule.CreateScheduleResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "创建日程成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		ScheduleID: scheduleModel.ID,
	}, nil
}
