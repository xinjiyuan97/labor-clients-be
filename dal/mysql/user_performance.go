package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UserPerformance相关数据库操作

// GetUserJobApplications 获取用户工作申请记录
func GetUserJobApplications(tx *gorm.DB, userID int64, offset, limit int) ([]*models.JobApplication, error) {
	if tx == nil {
		tx = DB
	}

	var applications []*models.JobApplication
	if err := tx.Where("worker_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&applications).Error; err != nil {
		utils.Errorf("获取用户工作申请记录失败: %v", err)
		return nil, err
	}

	return applications, nil
}

// CountUserJobApplications 统计用户工作申请数量
func CountUserJobApplications(tx *gorm.DB, userID int64) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.JobApplication{}).Where("worker_id = ?", userID).Count(&count).Error; err != nil {
		utils.Errorf("统计用户工作申请数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountUserCompletedJobs 统计用户完成的工作数量
func CountUserCompletedJobs(tx *gorm.DB, userID int64) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.JobApplication{}).Where("worker_id = ? AND status = ?", userID, "completed").Count(&count).Error; err != nil {
		utils.Errorf("统计用户完成工作数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetUserReviews 获取用户收到的评价
func GetUserReviews(tx *gorm.DB, userID int64, offset, limit int) ([]*models.Review, error) {
	if tx == nil {
		tx = DB
	}

	var reviews []*models.Review
	if err := tx.Where("worker_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&reviews).Error; err != nil {
		utils.Errorf("获取用户评价记录失败: %v", err)
		return nil, err
	}

	return reviews, nil
}

// CountUserReviews 统计用户收到的评价数量
func CountUserReviews(tx *gorm.DB, userID int64) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Review{}).Where("worker_id = ?", userID).Count(&count).Error; err != nil {
		utils.Errorf("统计用户评价数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetUserAverageRating 获取用户平均评分
func GetUserAverageRating(tx *gorm.DB, userID int64) (float64, error) {
	if tx == nil {
		tx = DB
	}

	var result struct {
		Average float64
	}

	if err := tx.Model(&models.Review{}).
		Where("worker_id = ?", userID).
		Select("AVG(rating) as average").
		Scan(&result).Error; err != nil {
		utils.Errorf("获取用户平均评分失败: %v", err)
		return 0, err
	}

	return result.Average, nil
}

// GetUserPerformanceStats 获取用户绩效统计
func GetUserPerformanceStats(tx *gorm.DB, userID int64) (totalApplications, completedJobs int64, successRate, averageRating float64, err error) {
	if tx == nil {
		tx = DB
	}

	// 总申请数
	totalApplications, err = CountUserJobApplications(tx, userID)
	if err != nil {
		return
	}

	// 完成工作数
	completedJobs, err = CountUserCompletedJobs(tx, userID)
	if err != nil {
		return
	}

	// 成功率
	if totalApplications > 0 {
		successRate = float64(completedJobs) / float64(totalApplications) * 100
	}

	// 平均评分
	averageRating, err = GetUserAverageRating(tx, userID)
	if err != nil {
		return
	}

	return
}
