package schedule

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/schedule"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UpdateScheduleStatusLogic 更新日程状态业务逻辑
func UpdateScheduleStatusLogic(req *schedule.UpdateScheduleStatusReq) (*schedule.UpdateScheduleStatusResp, error) {
	// 检查日程是否存在
	existingSchedule, err := mysql.GetScheduleByID(nil, req.ScheduleID)
	if err != nil {
		utils.Errorf("获取日程详情失败: %v", err)
		return &schedule.UpdateScheduleStatusResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if existingSchedule == nil {
		return &schedule.UpdateScheduleStatusResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "日程不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 验证状态
	validStatuses := []string{"pending", "in_progress", "completed", "cancelled"}
	isValidStatus := false
	for _, status := range validStatuses {
		if req.Status == status {
			isValidStatus = true
			break
		}
	}

	if !isValidStatus {
		return &schedule.UpdateScheduleStatusResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "无效的状态",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 更新日程状态
	err = mysql.Transaction(func(tx *gorm.DB) error {
		return mysql.UpdateScheduleStatus(tx, req.ScheduleID, req.Status)
	})

	if err != nil {
		utils.Errorf("更新日程状态失败: %v", err)
		return &schedule.UpdateScheduleStatusResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "更新日程状态失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &schedule.UpdateScheduleStatusResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "更新日程状态成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
