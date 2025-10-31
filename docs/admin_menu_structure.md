# 管理员后台菜单结构说明

## 概述

管理员后台菜单按照角色分配不同的访问权限，主要支持三种角色：
- **admin**: 系统管理员（超级管理员）
- **brand_admin**: 品牌管理员
- **store_admin**: 门店管理员

## 菜单结构

### 1. Dashboard (系统管理员专属)
- **路径**: `/admin/dashboard`
- **图标**: DashboardOutlined
- **权限**: `dashboard:view`
- **可见角色**: admin

### 2. 人员管理 (系统管理员专属)
- **路径**: `/admin/staff`
- **图标**: TeamOutlined
- **权限**: `staff:view`
- **可见角色**: admin
- **说明**: 管理所有系统人员

### 3. 品牌管理 (支持多角色)
- **图标**: ShopOutlined
- **权限**: `brands:view`
- **可见角色**: admin, brand_admin

#### 3.1 品牌管理
- **路径**: `/admin/brands`
- **图标**: BankOutlined
- **权限**: `brands:list`
- **可见角色**: admin
- **说明**: 品牌的增删改查

#### 3.2 品牌人员管理
- **路径**: `/admin/brands/staff`
- **图标**: TeamOutlined
- **权限**: `brands:staff`
- **可见角色**: admin, brand_admin
- **说明**: 管理品牌下的管理员和员工

#### 3.3 岗位管理
- **路径**: `/admin/brands/jobs`
- **图标**: SolutionOutlined
- **权限**: `brands:jobs`
- **可见角色**: admin, brand_admin, store_admin
- **说明**: 管理品牌和门店的岗位信息

### 4. 用户管理 (系统管理员专属)
- **路径**: `/admin/users`
- **图标**: UserOutlined
- **权限**: `users:view`
- **可见角色**: admin
- **说明**: 管理所有平台用户（打工者、雇主等）

### 5. 门店管理 (支持多角色)
- **图标**: HomeOutlined
- **权限**: `stores:view`
- **可见角色**: admin, brand_admin, store_admin

#### 5.1 门店列表
- **路径**: `/admin/stores`
- **图标**: ShopOutlined
- **权限**: `stores:list`
- **可见角色**: admin, brand_admin
- **说明**: 门店的增删改查

#### 5.2 门店人员管理
- **路径**: `/admin/stores/staff`
- **图标**: TeamOutlined
- **权限**: `stores:staff`
- **可见角色**: admin, brand_admin, store_admin
- **说明**: 管理门店的员工

#### 5.3 门店岗位
- **路径**: `/admin/stores/jobs`
- **图标**: ProfileOutlined
- **权限**: `stores:jobs`
- **可见角色**: admin, brand_admin, store_admin
- **说明**: 管理门店发布的岗位

### 6. 岗位管理 (支持多角色)
- **图标**: SolutionOutlined
- **权限**: `jobs:view`
- **可见角色**: admin, brand_admin, store_admin

#### 6.1 岗位列表
- **路径**: `/admin/jobs`
- **图标**: UnorderedListOutlined
- **权限**: `jobs:list`
- **可见角色**: admin, brand_admin, store_admin
- **说明**: 查看和管理所有岗位

#### 6.2 应聘管理
- **路径**: `/admin/jobs/applications`
- **图标**: FileDoneOutlined
- **权限**: `jobs:applications`
- **可见角色**: admin, brand_admin, store_admin
- **说明**: 管理岗位应聘信息

#### 6.3 岗位统计
- **路径**: `/admin/jobs/statistics`
- **图标**: BarChartOutlined
- **权限**: `jobs:statistics`
- **可见角色**: admin
- **说明**: 岗位数据统计分析

### 7. 财务管理 (系统管理员专属)
- **图标**: AccountBookOutlined
- **权限**: `finance:view`
- **可见角色**: admin

#### 7.1 结算管理
- **路径**: `/admin/finance/settlements`
- **图标**: TransactionOutlined
- **权限**: `finance:settlements`
- **可见角色**: admin
- **说明**: 管理财务结算

#### 7.2 收入统计
- **路径**: `/admin/finance/income`
- **图标**: DollarOutlined
- **权限**: `finance:income`
- **可见角色**: admin
- **说明**: 收入数据统计

### 8. 消息管理 (系统管理员专属)
- **图标**: MessageOutlined
- **权限**: `messages:view`
- **可见角色**: admin

#### 8.1 系统通知
- **路径**: `/admin/messages/notices`
- **图标**: NotificationOutlined
- **权限**: `messages:notices`
- **可见角色**: admin
- **说明**: 发送和管理系统通知

#### 8.2 消息模板
- **路径**: `/admin/messages/templates`
- **图标**: FileTextOutlined
- **权限**: `messages:templates`
- **可见角色**: admin
- **说明**: 管理消息模板

### 9. 系统管理 (系统管理员专属)
- **图标**: SettingOutlined
- **权限**: `system:view`
- **可见角色**: admin

#### 9.1 管理员管理
- **路径**: `/admin/admins`
- **图标**: UserOutlined
- **权限**: `admin:list`
- **可见角色**: admin
- **说明**: 管理系统管理员账号

#### 9.2 系统配置
- **路径**: `/admin/config`
- **图标**: ControlOutlined
- **权限**: `system:config`
- **可见角色**: admin
- **说明**: 系统参数配置

#### 9.3 系统日志
- **路径**: `/admin/logs`
- **图标**: FileTextOutlined
- **权限**: `system:logs`
- **可见角色**: admin
- **说明**: 查看系统操作日志

## 角色权限总结

### 系统管理员 (admin)
拥有所有菜单的访问权限，包括：
- Dashboard
- 人员管理
- 品牌管理（所有子菜单）
- 用户管理
- 门店管理（所有子菜单）
- 岗位管理（所有子菜单）
- 财务管理（所有子菜单）
- 消息管理（所有子菜单）
- 系统管理（所有子菜单）

### 品牌管理员 (brand_admin)
可以访问：
- 品牌管理组（部分子菜单）：
  - 品牌人员管理
  - 岗位管理
- 门店管理组（部分子菜单）：
  - 门店列表
  - 门店人员管理
  - 门店岗位
- 岗位管理组（部分子菜单）：
  - 岗位列表
  - 应聘管理

### 门店管理员 (store_admin)
可以访问：
- 品牌管理组 > 岗位管理
- 门店管理组（部分子菜单）：
  - 门店人员管理
  - 门店岗位
- 岗位管理组（部分子菜单）：
  - 岗位列表
  - 应聘管理

## 前端集成指南

### 1. 获取菜单接口
```
GET /api/v1/admin/menus
Authorization: Bearer {token}
```

### 2. 响应示例
```json
{
  "base": {
    "code": 200,
    "message": "success"
  },
  "menus": [
    {
      "menu_id": "1",
      "name": "dashboard",
      "label": "Dashboard",
      "path": "/admin/dashboard",
      "icon": "DashboardOutlined",
      "type": "menu",
      "sort_order": 1,
      "visible": true
    },
    {
      "menu_id": "3",
      "name": "brands",
      "label": "品牌管理",
      "icon": "ShopOutlined",
      "type": "group",
      "sort_order": 3,
      "visible": true,
      "children": [...]
    }
  ]
}
```

### 3. 前端渲染示例（React + Ant Design）
```jsx
import { Menu } from 'antd';
import {
  DashboardOutlined,
  TeamOutlined,
  ShopOutlined,
  BankOutlined,
  SolutionOutlined,
  UserOutlined,
  // ... 其他图标
} from '@ant-design/icons';

// 图标映射
const iconMap = {
  DashboardOutlined: <DashboardOutlined />,
  TeamOutlined: <TeamOutlined />,
  ShopOutlined: <ShopOutlined />,
  BankOutlined: <BankOutlined />,
  SolutionOutlined: <SolutionOutlined />,
  UserOutlined: <UserOutlined />,
  // ... 其他映射
};

// 将菜单数据转换为 Ant Design Menu 组件的 items
const convertToMenuItems = (menus) => {
  return menus.map(menu => {
    const item = {
      key: menu.path || menu.name,
      icon: iconMap[menu.icon],
      label: menu.label,
    };
    
    if (menu.children && menu.children.length > 0) {
      item.children = convertToMenuItems(menu.children);
    }
    
    return item;
  });
};

// 使用
const menuItems = convertToMenuItems(menusFromAPI);
<Menu items={menuItems} mode="inline" />;
```

## 注意事项

1. **菜单过滤**: 后端API会根据用户的角色自动过滤菜单，前端无需额外处理
2. **动态加载**: 建议在用户登录后获取菜单配置，缓存在本地存储
3. **权限刷新**: 用户角色变更后需要重新获取菜单配置
4. **图标映射**: 前端需要维护icon字符串到实际图标组件的映射关系

