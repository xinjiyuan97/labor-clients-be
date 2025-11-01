package mysql

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CreateSMSVerificationCode 创建短信验证码
func CreateSMSVerificationCode(ctx context.Context, phone, code string) error {
	tx := GetDB(ctx)
	
	// 设置过期时间为5分钟后
	expiresAt := time.Now().Add(5 * time.Minute)
	
	smsCode := &models.SMSVerificationCode{
		Phone:     phone,
		Code:      code,
		Status:    "unused",
		ExpiresAt: expiresAt,
	}
	
	if err := tx.Create(smsCode).Error; err != nil {
		utils.Errorf("创建短信验证码失败: %v", err)
		return err
	}

	utils.Infof("短信验证码创建成功, Phone: %s", phone)
	return nil
}

// GetSMSVerificationCode 获取短信验证码
func GetSMSVerificationCode(ctx context.Context, phone, code string) (*models.SMSVerificationCode, error) {
	tx := GetDB(ctx)

	var smsCode models.SMSVerificationCode
	if err := tx.Where("phone = ? AND code = ? AND status = ?", phone, code, "unused").First(&smsCode).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("查询短信验证码失败: %v", err)
		return nil, err
	}

	return &smsCode, nil
}

// MarkSMSVerificationCodeUsed 标记验证码为已使用
func MarkSMSVerificationCodeUsed(ctx context.Context, phone, code string) error {
	tx := GetDB(ctx)
	
	usedAt := time.Now()
	
	if err := tx.Model(&models.SMSVerificationCode{}).
		Where("phone = ? AND code = ?", phone, code).
		Updates(map[string]interface{}{
			"status":  "used",
			"used_at": usedAt,
		}).Error; err != nil {
		utils.Errorf("标记验证码已使用失败: %v", err)
		return err
	}

	return nil
}

// CleanExpiredSMSVerificationCodes 清理过期的验证码
func CleanExpiredSMSVerificationCodes(ctx context.Context) error {
	tx := GetDB(ctx)
	now := time.Now()
	
	if err := tx.Model(&models.SMSVerificationCode{}).
		Where("expires_at < ? AND status = ?", now, "unused").
		Update("status", "expired").Error; err != nil {
		utils.Errorf("清理过期验证码失败: %v", err)
		return err
	}

	return nil
}

// CheckRecentCodeExists 检查是否最近发送过验证码（防止频繁发送）
func CheckRecentCodeExists(ctx context.Context, phone string, minutes int) (bool, error) {
	tx := GetDB(ctx)
	
	// 检查指定分钟内是否有未使用的验证码
	since := time.Now().Add(-time.Duration(minutes) * time.Minute)
	
	var count int64
	if err := tx.Model(&models.SMSVerificationCode{}).
		Where("phone = ? AND status = ? AND created_at > ?", phone, "unused", since).
		Count(&count).Error; err != nil {
		utils.Errorf("检查最近验证码失败: %v", err)
		return false, err
	}

	return count > 0, nil
}

