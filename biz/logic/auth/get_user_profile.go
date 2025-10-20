package auth

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xinjiyuan97/labor-clients/biz/model/auth"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/middleware"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetUserProfileLogic 获取用户信息业务逻辑
func GetUserProfileLogic(ctx context.Context, c *app.RequestContext) (*auth.GetUserProfileResp, error) {
	// 根据用户ID查询用户信息
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		return &auth.GetUserProfileResp{
			Base: &common.BaseResp{
				Code:      401,
				Message:   "未登录",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	user, err := mysql.GetUserByID(nil, userID)
	if err != nil {
		utils.Errorf("查询用户失败: %v", err)
		return &auth.GetUserProfileResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if user == nil {
		return &auth.GetUserProfileResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "用户不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建用户基础信息
	userInfo := &common.UserInfo{
		UserID:   user.ID,
		Username: user.Username,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
		Role:     user.Role,
	}

	var workerInfo *common.WorkerInfo

	// 如果是零工角色，查询零工详细信息
	if user.Role == "worker" {
		worker, err := mysql.GetWorkerByUserID(nil, user.ID)
		if err != nil {
			utils.Errorf("查询零工信息失败: %v", err)
			return &auth.GetUserProfileResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}

		if worker != nil {
			workerInfo = &common.WorkerInfo{
				UserID:         worker.UserID,
				RealName:       worker.RealName,
				Gender:         string(worker.Gender),
				Age:            int32(worker.Age),
				Education:      worker.Education,
				Height:         worker.Height.InexactFloat64(),
				Introduction:   worker.Introduction,
				WorkExperience: worker.WorkExperience,
				ExpectedSalary: worker.ExpectedSalary.InexactFloat64(),
			}
		}
	}

	return &auth.GetUserProfileResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取用户信息成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		UserInfo:   userInfo,
		WorkerInfo: workerInfo,
	}, nil
}
