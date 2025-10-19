package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/xinjiyuan97/labor-clients/config"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

var (
	DB *gorm.DB
)

// InitMySQL 初始化MySQL数据库连接
func InitMySQL(cfg *config.DatabaseConfig) error {
	// 构建DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Charset)

	utils.Infof("MySQL连接信息: %+v", dsn)
	// 设置日志级别
	var logLevel logger.LogLevel
	switch cfg.LogLevel {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Warn
	}

	// 配置GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	// 连接数据库
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return fmt.Errorf("连接MySQL数据库失败: %v", err)
	}

	// 获取底层sql.DB对象进行连接池配置
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取数据库连接失败: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	// 测试连接
	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %v", err)
	}

	utils.Info("MySQL数据库连接成功")
	return nil
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}

// CloseMySQL 关闭MySQL连接
func CloseMySQL() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(models ...interface{}) error {
	if DB == nil {
		return fmt.Errorf("数据库连接未初始化")
	}

	err := DB.AutoMigrate(models...)
	if err != nil {
		utils.Errorf("数据库自动迁移失败: %v", err)
		return err
	}

	utils.Info("数据库自动迁移成功")
	return nil
}

// Transaction 执行数据库事务
func Transaction(fn func(*gorm.DB) error) error {
	if DB == nil {
		return fmt.Errorf("数据库连接未初始化")
	}

	return DB.Transaction(fn)
}

// HealthCheck 数据库健康检查
func HealthCheck() error {
	if DB == nil {
		return fmt.Errorf("数据库连接未初始化")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

// GetStats 获取数据库连接池统计信息
func GetStats() map[string]interface{} {
	if DB == nil {
		return nil
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil
	}

	stats := sqlDB.Stats()
	return map[string]interface{}{
		"max_open_connections": stats.MaxOpenConnections,
		"open_connections":     stats.OpenConnections,
		"in_use":               stats.InUse,
		"idle":                 stats.Idle,
		"wait_count":           stats.WaitCount,
		"wait_duration":        stats.WaitDuration.String(),
		"max_idle_closed":      stats.MaxIdleClosed,
		"max_idle_time_closed": stats.MaxIdleTimeClosed,
		"max_lifetime_closed":  stats.MaxLifetimeClosed,
	}
}

// User相关数据库操作

// CreateUser 创建用户
func CreateUser(tx *gorm.DB, user *models.User) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Create(user).Error; err != nil {
		utils.Errorf("创建用户失败: %v", err)
		return err
	}

	utils.Infof("用户创建成功, ID: %d", user.ID)
	return nil
}

// GetUserByPhone 根据手机号获取用户
func GetUserByPhone(tx *gorm.DB, phone string) (*models.User, error) {
	if tx == nil {
		tx = DB
	}

	var user models.User
	if err := tx.Where("phone = ?", phone).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据手机号查询用户失败: %v", err)
		return nil, err
	}

	return &user, nil
}

// GetUserByID 根据ID获取用户
func GetUserByID(tx *gorm.DB, userID int64) (*models.User, error) {
	if tx == nil {
		tx = DB
	}

	var user models.User
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据ID查询用户失败: %v", err)
		return nil, err
	}

	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(tx *gorm.DB, user *models.User) error {
	if tx == nil {
		tx = DB
	}

	if err := tx.Save(user).Error; err != nil {
		utils.Errorf("更新用户失败: %v", err)
		return err
	}

	utils.Infof("用户更新成功, ID: %d", user.ID)
	return nil
}

// CheckUserExistsByPhone 检查手机号是否已存在
func CheckUserExistsByPhone(tx *gorm.DB, phone string) (bool, error) {
	if tx == nil {
		tx = DB
	}

	var count int64
	if err := tx.Model(&models.User{}).Where("phone = ?", phone).Count(&count).Error; err != nil {
		utils.Errorf("检查用户手机号是否存在失败: %v", err)
		return false, err
	}

	return count > 0, nil
}

// Worker相关数据库操作

// GetWorkerByUserID 根据用户ID获取零工信息
func GetWorkerByUserID(tx *gorm.DB, userID int64) (*models.Worker, error) {
	if tx == nil {
		tx = DB
	}

	var worker models.Worker
	if err := tx.Where("user_id = ?", userID).First(&worker).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		utils.Errorf("根据用户ID查询零工信息失败: %v", err)
		return nil, err
	}

	return &worker, nil
}
