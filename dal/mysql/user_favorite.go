package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UserFavorite相关数据库操作

// CreateUserFavoriteJob 创建用户收藏工作
func CreateUserFavoriteJob(tx *gorm.DB, favorite *models.UserFavoriteJob) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(favorite).Error; err != nil {
		utils.Errorf("创建用户收藏工作失败: %v", err)
		return err
	}

	utils.Infof("用户收藏工作创建成功, UserID: %d, JobID: %d", favorite.UserID, favorite.JobID)
	return nil
}

// DeleteUserFavoriteJob 删除用户收藏工作
func DeleteUserFavoriteJob(tx *gorm.DB, userID, jobID int64) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Where("user_id = ? AND job_id = ?", userID, jobID).Delete(&models.UserFavoriteJob{}).Error; err != nil {
		utils.Errorf("删除用户收藏工作失败: %v", err)
		return err
	}

	utils.Infof("用户收藏工作删除成功, UserID: %d, JobID: %d", userID, jobID)
	return nil
}

// CheckUserFavoriteJob 检查用户是否已收藏工作
func CheckUserFavoriteJob(tx *gorm.DB, userID, jobID int64) (bool, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.UserFavoriteJob{}).Where("user_id = ? AND job_id = ?", userID, jobID).Count(&count).Error; err != nil {
		utils.Errorf("检查用户收藏工作失败: %v", err)
		return false, err
	}

	return count > 0, nil
}

// GetUserFavoriteJobs 获取用户收藏的工作列表
func GetUserFavoriteJobs(tx *gorm.DB, userID int64, offset, limit int) ([]*models.UserFavoriteJob, error) {
	if tx == nil {
		tx = DB
	}

	var favorites []*models.UserFavoriteJob
	if err := tx.Where("user_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&favorites).Error; err != nil {
		utils.Errorf("获取用户收藏工作列表失败: %v", err)
		return nil, err
	}

	return favorites, nil
}

// CountUserFavoriteJobs 统计用户收藏工作数量
func CountUserFavoriteJobs(tx *gorm.DB, userID int64) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.UserFavoriteJob{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		utils.Errorf("统计用户收藏工作数量失败: %v", err)
		return 0, err
	}

	return count, nil
}
