package mysql

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/models"
)

// CreateUserRole 创建用户角色关联
func CreateUserRole(ctx context.Context, userRole *models.UserRole) error {
	return DB.WithContext(ctx).Create(userRole).Error
}

// GetUserRolesByUserID 获取用户的所有角色
func GetUserRolesByUserID(ctx context.Context, userID int64) ([]*models.UserRole, error) {
	var roles []*models.UserRole
	err := DB.WithContext(ctx).Where("user_id = ? AND status = ?", userID, "active").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// GetUserRoleByID 根据ID获取用户角色
func GetUserRoleByID(ctx context.Context, roleID int64) (*models.UserRole, error) {
	var role models.UserRole
	err := DB.WithContext(ctx).Where("id = ?", roleID).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// UpdateUserRole 更新用户角色
func UpdateUserRole(ctx context.Context, roleID int64, updates map[string]interface{}) error {
	return DB.WithContext(ctx).Model(&models.UserRole{}).Where("id = ?", roleID).Updates(updates).Error
}

// DeleteUserRole 删除用户角色（软删除）
func DeleteUserRole(ctx context.Context, roleID int64) error {
	return DB.WithContext(ctx).Where("id = ?", roleID).Delete(&models.UserRole{}).Error
}

// CheckUserHasBrandAdminRole 检查用户是否为指定品牌的管理员
func CheckUserHasBrandAdminRole(ctx context.Context, userID, brandID int64) (bool, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&models.UserRole{}).
		Where("user_id = ? AND role_type = ? AND brand_id = ? AND status = ?",
			userID, "brand_admin", brandID, "active").
		Count(&count).Error
	return count > 0, err
}

// CheckUserHasStoreAdminRole 检查用户是否为指定门店的管理员
func CheckUserHasStoreAdminRole(ctx context.Context, userID, storeID int64) (bool, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&models.UserRole{}).
		Where("user_id = ? AND role_type = ? AND store_id = ? AND status = ?",
			userID, "store_admin", storeID, "active").
		Count(&count).Error
	return count > 0, err
}

// GetBrandAdminsByBrandID 获取品牌的所有管理员
func GetBrandAdminsByBrandID(ctx context.Context, brandID int64) ([]*models.UserRole, error) {
	var roles []*models.UserRole
	err := DB.WithContext(ctx).Where("brand_id = ? AND role_type = ? AND status = ?",
		brandID, "brand_admin", "active").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// GetStoreAdminsByStoreID 获取门店的所有管理员
func GetStoreAdminsByStoreID(ctx context.Context, storeID int64) ([]*models.UserRole, error) {
	var roles []*models.UserRole
	err := DB.WithContext(ctx).Where("store_id = ? AND role_type = ? AND status = ?",
		storeID, "store_admin", "active").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// DeleteUserRolesByUserID 删除用户的所有角色
func DeleteUserRolesByUserID(ctx context.Context, userID int64) error {
	return DB.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.UserRole{}).Error
}
