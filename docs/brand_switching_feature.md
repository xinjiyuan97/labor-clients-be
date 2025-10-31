# 品牌切换功能说明

## 功能概述

管理员信息接口 (`GetAdminInfo`) 现在支持通过 `brand_id` 参数来过滤角色和菜单，实现品牌切换功能。

当用户拥有多个品牌的管理员权限时，可以通过切换品牌来查看不同品牌的菜单和功能。

## 核心功能

### 1. 获取所有角色和菜单

**请求**:
```
GET /api/v1/admin/info
```

**说明**: 不提供 `brand_id` 参数时，返回用户的所有角色和完整菜单。

**适用场景**:
- 登录后的初始化
- 显示用户拥有的所有品牌
- 提供品牌选择器

### 2. 获取指定品牌的角色和菜单

**请求**:
```
GET /api/v1/admin/info?brand_id=10
```

**说明**: 提供 `brand_id` 参数时，只返回该品牌相关的角色和菜单。

**适用场景**:
- 用户选择特定品牌后
- 切换品牌上下文
- 只显示特定品牌的功能

## 使用流程

```
┌─────────────────┐
│  用户登录成功   │
└────────┬────────┘
         │
         ▼
┌─────────────────────────────┐
│ GET /api/v1/admin/info      │
│ (不带 brand_id 参数)        │
└────────┬────────────────────┘
         │
         ▼
┌──────────────────────────────┐
│ 返回所有角色和完整菜单       │
│ - employer (基础角色)        │
│ - brand_admin (品牌A)        │
│ - brand_admin (品牌B)        │
│ - store_admin (品牌A-门店1)  │
└────────┬─────────────────────┘
         │
         ▼
┌──────────────────────────────┐
│ 检测到多个品牌角色           │
│ 显示品牌选择器               │
└────────┬─────────────────────┘
         │
         ▼
┌──────────────────────────────┐
│ 用户选择品牌A                │
└────────┬─────────────────────┘
         │
         ▼
┌──────────────────────────────┐
│ GET /api/v1/admin/info       │
│ ?brand_id=10 (品牌A的ID)     │
└────────┬─────────────────────┘
         │
         ▼
┌──────────────────────────────┐
│ 返回品牌A相关的角色和菜单    │
│ - brand_admin (品牌A)        │
│ - store_admin (品牌A-门店1)  │
│ 菜单只包含品牌A相关的内容    │
└────────┬─────────────────────┘
         │
         ▼
┌──────────────────────────────┐
│ 前端更新显示                 │
│ - 当前品牌：品牌A            │
│ - 菜单：品牌A相关功能        │
│ - 页面内容：品牌A的数据      │
└──────────────────────────────┘
```

## 响应对比

### 不带 brand_id（返回所有角色）

```json
{
  "user_id": "100",
  "username": "张三",
  "base_role": "employer",
  "roles": [
    {
      "role_type": "employer"
    },
    {
      "role_type": "brand_admin",
      "role_id": "10",
      "brand_id": "1",
      "brand_name": "品牌A"
    },
    {
      "role_type": "brand_admin",
      "role_id": "11",
      "brand_id": "2",
      "brand_name": "品牌B"
    },
    {
      "role_type": "store_admin",
      "role_id": "12",
      "brand_id": "1",
      "brand_name": "品牌A",
      "store_id": "5",
      "store_name": "朝阳门店"
    }
  ],
  "menus": [
    // 所有角色的菜单（品牌A + 品牌B）
  ]
}
```

### 带 brand_id=1（只返回品牌A的角色）

```json
{
  "user_id": "100",
  "username": "张三",
  "base_role": "employer",
  "roles": [
    {
      "role_type": "brand_admin",
      "role_id": "10",
      "brand_id": "1",
      "brand_name": "品牌A"
    },
    {
      "role_type": "store_admin",
      "role_id": "12",
      "brand_id": "1",
      "brand_name": "品牌A",
      "store_id": "5",
      "store_name": "朝阳门店"
    }
  ],
  "menus": [
    // 只包含品牌A相关的菜单
  ]
}
```

**注意**: 
- 基础角色 `employer` 不会出现在返回结果中（因为指定了 brand_id）
- 品牌B的角色被过滤掉了
- 只返回品牌A的 brand_admin 和 store_admin 角色

## 前端实现示例

### React + Ant Design

```jsx
import { useState, useEffect } from 'react';
import { Select, Layout, Menu } from 'antd';

function AdminLayout() {
  const [adminInfo, setAdminInfo] = useState(null);
  const [currentBrandId, setCurrentBrandId] = useState(0);
  const [availableBrands, setAvailableBrands] = useState([]);

  // 初始化：获取所有角色
  useEffect(() => {
    fetchAdminInfo();
  }, []);

  // 获取管理员信息
  const fetchAdminInfo = async (brandId = 0) => {
    const url = brandId > 0 
      ? `/api/v1/admin/info?brand_id=${brandId}`
      : '/api/v1/admin/info';
    
    const response = await fetch(url, {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('token')}`
      }
    });
    
    const data = await response.json();
    
    if (data.base.code === 200) {
      setAdminInfo(data);
      
      // 如果是第一次加载（获取所有角色），提取可用品牌列表
      if (brandId === 0) {
        const brands = extractBrands(data.roles);
        setAvailableBrands(brands);
      }
    }
  };

  // 提取品牌列表
  const extractBrands = (roles) => {
    const brandsMap = new Map();
    roles.forEach(role => {
      if (role.brand_id && role.brand_id !== '0') {
        brandsMap.set(role.brand_id, {
          brand_id: role.brand_id,
          brand_name: role.brand_name
        });
      }
    });
    return Array.from(brandsMap.values());
  };

  // 切换品牌
  const handleBrandChange = (brandId) => {
    setCurrentBrandId(brandId);
    fetchAdminInfo(brandId);
  };

  if (!adminInfo) {
    return <div>Loading...</div>;
  }

  return (
    <Layout>
      <Layout.Header>
        <div style={{ display: 'flex', alignItems: 'center' }}>
          <span>当前品牌：</span>
          <Select 
            value={currentBrandId} 
            onChange={handleBrandChange}
            style={{ width: 200, marginLeft: 10 }}
          >
            <Select.Option value={0}>全部品牌</Select.Option>
            {availableBrands.map(brand => (
              <Select.Option key={brand.brand_id} value={brand.brand_id}>
                {brand.brand_name}
              </Select.Option>
            ))}
          </Select>
        </div>
      </Layout.Header>

      <Layout.Sider>
        <Menu items={convertMenus(adminInfo.menus)} />
      </Layout.Sider>

      <Layout.Content>
        {/* 页面内容 */}
      </Layout.Content>
    </Layout>
  );
}
```

### Vue 3

```vue
<template>
  <a-layout>
    <a-layout-header>
      <div class="brand-selector">
        <span>当前品牌：</span>
        <a-select 
          v-model:value="currentBrandId" 
          @change="handleBrandChange"
          style="width: 200px; margin-left: 10px"
        >
          <a-select-option :value="0">全部品牌</a-select-option>
          <a-select-option 
            v-for="brand in availableBrands" 
            :key="brand.brand_id"
            :value="brand.brand_id"
          >
            {{ brand.brand_name }}
          </a-select-option>
        </a-select>
      </div>
    </a-layout-header>

    <a-layout-sider>
      <a-menu :items="menuItems" />
    </a-layout-sider>

    <a-layout-content>
      <!-- 页面内容 -->
    </a-layout-content>
  </a-layout>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const adminInfo = ref(null);
const currentBrandId = ref(0);
const availableBrands = ref([]);
const menuItems = ref([]);

onMounted(() => {
  fetchAdminInfo();
});

const fetchAdminInfo = async (brandId = 0) => {
  const url = brandId > 0 
    ? `/api/v1/admin/info?brand_id=${brandId}`
    : '/api/v1/admin/info';
  
  const response = await fetch(url, {
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    }
  });
  
  const data = await response.json();
  
  if (data.base.code === 200) {
    adminInfo.value = data;
    menuItems.value = convertMenus(data.menus);
    
    if (brandId === 0) {
      availableBrands.value = extractBrands(data.roles);
    }
  }
};

const extractBrands = (roles) => {
  const brandsMap = new Map();
  roles.forEach(role => {
    if (role.brand_id && role.brand_id !== '0') {
      brandsMap.set(role.brand_id, {
        brand_id: role.brand_id,
        brand_name: role.brand_name
      });
    }
  });
  return Array.from(brandsMap.values());
};

const handleBrandChange = (brandId) => {
  currentBrandId.value = brandId;
  fetchAdminInfo(brandId);
};
</script>
```

## 注意事项

### 1. 权限验证

当用户切换品牌时，后端会验证该用户是否有该品牌的权限：

- **有权限**: 返回该品牌的角色和菜单
- **无权限**: 返回 403 错误："您不是该品牌的管理员"

### 2. 系统管理员

系统管理员（`admin` 角色）不受 `brand_id` 参数影响，始终返回所有菜单。

### 3. 菜单过滤

菜单是根据过滤后的角色列表动态生成的：

- 不带 `brand_id`: 根据所有角色的权限合并菜单
- 带 `brand_id`: 只根据该品牌相关的角色权限生成菜单

### 4. 状态管理

建议在前端维护以下状态：

- `adminInfo`: 当前的管理员信息
- `currentBrandId`: 当前选中的品牌ID（0表示全部）
- `availableBrands`: 可用的品牌列表（初始化时获取）
- `allRoles`: 用户的所有角色（用于品牌选择器）

## 应用场景

### 1. 多品牌管理平台

一个用户可能同时管理多个品牌，通过品牌切换功能可以：
- 分别查看每个品牌的数据
- 避免菜单混乱
- 提供更清晰的用户体验

### 2. 品牌隔离

确保用户在操作某个品牌时：
- 只能看到该品牌的菜单和功能
- 数据请求时自动带上 brand_id
- 避免误操作其他品牌的数据

### 3. 权限管理

清晰展示用户在每个品牌中的角色：
- 品牌A: brand_admin + store_admin
- 品牌B: store_admin
- 方便用户理解自己的权限范围

## 总结

品牌切换功能通过 `brand_id` 参数实现了灵活的角色和菜单过滤，使得：

1. **用户体验更好**: 多品牌用户可以专注于当前品牌
2. **权限更清晰**: 明确显示用户在每个品牌的角色
3. **功能更安全**: 后端验证品牌权限，防止越权访问
4. **实现更简单**: 前端只需一个下拉选择器即可切换

