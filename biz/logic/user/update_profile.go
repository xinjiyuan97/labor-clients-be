package user

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/user"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UpdateProfileLogic 更新用户信息业务逻辑
func UpdateProfileLogic(ctx context.Context, req *user.UpdateProfileReq, userID int64) (*user.UpdateProfileResp, error) {
	// 获取当前用户信息
	currentUser, err := mysql.GetUserByID(nil, userID)
	if err != nil {
		utils.Errorf("获取用户信息失败: %v", err)
		return &user.UpdateProfileResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if currentUser == nil {
		return &user.UpdateProfileResp{
			Base: &common.BaseResp{
				Code:      404,
				Message:   "用户不存在",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 使用事务更新用户信息
	err = mysql.Transaction(ctx, func(tx *gorm.DB) error {
		// 更新用户基础信息
		if req.Username != "" {
			currentUser.Username = req.Username
		}
		if req.Avatar != "" {
			currentUser.Avatar = req.Avatar
		}

		if err := mysql.UpdateUserProfile(tx, currentUser); err != nil {
			return err
		}

		// 如果是worker角色且有worker信息，更新或创建worker详细信息
		if currentUser.Role == "worker" && req.WorkerInfo != nil {
			workerExists, err := mysql.CheckWorkerExists(tx, userID)
			if err != nil {
				return err
			}

			if workerExists {
				// 更新现有worker信息
				worker, err := mysql.GetWorkerByUserID(tx, userID)
				if err != nil {
					return err
				}
				if worker != nil {
					worker.RealName = req.WorkerInfo.RealName
					worker.Gender = models.Gender(req.WorkerInfo.Gender)
					worker.Age = uint8(req.WorkerInfo.Age)
					worker.Education = req.WorkerInfo.Education
					worker.Height = utils.DecimalFromFloat(req.WorkerInfo.Height)
					worker.Introduction = req.WorkerInfo.Introduction
					worker.WorkExperience = req.WorkerInfo.WorkExperience
					worker.ExpectedSalary = utils.DecimalFromFloat(req.WorkerInfo.ExpectedSalary)

					if err := mysql.UpdateWorkerProfile(tx, worker); err != nil {
						return err
					}
				}
			} else {
				// 创建新的worker信息
				worker := &models.Worker{
					UserID:         userID,
					RealName:       req.WorkerInfo.RealName,
					Gender:         models.Gender(req.WorkerInfo.Gender),
					Age:            uint8(req.WorkerInfo.Age),
					Education:      req.WorkerInfo.Education,
					Height:         utils.DecimalFromFloat(req.WorkerInfo.Height),
					Introduction:   req.WorkerInfo.Introduction,
					WorkExperience: req.WorkerInfo.WorkExperience,
					ExpectedSalary: utils.DecimalFromFloat(req.WorkerInfo.ExpectedSalary),
				}

				if err := mysql.CreateWorkerProfile(tx, worker); err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		utils.Errorf("更新用户信息失败: %v", err)
		return &user.UpdateProfileResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 重新获取更新后的用户信息
	updatedUser, err := mysql.GetUserByID(nil, userID)
	if err != nil {
		utils.Errorf("获取更新后的用户信息失败: %v", err)
		return &user.UpdateProfileResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建响应
	userInfo := &common.UserInfo{
		UserID:   updatedUser.ID,
		Username: updatedUser.Username,
		Phone:    updatedUser.Phone,
		Avatar:   updatedUser.Avatar,
		Role:     updatedUser.Role,
	}

	var workerInfo *common.WorkerInfo
	if updatedUser.Role == "worker" {
		worker, err := mysql.GetWorkerByUserID(nil, userID)
		if err != nil {
			utils.Errorf("获取worker信息失败: %v", err)
		} else if worker != nil {
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

	return &user.UpdateProfileResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "更新用户信息成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		UserInfo:   userInfo,
		WorkerInfo: workerInfo,
	}, nil
}
