package upload

import (
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/upload"
	"github.com/xinjiyuan97/labor-clients/config"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// GetSignedURLLogic 获取签名URL业务逻辑
func GetSignedURLLogic(fileURL string, expireSeconds int64, cfg *config.OSSConfig) (*upload.GetSignedURLResp, error) {
	// 验证文件URL
	if fileURL == "" {
		utils.Errorf("文件URL不能为空")
		return &upload.GetSignedURLResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "文件URL不能为空",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 设置默认过期时间为1小时
	if expireSeconds <= 0 {
		expireSeconds = 3600
	}

	// 最大过期时间为7天
	maxExpireSeconds := int64(7 * 24 * 3600)
	if expireSeconds > maxExpireSeconds {
		expireSeconds = maxExpireSeconds
	}

	// 获取上传服务
	uploadService, err := utils.GetUploadService(cfg)
	if err != nil {
		utils.Errorf("获取上传服务失败: %v", err)
		return &upload.GetSignedURLResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "上传服务初始化失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 生成签名URL
	signedURL, err := uploadService.GetSignedURL(fileURL, expireSeconds)
	if err != nil {
		utils.Errorf("生成签名URL失败: %v", err)
		return &upload.GetSignedURLResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "生成签名URL失败: " + err.Error(),
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 计算过期时间
	expireTime := time.Now().Add(time.Duration(expireSeconds) * time.Second)

	return &upload.GetSignedURLResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "获取签名URL成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		SignedURL:     signedURL,
		ExpireSeconds: expireSeconds,
		ExpireTime:    expireTime.Format(time.RFC3339),
	}, nil
}
