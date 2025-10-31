package schedule

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/schedule"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// DeleteScheduleLogic 删除日程业务逻辑
func DeleteScheduleLogic(ctx context.Context, req *schedule.DeleteScheduleReq) (*schedule.DeleteScheduleResp, error) {
	// 检查日程是否存在
	existingSchedule, err := mysql.GetScheduleByID(nil, req.ScheduleID)
	if err != nil {
		utils.Errorf("获取日程详情失败: %v", err)
		return &schedule.DeleteScheduleResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if existingSchedule == nil {
		return &schedule.DeleteScheduleResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "日程不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 删除日程
	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		return mysql.DeleteSchedule(tx, req.ScheduleID)
	})

	if err != nil {
		utils.Errorf("删除日程失败: %v", err)
		return &schedule.DeleteScheduleResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "删除日程失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &schedule.DeleteScheduleResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "删除日程成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}, nil
}
