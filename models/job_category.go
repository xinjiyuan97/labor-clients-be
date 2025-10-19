package models

// JobCategory 岗位分类表
type JobCategory struct {
	BaseModel
	Name        string `json:"name" gorm:"column:name;type:varchar(50);not null;comment:分类名称"`
	Description string `json:"description" gorm:"column:description;type:text;comment:分类描述"`
	ParentID    int    `json:"parent_id" gorm:"column:parent_id;type:int;default:0;index;comment:父级分类ID"`
	SortOrder   int    `json:"sort_order" gorm:"column:sort_order;type:int;default:0;comment:排序"`
}

// TableName 指定表名
func (JobCategory) TableName() string {
	return "job_categories"
}
