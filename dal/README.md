# DAL (Data Access Layer) 数据访问层

这个目录包含了项目的所有数据访问层代码，包括 MySQL 和 Redis 的客户端封装。

## 目录结构

```
dal/
├── mysql/           # MySQL 数据访问层
│   └── mysql.go    # GORM 封装
├── redis/           # Redis 数据访问层
│   └── redis.go    # Redis 客户端封装
├── example/         # 使用示例
│   └── example.go  # 完整的使用示例
├── dal.go          # DAL 统一初始化
└── README.md       # 本文档
```

## 功能特性

### MySQL DAL (`dal/mysql`)

基于 GORM 的 MySQL 数据访问层，提供以下功能：

- **连接管理**: 自动连接池管理
- **自动迁移**: 支持数据库表结构自动迁移
- **事务支持**: 提供事务操作方法
- **健康检查**: 数据库连接状态检查
- **统计信息**: 连接池使用统计

#### 使用方法：

```go
import (
    "github.com/xinjiyuan97/labor-clients/config"
    "github.com/xinjiyuan97/labor-clients/dal/mysql"
    "github.com/xinjiyuan97/labor-clients/models"
)

// 初始化MySQL
err := mysql.InitMySQL(&cfg.Database)
if err != nil {
    log.Fatal(err)
}

// 获取数据库连接
db := mysql.GetDB()

// 自动迁移表结构
err = mysql.AutoMigrate(&models.User{}, &models.Brand{})
if err != nil {
    log.Fatal(err)
}

// 创建记录
user := &models.User{
    Username: "test",
    Phone: "13800138000",
    // ... 其他字段
}
err = db.Create(user).Error

// 查询记录
var users []models.User
err = db.Where("role = ?", "worker").Find(&users).Error

// 使用事务
err = mysql.Transaction(func(tx *gorm.DB) error {
    // 在事务中执行操作
    return tx.Create(user).Error
})

// 健康检查
err = mysql.HealthCheck()

// 获取统计信息
stats := mysql.GetStats()

// 关闭连接
err = mysql.CloseMySQL()
```

### Redis DAL (`dal/redis`)

基于 go-redis 的 Redis 数据访问层，提供以下功能：

- **连接管理**: 自动连接池管理
- **数据类型支持**: 字符串、哈希、列表、集合、有序集合
- **管道操作**: 支持管道和事务管道
- **健康检查**: Redis 连接状态检查
- **统计信息**: 连接池使用统计

#### 使用方法：

```go
import (
    "github.com/xinjiyuan97/labor-clients/config"
    "github.com/xinjiyuan97/labor-clients/dal/redis"
    "time"
)

// 初始化Redis
err := redis.InitRedis(&cfg.Redis)
if err != nil {
    log.Fatal(err)
}

// 基本操作
err = redis.Set("key", "value", time.Hour)
value, err := redis.Get("key")
err = redis.Del("key")

// 哈希操作
err = redis.HSet("hash", "field", "value")
value, err := redis.HGet("hash", "field")
fields, err := redis.HGetAll("hash")

// 列表操作
err = redis.LPush("list", "item1", "item2")
err = redis.RPush("list", "item3")
item, err := redis.LPop("list")
length, err := redis.LLen("list")

// 集合操作
err = redis.SAdd("set", "member1", "member2")
members, err := redis.SMembers("set")
exists, err := redis.SIsMember("set", "member1")

// 有序集合操作
err = redis.ZAdd("zset", &redis.Z{Score: 1.0, Member: "member1"})
members, err := redis.ZRange("zset", 0, -1)
count, err := redis.ZCard("zset")

// 计数器操作
count, err := redis.Incr("counter")
count, err := redis.IncrBy("counter", 5)

// 批量操作
err = redis.MSet("key1", "value1", "key2", "value2")
values, err := redis.MGet("key1", "key2")

// 管道操作
pipe := redis.Pipeline()
pipe.Set(ctx, "key1", "value1", 0)
pipe.Set(ctx, "key2", "value2", 0)
_, err = pipe.Exec(ctx)

// 健康检查
err = redis.HealthCheck()

// 获取统计信息
stats := redis.GetStats()

// 关闭连接
err = redis.CloseRedis()
```

### 统一初始化 (`dal.go`)

提供统一的数据访问层初始化和管理：

```go
import (
    "github.com/xinjiyuan97/labor-clients/config"
    "github.com/xinjiyuan97/labor-clients/dal"
)

// 初始化所有DAL
err := dal.InitDAL(cfg)
if err != nil {
    log.Fatal(err)
}

// 健康检查
err = dal.HealthCheck()

// 获取统计信息
stats := dal.GetStats()

// 关闭所有连接
err = dal.CloseDAL()
```

## 配置说明

### MySQL 配置

```yaml
database:
  driver: "mysql"
  host: "127.0.0.1"
  port: 3306
  username: "root"
  password: "password"
  database: "labor_clients"
  charset: "utf8mb4"
  max_open_conns: 100
  max_idle_conns: 10
  conn_max_lifetime: 3600
  log_level: "warn"
```

### Redis 配置

```yaml
redis:
  host: "127.0.0.1"
  port: 6379
  password: ""
  database: 0
  pool_size: 100
  min_idle_conns: 10
  dial_timeout: 5
  read_timeout: 3
  write_timeout: 3
```

## 依赖包

- `gorm.io/gorm` - GORM ORM 框架
- `gorm.io/driver/mysql` - MySQL 驱动
- `github.com/go-redis/redis/v8` - Redis 客户端

## 注意事项

1. **连接池配置**: 根据实际负载调整连接池参数
2. **错误处理**: 所有操作都应该检查错误返回值
3. **事务使用**: 在需要原子性的操作中使用事务
4. **资源清理**: 程序退出时调用 `CloseDAL()` 关闭连接
5. **健康检查**: 定期进行健康检查确保连接可用
6. **日志记录**: 重要操作会记录日志，便于调试和监控

## 示例

查看 `example/example.go` 文件获取完整的使用示例。
