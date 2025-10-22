package upload

import (
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/upload"
	"github.com/xinjiyuan97/labor-clients/config"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UploadImageLogic 上传图片业务逻辑
func UploadImageLogic(file multipart.File, header *multipart.FileHeader, uploadType string, cfg *config.OSSConfig) (*upload.UploadImageResp, error) {
	// 检查文件类型
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(header.Filename), "."))
	allowedImageExts := []string{"jpg", "jpeg", "png", "gif", "webp", "bmp", "svg"}

	isImage := false
	for _, allowedExt := range allowedImageExts {
		if ext == allowedExt {
			isImage = true
			break
		}
	}

	if !isImage {
		utils.Errorf("不支持的图片格式: %s", ext)
		return &upload.UploadImageResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "不支持的图片格式，仅支持: jpg, jpeg, png, gif, webp, bmp, svg",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取上传服务
	uploadService, err := utils.GetUploadService(cfg)
	if err != nil {
		utils.Errorf("获取上传服务失败: %v", err)
		return &upload.UploadImageResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "上传服务初始化失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 上传文件
	fileURL, err := uploadService.UploadFile(file, header, uploadType)
	if err != nil {
		utils.Errorf("上传图片失败: %v", err)
		return &upload.UploadImageResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "上传图片失败: " + err.Error(),
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &upload.UploadImageResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "图片上传成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		ImageURL: fileURL,
		FileName: header.Filename,
		FileSize: header.Size,
	}, nil
}
