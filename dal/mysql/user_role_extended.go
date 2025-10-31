package mysql

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/models"
)

// BrandAdminInfo 品牌管理员信息（包含用户信息）
type BrandAdminInfo struct {
	UserID    int64  `json:"user_id"`
	RoleID    int64  `json:"role_id"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	RoleType  string `json:"role_type"`
	BrandID   *int64 `json:"brand_id"`
	BrandName string `json:"brand_name"`
	StoreID   *int64 `json:"store_id"`
	StoreName string `json:"store_name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// GetBrandAdminsList 获取品牌管理员列表（包含用户详细信息）
func GetBrandAdminsList(ctx context.Context, brandID int64, roleType, status string, offset, limit int) ([]*BrandAdminInfo, int64, error) {
	var admins []*BrandAdminInfo
	var total int64

	// 构建查询
	query := DB.WithContext(ctx).Table("user_roles ur").
		Select(`ur.id as role_id, ur.user_id, u.username, u.phone, ur.role_type, 
			ur.brand_id, b.name as brand_name, 
			ur.store_id, s.name as store_name, 
			ur.status, ur.created_at`).
		Joins("INNER JOIN users u ON ur.user_id = u.id").
		Joins("LEFT JOIN brands b ON ur.brand_id = b.id").
		Joins("LEFT JOIN stores s ON ur.store_id = s.id").
		Where("ur.brand_id = ? AND ur.deleted_at IS NULL", brandID)

	// 添加过滤条件
	if roleType != "" {
		query = query.Where("ur.role_type = ?", roleType)
	}
	if status != "" {
		query = query.Where("ur.status = ?", status)
	}

	// 计数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := query.Offset(offset).Limit(limit).Order("ur.created_at DESC").Scan(&admins).Error; err != nil {
		return nil, 0, err
	}

	return admins, total, nil
}

// GetStoreAdminsByStoreIDWithUserInfo 获取门店管理员列表（包含用户信息）
func GetStoreAdminsByStoreIDWithUserInfo(ctx context.Context, storeID int64) ([]*BrandAdminInfo, error) {
	var admins []*BrandAdminInfo

	err := DB.WithContext(ctx).Table("user_roles ur").
		Select(`ur.id as role_id, ur.user_id, u.username, u.phone, ur.role_type, 
			ur.brand_id, b.name as brand_name, 
			ur.store_id, s.name as store_name, 
			ur.status, ur.created_at`).
		Joins("INNER JOIN users u ON ur.user_id = u.id").
		Joins("LEFT JOIN brands b ON ur.brand_id = b.id").
		Joins("LEFT JOIN stores s ON ur.store_id = s.id").
		Where("ur.store_id = ? AND ur.role_type = ? AND ur.status = ? AND ur.deleted_at IS NULL",
			storeID, "store_admin", "active").
		Order("ur.created_at DESC").
		Scan(&admins).Error

	if err != nil {
		return nil, err
	}

	return admins, nil
}

// CheckUserRoleExists 检查用户角色是否已存在
func CheckUserRoleExists(ctx context.Context, userID, brandID int64, roleType string, storeID *int64) (bool, error) {
	var count int64
	query := DB.WithContext(ctx).Model(&models.UserRole{}).
		Where("user_id = ? AND brand_id = ? AND role_type = ?", userID, brandID, roleType)

	if storeID != nil {
		query = query.Where("store_id = ?", *storeID)
	} else {
		query = query.Where("store_id IS NULL")
	}

	err := query.Count(&count).Error
	return count > 0, err
}
