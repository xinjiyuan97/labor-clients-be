# 单品牌绑定限制（临时逻辑）

## 概述

为了简化业务逻辑，避免过多的复杂度，我们实施了一个临时策略：**一个账号只能绑定一个品牌**。

## 限制说明

### 1. 用户-品牌绑定限制

- ✅ 一个用户只能成为一个品牌的管理员
- ✅ 用户可以在同一个品牌下拥有多个角色（brand_admin + store_admin）
- ❌ 用户不能同时管理多个不同的品牌

### 2. 适用场景

这个限制适用于以下角色：
- `brand_admin`（品牌管理员）
- `store_admin`（门店管理员）

**不受影响的角色：**
- `admin`（系统管理员）- 可以管理所有品牌
- `worker`（工人）
- `employer`（雇主）

## 实现细节

### 1. 创建品牌管理员时的检查

**文件**: `biz/logic/admin/create_brand_admin.go`

```go
// 检查该用户是否已经绑定了任何品牌
existingRoles, err := mysql.GetUserRolesByUserID(ctx, user.ID)

for _, role := range existingRoles {
    if role.Status == "active" && role.BrandID != nil {
        if *role.BrandID != req.BrandID {
            // 返回错误：该用户已绑定其他品牌
            return error("该用户已绑定其他品牌，一个账号只能绑定一个品牌")
        }
    }
}
```

**行为**：
- 如果用户未绑定任何品牌 → 允许创建
- 如果用户已绑定相同品牌 → 检查是否重复角色
- 如果用户已绑定其他品牌 → **拒绝，返回错误**

### 2. 创建门店管理员时的检查

**文件**: `biz/logic/admin/create_store_admin.go`

```go
// 同样的检查逻辑
existingRoles, err := mysql.GetUserRolesByUserID(ctx, user.ID)

for _, role := range existingRoles {
    if role.Status == "active" && role.BrandID != nil {
        if *role.BrandID != req.BrandID {
            return error("该用户已绑定其他品牌，一个账号只能绑定一个品牌")
        }
    }
}
```

**行为**：
- 只能为已绑定该品牌或未绑定任何品牌的用户创建门店管理员
- 不允许跨品牌分配门店管理员

### 3. 获取管理员信息时的处理

**文件**: `biz/logic/admin/get_admin_info.go`

```go
// 只返回第一个品牌的所有角色
var uniqueBrandID int64 = 0

for _, record := range roleRecords {
    if record.Status == "active" && record.BrandID != nil {
        if uniqueBrandID == 0 {
            uniqueBrandID = *record.BrandID
        } else if uniqueBrandID != *record.BrandID {
            continue  // 跳过其他品牌的角色
        }
        // 添加角色信息...
    }
}
```

**行为**：
- 自动返回用户绑定的唯一品牌信息
- 如果存在多个品牌角色（不应该发生），只返回第一个
- 不再需要前端传入`brand_id`参数

### 4. 接口变更

**修改前**：
```thrift
struct GetAdminInfoReq {
    1: i64 brand_id (api.query="brand_id");  // 可选：指定品牌ID
}
```

**修改后**：
```thrift
struct GetAdminInfoReq {
    // 临时逻辑：一个账号只能绑定一个品牌，不再需要brand_id参数
}
```

## 使用示例

### 场景1：创建品牌管理员（成功）

```bash
POST /api/v1/admin/brand-admins

{
  "phone": "13800138000",
  "real_name": "张三",
  "brand_id": 100,
  "role_type": "brand_admin"
}
```

**结果**：✅ 成功创建（假设用户未绑定任何品牌）

### 场景2：尝试绑定到第二个品牌（失败）

```bash
POST /api/v1/admin/brand-admins

{
  "phone": "13800138000",  // 已经是品牌100的管理员
  "real_name": "张三",
  "brand_id": 200,  // 尝试绑定到品牌200
  "role_type": "brand_admin"
}
```

**结果**：❌ 失败
```json
{
  "code": 400,
  "message": "该用户已绑定其他品牌，一个账号只能绑定一个品牌"
}
```

### 场景3：在同一品牌下添加门店管理员（成功）

```bash
POST /api/v1/admin/store-admins

{
  "phone": "13800138000",  // 已经是品牌100的brand_admin
  "real_name": "张三",
  "brand_id": 100,  // 同一品牌
  "store_id": 1001,
  "role_type": "store_admin"
}
```

**结果**：✅ 成功创建（同一品牌下可以有多个角色）

### 场景4：获取管理员信息

```bash
GET /api/v1/admin/info
```

**结果**：✅ 自动返回用户绑定的唯一品牌信息
```json
{
  "code": 200,
  "user_id": 1,
  "username": "张三",
  "phone": "13800138000",
  "base_role": "employer",
  "roles": [
    {
      "role_type": "employer"
    },
    {
      "role_type": "brand_admin",
      "role_id": 10,
      "brand_id": 100,
      "brand_name": "示例品牌"
    },
    {
      "role_type": "store_admin",
      "role_id": 11,
      "brand_id": 100,
      "brand_name": "示例品牌",
      "store_id": 1001,
      "store_name": "北京朝阳店"
    }
  ]
}
```

## 数据一致性保障

### 现有数据

对于已经存在的违反此规则的数据（一个用户有多个品牌角色）：
- `GetAdminInfo`会自动只返回第一个品牌的角色
- 不会影响系统运行
- 建议后续清理不符合规则的数据

### 新数据

从实施此逻辑开始：
- 所有新创建的角色都会经过检查
- 保证不会产生新的违规数据

## 优势

1. **简化业务逻辑**
   - 前端不需要实现品牌切换功能
   - 减少状态管理的复杂度
   - 降低用户混淆的可能性

2. **减少安全风险**
   - 避免跨品牌的权限混乱
   - 更清晰的数据隔离

3. **提升性能**
   - 减少查询和过滤逻辑
   - 降低接口复杂度

## 未来规划

如果未来需要支持一个账号管理多个品牌，需要进行以下改造：

1. **移除限制检查**
   - 删除`create_brand_admin.go`和`create_store_admin.go`中的品牌限制检查

2. **恢复brand_id参数**
   - 在`GetAdminInfoReq`中恢复`brand_id`参数
   - 前端实现品牌切换逻辑

3. **增强JWT claims**
   - 添加当前选中的brand_id到token中
   - 中间件支持brand上下文切换

4. **前端改造**
   - 添加品牌选择器组件
   - 实现品牌切换后的数据刷新

## 相关文件

### 核心逻辑文件
- `biz/logic/admin/create_brand_admin.go` - 品牌管理员创建
- `biz/logic/admin/create_store_admin.go` - 门店管理员创建
- `biz/logic/admin/get_admin_info.go` - 管理员信息获取

### 接口定义
- `idls/admin.thrift` - Thrift接口定义

### 数据库
- `schemas/user_roles.sql` - 用户角色表
- `models/user_role.go` - 用户角色模型

## 注意事项

⚠️ **重要提醒**：

1. 这是一个**临时策略**，是为了快速上线而采取的简化方案
2. 在代码中用注释标记了"临时逻辑"，便于未来识别和修改
3. 如果需要支持多品牌，请参考"未来规划"章节
4. 不要删除现有的多品牌支持基础设施（如user_roles表的设计）

## 更新日志

- 2025-10-24: 初始版本，实施单品牌绑定限制

