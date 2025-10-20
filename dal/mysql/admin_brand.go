package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Brand管理相关数据库操作

// CreateBrand 创建品牌
func CreateBrand(tx *gorm.DB, brand *models.Brand) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(brand).Error; err != nil {
		utils.Errorf("创建品牌失败: %v", err)
		return err
	}

	return nil
}

// UpdateBrand 更新品牌信息
func UpdateBrand(tx *gorm.DB, brand *models.Brand) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(brand).Where("id = ?", brand.ID).Updates(brand).Error; err != nil {
		utils.Errorf("更新品牌信息失败: %v", err)
		return err
	}

	return nil
}

// ReviewBrand 审核品牌
func ReviewBrand(tx *gorm.DB, brandID int64, authStatus string) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.Brand{}).Where("id = ?", brandID).Update("auth_status", authStatus).Error; err != nil {
		utils.Errorf("审核品牌失败: %v", err)
		return err
	}

	return nil
}

// BatchCreateBrands 批量创建品牌
func BatchCreateBrands(tx *gorm.DB, brands []*models.Brand) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.CreateInBatches(brands, 100).Error; err != nil {
		utils.Errorf("批量创建品牌失败: %v", err)
		return err
	}

	return nil
}

// GetBrandsForAdmin 获取品牌列表（管理员用）
func GetBrandsForAdmin(tx *gorm.DB, offset, limit int) ([]*models.Brand, error) {
	if tx == nil {
		tx = DB
	}

	var brands []*models.Brand
	if err := tx.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&brands).Error; err != nil {
		utils.Errorf("获取品牌列表失败: %v", err)
		return nil, err
	}

	return brands, nil
}

// CountBrandsForAdmin 统计品牌数量（管理员用）
func CountBrandsForAdmin(tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Brand{}).Count(&count).Error; err != nil {
		utils.Errorf("统计品牌数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetBrandStatistics 获取品牌统计信息
func GetBrandStatistics(tx *gorm.DB) (map[string]int64, error) {
	if tx == nil {
		tx = DB
	}

	stats := make(map[string]int64)

	// 总品牌数
	var total int64
	if err := tx.Model(&models.Brand{}).Count(&total).Error; err != nil {
		utils.Errorf("统计总品牌数失败: %v", err)
		return nil, err
	}
	stats["total"] = total

	// 待审核品牌数
	var pending int64
	if err := tx.Model(&models.Brand{}).Where("auth_status = ?", "pending").Count(&pending).Error; err != nil {
		utils.Errorf("统计待审核品牌数失败: %v", err)
		return nil, err
	}
	stats["pending"] = pending

	// 已通过品牌数
	var approved int64
	if err := tx.Model(&models.Brand{}).Where("auth_status = ?", "approved").Count(&approved).Error; err != nil {
		utils.Errorf("统计已通过品牌数失败: %v", err)
		return nil, err
	}
	stats["approved"] = approved

	// 已拒绝品牌数
	var rejected int64
	if err := tx.Model(&models.Brand{}).Where("auth_status = ?", "rejected").Count(&rejected).Error; err != nil {
		utils.Errorf("统计已拒绝品牌数失败: %v", err)
		return nil, err
	}
	stats["rejected"] = rejected

	return stats, nil
}
