package mysql

import (
	"context"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CreateThirdPartyBinding 创建第三方绑定
func CreateThirdPartyBinding(ctx context.Context, binding *models.ThirdPartyBinding) error {
	tx := GetDB(ctx)
	if err := tx.Create(binding).Error; err != nil {
		utils.Errorf("创建第三方绑定失败: %v", err)
		return err
	}

	utils.Infof("第三方绑定创建成功, UserID: %d, Platform: %s, OpenID: %s", binding.UserID, binding.Platform, binding.OpenID)
	return nil
}

// GetThirdPartyBindingByPlatformAndOpenID 根据平台和OpenID获取绑定信息
func GetThirdPartyBindingByPlatformAndOpenID(ctx context.Context, platform, openid string) (*models.ThirdPartyBinding, error) {
	tx := GetDB(ctx)

	var binding models.ThirdPartyBinding
	if err := tx.Where("platform = ? AND openid = ? AND status = ?", platform, openid, "active").First(&binding).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据平台和OpenID查询第三方绑定失败: %v", err)
		return nil, err
	}

	return &binding, nil
}

// GetThirdPartyBindingByUserID 根据用户ID获取绑定信息（可指定平台）
func GetThirdPartyBindingByUserID(ctx context.Context, userID int64, platform string) (*models.ThirdPartyBinding, error) {
	tx := GetDB(ctx)

	var binding models.ThirdPartyBinding
	query := tx.Where("user_id = ? AND status = ?", userID, "active")
	if platform != "" {
		query = query.Where("platform = ?", platform)
	}
	
	if err := query.First(&binding).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据用户ID查询第三方绑定失败: %v", err)
		return nil, err
	}

	return &binding, nil
}

// UpdateThirdPartyBinding 更新第三方绑定信息
func UpdateThirdPartyBinding(ctx context.Context, bindingID int64, updates map[string]interface{}) error {
	tx := GetDB(ctx)
	if err := tx.Model(&models.ThirdPartyBinding{}).Where("id = ?", bindingID).Updates(updates).Error; err != nil {
		utils.Errorf("更新第三方绑定失败: %v", err)
		return err
	}

	return nil
}

// CheckThirdPartyBindingExists 检查是否已存在绑定
func CheckThirdPartyBindingExists(ctx context.Context, platform, openid string, userID int64) (bool, error) {
	tx := GetDB(ctx)

	var count int64
	if err := tx.Model(&models.ThirdPartyBinding{}).Where("platform = ? AND openid = ? OR user_id = ?", platform, openid, userID).Count(&count).Error; err != nil {
		utils.Errorf("检查第三方绑定是否存在失败: %v", err)
		return false, err
	}

	return count > 0, nil
}
