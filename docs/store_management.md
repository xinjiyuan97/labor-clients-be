# 门店管理功能说明

## 概述

本次更新新增了品牌门店管理功能，支持品牌管理员和门店管理员两种新的角色类型。

## 主要变更

### 1. 数据库表结构变更

#### 新增门店表 (stores)
- `id`: 门店ID
- `brand_id`: 所属品牌ID
- `name`: 门店名称
- `address`: 门店地址
- `latitude`: 纬度
- `longitude`: 经度
- `contact_phone`: 联系电话
- `contact_person`: 联系人
- `description`: 门店描述
- `status`: 门店状态 (active/disabled)
- `created_at`: 创建时间
- `updated_at`: 更新时间
- `deleted_at`: 删除时间

#### 新增用户角色关联表 (user_roles)
- `id`: 角色关联ID
- `user_id`: 用户ID
- `role_type`: 角色类型 (brand_admin/store_admin)
- `brand_id`: 关联品牌ID
- `store_id`: 关联门店ID
- `status`: 角色状态 (active/disabled)
- `created_at`: 创建时间
- `updated_at`: 更新时间
- `deleted_at`: 删除时间

**设计说明**：
- 一个用户可以拥有多个角色（如同时管理多个品牌或多个门店）
- 品牌管理员角色：设置`role_type='brand_admin'`和`brand_id`
- 门店管理员角色：设置`role_type='store_admin'`、`brand_id`和`store_id`
- 支持独立启用/禁用每个角色，不影响用户的其他角色

#### 岗位表 (jobs) 更新
- 新增字段 `store_id`: 所属门店ID

### 2. 角色说明

#### 品牌管理员 (brand_admin)
- 可以创建和管理品牌下的所有门店
- 继承门店管理员的所有权限
- 可以管理品牌下的所有岗位信息

#### 门店管理员 (store_admin)
- 可以发布和下架所属门店的岗位信息
- 只能管理自己门店的岗位
- 不能创建或修改门店信息

### 3. API 接口

#### 门店管理接口

**获取门店列表**
- 接口: `GET /api/v1/admin/stores`
- 参数:
  - `page`: 页码 (必填)
  - `limit`: 每页数量 (必填)
  - `brand_id`: 品牌ID (可选)
  - `status`: 门店状态 (可选)
  - `name`: 门店名称 (可选)

**获取门店详情**
- 接口: `GET /api/v1/admin/stores/:store_id`
- 参数:
  - `store_id`: 门店ID (路径参数)

**创建门店**
- 接口: `POST /api/v1/admin/stores`
- 参数:
  - `brand_id`: 品牌ID (必填)
  - `name`: 门店名称 (必填)
  - `address`: 门店地址 (必填)
  - `latitude`: 纬度 (可选)
  - `longitude`: 经度 (可选)
  - `contact_phone`: 联系电话 (可选)
  - `contact_person`: 联系人 (可选)
  - `description`: 门店描述 (可选)

**更新门店**
- 接口: `PUT /api/v1/admin/stores/:store_id`
- 参数:
  - `store_id`: 门店ID (路径参数)
  - `name`: 门店名称 (可选)
  - `address`: 门店地址 (可选)
  - `latitude`: 纬度 (可选)
  - `longitude`: 经度 (可选)
  - `contact_phone`: 联系电话 (可选)
  - `contact_person`: 联系人 (可选)
  - `description`: 门店描述 (可选)
  - `status`: 门店状态 (可选)

**删除门店**
- 接口: `DELETE /api/v1/admin/stores/:store_id`
- 参数:
  - `store_id`: 门店ID (路径参数)

#### 用户管理接口更新

#### 用户角色管理接口（需要实现）

**创建用户角色关联**
- 接口: `POST /api/v1/admin/user-roles`（待实现）
- 参数:
  - `user_id`: 用户ID (必填)
  - `role_type`: 角色类型，"brand_admin" 或 "store_admin" (必填)
  - `brand_id`: 品牌ID (品牌管理员必填)
  - `store_id`: 门店ID (门店管理员必填)
  - `status`: 角色状态，默认 "active" (可选)

**获取用户的角色列表**
- 接口: `GET /api/v1/admin/user-roles?user_id={user_id}`（待实现）
- 参数:
  - `user_id`: 用户ID (必填)

**更新用户角色**
- 接口: `PUT /api/v1/admin/user-roles/:role_id`（待实现）
- 参数:
  - `role_id`: 角色ID (路径参数)
  - `status`: 角色状态 (可选)

**删除用户角色**
- 接口: `DELETE /api/v1/admin/user-roles/:role_id`（待实现）
- 参数:
  - `role_id`: 角色ID (路径参数)

**注意**: 以上用户角色管理接口需要根据业务需求进一步实现。目前已创建了相应的数据库表和DAL层方法。

### 4. 数据库迁移

执行以下SQL文件来更新数据库结构：
```bash
mysql -u username -p database_name < migrations/add_store_management.sql
```

### 5. 使用流程

#### 创建品牌管理员
1. 系统管理员调用 `POST /api/v1/admin/users` 创建用户账号
2. 调用用户角色接口为该用户分配品牌管理员角色
3. 设置 `role_type` 为 "brand_admin"，指定 `brand_id`

#### 创建门店
1. 品牌管理员或系统管理员调用 `POST /api/v1/admin/stores` 接口
2. 提供门店基本信息

#### 创建门店管理员
1. 品牌管理员或系统管理员调用 `POST /api/v1/admin/users` 创建用户账号
2. 调用用户角色接口为该用户分配门店管理员角色
3. 设置 `role_type` 为 "store_admin"，指定 `brand_id` 和 `store_id`

#### 门店管理员发布岗位
1. 门店管理员调用岗位发布接口
2. 设置 `brand_id` 和 `store_id`
3. 填写岗位详细信息

#### 一个用户管理多个品牌/门店
用户可以拥有多个角色关联记录，例如：
- 同一个用户可以是多个品牌的管理员
- 同一个用户可以是多个门店的管理员
- 同一个用户可以同时是品牌管理员和门店管理员

通过在 `user_roles` 表中创建多条记录即可实现。

### 6. 权限说明

- **系统管理员 (admin)**: 拥有所有权限
- **品牌管理员 (brand_admin)**: 
  - 可以管理所属品牌下的所有门店
  - 可以发布/下架所属品牌下的所有岗位
  - 可以查看所属品牌的所有数据
- **门店管理员 (store_admin)**: 
  - 只能发布/下架所属门店的岗位
  - 只能查看所属门店的数据
  - 不能创建或删除门店

### 7. 注意事项

1. 用户的基础角色 (`users.role`) 仍然是 worker、employer、admin 三种
2. 品牌管理员和门店管理员通过 `user_roles` 表来管理，不影响基础角色
3. 创建门店管理员角色时，必须同时指定 `brand_id` 和 `store_id`
4. 创建品牌管理员角色时，只需要指定 `brand_id`
5. 门店必须属于某个品牌，创建门店时 `brand_id` 为必填项
6. 岗位可以关联到门店，`store_id` 为可选项
7. 删除门店为软删除，不会真正删除数据
8. 门店状态分为 active（活跃）和 disabled（禁用）两种
9. 每个角色关联可以独立启用/禁用，不影响其他角色
10. 一个用户可以拥有多个品牌管理员或门店管理员角色

## 相关文件

### 数据库Schema
- `/schemas/stores.sql` - 门店表结构
- `/schemas/user_roles.sql` - 用户角色关联表结构（新增）
- `/schemas/users.sql` - 用户表结构
- `/schemas/jobs.sql` - 岗位表结构（已更新）

### Model
- `/models/store.go` - 门店模型
- `/models/user_role.go` - 用户角色关联模型（新增）
- `/models/user.go` - 用户模型
- `/models/job.go` - 岗位模型（已更新）

### DAL层
- `/dal/mysql/store.go` - 门店数据访问层
- `/dal/mysql/user_role.go` - 用户角色数据访问层（新增）

### Logic层
- `/biz/logic/admin/create_store.go` - 创建门店逻辑
- `/biz/logic/admin/get_store_list.go` - 获取门店列表逻辑
- `/biz/logic/admin/get_store_detail.go` - 获取门店详情逻辑
- `/biz/logic/admin/update_store.go` - 更新门店逻辑
- `/biz/logic/admin/delete_store.go` - 删除门店逻辑

### Handler层
- `/biz/handler/admin/create_store.go` - 创建门店处理器
- `/biz/handler/admin/get_store_list.go` - 获取门店列表处理器
- `/biz/handler/admin/get_store_detail.go` - 获取门店详情处理器
- `/biz/handler/admin/update_store.go` - 更新门店处理器
- `/biz/handler/admin/delete_store.go` - 删除门店处理器

### Thrift定义
- `/idls/admin.thrift` - Admin服务定义（已更新）

