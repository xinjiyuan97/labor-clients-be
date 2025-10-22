package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var GlobalConfig *Config

func SetGlobalConfig(config *Config) {
	GlobalConfig = config
}

func GetGlobalConfig() *Config {
	return GlobalConfig
}

// Config 应用配置结构
type Config struct {
	Server    ServerConfig    `yaml:"server"`
	Database  DatabaseConfig  `yaml:"database"`
	Redis     RedisConfig     `yaml:"redis"`
	OSS       OSSConfig       `yaml:"oss"`
	Log       LogConfig       `yaml:"log"`
	Auth      AuthConfig      `yaml:"auth"`
	Snowflake SnowflakeConfig `yaml:"snowflake"`
}

// ServerConfig Web服务器配置
type ServerConfig struct {
	Host         string `yaml:"host"`          // 服务器地址
	Port         int    `yaml:"port"`          // 服务器端口
	ReadTimeout  int    `yaml:"read_timeout"`  // 读取超时时间(秒)
	WriteTimeout int    `yaml:"write_timeout"` // 写入超时时间(秒)
	IdleTimeout  int    `yaml:"idle_timeout"`  // 空闲超时时间(秒)
	Mode         string `yaml:"mode"`          // 运行模式: debug, release, test
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string `yaml:"driver"`            // 数据库驱动
	Host            string `yaml:"host"`              // 数据库主机
	Port            int    `yaml:"port"`              // 数据库端口
	Username        string `yaml:"username"`          // 用户名
	Password        string `yaml:"password"`          // 密码
	Database        string `yaml:"database"`          // 数据库名
	Charset         string `yaml:"charset"`           // 字符集
	MaxOpenConns    int    `yaml:"max_open_conns"`    // 最大打开连接数
	MaxIdleConns    int    `yaml:"max_idle_conns"`    // 最大空闲连接数
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"` // 连接最大生命周期(秒)
	LogLevel        string `yaml:"log_level"`         // 日志级别: silent, error, warn, info
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host         string `yaml:"host"`           // Redis主机
	Port         int    `yaml:"port"`           // Redis端口
	Password     string `yaml:"password"`       // Redis密码
	Database     int    `yaml:"database"`       // Redis数据库
	PoolSize     int    `yaml:"pool_size"`      // 连接池大小
	MinIdleConns int    `yaml:"min_idle_conns"` // 最小空闲连接数
	DialTimeout  int    `yaml:"dial_timeout"`   // 连接超时时间(秒)
	ReadTimeout  int    `yaml:"read_timeout"`   // 读取超时时间(秒)
	WriteTimeout int    `yaml:"write_timeout"`  // 写入超时时间(秒)
}

// OSSConfig 对象存储配置
type OSSConfig struct {
	Provider    string   `yaml:"provider"`      // 存储提供商: aliyun, tencent, aws, volcengine, local
	AccessKey   string   `yaml:"access_key"`    // 访问密钥
	SecretKey   string   `yaml:"secret_key"`    // 秘密密钥
	Region      string   `yaml:"region"`        // 区域
	Bucket      string   `yaml:"bucket"`        // 存储桶名称
	Endpoint    string   `yaml:"endpoint"`      // 端点地址
	BaseURL     string   `yaml:"base_url"`      // 基础URL
	UploadPath  string   `yaml:"upload_path"`   // 上传路径
	MaxFileSize int64    `yaml:"max_file_size"` // 最大文件大小(字节)
	AllowedExts []string `yaml:"allowed_exts"`  // 允许的文件扩展名
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `yaml:"level"`       // 日志级别: debug, info, warn, error, fatal, panic
	Format     string `yaml:"format"`      // 日志格式: json, text
	Output     string `yaml:"output"`      // 输出方式: file, stdout, both
	FilePath   string `yaml:"file_path"`   // 日志文件路径
	MaxSize    int    `yaml:"max_size"`    // 单个日志文件最大大小(MB)
	MaxBackups int    `yaml:"max_backups"` // 保留的日志文件数量
	MaxAge     int    `yaml:"max_age"`     // 日志文件保留天数
	Compress   bool   `yaml:"compress"`    // 是否压缩旧日志文件
}

// AuthConfig 鉴权配置
type AuthConfig struct {
	JWTSecret     string `yaml:"jwt_secret"`     // JWT密钥
	JWTExpire     int    `yaml:"jwt_expire"`     // JWT过期时间(小时)
	RefreshExpire int    `yaml:"refresh_expire"` // 刷新令牌过期时间(天)
	Issuer        string `yaml:"issuer"`         // JWT签发者
	Audience      string `yaml:"audience"`       // JWT受众
	Algorithm     string `yaml:"algorithm"`      // 签名算法
}

// SnowflakeConfig 雪花算法配置
type SnowflakeConfig struct {
	NodeID int64 `yaml:"node_id"` // 节点ID
}

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	// 如果配置文件路径为空，使用默认路径
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("配置文件不存在: %s", configPath)
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

// SaveConfig 保存配置到文件
func SaveConfig(config *Config, configPath string) error {
	// 如果配置文件路径为空，使用默认路径
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	// 转换为YAML
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	// 写入文件
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}

	return nil
}

// GetDSN 获取数据库连接字符串
func (d *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		d.Username, d.Password, d.Host, d.Port, d.Database, d.Charset)
}

// GetRedisAddr 获取Redis地址
func (r *RedisConfig) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

// GetServerAddr 获取服务器地址
func (s *ServerConfig) GetServerAddr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.Server.Host == "" {
		c.Server.Host = "0.0.0.0"
	}
	if c.Server.Port == 0 {
		c.Server.Port = 8080
	}
	if c.Server.Mode == "" {
		c.Server.Mode = "release"
	}
	if c.Database.Driver == "" {
		c.Database.Driver = "mysql"
	}
	if c.Database.Host == "" {
		return fmt.Errorf("数据库主机地址不能为空")
	}
	if c.Database.Port == 0 {
		c.Database.Port = 3306
	}
	if c.Database.Username == "" {
		return fmt.Errorf("数据库用户名不能为空")
	}
	if c.Database.Database == "" {
		return fmt.Errorf("数据库名不能为空")
	}
	if c.Redis.Host == "" {
		c.Redis.Host = "127.0.0.1"
	}
	if c.Redis.Port == 0 {
		c.Redis.Port = 6379
	}
	if c.Auth.JWTSecret == "" {
		return fmt.Errorf("JWT密钥不能为空")
	}
	if c.Auth.JWTExpire == 0 {
		c.Auth.JWTExpire = 24 // 默认24小时
	}
	if c.Log.Level == "" {
		c.Log.Level = "info"
	}
	if c.Log.Format == "" {
		c.Log.Format = "json"
	}
	if c.Log.Output == "" {
		c.Log.Output = "stdout"
	}
	if c.Snowflake.NodeID == 0 {
		c.Snowflake.NodeID = 1 // 默认节点ID为1
	}

	return nil
}
