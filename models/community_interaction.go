package models

// PostLike 帖子点赞表
type PostLike struct {
	BaseModel
	PostID int64 `json:"post_id" gorm:"column:post_id;type:bigint;not null;index;comment:帖子ID"`
	UserID int64 `json:"user_id" gorm:"column:user_id;type:bigint;not null;index;comment:用户ID"`
}

// TableName 指定表名
func (PostLike) TableName() string {
	return "post_likes"
}

// PostComment 帖子评论表
type PostComment struct {
	BaseModel
	PostID   int64  `json:"post_id" gorm:"column:post_id;type:bigint;not null;index;comment:帖子ID"`
	UserID   int64  `json:"user_id" gorm:"column:user_id;type:bigint;not null;index;comment:用户ID"`
	Content  string `json:"content" gorm:"column:content;type:text;not null;comment:评论内容"`
	ParentID *int64 `json:"parent_id" gorm:"column:parent_id;type:bigint;index;comment:父评论ID"`
}

// TableName 指定表名
func (PostComment) TableName() string {
	return "post_comments"
}
