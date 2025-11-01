package dal

import (
	"github.com/xinjiyuan97/labor-clients/config"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/dal/redis"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// InitDAL 初始化所有数据访问层
func InitDAL(cfg *config.Config) error {
	// 初始化MySQL
	if cfg.Database != nil {
		if err := mysql.InitMySQL(cfg.Database); err != nil {
			utils.Errorf("初始化MySQL失败: %v", err)
			return err
		}
	}

	// 初始化Redis
	if cfg.Redis != nil {
		if err := redis.InitRedis(cfg.Redis); err != nil {
			utils.Errorf("初始化Redis失败: %v", err)
			return err
		}
	}

	utils.Info("数据访问层初始化成功")
	return nil
}

// CloseDAL 关闭所有数据访问层连接
func CloseDAL() error {
	var err error

	// 关闭Redis连接
	if redisErr := redis.CloseRedis(); redisErr != nil {
		utils.Errorf("关闭Redis连接失败: %v", redisErr)
		err = redisErr
	}

	// 关闭MySQL连接
	if mysqlErr := mysql.CloseMySQL(); mysqlErr != nil {
		utils.Errorf("关闭MySQL连接失败: %v", mysqlErr)
		err = mysqlErr
	}

	return err
}

// HealthCheck 健康检查
func HealthCheck() error {
	// 检查MySQL
	if err := mysql.HealthCheck(); err != nil {
		return err
	}

	// 检查Redis
	if err := redis.HealthCheck(); err != nil {
		return err
	}

	return nil
}

// GetStats 获取所有连接统计信息
func GetStats() map[string]interface{} {
	stats := make(map[string]interface{})

	// MySQL统计信息
	if mysqlStats := mysql.GetStats(); mysqlStats != nil {
		stats["mysql"] = mysqlStats
	}

	// Redis统计信息
	if redisStats := redis.GetStats(); redisStats != nil {
		stats["redis"] = redisStats
	}

	return stats
}
