package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Review相关数据库操作

// CreateReview 创建评价
func CreateReview(tx *gorm.DB, review *models.Review) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(review).Error; err != nil {
		utils.Errorf("创建评价失败: %v", err)
		return err
	}

	return nil
}

// GetReviewByID 根据ID获取评价详情
func GetReviewByID(tx *gorm.DB, reviewID int64) (*models.Review, error) {
	if tx == nil {
		tx = DB
	}

	var review models.Review
	if err := tx.Where("id = ?", reviewID).First(&review).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询评价详情失败: %v", err)
		return nil, err
	}

	return &review, nil
}

// GetReceivedReviews 获取收到的评价列表
func GetReceivedReviews(tx *gorm.DB, userID int64, reviewType string, offset, limit int) ([]*models.Review, error) {
	if tx == nil {
		tx = DB
	}

	var reviews []*models.Review
	query := tx

	// 根据评价类型确定查询条件
	if reviewType == "employer_to_worker" {
		// 零工收到的评价（雇主评价零工）
		query = query.Where("worker_id = ? AND review_type = ?", userID, reviewType)
	} else if reviewType == "worker_to_employer" {
		// 雇主收到的评价（零工评价雇主）
		query = query.Where("employer_id = ? AND review_type = ?", userID, reviewType)
	} else {
		// 所有收到的评价
		query = query.Where("worker_id = ? OR employer_id = ?", userID, userID)
	}

	if err := query.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&reviews).Error; err != nil {
		utils.Errorf("获取收到的评价列表失败: %v", err)
		return nil, err
	}

	return reviews, nil
}

// CountReceivedReviews 统计收到的评价数量
func CountReceivedReviews(tx *gorm.DB, userID int64, reviewType string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	query := tx.Model(&models.Review{})

	// 根据评价类型确定查询条件
	if reviewType == "employer_to_worker" {
		query = query.Where("worker_id = ? AND review_type = ?", userID, reviewType)
	} else if reviewType == "worker_to_employer" {
		query = query.Where("employer_id = ? AND review_type = ?", userID, reviewType)
	} else {
		query = query.Where("worker_id = ? OR employer_id = ?", userID, userID)
	}

	if err := query.Count(&count).Error; err != nil {
		utils.Errorf("统计收到的评价数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetGivenReviews 获取发出的评价列表
func GetGivenReviews(tx *gorm.DB, userID int64, reviewType string, offset, limit int) ([]*models.Review, error) {
	if tx == nil {
		tx = DB
	}

	var reviews []*models.Review
	query := tx

	// 根据评价类型确定查询条件
	if reviewType == "employer_to_worker" {
		// 雇主发出的评价（雇主评价零工）
		query = query.Where("employer_id = ? AND review_type = ?", userID, reviewType)
	} else if reviewType == "worker_to_employer" {
		// 零工发出的评价（零工评价雇主）
		query = query.Where("worker_id = ? AND review_type = ?", userID, reviewType)
	} else {
		// 所有发出的评价
		query = query.Where("employer_id = ? OR worker_id = ?", userID, userID)
	}

	if err := query.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&reviews).Error; err != nil {
		utils.Errorf("获取发出的评价列表失败: %v", err)
		return nil, err
	}

	return reviews, nil
}

// CountGivenReviews 统计发出的评价数量
func CountGivenReviews(tx *gorm.DB, userID int64, reviewType string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	query := tx.Model(&models.Review{})

	// 根据评价类型确定查询条件
	if reviewType == "employer_to_worker" {
		query = query.Where("employer_id = ? AND review_type = ?", userID, reviewType)
	} else if reviewType == "worker_to_employer" {
		query = query.Where("worker_id = ? AND review_type = ?", userID, reviewType)
	} else {
		query = query.Where("employer_id = ? OR worker_id = ?", userID, userID)
	}

	if err := query.Count(&count).Error; err != nil {
		utils.Errorf("统计发出的评价数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// UpdateReview 更新评价
func UpdateReview(tx *gorm.DB, review *models.Review) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Save(review).Error; err != nil {
		utils.Errorf("更新评价失败: %v", err)
		return err
	}

	return nil
}

// GetReviewByJobAndUsers 根据工作和用户获取评价
func GetReviewByJobAndUsers(tx *gorm.DB, jobID, employerID, workerID int64, reviewType string) (*models.Review, error) {
	if tx == nil {
		tx = DB
	}

	var review models.Review
	if err := tx.Where("job_id = ? AND employer_id = ? AND worker_id = ? AND review_type = ?",
		jobID, employerID, workerID, reviewType).First(&review).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据工作和用户查询评价失败: %v", err)
		return nil, err
	}

	return &review, nil
}
