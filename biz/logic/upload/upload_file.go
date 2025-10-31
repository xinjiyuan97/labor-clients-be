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

// UploadFileLogic 上传文件业务逻辑
func UploadFileLogic(file multipart.File, header *multipart.FileHeader, uploadType string, cfg *config.OSSConfig) (*upload.UploadFileResp, error) {
	// 获取文件扩展名和MIME类型
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(header.Filename), "."))

	// 确定文件类型
	fileType := determineFileType(ext)

	// 获取上传服务
	uploadService, err := utils.GetUploadService(cfg)
	if err != nil {
		utils.Errorf("获取上传服务失败: %v", err)
		return &upload.UploadFileResp{
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
		utils.Errorf("上传文件失败: %v", err)
		return &upload.UploadFileResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "上传文件失败: " + err.Error(),
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 生成签名URL用于展示（默认7天有效期）
	displayURL, err := uploadService.GetSignedURL(fileURL, 7*24*3600)
	if err != nil {
		utils.Warnf("生成签名URL失败: %v, 使用原始URL", err)
		displayURL = fileURL
	}

	return &upload.UploadFileResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "文件上传成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		FileURL:    fileURL,    // 原始地址（用于存储）
		DisplayURL: displayURL, // 签名后的展示地址
		FileName:   header.Filename,
		FileSize:   header.Size,
		FileType:   fileType,
	}, nil
}

// determineFileType 根据文件扩展名确定文件类型
func determineFileType(ext string) string {
	switch ext {
	case "jpg", "jpeg", "png", "gif", "webp", "bmp", "svg":
		return "image"
	case "pdf":
		return "pdf"
	case "doc", "docx":
		return "word"
	case "xls", "xlsx":
		return "excel"
	case "ppt", "pptx":
		return "powerpoint"
	case "txt":
		return "text"
	case "zip", "rar", "7z":
		return "archive"
	case "mp4", "avi", "mov", "wmv":
		return "video"
	case "mp3", "wav", "flac":
		return "audio"
	default:
		return "other"
	}
}
