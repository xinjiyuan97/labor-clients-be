package mysql

import (
	"math"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Job相关数据库操作

// GetJobs 获取工作列表
func GetJobs(tx *gorm.DB, offset, limit int) ([]*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var jobs []*models.Job
	if err := tx.Where("status = ?", "published").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&jobs).Error; err != nil {
		utils.Errorf("获取工作列表失败: %v", err)
		return nil, err
	}

	return jobs, nil
}

// GetJobsByCategory 根据分类获取工作列表
func GetJobsByCategory(tx *gorm.DB, categoryID int64, offset, limit int) ([]*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var jobs []*models.Job
	if err := tx.Where("category_id = ? AND status = ?", categoryID, "published").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&jobs).Error; err != nil {
		utils.Errorf("根据分类获取工作列表失败: %v", err)
		return nil, err
	}

	return jobs, nil
}

// GetJobsByType 根据类型获取工作列表
func GetJobsByType(tx *gorm.DB, jobType string, offset, limit int) ([]*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var jobs []*models.Job
	if err := tx.Where("job_type = ? AND status = ?", jobType, "published").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&jobs).Error; err != nil {
		utils.Errorf("根据类型获取工作列表失败: %v", err)
		return nil, err
	}

	return jobs, nil
}

// GetJobsBySalaryRange 根据薪资范围获取工作列表
func GetJobsBySalaryRange(tx *gorm.DB, salaryMin, salaryMax float64, offset, limit int) ([]*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var jobs []*models.Job
	query := tx.Where("status = ?", "published")

	if salaryMin > 0 {
		query = query.Where("salary >= ?", salaryMin)
	}
	if salaryMax > 0 {
		query = query.Where("salary <= ?", salaryMax)
	}

	if err := query.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&jobs).Error; err != nil {
		utils.Errorf("根据薪资范围获取工作列表失败: %v", err)
		return nil, err
	}

	return jobs, nil
}

// SearchJobs 搜索工作
func SearchJobs(tx *gorm.DB, keyword string, offset, limit int) ([]*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var jobs []*models.Job
	if err := tx.Where("status = ? AND (title LIKE ? OR description LIKE ? OR location LIKE ?)",
		"published", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&jobs).Error; err != nil {
		utils.Errorf("搜索工作失败: %v", err)
		return nil, err
	}

	return jobs, nil
}

// SearchJobsWithFilters 带过滤条件的搜索工作
func SearchJobsWithFilters(tx *gorm.DB, keyword, location string, salaryMin, salaryMax float64, offset, limit int) ([]*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var jobs []*models.Job
	query := tx.Where("status = ?", "published")

	if keyword != "" {
		query = query.Where("(title LIKE ? OR description LIKE ?)", "%"+keyword+"%", "%"+keyword+"%")
	}
	if location != "" {
		query = query.Where("location LIKE ?", "%"+location+"%")
	}
	if salaryMin > 0 {
		query = query.Where("salary >= ?", salaryMin)
	}
	if salaryMax > 0 {
		query = query.Where("salary <= ?", salaryMax)
	}

	if err := query.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&jobs).Error; err != nil {
		utils.Errorf("带过滤条件搜索工作失败: %v", err)
		return nil, err
	}

	return jobs, nil
}

// GetJobByID 根据ID获取工作详情
func GetJobByID(tx *gorm.DB, jobID int64) (*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var job models.Job
	if err := tx.Where("id = ? AND status = ?", jobID, "published").First(&job).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询工作详情失败: %v", err)
		return nil, err
	}

	return &job, nil
}

// GetJobsNearby 获取附近的工作
func GetJobsNearby(tx *gorm.DB, latitude, longitude float64, distance float64, offset, limit int) ([]*models.Job, error) {
	if tx == nil {
		tx = DB
	}

	var jobs []*models.Job
	// 使用简单的矩形范围查询，实际项目中可以使用更精确的地理位置查询
	latRange := distance / 111.0 // 1度约等于111公里
	lngRange := distance / (111.0 * math.Cos(latitude*math.Pi/180))

	if err := tx.Where("status = ? AND latitude BETWEEN ? AND ? AND longitude BETWEEN ? AND ?",
		"published", latitude-latRange, latitude+latRange, longitude-lngRange, longitude+lngRange).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&jobs).Error; err != nil {
		utils.Errorf("获取附近工作失败: %v", err)
		return nil, err
	}

	return jobs, nil
}

// CountJobs 统计工作数量
func CountJobs(tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Job{}).Where("status = ?", "published").Count(&count).Error; err != nil {
		utils.Errorf("统计工作数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountJobsByCategory 根据分类统计工作数量
func CountJobsByCategory(tx *gorm.DB, categoryID int64) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Job{}).Where("category_id = ? AND status = ?", categoryID, "published").Count(&count).Error; err != nil {
		utils.Errorf("根据分类统计工作数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountJobsByType 根据类型统计工作数量
func CountJobsByType(tx *gorm.DB, jobType string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Job{}).Where("job_type = ? AND status = ?", jobType, "published").Count(&count).Error; err != nil {
		utils.Errorf("根据类型统计工作数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountSearchJobs 统计搜索结果数量
func CountSearchJobs(tx *gorm.DB, keyword string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Job{}).Where("status = ? AND (title LIKE ? OR description LIKE ? OR location LIKE ?)",
		"published", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Count(&count).Error; err != nil {
		utils.Errorf("统计搜索结果数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountSearchJobsWithFilters 统计带过滤条件的搜索结果数量
func CountSearchJobsWithFilters(tx *gorm.DB, keyword, location string, salaryMin, salaryMax float64) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	query := tx.Model(&models.Job{}).Where("status = ?", "published")

	if keyword != "" {
		query = query.Where("(title LIKE ? OR description LIKE ?)", "%"+keyword+"%", "%"+keyword+"%")
	}
	if location != "" {
		query = query.Where("location LIKE ?", "%"+location+"%")
	}
	if salaryMin > 0 {
		query = query.Where("salary >= ?", salaryMin)
	}
	if salaryMax > 0 {
		query = query.Where("salary <= ?", salaryMax)
	}

	if err := query.Count(&count).Error; err != nil {
		utils.Errorf("统计带过滤条件的搜索结果数量失败: %v", err)
		return 0, err
	}

	return count, nil
}
