package mysql

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// System相关数据库操作

// GetSystemConfig 获取系统配置
func GetSystemConfig(tx *gorm.DB) (map[string]string, error) {
	if tx == nil {
		tx = DB
	}

	// 这里可以从数据库的配置表中获取配置
	// 暂时返回默认配置
	config := map[string]string{
		"app_name":        "零工客户端",
		"app_version":     "1.0.0",
		"min_version":     "1.0.0",
		"force_update":    "false",
		"maintenance":     "false",
		"maintenance_msg": "系统维护中，请稍后再试",
		"contact_phone":   "400-123-4567",
		"contact_email":   "support@example.com",
		"privacy_policy":  "https://example.com/privacy",
		"terms_service":   "https://example.com/terms",
	}

	return config, nil
}

// CreateFeedback 创建反馈
func CreateFeedback(tx *gorm.DB, feedback *models.Feedback) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(feedback).Error; err != nil {
		utils.Errorf("创建反馈失败: %v", err)
		return err
	}

	return nil
}

// GetNotices 获取通知列表
func GetNotices(tx *gorm.DB, offset, limit int) ([]*models.Notice, error) {
	if tx == nil {
		tx = DB
	}

	var notices []*models.Notice
	if err := tx.Where("status = ?", "published").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&notices).Error; err != nil {
		utils.Errorf("获取通知列表失败: %v", err)
		return nil, err
	}

	return notices, nil
}

// GetNoticeByID 根据ID获取通知详情
func GetNoticeByID(tx *gorm.DB, noticeID int64) (*models.Notice, error) {
	if tx == nil {
		tx = DB
	}

	var notice models.Notice
	if err := tx.Where("id = ? AND status = ?", noticeID, "published").First(&notice).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询通知详情失败: %v", err)
		return nil, err
	}

	return &notice, nil
}

// CountNotices 统计通知数量
func CountNotices(tx *gorm.DB) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Notice{}).Where("status = ?", "published").Count(&count).Error; err != nil {
		utils.Errorf("统计通知数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetVersionInfo 获取版本信息
func GetVersionInfo(tx *gorm.DB) (*models.VersionInfo, error) {
	if tx == nil {
		tx = DB
	}

	// 这里可以从数据库的版本表中获取版本信息
	// 暂时返回默认版本信息
	versionInfo := &models.VersionInfo{
		Version:     "1.0.0",
		BuildNumber: "100",
		MinVersion:  "1.0.0",
		ForceUpdate: false,
		UpdateURL:   "https://example.com/download",
		UpdateNote:  "修复了一些已知问题，提升了应用稳定性",
		CreatedAt:   time.Now(),
	}

	return versionInfo, nil
}
