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

// UploadCertFileLogic 上传认证文件业务逻辑
func UploadCertFileLogic(file multipart.File, header *multipart.FileHeader, certType string, cfg *config.OSSConfig) (*upload.UploadCertFileResp, error) {
	// 验证证书类型
	validCertTypes := []string{"id_card", "passport", "driver_license", "business_license", "qualification_cert", "health_cert", "other"}
	isValid := false
	for _, validType := range validCertTypes {
		if certType == validType {
			isValid = true
			break
		}
	}

	if !isValid {
		utils.Errorf("不支持的证书类型: %s", certType)
		return &upload.UploadCertFileResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "不支持的证书类型，支持的类型: id_card, passport, driver_license, business_license, qualification_cert, health_cert, other",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 检查文件类型（证书文件通常是图片或PDF）
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(header.Filename), "."))
	allowedExts := []string{"jpg", "jpeg", "png", "pdf"}

	isAllowed := false
	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		utils.Errorf("不支持的证书文件格式: %s", ext)
		return &upload.UploadCertFileResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "不支持的证书文件格式，仅支持: jpg, jpeg, png, pdf",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 获取上传服务
	uploadService, err := utils.GetUploadService(cfg)
	if err != nil {
		utils.Errorf("获取上传服务失败: %v", err)
		return &upload.UploadCertFileResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "上传服务初始化失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	// 上传文件，使用证书类型作为uploadType
	uploadType := "cert/" + certType
	fileURL, err := uploadService.UploadFile(file, header, uploadType)
	if err != nil {
		utils.Errorf("上传证书文件失败: %v", err)
		return &upload.UploadCertFileResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "上传证书文件失败: " + err.Error(),
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

	return &upload.UploadCertFileResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "证书文件上传成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		FileURL:    fileURL,    // 原始地址（用于存储）
		DisplayURL: displayURL, // 签名后的展示地址
		CertType:   certType,
		FileName:   header.Filename,
		FileSize:   header.Size,
	}, nil
}
