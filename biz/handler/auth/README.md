# Auth Handlers 实现说明

## 概述

本目录下的所有auth handlers已经完成实现，使用JWT进行鉴权，数据库操作在mysql中实现，事务操作使用gorm的闭包transaction实现。

## 实现的功能

### 1. 用户注册 (`register.go`)
- **路由**: `POST /api/v1/auth/register`
- **功能**: 用户注册，支持worker和employer两种角色
- **请求参数**:
  - `phone`: 手机号
  - `password`: 密码
  - `role`: 角色 (worker/employer)
  - `username`: 用户名
- **响应**: 返回用户ID、JWT token和过期时间

### 2. 用户登录 (`login.go`)
- **路由**: `POST /api/v1/auth/login`
- **功能**: 用户登录验证
- **请求参数**:
  - `phone`: 手机号
  - `password`: 密码
- **响应**: 返回用户ID、JWT token和过期时间

### 3. 用户登出 (`logout.go`)
- **路由**: `POST /api/v1/auth/logout`
- **功能**: 用户登出
- **请求参数**:
  - `Authorization`: Bearer token (header)
- **响应**: 登出成功消息

### 4. 刷新Token (`refresh_token.go`)
- **路由**: `POST /api/v1/auth/refresh`
- **功能**: 刷新JWT token
- **请求参数**:
  - `refresh_token`: 刷新token
- **响应**: 返回新的JWT token和过期时间

### 5. 获取用户信息 (`get_user_profile.go`)
- **路由**: `GET /api/v1/auth/profile`
- **功能**: 获取用户详细信息
- **请求参数**:
  - `user_id`: 用户ID (path参数)
- **响应**: 返回用户基础信息和零工详细信息(如果是worker角色)

## 技术实现

### JWT工具类 (`utils/jwt.go`)
- 使用 `github.com/golang-jwt/jwt/v5` 实现JWT token生成和验证
- 支持access token和refresh token
- Token过期时间: 24小时
- Refresh token过期时间: 7天

### 密码加密 (`utils/password.go`)
- 使用 `golang.org/x/crypto/bcrypt` 进行密码哈希加密
- 提供密码加密和验证功能

### 数据库操作 (`dal/mysql/mysql.go`)
- 使用GORM进行数据库操作
- 支持事务操作
- 提供用户和零工信息的CRUD操作

### 业务逻辑层 (`biz/logic/auth/`)
- 分离业务逻辑和HTTP处理
- 使用事务确保数据一致性
- 统一的错误处理和响应格式

### JWT中间件 (`middleware/jwt.go`)
- 提供JWT鉴权中间件
- 自动验证Authorization header
- 将用户信息存储到context中供后续使用

## 使用示例

### 注册用户
```bash
curl -X POST http://localhost:8888/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "13800138000",
    "password": "password123",
    "role": "worker",
    "username": "testuser"
  }'
```

### 用户登录
```bash
curl -X POST http://localhost:8888/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "13800138000",
    "password": "password123"
  }'
```

### 获取用户信息
```bash
curl -X GET http://localhost:8888/api/v1/auth/profile/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 注意事项

1. JWT密钥目前硬编码在代码中，生产环境应该从配置文件读取
2. 登出功能目前只是验证token，没有实现token黑名单机制
3. 所有API都返回统一的响应格式，包含code、message和timestamp字段
4. 数据库操作都支持事务，确保数据一致性
5. 密码使用bcrypt加密，安全性较高
