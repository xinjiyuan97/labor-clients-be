package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Brand相关数据库操作

// GetBrands 获取品牌列表
func GetBrands(tx *gorm.DB, offset, limit int) ([]*models.Brand, error) {
	if tx == nil {
		tx = DB
	}

	var brands []*models.Brand
	if err := tx.Where("auth_status = ?", "approved").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&brands).Error; err != nil {
		utils.Errorf("获取品牌列表失败: %v", err)
		return nil, err
	}

	return brands, nil
}

// GetBrandsByName 根据名称搜索品牌
func GetBrandsByName(tx *gorm.DB, name string, offset, limit int) ([]*models.Brand, error) {
	if tx == nil {
		tx = DB
	}

	var brands []*models.Brand
	if err := tx.Where("auth_status = ? AND name LIKE ?", "approved", "%"+name+"%").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&brands).Error; err != nil {
		utils.Errorf("根据名称搜索品牌失败: %v", err)
		return nil, err
	}

	return brands, nil
}

// GetBrandsByAuthStatus 根据认证状态获取品牌列表
func GetBrandsByAuthStatus(tx *gorm.DB, authStatus string, offset, limit int) ([]*models.Brand, error) {
	if tx == nil {
		tx = DB
	}

	var brands []*models.Brand
	if err := tx.Where("auth_status = ?", authStatus).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&brands).Error; err != nil {
		utils.Errorf("根据认证状态获取品牌列表失败: %v", err)
		return nil, err
	}

	return brands, nil
}

// GetBrandByID 根据ID获取品牌详情
func GetBrandByID(tx *gorm.DB, brandID int64) (*models.Brand, error) {
	if tx == nil {
		tx = DB
	}

	var brand models.Brand
	if err := tx.Where("id = ?", brandID).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询品牌详情失败: %v", err)
		return nil, err
	}

	return &brand, nil
}

// CountBrands 统计品牌数量
func CountBrands(tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Brand{}).Where("auth_status = ?", "approved").Count(&count).Error; err != nil {
		utils.Errorf("统计品牌数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountBrandsByName 根据名称统计品牌数量
func CountBrandsByName(tx *gorm.DB, name string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Brand{}).Where("auth_status = ? AND name LIKE ?", "approved", "%"+name+"%").Count(&count).Error; err != nil {
		utils.Errorf("根据名称统计品牌数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountBrandsByAuthStatus 根据认证状态统计品牌数量
func CountBrandsByAuthStatus(tx *gorm.DB, authStatus string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Brand{}).Where("auth_status = ?", authStatus).Count(&count).Error; err != nil {
		utils.Errorf("根据认证状态统计品牌数量失败: %v", err)
		return 0, err
	}

	return count, nil
}
