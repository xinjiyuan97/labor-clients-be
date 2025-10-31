package mysql

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/models"
)

// CreateStore 创建门店
func CreateStore(ctx context.Context, store *models.Store) error {
	return DB.WithContext(ctx).Create(store).Error
}

// GetStoreByID 根据ID获取门店信息
func GetStoreByID(ctx context.Context, storeID int64) (*models.Store, error) {
	var store models.Store
	err := DB.WithContext(ctx).Where("id = ?", storeID).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

// GetStoreList 获取门店列表
func GetStoreList(ctx context.Context, brandID *int64, status, name string, offset, limit int) ([]*models.Store, int64, error) {
	var stores []*models.Store
	var total int64

	query := DB.WithContext(ctx).Model(&models.Store{})

	if brandID != nil && *brandID > 0 {
		query = query.Where("brand_id = ?", *brandID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&stores).Error; err != nil {
		return nil, 0, err
	}

	return stores, total, nil
}

// UpdateStore 更新门店信息
func UpdateStore(ctx context.Context, storeID int64, updates map[string]interface{}) error {
	return DB.WithContext(ctx).Model(&models.Store{}).Where("id = ?", storeID).Updates(updates).Error
}

// DeleteStore 删除门店（软删除）
func DeleteStore(ctx context.Context, storeID int64) error {
	return DB.WithContext(ctx).Where("id = ?", storeID).Delete(&models.Store{}).Error
}

// GetStoresByBrandID 根据品牌ID获取所有门店
func GetStoresByBrandID(ctx context.Context, brandID int64) ([]*models.Store, error) {
	var stores []*models.Store
	err := DB.WithContext(ctx).Where("brand_id = ? AND status = ?", brandID, "active").Find(&stores).Error
	if err != nil {
		return nil, err
	}
	return stores, nil
}

// CheckStoreExists 检查门店是否存在
func CheckStoreExists(ctx context.Context, storeID int64) (bool, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&models.Store{}).Where("id = ?", storeID).Count(&count).Error
	return count > 0, err
}

// CheckStoreBelongsToBrand 检查门店是否属于指定品牌
func CheckStoreBelongsToBrand(ctx context.Context, storeID, brandID int64) (bool, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&models.Store{}).Where("id = ? AND brand_id = ?", storeID, brandID).Count(&count).Error
	return count > 0, err
}
