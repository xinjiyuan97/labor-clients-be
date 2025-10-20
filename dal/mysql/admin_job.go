package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Job管理相关数据库操作

// GetJobsForAdmin 获取工作列表（管理员用）
func GetJobsForAdmin(tx *gorm.DB, offset, limit int) ([]*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var jobs []*models.Job
	if err := tx.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&jobs).Error; err != nil {
		utils.Errorf("获取工作列表失败: %v", err)
		return nil, err
	}

	return jobs, nil
}

// GetJobsByStatus 根据状态获取工作列表
func GetJobsByStatus(tx *gorm.DB, status string, offset, limit int) ([]*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var jobs []*models.Job
	if err := tx.Where("status = ?", status).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&jobs).Error; err != nil {
		utils.Errorf("根据状态获取工作列表失败: %v", err)
		return nil, err
	}

	return jobs, nil
}

// ReviewJob 审核工作
func ReviewJob(tx *gorm.DB, jobID int64, status string) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.Job{}).Where("id = ?", jobID).Update("status", status).Error; err != nil {
		utils.Errorf("审核工作失败: %v", err)
		return err
	}

	return nil
}

// CountJobsForAdmin 统计工作数量（管理员用）
func CountJobsForAdmin(tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Job{}).Count(&count).Error; err != nil {
		utils.Errorf("统计工作数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountJobsByStatus 根据状态统计工作数量
func CountJobsByStatus(tx *gorm.DB, status string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Job{}).Where("status = ?", status).Count(&count).Error; err != nil {
		utils.Errorf("根据状态统计工作数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetJobStatistics 获取工作统计信息
func GetJobStatistics(tx *gorm.DB) (map[string]int64, error) {
	if tx == nil {
		tx = DB
	}

	stats := make(map[string]int64)

	// 总工作数
	var total int64
	if err := tx.Model(&models.Job{}).Count(&total).Error; err != nil {
		utils.Errorf("统计总工作数失败: %v", err)
		return nil, err
	}
	stats["total"] = total

	// 草稿工作数
	var draft int64
	if err := tx.Model(&models.Job{}).Where("status = ?", "draft").Count(&draft).Error; err != nil {
		utils.Errorf("统计草稿工作数失败: %v", err)
		return nil, err
	}
	stats["draft"] = draft

	// 已发布工作数
	var published int64
	if err := tx.Model(&models.Job{}).Where("status = ?", "published").Count(&published).Error; err != nil {
		utils.Errorf("统计已发布工作数失败: %v", err)
		return nil, err
	}
	stats["published"] = published

	// 已招满工作数
	var filled int64
	if err := tx.Model(&models.Job{}).Where("status = ?", "filled").Count(&filled).Error; err != nil {
		utils.Errorf("统计已招满工作数失败: %v", err)
		return nil, err
	}
	stats["filled"] = filled

	// 已完成工作数
	var completed int64
	if err := tx.Model(&models.Job{}).Where("status = ?", "completed").Count(&completed).Error; err != nil {
		utils.Errorf("统计已完成工作数失败: %v", err)
		return nil, err
	}
	stats["completed"] = completed

	// 已取消工作数
	var cancelled int64
	if err := tx.Model(&models.Job{}).Where("status = ?", "cancelled").Count(&cancelled).Error; err != nil {
		utils.Errorf("统计已取消工作数失败: %v", err)
		return nil, err
	}
	stats["cancelled"] = cancelled

	return stats, nil
}
