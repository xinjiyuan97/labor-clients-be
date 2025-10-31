# 获取管理员信息 API 文档

## 概述

`GetAdminInfo` 接口用于获取当前登录管理员的完整信息，包括用户基本信息、所有角色列表、所属品牌/门店信息以及可访问的菜单列表。

此接口是对原有 `GetMenus` 接口的增强版本，提供了更完整的管理员上下文信息，适合在管理后台初始化时调用。

## API 信息

### 请求

**接口**: `GET /api/v1/admin/info`

**认证**: 需要 JWT Token

**请求头**:
```
Authorization: Bearer {token}
```

**查询参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| brand_id | int64 | 否 | 品牌ID。如果提供，则只返回该品牌相关的角色和菜单 |

**请求示例**:
```bash
# 获取所有角色和菜单
GET /api/v1/admin/info

# 获取指定品牌的角色和菜单
GET /api/v1/admin/info?brand_id=10
```

### 响应

**状态码**: `200 OK`

**响应体**:

```json
{
  "base": {
    "code": 200,
    "message": "success"
  },
  "user_id": "123",
  "username": "张三",
  "phone": "13800138000",
  "avatar": "https://example.com/avatar.jpg",
  "base_role": "employer",
  "roles": [
    {
      "role_type": "employer",
      "role_id": "0",
      "brand_id": "0",
      "brand_name": "",
      "store_id": "0",
      "store_name": ""
    },
    {
      "role_type": "brand_admin",
      "role_id": "456",
      "brand_id": "10",
      "brand_name": "某某品牌",
      "store_id": "0",
      "store_name": ""
    },
    {
      "role_type": "store_admin",
      "role_id": "789",
      "brand_id": "10",
      "brand_name": "某某品牌",
      "store_id": "20",
      "store_name": "朝阳门店"
    }
  ],
  "menus": [
    {
      "menu_id": "3",
      "name": "brands",
      "label": "品牌管理",
      "icon": "ShopOutlined",
      "type": "group",
      "sort_order": 3,
      "visible": true,
      "disabled": false,
      "children": [
        {
          "menu_id": "3-2",
          "name": "brand-staff",
          "label": "品牌人员管理",
          "path": "/admin/brands/staff",
          "icon": "TeamOutlined",
          "parent_id": "3",
          "type": "menu",
          "sort_order": 2,
          "visible": true,
          "disabled": false
        }
      ]
    }
  ]
}
```

## 响应字段说明

### 基础信息

| 字段 | 类型 | 说明 |
|------|------|------|
| `user_id` | int64 | 用户ID |
| `username` | string | 用户名 |
| `phone` | string | 手机号 |
| `avatar` | string | 头像URL |
| `base_role` | string | 基础角色（users表的role字段）：admin, employer, worker |

### 角色列表 (roles)

| 字段 | 类型 | 说明 |
|------|------|------|
| `role_type` | string | 角色类型：admin, employer, worker, brand_admin, store_admin |
| `role_id` | int64 | 扩展角色ID（user_roles表的ID，基础角色为0） |
| `brand_id` | int64 | 关联的品牌ID（如果有） |
| `brand_name` | string | 品牌名称（如果有） |
| `store_id` | int64 | 关联的门店ID（如果有） |
| `store_name` | string | 门店名称（如果有） |

### 菜单列表 (menus)

与原有的 `GetMenus` 接口返回的菜单结构相同，根据用户的所有角色进行过滤。

## 使用场景

### 1. 管理后台初始化

在用户登录成功后，前端可以调用此接口获取完整的用户上下文信息：

```javascript
// 登录后获取管理员信息（获取所有角色）
async function initAdminInfo() {
  const response = await fetch('/api/v1/admin/info', {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  });
  
  const data = await response.json();
  
  // 保存用户信息
  setUserInfo({
    userId: data.user_id,
    username: data.username,
    phone: data.phone,
    avatar: data.avatar,
    baseRole: data.base_role
  });
  
  // 保存角色信息
  setRoles(data.roles);
  
  // 渲染菜单
  setMenus(data.menus);
  
  // 如果用户有多个品牌，提供品牌选择器
  const brandRoles = data.roles.filter(r => r.brand_id && r.brand_id !== '0');
  if (brandRoles.length > 1) {
    showBrandSelector(brandRoles);
  }
}
```

### 2. 品牌切换（重点）

如果用户拥有多个品牌的管理员角色，可以通过 `brand_id` 参数切换查看不同品牌的菜单：

```javascript
// 品牌切换功能
async function switchBrand(brandId) {
  const response = await fetch(`/api/v1/admin/info?brand_id=${brandId}`, {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  });
  
  const data = await response.json();
  
  if (data.base.code === 200) {
    // 更新当前选中的品牌
    setCurrentBrand(brandId);
    
    // 更新角色信息（只包含该品牌的角色）
    setRoles(data.roles);
    
    // 更新菜单（只包含该品牌相关的菜单）
    setMenus(data.menus);
    
    // 刷新页面内容
    refreshPageContent();
  } else if (data.base.code === 403) {
    // 没有该品牌的权限
    showError('您不是该品牌的管理员');
  }
}

// 品牌选择器组件
function BrandSelector({ roles }) {
  // 提取所有品牌
  const brands = roles
    .filter(r => r.brand_id && r.brand_id !== '0')
    .reduce((acc, role) => {
      if (!acc.find(b => b.brand_id === role.brand_id)) {
        acc.push({
          brand_id: role.brand_id,
          brand_name: role.brand_name
        });
      }
      return acc;
    }, []);
  
  return (
    <Select onChange={(brandId) => switchBrand(brandId)}>
      <Option value="0">全部品牌</Option>
      {brands.map(brand => (
        <Option key={brand.brand_id} value={brand.brand_id}>
          {brand.brand_name}
        </Option>
      ))}
    </Select>
  );
}
```

### 3. 角色切换的使用流程

```
1. 用户登录成功
   ↓
2. 调用 GET /api/v1/admin/info（不带参数）
   ↓
3. 返回所有角色和完整菜单
   ↓
4. 前端检测到用户有多个品牌角色
   ↓
5. 显示品牌选择器
   ↓
6. 用户选择某个品牌
   ↓
7. 调用 GET /api/v1/admin/info?brand_id=10
   ↓
8. 返回该品牌相关的角色和菜单
   ↓
9. 前端更新显示，只显示该品牌相关的功能
```

### 4. 权限判断

前端可以根据 `roles` 列表进行细粒度的权限控制：

```javascript
function hasPermission(requiredRole) {
  return adminInfo.roles.some(r => r.role_type === requiredRole);
}

// 使用示例
if (hasPermission('brand_admin')) {
  // 显示品牌管理功能
}
```

## 不同角色的响应示例

### 系统管理员 (admin)

```json
{
  "user_id": "1",
  "username": "系统管理员",
  "phone": "13800000000",
  "base_role": "admin",
  "roles": [
    {
      "role_type": "admin"
    }
  ],
  "menus": [
    // 所有菜单
  ]
}
```

### 品牌管理员 (brand_admin)

```json
{
  "user_id": "100",
  "username": "张三",
  "phone": "13800138000",
  "base_role": "employer",
  "roles": [
    {
      "role_type": "employer"
    },
    {
      "role_type": "brand_admin",
      "role_id": "10",
      "brand_id": "5",
      "brand_name": "某某品牌"
    }
  ],
  "menus": [
    // 品牌管理相关菜单
  ]
}
```

### 门店管理员 (store_admin)

```json
{
  "user_id": "200",
  "username": "李四",
  "phone": "13800138001",
  "base_role": "employer",
  "roles": [
    {
      "role_type": "employer"
    },
    {
      "role_type": "store_admin",
      "role_id": "20",
      "brand_id": "5",
      "brand_name": "某某品牌",
      "store_id": "15",
      "store_name": "朝阳门店"
    }
  ],
  "menus": [
    // 门店管理相关菜单
  ]
}
```

### 多角色用户

一个用户可以同时拥有多个品牌/门店的管理员角色：

```json
{
  "user_id": "300",
  "username": "王五",
  "phone": "13800138002",
  "base_role": "employer",
  "roles": [
    {
      "role_type": "employer"
    },
    {
      "role_type": "brand_admin",
      "role_id": "30",
      "brand_id": "5",
      "brand_name": "品牌A"
    },
    {
      "role_type": "brand_admin",
      "role_id": "31",
      "brand_id": "6",
      "brand_name": "品牌B"
    },
    {
      "role_type": "store_admin",
      "role_id": "32",
      "brand_id": "5",
      "brand_name": "品牌A",
      "store_id": "10",
      "store_name": "海淀门店"
    }
  ],
  "menus": [
    // 根据所有角色合并后的菜单
  ]
}
```

## 错误码

| 错误码 | 错误信息 | 说明 |
|--------|----------|------|
| 401 | 未登录 | Token 无效或过期 |
| 403 | 您不是该品牌的管理员 | 提供了 brand_id 但用户不是该品牌的管理员 |
| 404 | 用户不存在 | 用户ID对应的用户记录不存在 |
| 500 | 菜单配置未加载 | 服务器菜单配置文件加载失败 |

## 与 GetMenus 接口的对比

### GetMenus (原接口)

- **路径**: `/api/v1/admin/menus`
- **功能**: 只返回菜单列表
- **适用场景**: 单独刷新菜单

### GetAdminInfo (新接口)

- **路径**: `/api/v1/admin/info`
- **功能**: 返回用户信息 + 角色列表 + 菜单列表
- **适用场景**: 
  - 管理后台初始化
  - 获取完整的用户上下文
  - 实现角色切换功能
  - 前端权限判断

## 推荐使用方式

1. **登录后**: 调用 `GetAdminInfo` 获取完整信息并初始化前端状态
2. **菜单刷新**: 如果只需要刷新菜单，可以继续使用 `GetMenus`
3. **权限验证**: 基于 `GetAdminInfo` 返回的 `roles` 进行前端权限控制

## 前端集成示例（React）

```jsx
import { useState, useEffect } from 'react';
import { Menu, Avatar, Dropdown } from 'antd';

function AdminLayout() {
  const [adminInfo, setAdminInfo] = useState(null);
  
  useEffect(() => {
    // 初始化时获取管理员信息
    fetchAdminInfo();
  }, []);
  
  const fetchAdminInfo = async () => {
    const response = await fetch('/api/v1/admin/info', {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('token')}`
      }
    });
    const data = await response.json();
    setAdminInfo(data);
  };
  
  if (!adminInfo) {
    return <div>Loading...</div>;
  }
  
  return (
    <Layout>
      <Header>
        <Dropdown overlay={
          <Menu>
            <Menu.Item key="profile">个人信息</Menu.Item>
            <Menu.Item key="logout">退出登录</Menu.Item>
          </Menu>
        }>
          <div>
            <Avatar src={adminInfo.avatar} />
            <span>{adminInfo.username}</span>
          </div>
        </Dropdown>
      </Header>
      
      <Sider>
        <Menu items={convertMenus(adminInfo.menus)} />
      </Sider>
      
      <Content>
        {/* 根据角色显示不同内容 */}
        {adminInfo.roles.some(r => r.role_type === 'brand_admin') && (
          <BrandManagement brands={adminInfo.roles.filter(r => r.role_type === 'brand_admin')} />
        )}
      </Content>
    </Layout>
  );
}
```

## 注意事项

1. **缓存策略**: 建议将管理员信息缓存在内存中，避免频繁请求
2. **刷新时机**: 
   - 登录后必须调用
   - 角色变更后需要重新调用
   - Token 刷新后建议重新调用
3. **安全性**: 此接口包含敏感信息，必须进行 JWT 认证
4. **性能**: 一次请求获取所有必要信息，减少前端多次请求

