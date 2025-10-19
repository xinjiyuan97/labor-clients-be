package models

// CommunityPost 社区帖子表
type CommunityPost struct {
	BaseModel
	AuthorID  int64  `json:"author_id" gorm:"column:author_id;type:bigint;not null;index;comment:作者用户ID"`
	Title     string `json:"title" gorm:"column:title;type:varchar(200);not null;comment:帖子标题"`
	Content   string `json:"content" gorm:"column:content;type:text;not null;comment:帖子内容"`
	PostType  string `json:"post_type" gorm:"column:post_type;type:varchar(20);default:discussion;index;comment:帖子类型"`
	ViewCount int    `json:"view_count" gorm:"column:view_count;type:int;default:0;comment:浏览数"`
	LikeCount int    `json:"like_count" gorm:"column:like_count;type:int;default:0;comment:点赞数"`
	Status    string `json:"status" gorm:"column:status;type:enum('draft','published','deleted');default:published;comment:状态"`
}

// TableName 指定表名
func (CommunityPost) TableName() string {
	return "community_posts"
}
