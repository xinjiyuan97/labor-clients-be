package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"github.com/xinjiyuan97/labor-clients/config"
)

var (
	// tosClientInstance TOS客户端实例（用于避免循环导入）
	tosClientInstance *tos.ClientV2
)

// SetTOSClient 设置TOS客户端实例
func SetTOSClient(client *tos.ClientV2) {
	tosClientInstance = client
}

// GetTOSClient 获取TOS客户端实例
func GetTOSClient() *tos.ClientV2 {
	return tosClientInstance
}

// UploadService 文件上传服务接口
type UploadService interface {
	// UploadFile 上传文件
	UploadFile(file multipart.File, header *multipart.FileHeader, uploadType string) (string, error)
	// DeleteFile 删除文件
	DeleteFile(fileURL string) error
	// GetSignedURL 获取文件的预签名URL
	GetSignedURL(fileURL string, expireSeconds int64) (string, error)
}

// TOSUploadService 火山引擎TOS上传服务实现
type TOSUploadService struct {
	cfg *config.OSSConfig
}

// NewTOSUploadService 创建TOS上传服务
func NewTOSUploadService(cfg *config.OSSConfig) *TOSUploadService {
	return &TOSUploadService{
		cfg: cfg,
	}
}

// UploadFile 上传文件到TOS
func (s *TOSUploadService) UploadFile(file multipart.File, header *multipart.FileHeader, uploadType string) (string, error) {
	// 获取TOS客户端（避免循环导入，直接通过包名访问）
	// 注意：这里需要在初始化时确保TOS客户端已经创建
	client := GetTOSClient()
	if client == nil {
		return "", fmt.Errorf("TOS client is not initialized")
	}

	// 检查文件大小
	if header.Size > s.cfg.MaxFileSize {
		return "", fmt.Errorf("文件大小超过限制: %d bytes, 最大允许: %d bytes", header.Size, s.cfg.MaxFileSize)
	}

	// 检查文件扩展名
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(header.Filename), "."))
	if !s.isAllowedExt(ext) {
		return "", fmt.Errorf("不支持的文件类型: %s", ext)
	}

	// 生成对象key
	objectKey := s.generateObjectKey(header.Filename, uploadType)

	// 读取文件内容
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("读取文件失败: %v", err)
	}

	// 上传到TOS
	input := &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{
			Bucket: s.cfg.Bucket,
			Key:    objectKey,
		},
		Content: bytes.NewReader(fileContent),
	}

	_, err = client.PutObjectV2(context.Background(), input)
	if err != nil {
		return "", fmt.Errorf("上传文件到TOS失败: %v", err)
	}

	// 构建文件URL
	fileURL := s.buildFileURL(objectKey)
	log.Printf("文件上传成功: %s, size: %d bytes", fileURL, header.Size)
	return fileURL, nil
}

// DeleteFile 从TOS删除文件
func (s *TOSUploadService) DeleteFile(fileURL string) error {
	client := GetTOSClient()
	if client == nil {
		return fmt.Errorf("TOS client is not initialized")
	}

	// 从URL中提取object key
	objectKey := s.extractObjectKey(fileURL)
	if objectKey == "" {
		return fmt.Errorf("无效的文件URL: %s", fileURL)
	}

	// 从TOS删除
	input := &tos.DeleteObjectV2Input{
		Bucket: s.cfg.Bucket,
		Key:    objectKey,
	}

	_, err := client.DeleteObjectV2(context.Background(), input)
	if err != nil {
		return fmt.Errorf("从TOS删除文件失败: %v", err)
	}

	return nil
}

// GetSignedURL 获取文件的预签名URL
func (s *TOSUploadService) GetSignedURL(fileURL string, expireSeconds int64) (string, error) {
	client := GetTOSClient()
	if client == nil {
		return "", fmt.Errorf("TOS client is not initialized")
	}

	// 从URL中提取object key
	objectKey := s.extractObjectKey(fileURL)
	if objectKey == "" {
		return "", fmt.Errorf("无效的文件URL: %s", fileURL)
	}

	// 默认过期时间为1小时
	if expireSeconds <= 0 {
		expireSeconds = 3600
	}

	// 生成预签名URL
	output, err := client.PreSignedURL(&tos.PreSignedURLInput{
		Bucket:  s.cfg.Bucket,
		Key:     objectKey,
		Expires: expireSeconds,
	})
	if err != nil {
		return "", fmt.Errorf("生成签名URL失败: %v", err)
	}

	log.Printf("生成签名URL成功: %s, 过期时间: %d秒", objectKey, expireSeconds)
	return output.SignedUrl, nil
}

// isAllowedExt 检查文件扩展名是否允许
func (s *TOSUploadService) isAllowedExt(ext string) bool {
	if len(s.cfg.AllowedExts) == 0 {
		return true // 如果没有配置，则允许所有类型
	}

	for _, allowedExt := range s.cfg.AllowedExts {
		if strings.EqualFold(ext, allowedExt) {
			return true
		}
	}
	return false
}

// generateObjectKey 生成对象key
func (s *TOSUploadService) generateObjectKey(filename string, uploadType string) string {
	ext := filepath.Ext(filename)
	timestamp := time.Now().Format("20060102150405")

	// 使用雪花算法生成唯一ID
	id := GenerateID()

	// 构建对象key: uploads/uploadType/20240122/timestamp_id.ext
	var objectKey string
	if uploadType != "" {
		objectKey = path.Join(s.cfg.UploadPath, uploadType, time.Now().Format("20060102"), fmt.Sprintf("%s_%d%s", timestamp, id, ext))
	} else {
		objectKey = path.Join(s.cfg.UploadPath, time.Now().Format("20060102"), fmt.Sprintf("%s_%d%s", timestamp, id, ext))
	}

	return objectKey
}

// buildFileURL 构建文件URL
func (s *TOSUploadService) buildFileURL(objectKey string) string {
	return strings.TrimSuffix(s.cfg.BaseURL, "/") + "/" + objectKey
}

// extractObjectKey 从URL中提取对象key
func (s *TOSUploadService) extractObjectKey(fileURL string) string {
	baseURL := strings.TrimSuffix(s.cfg.BaseURL, "/")
	if !strings.HasPrefix(fileURL, baseURL) {
		return ""
	}
	return strings.TrimPrefix(fileURL, baseURL+"/")
}

// LocalUploadService 本地上传服务实现（用于开发测试）
type LocalUploadService struct {
	cfg *config.OSSConfig
}

// NewLocalUploadService 创建本地上传服务
func NewLocalUploadService(cfg *config.OSSConfig) *LocalUploadService {
	return &LocalUploadService{
		cfg: cfg,
	}
}

// UploadFile 上传文件到本地
func (s *LocalUploadService) UploadFile(file multipart.File, header *multipart.FileHeader, uploadType string) (string, error) {
	// TODO: 实现本地文件上传逻辑
	// 这里简化处理，返回一个模拟的URL
	filename := header.Filename
	mockURL := fmt.Sprintf("%s/uploads/%s/%s", s.cfg.BaseURL, uploadType, filename)
	return mockURL, nil
}

// DeleteFile 从本地删除文件
func (s *LocalUploadService) DeleteFile(fileURL string) error {
	// TODO: 实现本地文件删除逻辑
	return nil
}

// GetSignedURL 获取文件的预签名URL（本地实现直接返回原URL）
func (s *LocalUploadService) GetSignedURL(fileURL string, expireSeconds int64) (string, error) {
	// 本地模式直接返回原URL
	return fileURL, nil
}

// GetUploadService 获取上传服务实例
func GetUploadService(cfg *config.OSSConfig) (UploadService, error) {
	switch cfg.Provider {
	case "volcengine":
		return NewTOSUploadService(cfg), nil
	case "local":
		return NewLocalUploadService(cfg), nil
	default:
		return nil, fmt.Errorf("不支持的存储提供商: %s", cfg.Provider)
	}
}
