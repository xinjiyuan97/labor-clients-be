package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UserProfile相关数据库操作

// UpdateUserProfile 更新用户基础信息
func UpdateUserProfile(tx *gorm.DB, user *models.User) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Save(user).Error; err != nil {
		utils.Errorf("更新用户基础信息失败: %v", err)
		return err
	}

	utils.Infof("用户基础信息更新成功, ID: %d", user.ID)
	return nil
}

// UpdateWorkerProfile 更新零工详细信息
func UpdateWorkerProfile(tx *gorm.DB, worker *models.Worker) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Save(worker).Error; err != nil {
		utils.Errorf("更新零工信息失败: %v", err)
		return err
	}

	utils.Infof("零工信息更新成功, UserID: %d", worker.UserID)
	return nil
}

// CreateWorkerProfile 创建零工详细信息
func CreateWorkerProfile(tx *gorm.DB, worker *models.Worker) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(worker).Error; err != nil {
		utils.Errorf("创建零工信息失败: %v", err)
		return err
	}

	utils.Infof("零工信息创建成功, UserID: %d", worker.UserID)
	return nil
}

// GetWorkerByUserID 根据用户ID获取零工信息
func GetWorkerByUserID(tx *gorm.DB, userID int64) (*models.Worker, error) {
	if tx == nil {
		tx = DB
	}

	var worker models.Worker
	if err := tx.Where("user_id = ?", userID).First(&worker).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据用户ID查询零工信息失败: %v", err)
		return nil, err
	}

	return &worker, nil
}

// CheckWorkerExists 检查零工信息是否存在
func CheckWorkerExists(tx *gorm.DB, userID int64) (bool, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Worker{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		utils.Errorf("检查零工信息是否存在失败: %v", err)
		return false, err
	}

	return count > 0, nil
}
