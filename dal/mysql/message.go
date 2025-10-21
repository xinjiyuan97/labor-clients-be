package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Message相关数据库操作

// CreateMessage 创建消息
func CreateMessage(tx *gorm.DB, message *models.Message) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(message).Error; err != nil {
		utils.Errorf("创建消息失败: %v", err)
		return err
	}

	return nil
}

// GetMessageByID 根据ID获取消息详情
func GetMessageByID(tx *gorm.DB, messageID int64) (*models.Message, error) {
	if tx == nil {
		tx = DB
	}

	var message models.Message
	if err := tx.Where("id = ?", messageID).First(&message).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询消息详情失败: %v", err)
		return nil, err
	}

	return &message, nil
}

// GetMessages 获取消息列表
func GetMessages(tx *gorm.DB, userID int64, msgCategory string, offset, limit int) ([]*models.Message, error) {
	if tx == nil {
		tx = DB
	}

	var messages []*models.Message
	query := tx.Where("to_user = ?", userID)

	if msgCategory != "" {
		query = query.Where("msg_category = ?", msgCategory)
	}

	if err := query.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&messages).Error; err != nil {
		utils.Errorf("获取消息列表失败: %v", err)
		return nil, err
	}

	return messages, nil
}

// CountMessages 统计消息数量
func CountMessages(tx *gorm.DB, userID int64, msgCategory string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	query := tx.Model(&models.Message{}).Where("to_user = ?", userID)

	if msgCategory != "" {
		query = query.Where("msg_category = ?", msgCategory)
	}

	if err := query.Count(&count).Error; err != nil {
		utils.Errorf("统计消息数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// MarkMessageRead 标记消息为已读
func MarkMessageRead(tx *gorm.DB, messageID int64) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.Message{}).
		Where("id = ?", messageID).
		Update("is_read", true).Error; err != nil {
		utils.Errorf("标记消息已读失败: %v", err)
		return err
	}

	return nil
}

// BatchMarkMessagesRead 批量标记消息为已读
func BatchMarkMessagesRead(tx *gorm.DB, messageIDs []int64) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.Message{}).
		Where("id IN ?", messageIDs).
		Update("is_read", true).Error; err != nil {
		utils.Errorf("批量标记消息已读失败: %v", err)
		return err
	}

	return nil
}

// GetUnreadCount 获取未读消息数量
func GetUnreadCount(tx *gorm.DB, userID int64, msgCategory string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	query := tx.Model(&models.Message{}).
		Where("to_user = ? AND is_read = ?", userID, false)

	if msgCategory != "" {
		query = query.Where("msg_category = ?", msgCategory)
	}

	if err := query.Count(&count).Error; err != nil {
		utils.Errorf("获取未读消息数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetMessagesByCategory 根据分类获取消息
func GetMessagesByCategory(tx *gorm.DB, userID int64, msgCategory string, offset, limit int) ([]*models.Message, error) {
	if tx == nil {
		tx = DB
	}

	var messages []*models.Message
	if err := tx.Where("to_user = ? AND msg_category = ?", userID, msgCategory).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&messages).Error; err != nil {
		utils.Errorf("根据分类获取消息失败: %v", err)
		return nil, err
	}

	return messages, nil
}

// CountMessagesByCategory 统计分类消息数量
func CountMessagesByCategory(tx *gorm.DB, userID int64, msgCategory string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Message{}).
		Where("to_user = ? AND msg_category = ?", userID, msgCategory).
		Count(&count).Error; err != nil {
		utils.Errorf("统计分类消息数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// GetUnreadCountByCategory 根据分类获取未读消息数量
func GetUnreadCountByCategory(tx *gorm.DB, userID int64, msgCategory string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.Message{}).
		Where("to_user = ? AND msg_category = ? AND is_read = ?", userID, msgCategory, false).
		Count(&count).Error; err != nil {
		utils.Errorf("根据分类获取未读消息数量失败: %v", err)
		return 0, err
	}

	return count, nil
}
