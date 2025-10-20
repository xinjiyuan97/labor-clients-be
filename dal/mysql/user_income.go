package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UserIncome相关数据库操作

// GetUserPayments 获取用户支付记录
func GetUserPayments(tx *gorm.DB, userID int64, offset, limit int) ([]*models.Payment, error) {
	if tx == nil {
		tx = DB
	}

	var payments []*models.Payment
	if err := tx.Where("worker_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&payments).Error; err != nil {
		utils.Errorf("获取用户支付记录失败: %v", err)
		return nil, err
	}

	return payments, nil
}

// GetUserPaymentsByPeriod 根据时间段获取用户支付记录
func GetUserPaymentsByPeriod(tx *gorm.DB, userID int64, startDate, endDate string, offset, limit int) ([]*models.Payment, error) {
	if tx == nil {
		tx = DB
	}

	var payments []*models.Payment
	query := tx.Where("worker_id = ?", userID)

	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	if err := query.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&payments).Error; err != nil {
		utils.Errorf("根据时间段获取用户支付记录失败: %v", err)
		return nil, err
	}

	return payments, nil
}

// GetUserPaymentsByStatus 根据状态获取用户支付记录
func GetUserPaymentsByStatus(tx *gorm.DB, userID int64, status string, offset, limit int) ([]*models.Payment, error) {
	if tx == nil {
		tx = DB
	}

	var payments []*models.Payment
	query := tx.Where("worker_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&payments).Error; err != nil {
		utils.Errorf("根据状态获取用户支付记录失败: %v", err)
		return nil, err
	}

	return payments, nil
}

// CountUserPayments 统计用户支付记录数量
func CountUserPayments(tx *gorm.DB, userID int64) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Payment{}).Where("worker_id = ?", userID).Count(&count).Error; err != nil {
		utils.Errorf("统计用户支付记录数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountUserPaymentsByPeriod 根据时间段统计用户支付记录数量
func CountUserPaymentsByPeriod(tx *gorm.DB, userID int64, startDate, endDate string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	query := tx.Model(&models.Payment{}).Where("worker_id = ?", userID)

	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	if err := query.Count(&count).Error; err != nil {
		utils.Errorf("根据时间段统计用户支付记录数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountUserPaymentsByStatus 根据状态统计用户支付记录数量
func CountUserPaymentsByStatus(tx *gorm.DB, userID int64, status string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	query := tx.Model(&models.Payment{}).Where("worker_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&count).Error; err != nil {
		utils.Errorf("根据状态统计用户支付记录数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetUserIncomeStats 获取用户收入统计
func GetUserIncomeStats(tx *gorm.DB, userID int64, startDate, endDate string) (totalIncome, pendingIncome, paidIncome float64, err error) {
	if tx == nil {
		tx = DB
	}

	query := tx.Model(&models.Payment{}).Where("worker_id = ?", userID)

	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	// 总收入
	var totalResult struct {
		Total float64
	}
	if err = query.Select("SUM(amount) as total").Scan(&totalResult).Error; err != nil {
		utils.Errorf("获取总收入失败: %v", err)
		return
	}
	totalIncome = totalResult.Total

	// 待支付收入
	var pendingResult struct {
		Total float64
	}
	if err = query.Where("status IN ?", []string{"pending", "processing"}).Select("SUM(amount) as total").Scan(&pendingResult).Error; err != nil {
		utils.Errorf("获取待支付收入失败: %v", err)
		return
	}
	pendingIncome = pendingResult.Total

	// 已支付收入
	var paidResult struct {
		Total float64
	}
	if err = query.Where("status = ?", "completed").Select("SUM(amount) as total").Scan(&paidResult).Error; err != nil {
		utils.Errorf("获取已支付收入失败: %v", err)
		return
	}
	paidIncome = paidResult.Total

	return
}
