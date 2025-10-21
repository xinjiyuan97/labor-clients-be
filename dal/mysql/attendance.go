package mysql

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Attendance相关数据库操作

// CreateAttendanceRecord 创建考勤记录
func CreateAttendanceRecord(tx *gorm.DB, record *models.AttendanceRecord) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(record).Error; err != nil {
		utils.Errorf("创建考勤记录失败: %v", err)
		return err
	}

	return nil
}

// GetAttendanceRecordByID 根据ID获取考勤记录详情
func GetAttendanceRecordByID(tx *gorm.DB, recordID int64) (*models.AttendanceRecord, error) {
	if tx == nil {
		tx = DB
	}

	var record models.AttendanceRecord
	if err := tx.Where("id = ?", recordID).First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询考勤记录详情失败: %v", err)
		return nil, err
	}

	return &record, nil
}

// GetAttendanceRecords 获取考勤记录列表
func GetAttendanceRecords(tx *gorm.DB, workerID int64, jobID *int64, startDate, endDate string, offset, limit int) ([]*models.AttendanceRecord, error) {
	if tx == nil {
		tx = DB
	}

	var records []*models.AttendanceRecord
	query := tx.Where("worker_id = ?", workerID)

	if jobID != nil && *jobID > 0 {
		query = query.Where("job_id = ?", *jobID)
	}

	if startDate != "" {
		query = query.Where("check_in >= ?", startDate)
	}

	if endDate != "" {
		query = query.Where("check_in <= ?", endDate)
	}

	if err := query.Offset(offset).
		Limit(limit).
		Order("check_in DESC").
		Find(&records).Error; err != nil {
		utils.Errorf("获取考勤记录列表失败: %v", err)
		return nil, err
	}

	return records, nil
}

// CountAttendanceRecords 统计考勤记录数量
func CountAttendanceRecords(tx *gorm.DB, workerID int64, jobID *int64, startDate, endDate string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	query := tx.Model(&models.AttendanceRecord{}).Where("worker_id = ?", workerID)

	if jobID != nil && *jobID > 0 {
		query = query.Where("job_id = ?", *jobID)
	}

	if startDate != "" {
		query = query.Where("check_in >= ?", startDate)
	}

	if endDate != "" {
		query = query.Where("check_in <= ?", endDate)
	}

	if err := query.Count(&count).Error; err != nil {
		utils.Errorf("统计考勤记录数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// UpdateAttendanceRecord 更新考勤记录
func UpdateAttendanceRecord(tx *gorm.DB, record *models.AttendanceRecord) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Save(record).Error; err != nil {
		utils.Errorf("更新考勤记录失败: %v", err)
		return err
	}

	return nil
}

// GetTodayAttendanceRecord 获取今日考勤记录
func GetTodayAttendanceRecord(tx *gorm.DB, workerID, jobID int64) (*models.AttendanceRecord, error) {
	if tx == nil {
		tx = DB
	}

	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	var record models.AttendanceRecord
	if err := tx.Where("worker_id = ? AND job_id = ? AND check_in >= ? AND check_in < ?",
		workerID, jobID, today, tomorrow).First(&record).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("获取今日考勤记录失败: %v", err)
		return nil, err
	}

	return &record, nil
}

// CalculateWorkHours 计算工作时长
func CalculateWorkHours(checkIn, checkOut time.Time) decimal.Decimal {
	duration := checkOut.Sub(checkIn)
	hours := duration.Hours()
	return decimal.NewFromFloat(hours).Round(2)
}
