package community

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/community"
)

// GetCoursesLogic 获取课程列表业务逻辑
func GetCoursesLogic(req *community.GetCoursesReq) (*community.GetCoursesResp, error) {
	// 这里返回模拟数据，实际应该从数据库获取
	// 可以使用类似帖子列表的方式实现
	var courseInfos []*common.CommunityPostInfo

	// 模拟一些课程数据
	courseInfos = append(courseInfos, &common.CommunityPostInfo{
		PostID:    1,
		Title:     "零工技能培训基础课程",
		Content:   "学习基本的零工技能和注意事项",
		AuthorID:  0,
		PostType:  "course",
		ViewCount: 1000,
		LikeCount: 100,
		CreatedAt: time.Now().Format(time.RFC3339),
	})

	courseInfos = append(courseInfos, &common.CommunityPostInfo{
		PostID:    2,
		Title:     "职业安全培训",
		Content:   "了解工作场所的安全知识和防护措施",
		AuthorID:  0,
		PostType:  "course",
		ViewCount: 800,
		LikeCount: 80,
		CreatedAt: time.Now().Format(time.RFC3339),
	})

	return &community.GetCoursesResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取课程列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		Courses: courseInfos,
	}, nil
}
