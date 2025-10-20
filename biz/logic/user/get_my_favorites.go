package user

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/user"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetMyFavoritesLogic 获取我的收藏列表业务逻辑
func GetMyFavoritesLogic(req *user.GetMyFavoritesReq, userID int64) (*user.GetMyFavoritesResp, error) {
	// 设置默认分页参数
	page := 1
	limit := 10
	if req.PageReq != nil {
		if req.PageReq.Page > 0 {
			page = int(req.PageReq.Page)
		}
		if req.PageReq.Limit > 0 {
			limit = int(req.PageReq.Limit)
		}
	}

	offset := (page - 1) * limit

	// 获取收藏列表
	favorites, err := mysql.GetUserFavoriteJobs(nil, userID, offset, limit)
	if err != nil {
		utils.Errorf("获取收藏列表失败: %v", err)
		return &user.GetMyFavoritesResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取总数
	total, err := mysql.CountUserFavoriteJobs(nil, userID)
	if err != nil {
		utils.Errorf("获取收藏总数失败: %v", err)
		return &user.GetMyFavoritesResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 构建响应数据
	var favoriteInfos []*common.UserFavoriteJobInfo
	for _, favorite := range favorites {
		// 这里应该根据JobID获取工作详细信息
		// 由于没有Job模型，这里简化处理
		favoriteInfo := &common.UserFavoriteJobInfo{
			FavoriteID: favorite.ID,
			JobID:      favorite.JobID,
			UserID:     favorite.UserID,
			CreatedAt:  favorite.CreatedAt.Format(time.RFC3339),
			// 这里应该添加工作详细信息，如标题、薪资等
		}
		favoriteInfos = append(favoriteInfos, favoriteInfo)
	}

	// 构建分页响应
	pageResp := &common.PageResp{
		Page:  int32(page),
		Limit: int32(limit),
		Total: int32(total),
	}

	return &user.GetMyFavoritesResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取收藏列表成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		PageResp:  pageResp,
		Favorites: favoriteInfos,
	}, nil
}
