# 门店管理功能实现总结

## 已完成的工作

### 1. 数据库设计 ✅
- ✅ 创建门店表 (`stores`)
- ✅ 创建用户角色关联表 (`user_roles`)
- ✅ 更新岗位表，添加 `store_id` 字段
- ✅ 保持用户表简洁，角色扩展通过独立表实现
- ✅ 创建数据库迁移SQL文件

### 2. Model层 ✅
- ✅ `models/store.go` - 门店模型
- ✅ `models/user_role.go` - 用户角色关联模型
- ✅ 更新 `models/user.go` - 移除了brand_id和store_id字段
- ✅ 更新 `models/job.go` - 添加store_id字段

### 3. DAL层 ✅
- ✅ `dal/mysql/store.go` - 门店CRUD操作
  - CreateStore
  - GetStoreByID
  - GetStoreList
  - UpdateStore
  - DeleteStore
  - GetStoresByBrandID
  - CheckStoreExists
  - CheckStoreBelongsToBrand

- ✅ `dal/mysql/user_role.go` - 用户角色CRUD操作
  - CreateUserRole
  - GetUserRolesByUserID
  - GetUserRoleByID
  - UpdateUserRole
  - DeleteUserRole
  - CheckUserHasBrandAdminRole
  - CheckUserHasStoreAdminRole
  - GetBrandAdminsByBrandID
  - GetStoreAdminsByStoreID
  - DeleteUserRolesByUserID

### 4. Logic层 ✅
- ✅ `biz/logic/admin/create_store.go`
- ✅ `biz/logic/admin/get_store_list.go`
- ✅ `biz/logic/admin/get_store_detail.go`
- ✅ `biz/logic/admin/update_store.go`
- ✅ `biz/logic/admin/delete_store.go`

### 5. Handler层 ✅
- ✅ `biz/handler/admin/create_store.go`
- ✅ `biz/handler/admin/get_store_list.go`
- ✅ `biz/handler/admin/get_store_detail.go`
- ✅ `biz/handler/admin/update_store.go`
- ✅ `biz/handler/admin/delete_store.go`
- ✅ `biz/handler/admin/assign_store_admin.go` (占位实现)
- ✅ `biz/handler/admin/remove_store_admin.go` (占位实现)

### 6. Thrift定义 ✅
- ✅ 更新 `idls/admin.thrift`
  - 添加门店相关的结构体定义
  - 添加门店管理相关的接口
  - 更新用户管理接口

### 7. 路由配置 ✅
- ✅ 更新 `biz/router/admin/middleware.go`
- ✅ 路由已自动生成

### 8. 中间件支持 ✅
- ✅ `middleware/jwt.go` - 添加新角色支持
  - RequireBrandAdminRole
  - RequireStoreAdminRole
  - RequireAnyRole
  - RequireAdminOrBrandAdmin
  - RequireAdminOrBrandAdminOrStoreAdmin
  - GetBrandIDFromContext
  - GetStoreIDFromContext

### 9. JWT增强 ✅
- ✅ `utils/jwt.go` - 支持BrandID和StoreID
  - 更新Claims结构
  - GenerateTokenWithExtra
  - GenerateRefreshTokenWithExtra

### 10. 文档 ✅
- ✅ `docs/store_management.md` - 完整功能说明
- ✅ `migrations/add_store_management.sql` - 数据库迁移脚本

## 需要进一步完善的功能

### 1. 用户角色管理接口 🔲
虽然已经创建了DAL层，但还需要实现以下接口：
- POST /api/v1/admin/user-roles - 创建用户角色关联
- GET /api/v1/admin/user-roles - 获取用户角色列表
- PUT /api/v1/admin/user-roles/:role_id - 更新用户角色
- DELETE /api/v1/admin/user-roles/:role_id - 删除用户角色

需要创建：
- `idls/admin.thrift` 中的用户角色相关定义
- `biz/logic/admin/` 中的用户角色逻辑
- `biz/handler/admin/` 中的用户角色处理器

### 2. 登录逻辑更新 🔲
需要在用户登录时，查询user_roles表，如果用户有品牌管理员或门店管理员角色，将相关信息加入JWT token中。

修改文件：
- `biz/logic/auth/login.go` 或相应的登录处理逻辑

### 3. 权限验证增强 🔲
在涉及品牌和门店的操作中，需要验证用户是否有相应权限：
- 品牌管理员只能管理自己的品牌
- 门店管理员只能管理自己的门店

可能需要在以下地方添加权限检查：
- 门店CRUD操作
- 岗位发布操作
- 数据查询操作

### 4. 测试 🔲
- 单元测试
- 集成测试
- API测试

## 核心设计思想

### 为什么使用独立的user_roles表？

1. **灵活性**：一个用户可以拥有多个角色
   - 同时管理多个品牌
   - 同时管理多个门店
   - 同时是品牌管理员和门店管理员

2. **可扩展性**：未来可以轻松添加新的角色类型
   - 不需要修改用户表结构
   - 新增角色类型只需要在enum中添加

3. **独立性**：每个角色可以独立启用/禁用
   - 不影响用户的基础账号
   - 不影响用户的其他角色

4. **清晰性**：基础用户信息和角色权限分离
   - users表保持简洁
   - 角色信息集中管理

## 使用示例

### 场景1：创建品牌管理员

```bash
# 1. 创建用户账号
POST /api/v1/admin/users
{
  "phone": "13800138000",
  "real_name": "张三",
  "role": "admin",
  "password": "123456"
}
# 返回: { "user_id": 100 }

# 2. 为用户分配品牌管理员角色
POST /api/v1/admin/user-roles
{
  "user_id": 100,
  "role_type": "brand_admin",
  "brand_id": 1,
  "status": "active"
}
```

### 场景2：创建门店管理员

```bash
# 1. 创建用户账号（同上）
POST /api/v1/admin/users
{
  "phone": "13800138001",
  "real_name": "李四",
  "role": "admin",
  "password": "123456"
}
# 返回: { "user_id": 101 }

# 2. 为用户分配门店管理员角色
POST /api/v1/admin/user-roles
{
  "user_id": 101,
  "role_type": "store_admin",
  "brand_id": 1,
  "store_id": 5,
  "status": "active"
}
```

### 场景3：一个用户管理多个品牌

```bash
# 为同一个用户分配多个品牌管理员角色
POST /api/v1/admin/user-roles
{
  "user_id": 100,
  "role_type": "brand_admin",
  "brand_id": 2,
  "status": "active"
}

POST /api/v1/admin/user-roles
{
  "user_id": 100,
  "role_type": "brand_admin",
  "brand_id": 3,
  "status": "active"
}
```

## 数据库迁移

执行以下命令应用数据库变更：

```bash
cd /Users/jiyuanxin/work/src/github.com/xinjiyuan97/labor-clients-be
mysql -u your_username -p your_database < migrations/add_store_management.sql
```

## 下一步建议

1. **实现用户角色管理API**
   - 优先级：高
   - 工作量：中等
   - 依赖：无

2. **更新登录逻辑**
   - 优先级：高
   - 工作量：小
   - 依赖：无

3. **添加权限验证**
   - 优先级：中
   - 工作量：中等
   - 依赖：用户角色管理API

4. **编写测试**
   - 优先级：中
   - 工作量：大
   - 依赖：所有功能完成

## 相关文件清单

```
schemas/
  ├── stores.sql                    # 门店表
  ├── user_roles.sql                # 用户角色关联表
  ├── users.sql                     # 用户表（未改动）
  └── jobs.sql                      # 岗位表（添加store_id）

models/
  ├── store.go                      # 门店模型
  ├── user_role.go                  # 用户角色模型
  ├── user.go                       # 用户模型
  └── job.go                        # 岗位模型

dal/mysql/
  ├── store.go                      # 门店数据访问
  └── user_role.go                  # 用户角色数据访问

biz/logic/admin/
  ├── create_store.go
  ├── get_store_list.go
  ├── get_store_detail.go
  ├── update_store.go
  └── delete_store.go

biz/handler/admin/
  ├── create_store.go
  ├── get_store_list.go
  ├── get_store_detail.go
  ├── update_store.go
  ├── delete_store.go
  ├── assign_store_admin.go         # 占位
  └── remove_store_admin.go         # 占位

idls/
  └── admin.thrift                  # 更新了门店相关定义

middleware/
  └── jwt.go                        # 添加新角色支持

utils/
  └── jwt.go                        # JWT支持BrandID和StoreID

migrations/
  └── add_store_management.sql      # 数据库迁移脚本

docs/
  ├── store_management.md           # 功能说明文档
  └── store_management_summary.md   # 本文档
```

