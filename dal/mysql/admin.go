package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Admin相关数据库操作

// GetAdmins 获取管理员列表
func GetAdmins(tx *gorm.DB, offset, limit int) ([]*models.User, error) {
	if tx == nil {
		tx = DB
	}

	var admins []*models.User
	if err := tx.Where("role = ?", "admin").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&admins).Error; err != nil {
		utils.Errorf("获取管理员列表失败: %v", err)
		return nil, err
	}

	return admins, nil
}

// GetAdminByID 根据ID获取管理员
func GetAdminByID(tx *gorm.DB, adminID int64) (*models.User, error) {
	if tx == nil {
		tx = DB
	}

	var admin models.User
	if err := tx.Where("id = ? AND role = ?", adminID, "admin").First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询管理员失败: %v", err)
		return nil, err
	}

	return &admin, nil
}

// UpdateAdmin 更新管理员信息
func UpdateAdmin(tx *gorm.DB, admin *models.User) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(admin).Where("id = ? AND role = ?", admin.ID, "admin").Updates(admin).Error; err != nil {
		utils.Errorf("更新管理员信息失败: %v", err)
		return err
	}

	return nil
}

// CountAdmins 统计管理员数量
func CountAdmins(tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.User{}).Where("role = ?", "admin").Count(&count).Error; err != nil {
		utils.Errorf("统计管理员数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetUsers 获取用户列表
func GetUsers(tx *gorm.DB, offset, limit int) ([]*models.User, error) {
	if tx == nil {
		tx = DB
	}

	var users []*models.User
	if err := tx.Where("role IN ?", []string{"worker", "employer"}).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&users).Error; err != nil {
		utils.Errorf("获取用户列表失败: %v", err)
		return nil, err
	}

	return users, nil
}

// GetUsersByRole 根据角色获取用户列表
func GetUsersByRole(tx *gorm.DB, role string, offset, limit int) ([]*models.User, error) {
	if tx == nil {
		tx = DB
	}

	var users []*models.User
	if err := tx.Where("role = ?", role).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&users).Error; err != nil {
		utils.Errorf("根据角色获取用户列表失败: %v", err)
		return nil, err
	}

	return users, nil
}

// ResetUserPassword 重置用户密码
func ResetUserPassword(tx *gorm.DB, userID int64, newPassword string) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.User{}).Where("id = ?", userID).Update("password", newPassword).Error; err != nil {
		utils.Errorf("重置用户密码失败: %v", err)
		return err
	}

	return nil
}

// CountUsers 统计用户数量
func CountUsers(tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.User{}).Where("role IN ?", []string{"worker", "employer"}).Count(&count).Error; err != nil {
		utils.Errorf("统计用户数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// CountUsersByRole 根据角色统计用户数量
func CountUsersByRole(tx *gorm.DB, role string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.User{}).Where("role = ?", role).Count(&count).Error; err != nil {
		utils.Errorf("根据角色统计用户数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(tx *gorm.DB, username string) (*models.User, error) {
	if tx == nil {
		tx = DB
	}

	var user models.User
	if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据用户名查询用户失败: %v", err)
		return nil, err
	}

	return &user, nil
}
