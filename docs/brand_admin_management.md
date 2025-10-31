# 品牌用户角色管理API文档

## 概述

本模块提供了按品牌获取和管理管理员的API接口，支持品牌管理员和门店管理员的增删改查操作。

## API接口列表

### 1. 获取品牌管理员列表

**接口**: `GET /api/v1/admin/brands/:brand_id/admins`

**描述**: 获取指定品牌下的所有管理员列表（包括品牌管理员和门店管理员）

**请求参数**:

| 参数 | 类型 | 位置 | 必填 | 说明 |
|------|------|------|------|------|
| brand_id | int64 | path | 是 | 品牌ID |
| page | int32 | query | 是 | 页码（从1开始） |
| limit | int32 | query | 是 | 每页数量（1-100） |
| role_type | string | query | 否 | 角色类型过滤：brand_admin, store_admin |
| status | string | query | 否 | 状态过滤：active, disabled |

**请求示例**:
```bash
GET /api/v1/admin/brands/1/admins?page=1&limit=10&role_type=brand_admin&status=active
```

**响应示例**:
```json
{
  "base": {
    "code": 0,
    "message": "success"
  },
  "page_info": {
    "total": 15,
    "page": 1,
    "limit": 10
  },
  "admins": [
    {
      "user_id": "100",
      "role_id": "1",
      "username": "张三",
      "phone": "13800138000",
      "role_type": "brand_admin",
      "brand_id": "1",
      "brand_name": "某某品牌",
      "store_id": "0",
      "store_name": "",
      "status": "active",
      "created_at": "2025-01-01 10:00:00"
    },
    {
      "user_id": "101",
      "role_id": "2",
      "username": "李四",
      "phone": "13800138001",
      "role_type": "store_admin",
      "brand_id": "1",
      "brand_name": "某某品牌",
      "store_id": "5",
      "store_name": "朝阳门店",
      "status": "active",
      "created_at": "2025-01-02 14:30:00"
    }
  ]
}
```

### 2. 创建品牌管理员

**接口**: `POST /api/v1/admin/brands/admins`

**描述**: 通过手机号创建品牌管理员。如果手机号未注册，系统会自动创建用户（默认密码123456）

**请求参数**:

| 参数 | 类型 | 位置 | 必填 | 说明 |
|------|------|------|------|------|
| phone | string | body | 是 | 手机号 |
| brand_id | int64 | body | 是 | 品牌ID |
| role_type | string | body | 是 | 角色类型（通常为brand_admin） |
| real_name | string | body | 否 | 真实姓名（如果提供，将作为用户名；否则使用手机号） |

**请求示例**:
```json
{
  "phone": "13800138000",
  "brand_id": "1",
  "role_type": "brand_admin",
  "real_name": "张三"
}
```

**响应示例**:
```json
{
  "base": {
    "code": 0,
    "message": "success"
  },
  "role_id": "123"
}
```

**错误响应**:
```json
{
  "base": {
    "code": 404,
    "message": "用户不存在"
  }
}
```

```json
{
  "base": {
    "code": 400,
    "message": "该用户已是该品牌的管理员"
  }
}
```

### 3. 创建门店管理员

**接口**: `POST /api/v1/admin/stores/admins`

**描述**: 通过手机号创建门店管理员。如果手机号未注册，系统会自动创建用户（默认密码123456）

**请求参数**:

| 参数 | 类型 | 位置 | 必填 | 说明 |
|------|------|------|------|------|
| phone | string | body | 是 | 手机号 |
| brand_id | int64 | body | 是 | 品牌ID |
| store_id | int64 | body | 是 | 门店ID |
| real_name | string | body | 否 | 真实姓名（如果提供，将作为用户名；否则使用手机号） |

**请求示例**:
```json
{
  "phone": "13800138001",
  "brand_id": "1",
  "store_id": "5",
  "real_name": "李四"
}
```

**响应示例**:
```json
{
  "base": {
    "code": 0,
    "message": "success"
  },
  "role_id": "124"
}
```

**错误响应**:
```json
{
  "base": {
    "code": 400,
    "message": "门店不属于该品牌"
  }
}
```

```json
{
  "base": {
    "code": 400,
    "message": "该用户已是该门店的管理员"
  }
}
```

### 4. 删除品牌/门店管理员

**接口**: `DELETE /api/v1/admin/brand-admins/:role_id`

**描述**: 删除指定的品牌或门店管理员角色

**请求参数**:

| 参数 | 类型 | 位置 | 必填 | 说明 |
|------|------|------|------|------|
| role_id | int64 | path | 是 | 角色ID |

**请求示例**:
```bash
DELETE /api/v1/admin/brand-admins/123
```

**响应示例**:
```json
{
  "base": {
    "code": 0,
    "message": "success"
  }
}
```

### 5. 更新品牌/门店管理员状态

**接口**: `PUT /api/v1/admin/brand-admins/:role_id/status`

**描述**: 启用或禁用品牌/门店管理员角色

**请求参数**:

| 参数 | 类型 | 位置 | 必填 | 说明 |
|------|------|------|------|------|
| role_id | int64 | path | 是 | 角色ID |
| status | string | body | 是 | 状态：active, disabled |

**请求示例**:
```json
{
  "status": "disabled"
}
```

**响应示例**:
```json
{
  "base": {
    "code": 0,
    "message": "success"
  }
}
```

## 数据模型

### BrandAdminInfo

| 字段 | 类型 | 说明 |
|------|------|------|
| user_id | int64 | 用户ID |
| role_id | int64 | 角色ID（用于删除/更新） |
| username | string | 用户名 |
| phone | string | 手机号 |
| role_type | string | 角色类型：brand_admin, store_admin |
| brand_id | int64 | 品牌ID |
| brand_name | string | 品牌名称 |
| store_id | int64 | 门店ID（仅门店管理员） |
| store_name | string | 门店名称（仅门店管理员） |
| status | string | 状态：active, disabled |
| created_at | string | 创建时间 |

## 使用场景

### 场景1：查看品牌下所有管理员

```bash
# 系统管理员查看某个品牌的所有管理员
GET /api/v1/admin/brands/1/admins?page=1&limit=20
```

### 场景2：只查看品牌管理员

```bash
# 过滤只看品牌管理员
GET /api/v1/admin/brands/1/admins?page=1&limit=20&role_type=brand_admin
```

### 场景3：只查看门店管理员

```bash
# 过滤只看门店管理员
GET /api/v1/admin/brands/1/admins?page=1&limit=20&role_type=store_admin
```

### 场景4：为用户分配品牌管理员角色

```bash
# 1. 先创建用户（如果还没有）
POST /api/v1/admin/users
{
  "phone": "13800138000",
  "real_name": "张三",
  "role": "admin",
  "password": "123456"
}
# 返回: { "user_id": "100" }

# 2. 为用户分配品牌管理员角色
POST /api/v1/admin/brands/admins
{
  "user_id": "100",
  "brand_id": "1",
  "role_type": "brand_admin"
}
```

### 场景5：为用户分配门店管理员角色

```bash
# 1. 先确保门店存在
POST /api/v1/admin/stores
{
  "brand_id": "1",
  "name": "朝阳门店",
  "address": "北京市朝阳区..."
}
# 返回: { "store_id": "5" }

# 2. 为用户分配门店管理员角色
POST /api/v1/admin/stores/admins
{
  "user_id": "101",
  "brand_id": "1",
  "store_id": "5"
}
```

### 场景6：临时禁用某个管理员

```bash
# 禁用管理员（不删除，只是禁用）
PUT /api/v1/admin/brand-admins/123/status
{
  "status": "disabled"
}

# 恢复管理员
PUT /api/v1/admin/brand-admins/123/status
{
  "status": "active"
}
```

### 场景7：移除管理员角色

```bash
# 完全移除管理员角色
DELETE /api/v1/admin/brand-admins/123
```

## 权限控制

建议的权限配置：

- **系统管理员(admin)**: 可以执行所有操作
- **品牌管理员(brand_admin)**: 
  - 可以查看自己品牌的管理员
  - 可以为自己品牌添加门店管理员
  - 可以禁用/删除门店管理员
- **门店管理员(store_admin)**: 只读权限，查看同品牌的管理员

## 业务规则

### 1. 角色唯一性

- 同一个用户不能重复分配相同的品牌管理员角色
- 同一个用户可以是多个品牌的管理员
- 同一个用户可以是多个门店的管理员

### 2. 数据验证

- 创建角色前会验证用户、品牌、门店是否存在
- 门店管理员的门店必须属于指定的品牌
- 状态只能是 `active` 或 `disabled`

### 3. 软删除

- 删除角色使用软删除，数据保留在数据库中
- deleted_at字段标记删除时间

## 数据库设计

### user_roles表

```sql
CREATE TABLE IF NOT EXISTS user_roles (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    role_type ENUM('brand_admin', 'store_admin') NOT NULL,
    brand_id BIGINT NULL,
    store_id BIGINT NULL,
    status ENUM('active', 'disabled') NOT NULL DEFAULT 'active',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_user_id (user_id),
    INDEX idx_role_type (role_type),
    INDEX idx_brand_id (brand_id),
    INDEX idx_store_id (store_id),
    INDEX idx_status (status),
    UNIQUE KEY unique_user_brand_store (user_id, brand_id, store_id, deleted_at)
);
```

## 相关文件

### DAL层
- `/dal/mysql/user_role.go` - 基础CRUD操作
- `/dal/mysql/user_role_extended.go` - 扩展查询（带用户信息）

### Logic层
- `/biz/logic/admin/get_brand_admins.go` - 获取管理员列表
- `/biz/logic/admin/create_brand_admin.go` - 创建品牌管理员
- `/biz/logic/admin/create_store_admin.go` - 创建门店管理员
- `/biz/logic/admin/delete_brand_admin.go` - 删除管理员
- `/biz/logic/admin/update_brand_admin_status.go` - 更新状态

### Handler层
- `/biz/handler/admin/get_brand_admins.go`
- `/biz/handler/admin/create_brand_admin.go`
- `/biz/handler/admin/create_store_admin.go`
- `/biz/handler/admin/delete_brand_admin.go`
- `/biz/handler/admin/update_brand_admin_status.go`

### Thrift定义
- `/idls/admin.thrift` - API接口定义

## 注意事项

1. **角色关联**: user_roles表记录的是用户的扩展角色，users表的role字段保持不变
2. **多角色**: 同一个用户可以拥有多个角色记录，例如同时管理3个品牌
3. **权限验证**: 前端API调用需要JWT认证，建议在中间件中验证用户是否有权限操作指定品牌
4. **数据一致性**: 删除品牌或门店时，建议级联处理相关的user_roles记录

## 测试建议

### 1. 功能测试

- 测试创建品牌管理员
- 测试创建门店管理员
- 测试分页查询
- 测试角色过滤
- 测试状态更新
- 测试删除操作

### 2. 边界测试

- 测试重复分配角色
- 测试不存在的用户/品牌/门店
- 测试门店不属于品牌的情况
- 测试无效的状态值

### 3. 权限测试

- 测试不同角色的访问权限
- 测试跨品牌的操作限制

