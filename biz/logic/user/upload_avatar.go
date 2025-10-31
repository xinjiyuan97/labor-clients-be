package user

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/user"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UploadAvatarLogic 上传头像业务逻辑
func UploadAvatarLogic(ctx context.Context, req *user.UploadAvatarReq, userID int64) (*user.UploadAvatarResp, error) {
	// 这里应该实现文件上传逻辑，将文件保存到文件服务器或云存储
	// 目前简化处理，直接返回一个模拟的URL
	avatarURL := "https://example.com/avatars/" + req.AvatarFile

	// 使用事务更新用户头像
	err := mysql.Transaction(ctx, func(tx *gorm.DB) error {
		// 获取当前用户信息
		currentUser, err := mysql.GetUserByID(tx, userID)
		if err != nil {
			return err
		}
		if currentUser == nil {
			return errors.New("用户不存在")
		}

		// 更新头像URL
		currentUser.Avatar = avatarURL
		return mysql.UpdateUserProfile(tx, currentUser)
	})

	if err != nil {
		utils.Errorf("上传头像失败: %v", err)
		return &user.UploadAvatarResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &user.UploadAvatarResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "头像上传成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		AvatarURL: avatarURL,
	}, nil
}
