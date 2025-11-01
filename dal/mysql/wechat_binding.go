package mysql

import (
	"context"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CreateWeChatBinding 创建微信绑定
func CreateWeChatBinding(ctx context.Context, binding *models.WeChatBinding) error {
	tx := GetDB(ctx)
	if err := tx.Create(binding).Error; err != nil {
		utils.Errorf("创建微信绑定失败: %v", err)
		return err
	}

	utils.Infof("微信绑定创建成功, UserID: %d, OpenID: %s", binding.UserID, binding.OpenID)
	return nil
}

// GetWeChatBindingByOpenID 根据OpenID获取绑定信息
func GetWeChatBindingByOpenID(ctx context.Context, openid string) (*models.WeChatBinding, error) {
	tx := GetDB(ctx)

	var binding models.WeChatBinding
	if err := tx.Where("openid = ? AND status = ?", openid, "active").First(&binding).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据OpenID查询微信绑定失败: %v", err)
		return nil, err
	}

	return &binding, nil
}

// GetWeChatBindingByUserID 根据用户ID获取绑定信息
func GetWeChatBindingByUserID(ctx context.Context, userID int64) (*models.WeChatBinding, error) {
	tx := GetDB(ctx)

	var binding models.WeChatBinding
	if err := tx.Where("user_id = ? AND status = ?", userID, "active").First(&binding).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据用户ID查询微信绑定失败: %v", err)
		return nil, err
	}

	return &binding, nil
}

// UpdateWeChatBinding 更新微信绑定信息
func UpdateWeChatBinding(ctx context.Context, bindingID int64, updates map[string]interface{}) error {
	tx := GetDB(ctx)
	if err := tx.Model(&models.WeChatBinding{}).Where("id = ?", bindingID).Updates(updates).Error; err != nil {
		utils.Errorf("更新微信绑定失败: %v", err)
		return err
	}

	return nil
}

// CheckWeChatBindingExists 检查是否已存在绑定
func CheckWeChatBindingExists(ctx context.Context, openid string, userID int64) (bool, error) {
	tx := GetDB(ctx)

	var count int64
	if err := tx.Model(&models.WeChatBinding{}).Where("openid = ? OR user_id = ?", openid, userID).Count(&count).Error; err != nil {
		utils.Errorf("检查微信绑定是否存在失败: %v", err)
		return false, err
	}

	return count > 0, nil
}

