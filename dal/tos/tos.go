package tos

import (
	"fmt"
	"log"

	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"github.com/xinjiyuan97/labor-clients/config"
)

var (
	// TOSClient TOS客户端实例
	TOSClient *tos.ClientV2
)

// Init 初始化TOS客户端
func Init(cfg *config.OSSConfig) error {
	if cfg.Provider != "volcengine" {
		log.Println("OSS provider is not volcengine, skipping TOS initialization")
		return nil
	}

	if cfg.AccessKey == "" || cfg.SecretKey == "" {
		return fmt.Errorf("TOS access_key or secret_key is empty")
	}

	if cfg.Region == "" {
		return fmt.Errorf("TOS region is empty")
	}

	if cfg.Bucket == "" {
		return fmt.Errorf("TOS bucket is empty")
	}

	// 创建TOS客户端
	client, err := tos.NewClientV2(cfg.Endpoint, tos.WithRegion(cfg.Region), tos.WithCredentials(tos.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey)))
	if err != nil {
		return fmt.Errorf("failed to create TOS client: %v", err)
	}

	TOSClient = client

	// 将客户端设置到utils包中（避免循环导入）
	// 这个调用在编译时会产生循环导入错误，需要在main.go中手动设置
	// utils.SetTOSClient(client)

	log.Printf("TOS client initialized successfully, endpoint: %s, region: %s, bucket: %s", cfg.Endpoint, cfg.Region, cfg.Bucket)
	return nil
}

// Close 关闭TOS客户端
func Close() {
	if TOSClient != nil {
		log.Println("TOS client closed")
		TOSClient = nil
	}
}

// GetClient 获取TOS客户端实例
func GetClient() *tos.ClientV2 {
	return TOSClient
}
