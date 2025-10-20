package mysql

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Statistics相关数据库操作

// GetUserStatistics 获取用户统计信息
func GetUserStatistics(tx *gorm.DB) (map[string]int64, error) {
	if tx == nil {
		tx = DB
	}

	stats := make(map[string]int64)

	// 总用户数
	var total int64
	if err := tx.Model(&models.User{}).Where("role IN ?", []string{"worker", "employer"}).Count(&total).Error; err != nil {
		utils.Errorf("统计总用户数失败: %v", err)
		return nil, err
	}
	stats["total"] = total

	// 工人数
	var workers int64
	if err := tx.Model(&models.User{}).Where("role = ?", "worker").Count(&workers).Error; err != nil {
		utils.Errorf("统计工人数失败: %v", err)
		return nil, err
	}
	stats["workers"] = workers

	// 雇主数
	var employers int64
	if err := tx.Model(&models.User{}).Where("role = ?", "employer").Count(&employers).Error; err != nil {
		utils.Errorf("统计雇主数失败: %v", err)
		return nil, err
	}
	stats["employers"] = employers

	// 今日新增用户数
	today := time.Now().Format("2006-01-02")
	var todayNew int64
	if err := tx.Model(&models.User{}).Where("role IN ? AND DATE(created_at) = ?", []string{"worker", "employer"}, today).Count(&todayNew).Error; err != nil {
		utils.Errorf("统计今日新增用户数失败: %v", err)
		return nil, err
	}
	stats["today_new"] = todayNew

	// 本周新增用户数
	weekStart := time.Now().AddDate(0, 0, -int(time.Now().Weekday())).Format("2006-01-02")
	var weekNew int64
	if err := tx.Model(&models.User{}).Where("role IN ? AND DATE(created_at) >= ?", []string{"worker", "employer"}, weekStart).Count(&weekNew).Error; err != nil {
		utils.Errorf("统计本周新增用户数失败: %v", err)
		return nil, err
	}
	stats["week_new"] = weekNew

	// 本月新增用户数
	monthStart := time.Now().AddDate(0, 0, -time.Now().Day()+1).Format("2006-01-02")
	var monthNew int64
	if err := tx.Model(&models.User{}).Where("role IN ? AND DATE(created_at) >= ?", []string{"worker", "employer"}, monthStart).Count(&monthNew).Error; err != nil {
		utils.Errorf("统计本月新增用户数失败: %v", err)
		return nil, err
	}
	stats["month_new"] = monthNew

	return stats, nil
}

// GetIncomeStatistics 获取收入统计信息
func GetIncomeStatistics(tx *gorm.DB) (map[string]float64, error) {
	if tx == nil {
		tx = DB
	}

	stats := make(map[string]float64)

	// 总收入
	var totalResult struct {
		Total float64 `json:"total"`
	}
	if err := tx.Model(&models.Payment{}).Select("SUM(amount) as total").Scan(&totalResult).Error; err != nil {
		utils.Errorf("统计总收入失败: %v", err)
		return nil, err
	}
	stats["total"] = totalResult.Total

	// 已支付收入
	var paidResult struct {
		Total float64 `json:"total"`
	}
	if err := tx.Model(&models.Payment{}).Where("status = ?", "paid").Select("SUM(amount) as total").Scan(&paidResult).Error; err != nil {
		utils.Errorf("统计已支付收入失败: %v", err)
		return nil, err
	}
	stats["paid"] = paidResult.Total

	// 待支付收入
	var pendingResult struct {
		Total float64 `json:"total"`
	}
	if err := tx.Model(&models.Payment{}).Where("status = ?", "pending").Select("SUM(amount) as total").Scan(&pendingResult).Error; err != nil {
		utils.Errorf("统计待支付收入失败: %v", err)
		return nil, err
	}
	stats["pending"] = pendingResult.Total

	// 今日收入
	today := time.Now().Format("2006-01-02")
	var todayResult struct {
		Total float64 `json:"total"`
	}
	if err := tx.Model(&models.Payment{}).Where("status = ? AND DATE(created_at) = ?", "paid", today).Select("SUM(amount) as total").Scan(&todayResult).Error; err != nil {
		utils.Errorf("统计今日收入失败: %v", err)
		return nil, err
	}
	stats["today"] = todayResult.Total

	// 本周收入
	weekStart := time.Now().AddDate(0, 0, -int(time.Now().Weekday())).Format("2006-01-02")
	var weekResult struct {
		Total float64 `json:"total"`
	}
	if err := tx.Model(&models.Payment{}).Where("status = ? AND DATE(created_at) >= ?", "paid", weekStart).Select("SUM(amount) as total").Scan(&weekResult).Error; err != nil {
		utils.Errorf("统计本周收入失败: %v", err)
		return nil, err
	}
	stats["week"] = weekResult.Total

	// 本月收入
	monthStart := time.Now().AddDate(0, 0, -time.Now().Day()+1).Format("2006-01-02")
	var monthResult struct {
		Total float64 `json:"total"`
	}
	if err := tx.Model(&models.Payment{}).Where("status = ? AND DATE(created_at) >= ?", "paid", monthStart).Select("SUM(amount) as total").Scan(&monthResult).Error; err != nil {
		utils.Errorf("统计本月收入失败: %v", err)
		return nil, err
	}
	stats["month"] = monthResult.Total

	return stats, nil
}
