# 短信验证码登录和第三方账号绑定功能文档

## 功能概述

新增了三个核心功能：
1. **发送短信验证码** - 用于手机号验证
2. **短信验证码登录** - 使用手机号和验证码登录
3. **第三方账号绑定** - 将第三方平台账号（微信、支付宝等）与现有账号绑定

## 数据库设计

### third_party_bindings 表

记录第三方平台账号与用户账号的绑定关系。支持多个平台（微信、支付宝、QQ等）。

**表结构**：
```sql
CREATE TABLE IF NOT EXISTS third_party_bindings (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    platform VARCHAR(50) NOT NULL COMMENT '第三方平台',
    openid VARCHAR(255) NOT NULL COMMENT '第三方平台OpenID',
    unionid VARCHAR(255) COMMENT '第三方平台UnionID',
    appid VARCHAR(100) NOT NULL COMMENT '应用AppID',
    nickname VARCHAR(100) COMMENT '第三方平台昵称',
    avatar VARCHAR(500) COMMENT '第三方平台头像',
    status ENUM('active', 'disabled') NOT NULL DEFAULT 'active' COMMENT '绑定状态',
    last_login_at DATETIME COMMENT '最后登录时间',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_user_id (user_id),
    INDEX idx_platform (platform),
    UNIQUE KEY idx_platform_openid (platform, openid),
    INDEX idx_unionid (unionid),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
```

**索引**：
- `user_id` - 用户ID索引
- `platform` - 平台索引
- `idx_platform_openid` - 平台和OpenID联合唯一索引
- `unionid` - UnionID索引

**支持的平台**：
- `wechat` - 微信小程序/公众号
- `alipay` - 支付宝小程序
- `qq` - QQ登录
- `apple` - Apple登录
- 其他平台可按需扩展

## API接口

### 1. 发送短信验证码

**接口**: `POST /api/v1/auth/send-sms-code`

**请求**:
```json
{
  "phone": "13800138000"
}
```

**响应**:
```json
{
  "base": {
    "code": 200,
    "message": "验证码发送成功",
    "timestamp": "2025-01-23T10:00:00Z"
  },
  "code": "123456",
  "expires_in": 300
}
```

**说明**:
- 验证码有效期5分钟
- 验证码存储在Redis中，key格式：`sms_code:{phone}`
- 开发环境会在响应中返回验证码（方便测试）
- 生产环境需要集成真实短信服务提供商

**TODO**: 集成真实短信服务API（阿里云、腾讯云等）

### 2. 短信验证码登录

**接口**: `POST /api/v1/auth/login-with-sms`

**请求**:
```json
{
  "phone": "13800138000",
  "code": "123456"
}
```

**响应**:
```json
{
  "base": {
    "code": 200,
    "message": "登录成功",
    "timestamp": "2025-01-23T10:00:00Z"
  },
  "user_id": "123456",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2025-01-24T10:00:00Z"
}
```

**错误响应**:
- 验证码错误：`code: 400, message: "验证码错误"`
- 验证码过期：`code: 400, message: "验证码已过期或无效"`
- 用户不存在：`code: 404, message: "用户不存在，请先注册"`
- 账号禁用：`code: 403, message: "账号已被禁用"`

**流程**:
1. 验证短信验证码
2. 从Redis删除已使用的验证码
3. 查询用户是否存在
4. 检查账号状态
5. 生成JWT token
6. 返回登录结果

### 3. 第三方登录绑定

**接口**: `POST /api/v1/auth/third-party-bind`

**请求**:
```json
{
  "platform": "wechat",
  "openid": "oxxxxxxxxxxxxxxxxxxxxxx",
  "unionid": "oxxxxxxxxxxxxxxxxxxxxxx",
  "appid": "wx1234567890abcdef",
  "phone": "13800138000",
  "code": "123456",
  "nickname": "第三方昵称",
  "avatar": "https://example.com/avatar.jpg"
}
```

**响应**:
```json
{
  "base": {
    "code": 200,
    "message": "登录成功",
    "timestamp": "2025-01-23T10:00:00Z"
  },
  "is_new_user": true,
  "user_id": "123456",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2025-01-24T10:00:00Z"
}
```

**平台参数说明**:
- `wechat` - 微信小程序/公众号
- `alipay` - 支付宝小程序
- `qq` - QQ登录
- `apple` - Apple登录

**错误响应**:
- 验证码错误：`code: 400`
- 验证码过期：`code: 400`
- 系统错误：`code: 500`

**流程**:
1. 验证短信验证码
2. 删除已使用的验证码
3. 检查该平台是否已有绑定
4. 查询或创建用户
   - 如果用户不存在，创建新用户（角色默认为worker）
   - 如果用户存在，使用现有用户
5. 更新或创建第三方绑定
6. 生成JWT token
7. 返回绑定结果

**特性**:
- 支持新用户自动注册
- 支持已有账号绑定第三方平台
- 同一个用户可以绑定多个第三方平台
- 如果该平台已绑定，更新绑定信息
- 记录最后登录时间

## 使用场景

### 场景1: 新用户微信登录

1. 用户在小程序中选择"微信登录"
2. 获取微信openid等信息
3. 填写手机号
4. 请求发送验证码
5. 输入验证码
6. 调用第三方绑定接口（platform: "wechat"）
7. 系统自动创建账号并返回token

### 场景2: 已有账号绑定微信

1. 用户通过手机号+短信验证码登录
2. 在小程序中获取微信openid
3. 调用第三方绑定接口绑定openid
4. 后续可使用微信快速登录

### 场景3: 短信验证码登录

1. 用户输入手机号
2. 请求发送验证码
3. 输入验证码
4. 调用短信验证码登录接口
5. 系统返回token

### 场景4: 绑定多个第三方平台

用户可以同时绑定多个第三方平台：
```json
// 绑定微信
{
  "platform": "wechat",
  "openid": "wechat_openid_123",
  "phone": "13800138000",
  "code": "123456"
}

// 绑定支付宝
{
  "platform": "alipay",
  "openid": "alipay_openid_456",
  "phone": "13800138000",
  "code": "789012"
}
```

## 数据模型

### ThirdPartyBinding 模型

```go
type ThirdPartyBinding struct {
    BaseModel
    UserID      int64  // 用户ID
    Platform    string // 第三方平台
    OpenID      string // 第三方平台OpenID
    UnionID     string // 第三方平台UnionID
    AppID       string // 应用AppID
    Nickname    string // 第三方平台昵称
    Avatar      string // 第三方平台头像
    Status      string // 绑定状态
    LastLoginAt string // 最后登录时间
}
```

## Redis存储

**验证码存储**：
- Key: `sms_code:{phone}`
- Value: 6位数字验证码
- TTL: 5分钟
- 使用后自动删除

**示例**:
```
Key: sms_code:13800138000
Value: 123456
TTL: 300s
```

## 代码结构

### 数据库访问层
- `dal/mysql/third_party_binding.go` - 第三方绑定CRUD操作

### 业务逻辑层
- `biz/logic/auth/send_sms_code.go` - 发送验证码逻辑
- `biz/logic/auth/login_with_sms_code.go` - 短信登录逻辑
- `biz/logic/auth/third_party_bind.go` - 第三方绑定逻辑

### Handler层
- `biz/handler/auth/send_smscode.go` - 发送验证码接口
- `biz/handler/auth/login_with_smscode.go` - 短信登录接口
- `biz/handler/auth/third_party_login_bind.go` - 第三方绑定接口

### 数据模型
- `models/third_party_binding.go` - 第三方绑定模型
- `schemas/third_party_bindings.sql` - 数据库表结构

## 待完善功能

### 短信服务集成

当前验证码仅打印到日志，生产环境需要集成真实短信服务：

```go
// TODO: 在 biz/logic/auth/send_sms_code.go 中
// 需要集成以下服务商之一：
// 1. 阿里云短信服务
// 2. 腾讯云短信服务
// 3. 云片网
// 4. 其他短信服务提供商
```

### 验证码发送频率限制

建议添加：
- 同一手机号1分钟内最多发送1次
- 同一手机号1小时内最多发送5次
- 使用Redis计数器实现

### 安全增强

1. **验证码复杂度**：当前为6位数字，可增加复杂度
2. **图形验证码**：在发送短信前验证图形验证码
3. **IP频率限制**：限制同一IP的请求频率
4. **账号安全**：新用户默认密码需要改进

## 测试

### 1. 发送验证码测试

```bash
curl -X POST http://localhost:8888/api/v1/auth/send-sms-code \
  -H "Content-Type: application/json" \
  -d '{"phone": "13800138000"}'
```

### 2. 短信登录测试

```bash
curl -X POST http://localhost:8888/api/v1/auth/login-with-sms \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "13800138000",
    "code": "123456"
  }'
```

### 3. 第三方绑定测试

**微信绑定**:
```bash
curl -X POST http://localhost:8888/api/v1/auth/third-party-bind \
  -H "Content-Type: application/json" \
  -d '{
    "platform": "wechat",
    "openid": "test_openid_12345",
    "unionid": "test_unionid_12345",
    "appid": "wx1234567890abcdef",
    "phone": "13800138000",
    "code": "123456",
    "nickname": "微信用户",
    "avatar": "https://example.com/avatar.jpg"
  }'
```

**支付宝绑定**:
```bash
curl -X POST http://localhost:8888/api/v1/auth/third-party-bind \
  -H "Content-Type: application/json" \
  -d '{
    "platform": "alipay",
    "openid": "alipay_openid_67890",
    "unionid": "alipay_unionid_67890",
    "appid": "alipay_app_id",
    "phone": "13800138000",
    "code": "123456",
    "nickname": "支付宝用户",
    "avatar": "https://example.com/alipay_avatar.jpg"
  }'
```

## 数据库迁移

执行数据库迁移创建新表：

```bash
# 方式1: 通过命令行
./labor-clients -mode migrate -env prod

# 方式2: 手动执行SQL
mysql -u root -p labors < schemas/third_party_bindings.sql
```

## 日志输出

所有操作都有详细的日志记录：

```json
{
  "level": "info",
  "msg": "发送短信验证码",
  "phone": "13800138000",
  "code": "123456",
  "key": "sms_code:13800138000",
  "time": "2025-01-23T10:00:00Z"
}

{
  "level": "info",
  "msg": "短信验证码登录成功",
  "user_id": 123456,
  "phone": "13800138000",
  "role": "worker",
  "time": "2025-01-23T10:00:00Z"
}

{
  "level": "info",
  "msg": "第三方登录绑定成功",
  "user_id": 123456,
  "phone": "13800138000",
  "platform": "wechat",
  "openid": "test_openid_12345",
  "is_new_user": true,
  "time": "2025-01-23T10:00:00Z"
}
```

## 注意事项

1. **短信服务**：需要集成真实的短信服务提供商
2. **验证码安全**：生产环境应从响应中移除验证码字段
3. **用户创建**：新用户默认角色为worker，需要根据业务调整
4. **默认密码**：当前使用固定默认密码，建议改进
5. **绑定唯一性**：同一个平台+openid只能绑定一个账号
6. **软删除**：支持软删除绑定记录
7. **平台扩展**：支持新增其他第三方平台，只需传入不同的platform参数

## 相关文档

- [微信云托管部署文档](./wechat_cloud_deployment.md)
- [数据库设计文档](./database_mapping.md)
- [API设计文档](./api_design.md)
