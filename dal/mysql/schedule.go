package mysql

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Schedule相关数据库操作

// CreateSchedule 创建日程
func CreateSchedule(tx *gorm.DB, schedule *models.Schedule) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(schedule).Error; err != nil {
		utils.Errorf("创建日程失败: %v", err)
		return err
	}

	return nil
}

// GetSchedules 获取日程列表
func GetSchedules(tx *gorm.DB, workerID int64, offset, limit int) ([]*models.Schedule, error) {
	if tx == nil {
		tx = DB
	}

	var schedules []*models.Schedule
	if err := tx.Where("worker_id = ?", workerID).
		Offset(offset).
		Limit(limit).
		Order("start_time ASC").
		Find(&schedules).Error; err != nil {
		utils.Errorf("获取日程列表失败: %v", err)
		return nil, err
	}

	return schedules, nil
}

// GetSchedulesByDate 根据日期获取日程列表
func GetSchedulesByDate(tx *gorm.DB, workerID int64, date time.Time, offset, limit int) ([]*models.Schedule, error) {
	if tx == nil {
		tx = DB
	}

	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var schedules []*models.Schedule
	if err := tx.Where("worker_id = ? AND start_time >= ? AND start_time < ?", workerID, startOfDay, endOfDay).
		Offset(offset).
		Limit(limit).
		Order("start_time ASC").
		Find(&schedules).Error; err != nil {
		utils.Errorf("根据日期获取日程列表失败: %v", err)
		return nil, err
	}

	return schedules, nil
}

// GetSchedulesByStatus 根据状态获取日程列表
func GetSchedulesByStatus(tx *gorm.DB, workerID int64, status string, offset, limit int) ([]*models.Schedule, error) {
	if tx == nil {
		tx = DB
	}

	var schedules []*models.Schedule
	if err := tx.Where("worker_id = ? AND status = ?", workerID, status).
		Offset(offset).
		Limit(limit).
		Order("start_time ASC").
		Find(&schedules).Error; err != nil {
		utils.Errorf("根据状态获取日程列表失败: %v", err)
		return nil, err
	}

	return schedules, nil
}

// GetScheduleByID 根据ID获取日程详情
func GetScheduleByID(tx *gorm.DB, scheduleID int64) (*models.Schedule, error) {
	if tx == nil {
		tx = DB
	}

	var schedule models.Schedule
	if err := tx.Where("id = ?", scheduleID).First(&schedule).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询日程详情失败: %v", err)
		return nil, err
	}

	return &schedule, nil
}

// UpdateSchedule 更新日程信息
func UpdateSchedule(tx *gorm.DB, schedule *models.Schedule) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(schedule).Where("id = ?", schedule.ID).Updates(schedule).Error; err != nil {
		utils.Errorf("更新日程信息失败: %v", err)
		return err
	}

	return nil
}

// UpdateScheduleStatus 更新日程状态
func UpdateScheduleStatus(tx *gorm.DB, scheduleID int64, status string) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.Schedule{}).Where("id = ?", scheduleID).Update("status", status).Error; err != nil {
		utils.Errorf("更新日程状态失败: %v", err)
		return err
	}

	return nil
}

// DeleteSchedule 删除日程
func DeleteSchedule(tx *gorm.DB, scheduleID int64) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Where("id = ?", scheduleID).Delete(&models.Schedule{}).Error; err != nil {
		utils.Errorf("删除日程失败: %v", err)
		return err
	}

	return nil
}

// CountSchedules 统计日程数量
func CountSchedules(tx *gorm.DB, workerID int64) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Schedule{}).Where("worker_id = ?", workerID).Count(&count).Error; err != nil {
		utils.Errorf("统计日程数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountSchedulesByDate 根据日期统计日程数量
func CountSchedulesByDate(tx *gorm.DB, workerID int64, date time.Time) (int64, error) {
	if tx == nil {
		tx = DB
	}

	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var count int64
	if err := tx.Model(&models.Schedule{}).Where("worker_id = ? AND start_time >= ? AND start_time < ?", workerID, startOfDay, endOfDay).Count(&count).Error; err != nil {
		utils.Errorf("根据日期统计日程数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountSchedulesByStatus 根据状态统计日程数量
func CountSchedulesByStatus(tx *gorm.DB, workerID int64, status string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Schedule{}).Where("worker_id = ? AND status = ?", workerID, status).Count(&count).Error; err != nil {
		utils.Errorf("根据状态统计日程数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetSchedulesByDateRange 根据日期范围获取日程列表
func GetSchedulesByDateRange(tx *gorm.DB, workerID int64, startDate, endDate time.Time, offset, limit int) ([]*models.Schedule, error) {
	if tx == nil {
		tx = DB
	}

	var schedules []*models.Schedule
	if err := tx.Where("worker_id = ? AND start_time >= ? AND start_time <= ?", workerID, startDate, endDate).
		Offset(offset).
		Limit(limit).
		Order("start_time ASC").
		Find(&schedules).Error; err != nil {
		utils.Errorf("根据日期范围获取日程列表失败: %v", err)
		return nil, err
	}

	return schedules, nil
}

// CountSchedulesByDateRange 根据日期范围统计日程数量
func CountSchedulesByDateRange(tx *gorm.DB, workerID int64, startDate, endDate time.Time) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Schedule{}).Where("worker_id = ? AND start_time >= ? AND start_time <= ?", workerID, startDate, endDate).Count(&count).Error; err != nil {
		utils.Errorf("根据日期范围统计日程数量失败: %v", err)
		return 0, err
	}

	return count, nil
}
