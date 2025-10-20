package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// JobCategory相关数据库操作

// GetJobCategories 获取工作分类列表
func GetJobCategories(tx *gorm.DB) ([]*models.JobCategory, error) {
	if tx == nil {
		tx = DB
	}

	var categories []*models.JobCategory
	if err := tx.Order("sort_order ASC, created_at ASC").Find(&categories).Error; err != nil {
		utils.Errorf("获取工作分类列表失败: %v", err)
		return nil, err
	}

	return categories, nil
}

// GetJobCategoriesByParent 根据父级ID获取工作分类列表
func GetJobCategoriesByParent(tx *gorm.DB, parentID int) ([]*models.JobCategory, error) {
	if tx == nil {
		tx = DB
	}

	var categories []*models.JobCategory
	if err := tx.Where("parent_id = ?", parentID).
		Order("sort_order ASC, created_at ASC").
		Find(&categories).Error; err != nil {
		utils.Errorf("根据父级ID获取工作分类列表失败: %v", err)
		return nil, err
	}

	return categories, nil
}

// GetJobCategoryByID 根据ID获取工作分类详情
func GetJobCategoryByID(tx *gorm.DB, categoryID int64) (*models.JobCategory, error) {
	if tx == nil {
		tx = DB
	}

	var category models.JobCategory
	if err := tx.Where("id = ?", categoryID).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询工作分类详情失败: %v", err)
		return nil, err
	}

	return &category, nil
}

// GetTopLevelJobCategories 获取顶级工作分类列表
func GetTopLevelJobCategories(tx *gorm.DB) ([]*models.JobCategory, error) {
	if tx == nil {
		tx = DB
	}

	var categories []*models.JobCategory
	if err := tx.Where("parent_id = 0").
		Order("sort_order ASC, created_at ASC").
		Find(&categories).Error; err != nil {
		utils.Errorf("获取顶级工作分类列表失败: %v", err)
		return nil, err
	}

	return categories, nil
}

// CountJobCategories 统计工作分类数量
func CountJobCategories(tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.JobCategory{}).Count(&count).Error; err != nil {
		utils.Errorf("统计工作分类数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountJobCategoriesByParent 根据父级ID统计工作分类数量
func CountJobCategoriesByParent(tx *gorm.DB, parentID int) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.JobCategory{}).Where("parent_id = ?", parentID).Count(&count).Error; err != nil {
		utils.Errorf("根据父级ID统计工作分类数量失败: %v", err)
		return 0, err
	}

	return count, nil
}
