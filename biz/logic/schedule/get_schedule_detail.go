package schedule

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/schedule"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetScheduleDetailLogic 获取日程详情业务逻辑
func GetScheduleDetailLogic(req *schedule.GetScheduleDetailReq) (*schedule.GetScheduleDetailResp, error) {
	// 获取日程详情
	scheduleModel, err := mysql.GetScheduleByID(nil, req.ScheduleID)
	if err != nil {
		utils.Errorf("获取日程详情失败: %v", err)
		return &schedule.GetScheduleDetailResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if scheduleModel == nil {
		return &schedule.GetScheduleDetailResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "日程不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建日程信息
	scheduleInfo := scheduleModel.ToThrift()

	return &schedule.GetScheduleDetailResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取日程详情成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Schedule: scheduleInfo,
	}, nil
}
