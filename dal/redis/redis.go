package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/xinjiyuan97/labor-clients/config"
	"github.com/xinjiyuan97/labor-clients/utils"
)

var (
	Client *redis.Client
	ctx    = context.Background()
)

// InitRedis 初始化Redis客户端
func InitRedis(cfg *config.RedisConfig) error {
	// 创建Redis客户端
	Client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.Database,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		DialTimeout:  time.Duration(cfg.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	})

	// 测试连接
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("连接Redis失败: %v", err)
	}

	utils.Info("Redis连接成功")
	return nil
}

// GetClient 获取Redis客户端
func GetClient() *redis.Client {
	return Client
}

// GetContext 获取Redis上下文
func GetContext() context.Context {
	return ctx
}

// CloseRedis 关闭Redis连接
func CloseRedis() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}

// HealthCheck Redis健康检查
func HealthCheck() error {
	if Client == nil {
		return fmt.Errorf("Redis客户端未初始化")
	}

	_, err := Client.Ping(ctx).Result()
	return err
}

// GetStats 获取Redis连接池统计信息
func GetStats() map[string]interface{} {
	if Client == nil {
		return nil
	}

	poolStats := Client.PoolStats()
	return map[string]interface{}{
		"hits":        poolStats.Hits,
		"misses":      poolStats.Misses,
		"timeouts":    poolStats.Timeouts,
		"total_conns": poolStats.TotalConns,
		"idle_conns":  poolStats.IdleConns,
		"stale_conns": poolStats.StaleConns,
	}
}

// Set 设置键值对
func Set(key string, value interface{}, expiration time.Duration) error {
	return Client.Set(ctx, key, value, expiration).Err()
}

// Get 获取值
func Get(key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

// Del 删除键
func Del(keys ...string) error {
	return Client.Del(ctx, keys...).Err()
}

// Exists 检查键是否存在
func Exists(keys ...string) (int64, error) {
	return Client.Exists(ctx, keys...).Result()
}

// Expire 设置键的过期时间
func Expire(key string, expiration time.Duration) error {
	return Client.Expire(ctx, key, expiration).Err()
}

// TTL 获取键的剩余过期时间
func TTL(key string) (time.Duration, error) {
	return Client.TTL(ctx, key).Result()
}

// HSet 设置哈希字段
func HSet(key string, values ...interface{}) error {
	return Client.HSet(ctx, key, values...).Err()
}

// HGet 获取哈希字段值
func HGet(key, field string) (string, error) {
	return Client.HGet(ctx, key, field).Result()
}

// HGetAll 获取所有哈希字段
func HGetAll(key string) (map[string]string, error) {
	return Client.HGetAll(ctx, key).Result()
}

// HDel 删除哈希字段
func HDel(key string, fields ...string) error {
	return Client.HDel(ctx, key, fields...).Err()
}

// LPush 从列表左侧推入元素
func LPush(key string, values ...interface{}) error {
	return Client.LPush(ctx, key, values...).Err()
}

// RPush 从列表右侧推入元素
func RPush(key string, values ...interface{}) error {
	return Client.RPush(ctx, key, values...).Err()
}

// LPop 从列表左侧弹出元素
func LPop(key string) (string, error) {
	return Client.LPop(ctx, key).Result()
}

// RPop 从列表右侧弹出元素
func RPop(key string) (string, error) {
	return Client.RPop(ctx, key).Result()
}

// LLen 获取列表长度
func LLen(key string) (int64, error) {
	return Client.LLen(ctx, key).Result()
}

// SAdd 向集合添加成员
func SAdd(key string, members ...interface{}) error {
	return Client.SAdd(ctx, key, members...).Err()
}

// SRem 从集合移除成员
func SRem(key string, members ...interface{}) error {
	return Client.SRem(ctx, key, members...).Err()
}

// SIsMember 检查成员是否在集合中
func SIsMember(key string, member interface{}) (bool, error) {
	return Client.SIsMember(ctx, key, member).Result()
}

// SMembers 获取集合所有成员
func SMembers(key string) ([]string, error) {
	return Client.SMembers(ctx, key).Result()
}

// ZAdd 向有序集合添加成员
func ZAdd(key string, members ...*redis.Z) error {
	return Client.ZAdd(ctx, key, members...).Err()
}

// ZRem 从有序集合移除成员
func ZRem(key string, members ...interface{}) error {
	return Client.ZRem(ctx, key, members...).Err()
}

// ZRange 获取有序集合指定范围的成员
func ZRange(key string, start, stop int64) ([]string, error) {
	return Client.ZRange(ctx, key, start, stop).Result()
}

// ZRangeByScore 根据分数范围获取有序集合成员
func ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	return Client.ZRangeByScore(ctx, key, opt).Result()
}

// ZCard 获取有序集合成员数量
func ZCard(key string) (int64, error) {
	return Client.ZCard(ctx, key).Result()
}

// Incr 递增计数器
func Incr(key string) (int64, error) {
	return Client.Incr(ctx, key).Result()
}

// IncrBy 按指定值递增计数器
func IncrBy(key string, value int64) (int64, error) {
	return Client.IncrBy(ctx, key, value).Result()
}

// Decr 递减计数器
func Decr(key string) (int64, error) {
	return Client.Decr(ctx, key).Result()
}

// DecrBy 按指定值递减计数器
func DecrBy(key string, value int64) (int64, error) {
	return Client.DecrBy(ctx, key, value).Result()
}

// MSet 批量设置键值对
func MSet(pairs ...interface{}) error {
	return Client.MSet(ctx, pairs...).Err()
}

// MGet 批量获取值
func MGet(keys ...string) ([]interface{}, error) {
	return Client.MGet(ctx, keys...).Result()
}

// Keys 获取匹配模式的所有键
func Keys(pattern string) ([]string, error) {
	return Client.Keys(ctx, pattern).Result()
}

// Scan 扫描键
func Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return Client.Scan(ctx, cursor, match, count).Result()
}

// Pipeline 创建管道
func Pipeline() redis.Pipeliner {
	return Client.Pipeline()
}

// TxPipeline 创建事务管道
func TxPipeline() redis.Pipeliner {
	return Client.TxPipeline()
}
