package schedule

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/schedule"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/middleware"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetTodayScheduleLogic 获取今日日程业务逻辑
func GetTodayScheduleLogic(c *app.RequestContext, req *schedule.GetTodayScheduleReq) (*schedule.GetTodayScheduleResp, error) {
	// 获取今日日期
	var targetDate time.Time
	if req.Date != "" {
		var err error
		targetDate, err = time.Parse("2006-01-02", req.Date)
		if err != nil {
			utils.Errorf("解析日期失败: %v", err)
			return &schedule.GetTodayScheduleResp{
				Base: &common.BaseResp{
					Code:      400,
					Message:   "日期格式错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
	} else {
		targetDate = time.Now()
	}

	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		return &schedule.GetTodayScheduleResp{
			Base: &common.BaseResp{
				Code:      401,
				Message:   "未登录",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}
	// 获取今日日程
	schedules, err := mysql.GetSchedulesByDate(nil, userID, targetDate, 0, -1)
	if err != nil {
		utils.Errorf("获取今日日程失败: %v", err)
		return &schedule.GetTodayScheduleResp{
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
		scheduleInfo := scheduleModel.ToThrift()
		scheduleInfos = append(scheduleInfos, scheduleInfo)
	}

	return &schedule.GetTodayScheduleResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取今日日程成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Schedules: scheduleInfos,
	}, nil
}
