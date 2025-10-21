package mysql

import (
	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// Community相关数据库操作

// CreateCommunityPost 创建帖子
func CreateCommunityPost(tx *gorm.DB, post *models.CommunityPost) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(post).Error; err != nil {
		utils.Errorf("创建帖子失败: %v", err)
		return err
	}

	return nil
}

// GetCommunityPostByID 根据ID获取帖子详情
func GetCommunityPostByID(tx *gorm.DB, postID int64) (*models.CommunityPost, error) {
	if tx == nil {
		tx = DB
	}

	var post models.CommunityPost
	if err := tx.Where("id = ? AND status != ?", postID, "deleted").First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询帖子详情失败: %v", err)
		return nil, err
	}

	return &post, nil
}

// GetCommunityPosts 获取帖子列表
func GetCommunityPosts(tx *gorm.DB, postType string, offset, limit int) ([]*models.CommunityPost, error) {
	if tx == nil {
		tx = DB
	}

	var posts []*models.CommunityPost
	query := tx.Where("status = ?", "published")

	if postType != "" {
		query = query.Where("post_type = ?", postType)
	}

	if err := query.Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&posts).Error; err != nil {
		utils.Errorf("获取帖子列表失败: %v", err)
		return nil, err
	}

	return posts, nil
}

// CountCommunityPosts 统计帖子数量
func CountCommunityPosts(tx *gorm.DB, postType string) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	query := tx.Model(&models.CommunityPost{}).Where("status = ?", "published")

	if postType != "" {
		query = query.Where("post_type = ?", postType)
	}

	if err := query.Count(&count).Error; err != nil {
		utils.Errorf("统计帖子数量失败: %v", err)
		return 0, err
	}

	return count, nil
}

// UpdateCommunityPost 更新帖子
func UpdateCommunityPost(tx *gorm.DB, post *models.CommunityPost) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Save(post).Error; err != nil {
		utils.Errorf("更新帖子失败: %v", err)
		return err
	}

	return nil
}

// DeleteCommunityPost 删除帖子（软删除）
func DeleteCommunityPost(tx *gorm.DB, postID int64) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.CommunityPost{}).
		Where("id = ?", postID).
		Update("status", "deleted").Error; err != nil {
		utils.Errorf("删除帖子失败: %v", err)
		return err
	}

	return nil
}

// IncrementPostViewCount 增加帖子浏览数
func IncrementPostViewCount(tx *gorm.DB, postID int64) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.CommunityPost{}).
		Where("id = ?", postID).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error; err != nil {
		utils.Errorf("增加帖子浏览数失败: %v", err)
		return err
	}

	return nil
}

// IncrementPostLikeCount 增加帖子点赞数
func IncrementPostLikeCount(tx *gorm.DB, postID int64) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.CommunityPost{}).
		Where("id = ?", postID).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		utils.Errorf("增加帖子点赞数失败: %v", err)
		return err
	}

	return nil
}

// DecrementPostLikeCount 减少帖子点赞数
func DecrementPostLikeCount(tx *gorm.DB, postID int64) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Model(&models.CommunityPost{}).
		Where("id = ?", postID).
		UpdateColumn("like_count", gorm.Expr("CASE WHEN like_count > 0 THEN like_count - 1 ELSE 0 END")).Error; err != nil {
		utils.Errorf("减少帖子点赞数失败: %v", err)
		return err
	}

	return nil
}

// CreatePostLike 创建点赞记录
func CreatePostLike(tx *gorm.DB, like *models.PostLike) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(like).Error; err != nil {
		utils.Errorf("创建点赞记录失败: %v", err)
		return err
	}

	return nil
}

// DeletePostLike 删除点赞记录
func DeletePostLike(tx *gorm.DB, postID, userID int64) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Where("post_id = ? AND user_id = ?", postID, userID).
		Delete(&models.PostLike{}).Error; err != nil {
		utils.Errorf("删除点赞记录失败: %v", err)
		return err
	}

	return nil
}

// CheckPostLike 检查是否已点赞
func CheckPostLike(tx *gorm.DB, postID, userID int64) (bool, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.PostLike{}).
		Where("post_id = ? AND user_id = ?", postID, userID).
		Count(&count).Error; err != nil {
		utils.Errorf("检查点赞记录失败: %v", err)
		return false, err
	}

	return count > 0, nil
}

// CreatePostComment 创建评论
func CreatePostComment(tx *gorm.DB, comment *models.PostComment) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(comment).Error; err != nil {
		utils.Errorf("创建评论失败: %v", err)
		return err
	}

	return nil
}

// GetPostComments 获取帖子评论列表
func GetPostComments(tx *gorm.DB, postID int64, offset, limit int) ([]*models.PostComment, error) {
	if tx == nil {
		tx = DB
	}

	var comments []*models.PostComment
	if err := tx.Where("post_id = ?", postID).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&comments).Error; err != nil {
		utils.Errorf("获取评论列表失败: %v", err)
		return nil, err
	}

	return comments, nil
}

// CountPostComments 统计评论数量
func CountPostComments(tx *gorm.DB, postID int64) (int64, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.PostComment{}).
		Where("post_id = ?", postID).
		Count(&count).Error; err != nil {
		utils.Errorf("统计评论数量失败: %v", err)
		return 0, err
	}

	return count, nil
}
