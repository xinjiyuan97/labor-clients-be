package utils

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger *logrus.Logger
)

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

// InitLogger 初始化全局日志
func InitLogger(config *LogConfig) error {
	Logger = logrus.New()

	// 设置日志级别
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)
	Logger.SetReportCaller(true)
	// 设置日志格式
	switch config.Format {
	case "json":
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	case "text":
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	default:
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	}

	// 设置输出
	switch config.Output {
	case "file":
		// 确保日志目录存在
		if config.FilePath != "" {
			dir := filepath.Dir(config.FilePath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		Logger.SetOutput(&lumberjack.Logger{
			Filename:   config.FilePath,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		})
	case "stdout":
		Logger.SetOutput(os.Stdout)
	case "both":
		// 同时输出到文件和控制台
		if config.FilePath != "" {
			dir := filepath.Dir(config.FilePath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}

		// 创建文件输出
		fileOutput := &lumberjack.Logger{
			Filename:   config.FilePath,
			MaxSize:    config.MaxSize,
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,
			Compress:   config.Compress,
		}

		// 使用MultiWriter同时输出到文件和控制台
		Logger.SetOutput(&MultiWriter{
			File:    fileOutput,
			Console: os.Stdout,
		})
	default:
		Logger.SetOutput(os.Stdout)
	}

	return nil
}

// MultiWriter 实现同时写入文件和控制台的Writer
type MultiWriter struct {
	File    *lumberjack.Logger
	Console *os.File
}

func (mw *MultiWriter) Write(p []byte) (n int, err error) {
	// 写入文件
	n1, err1 := mw.File.Write(p)
	// 写入控制台
	n2, err2 := mw.Console.Write(p)

	// 返回较小的写入字节数和错误
	if err1 != nil {
		return n1, err1
	}
	if err2 != nil {
		return n2, err2
	}

	if n1 < n2 {
		return n1, nil
	}
	return n2, nil
}

// GetLogger 获取全局日志实例
func GetLogger() *logrus.Logger {
	if Logger == nil {
		// 如果未初始化，使用默认配置
		config := &LogConfig{
			Level:  "info",
			Format: "json",
			Output: "stdout",
		}
		InitLogger(config)
	}
	return Logger
}

// LogWithFields 创建带字段的日志条目
func LogWithFields(fields logrus.Fields) *logrus.Entry {
	return GetLogger().WithFields(fields)
}

// LogWithField 创建带单个字段的日志条目
func LogWithField(key string, value interface{}) *logrus.Entry {
	return GetLogger().WithField(key, value)
}

// LogWithError 创建带错误的日志条目
func LogWithError(err error) *logrus.Entry {
	return GetLogger().WithError(err)
}

// 便捷的日志方法
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debugf(format, args...)
}

func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

func Infof(format string, args ...interface{}) {
	GetLogger().Infof(format, args...)
}

func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	GetLogger().Warnf(format, args...)
}

func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	GetLogger().Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	GetLogger().Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	GetLogger().Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	GetLogger().Panicf(format, args...)
}
