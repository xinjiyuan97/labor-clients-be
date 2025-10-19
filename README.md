# 零工APP后端服务

基于 Hertz 框架开发的零工APP后端服务，支持灵活就业人员的岗位匹配、日程管理和社区互动功能。

## 功能特性

- 🚀 基于 Hertz 高性能 Web 框架
- 🗄️ 支持 MySQL 和 Redis 数据存储
- 📝 完整的 API 接口设计
- 🔐 JWT 认证和权限管理
- 📊 日志记录和监控
- 🐳 Docker 容器化支持
- 🛠️ 命令行工具支持

## 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 6.0+
- Make (可选)

### 安装依赖

```bash
# 安装 Go 依赖
go mod download
go mod tidy

# 安装 hz 工具 (用于 IDL 代码生成)
go install github.com/cloudwego/hertz/cmd/hz@latest
```

### 配置设置

1. 复制配置文件：
```bash
cp config/config.example.yaml config/config.yaml
```

2. 修改配置文件中的数据库和 Redis 连接信息

### 运行方式

#### 方式一：使用启动脚本 (推荐)

```bash
# 启动开发环境服务器
./start.sh server example

# 启动生产环境服务器
./start.sh server prod

# 执行数据库迁移
./start.sh migrate example

# 编译项目
./start.sh build

# 运行测试
./start.sh test

# 查看帮助
./start.sh help
```

#### 方式二：使用 Makefile

```bash
# 开发环境完整流程
make dev

# 启动服务器 (开发环境)
make run

# 启动服务器 (生产环境)
make run-prod

# 执行数据库迁移
make migrate

# 执行数据库迁移 (生产环境)
make migrate-prod

# 编译项目
make build

# 运行测试
make test

# 清理构建文件
make clean
```

#### 方式三：直接使用 Go 命令

```bash
# 启动服务器 (开发环境)
go run . -mode server -env example

# 启动服务器 (生产环境)
go run . -mode server -env prod

# 执行数据库迁移
go run . -mode migrate -env example

# 查看帮助信息
go run . -help

# 查看版本信息
go run . -version
```

## 命令行参数

| 参数 | 说明 | 默认值 | 示例 |
|------|------|--------|------|
| `-mode` | 运行模式: server(启动服务器) \| migrate(数据库迁移) | server | `-mode server` |
| `-env` | 环境配置: example \| prod | example | `-env prod` |
| `-help` | 显示帮助信息 | - | `-help` |
| `-version` | 显示版本信息 | - | `-version` |

## 运行模式说明

### Server 模式

启动 HTTP 服务器，提供 API 服务。

```bash
# 启动开发环境服务器
go run . -mode server -env example

# 启动生产环境服务器
go run . -mode server -env prod
```

### Migrate 模式

执行数据库迁移，创建和更新数据库表结构。

```bash
# 执行开发环境数据库迁移
go run . -mode migrate -env example

# 执行生产环境数据库迁移
go run . -mode migrate -env prod
```

## 配置文件

项目支持多环境配置：

- `config/config.yaml` - 默认配置
- `config/config.example.yaml` - 开发环境配置示例
- `config/config.prod.yaml` - 生产环境配置

### 配置示例

```yaml
# Web服务器配置
server:
  host: "0.0.0.0"
  port: 8080
  mode: "debug"

# 数据库配置
database:
  host: "127.0.0.1"
  port: 3306
  username: "root"
  password: "password"
  database: "labor_clients"

# Redis配置
redis:
  host: "127.0.0.1"
  port: 6379
  password: ""
  database: 0

# 日志配置
log:
  level: "info"
  format: "json"
  output: "both"
  file_path: "./logs/app.log"
```

## API 文档

项目提供了完整的 API 接口文档：

- [API 设计文档](docs/api_design.md)
- [数据库字段映射](docs/database_mapping.md)
- [API 路由汇总](docs/api_routes.md)

## 项目结构

```
├── biz/                    # 业务逻辑层
│   ├── handler/           # 处理器
│   ├── model/             # 数据模型
│   └── router/            # 路由
├── config/                # 配置文件
├── dal/                   # 数据访问层
│   ├── mysql/            # MySQL 操作
│   └── redis/            # Redis 操作
├── docs/                  # 文档
├── idls/                  # Thrift IDL 定义
├── models/                # 数据模型
├── schemas/               # 数据库表结构
├── utils/                 # 工具包
├── main.go               # 主程序入口
├── Makefile              # 构建脚本
└── start.sh              # 启动脚本
```

## 开发指南

### 添加新的 API 接口

1. 在 `idls/` 目录下定义 Thrift IDL
2. 使用 `hz` 工具生成代码
3. 实现业务逻辑
4. 更新文档

### 数据库迁移

当修改数据模型时，需要执行数据库迁移：

```bash
# 开发环境
make migrate

# 生产环境
make migrate-prod
```

### 代码检查

```bash
# 格式化代码
make fmt

# 代码检查
make lint

# 运行测试
make test
```

## Docker 支持

### 构建镜像

```bash
make docker-build
```

### 运行容器

```bash
make docker-run
```

## 部署指南

### 生产环境部署

1. 配置生产环境配置文件
2. 编译项目
3. 执行数据库迁移
4. 启动服务

```bash
# 编译项目
make build

# 执行数据库迁移
make migrate-prod

# 启动生产服务
./output/bin/labor-clients-be -mode server -env prod
```

### 系统服务

可以将应用安装为系统服务：

```bash
# 安装
make install

# 卸载
make uninstall
```

## 监控和日志

### 日志查看

```bash
# 实时查看日志
make logs

# 查看日志文件
tail -f logs/app.log
```

### 健康检查

```bash
# 检查服务状态
make health

# 手动健康检查
curl http://localhost:8080/ping
```

## 故障排除

### 常见问题

1. **配置文件不存在**
   ```bash
   # 复制示例配置文件
   cp config/config.example.yaml config/config.yaml
   ```

2. **数据库连接失败**
   - 检查数据库配置
   - 确认数据库服务运行状态
   - 验证用户权限

3. **端口占用**
   ```bash
   # 查看端口占用
   lsof -i :8080
   
   # 修改配置文件中的端口
   ```

### 调试模式

```bash
# 启用调试日志
export LOG_LEVEL=debug
go run . -mode server -env example
```

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License

## 联系方式

- 项目地址: https://github.com/xinjiyuan97/labor-clients-be
- 问题反馈: https://github.com/xinjiyuan97/labor-clients-be/issues
