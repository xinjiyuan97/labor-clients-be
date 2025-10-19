package main

import (
	"fmt"
	"log"

	"github.com/xinjiyuan97/labor-clients/config"
	"github.com/xinjiyuan97/labor-clients/utils"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 验证配置
	if err := cfg.Validate(); err != nil {
		log.Fatalf("配置验证失败: %v", err)
	}

	// 初始化雪花算法
	if err := utils.InitSnowflake(cfg.Snowflake.NodeID); err != nil {
		log.Fatalf("初始化雪花算法失败: %v", err)
	}

	// 初始化日志
	if err := utils.InitLogger(&utils.LogConfig{
		Level:      cfg.Log.Level,
		Format:     cfg.Log.Format,
		Output:     cfg.Log.Output,
		FilePath:   cfg.Log.FilePath,
		MaxSize:    cfg.Log.MaxSize,
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,
		Compress:   cfg.Log.Compress,
	}); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}

	// 使用雪花算法生成ID
	id := utils.GenerateID()
	utils.Infof("生成的雪花ID: %d", id)

	// 使用JSON工具
	testData := map[string]interface{}{
		"id":     id,
		"name":   "测试用户",
		"email":  "test@example.com",
		"age":    25,
		"active": true,
	}

	jsonStr := utils.ToJSON(testData)
	utils.Infof("JSON字符串: %s", jsonStr)

	prettyJSON := utils.ToPrettyJSON(testData)
	utils.Infof("格式化JSON:\n%s", prettyJSON)

	// 使用日志记录配置信息
	utils.LogWithFields(map[string]interface{}{
		"server_host":   cfg.Server.Host,
		"server_port":   cfg.Server.Port,
		"database_host": cfg.Database.Host,
		"redis_host":    cfg.Redis.Host,
		"log_level":     cfg.Log.Level,
	}).Info("应用配置信息")

	// 测试不同级别的日志
	utils.Debug("这是一条调试日志")
	utils.Info("这是一条信息日志")
	utils.Warn("这是一条警告日志")
	utils.Error("这是一条错误日志")

	fmt.Println("示例程序运行完成！")
}
