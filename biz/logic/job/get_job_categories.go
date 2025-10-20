package job

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/job"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetJobCategoriesLogic 获取工作分类业务逻辑
func GetJobCategoriesLogic(req *job.GetJobCategoriesReq) (*job.GetJobCategoriesResp, error) {
	// 获取工作分类
	var categories []*models.JobCategory
	var err error

	if req.ParentID > 0 {
		// 根据父级ID获取子分类
		categories, err = mysql.GetJobCategoriesByParent(nil, int(req.ParentID))
		if err != nil {
			utils.Errorf("根据父级ID获取工作分类失败: %v", err)
			return &job.GetJobCategoriesResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
	} else {
		// 获取顶级分类
		categories, err = mysql.GetTopLevelJobCategories(nil)
		if err != nil {
			utils.Errorf("获取顶级工作分类失败: %v", err)
			return &job.GetJobCategoriesResp{
				Base: &common.BaseResp{
					Code:      500,
					Message:   "系统错误",
					Timestamp: time.Now().Format(time.RFC3339),
				},
			}, nil
		}
	}

	// 构建分类信息
	var categoryInfos []*common.JobCategoryInfo
	for _, category := range categories {
		categoryInfo := &common.JobCategoryInfo{
			CategoryID:  int32(category.ID),
			Name:        category.Name,
			Description: category.Description,
			ParentID:    int32(category.ParentID),
			SortOrder:   int32(category.SortOrder),
		}
		categoryInfos = append(categoryInfos, categoryInfo)
	}

	return &job.GetJobCategoriesResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取工作分类成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Categories: categoryInfos,
	}, nil
}
