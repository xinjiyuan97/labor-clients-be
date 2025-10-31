# 菜单权限管理 - 快速开始

## 🚀 快速开始

### 1. 应用已集成

菜单权限管理功能已经集成到应用中，无需额外配置即可使用。

### 2. 启动应用

```bash
# 启动服务器
./start.sh

# 或直接运行
./hertz_service
```

启动成功后会看到：
```
菜单配置加载成功，共 7 个顶级菜单
```

### 3. 获取菜单

```bash
# 使用JWT token获取当前用户的菜单
curl -X GET \
  http://localhost:8080/api/v1/admin/menus \
  -H 'Authorization: Bearer YOUR_JWT_TOKEN'
```

## 📋 菜单配置文件

配置文件位置: `conf/menus.yaml`

### 配置示例

```yaml
menus:
  - menu_id: "3"
    name: "brands"
    label: "品牌管理"
    icon: "ShopOutlined"
    parent_id: null
    type: "group"
    sort_order: 3
    visible: true
    disabled: false
    permission: "brands:view"
    roles: ["admin", "brand_admin"]  # 允许admin和brand_admin访问
    children:
      - menu_id: "3-1"
        name: "brand-list"
        label: "品牌列表"
        path: "/admin/brands"
        icon: "BankOutlined"
        parent_id: "3"
        type: "menu"
        sort_order: 1
        visible: true
        disabled: false
        permission: "brands:list"
        roles: ["admin"]  # 仅允许admin访问
```

## 🔑 角色说明

| 角色 | 权限范围 |
|------|---------|
| `admin` | 系统管理员，拥有所有权限 |
| `brand_admin` | 品牌管理员，可管理品牌和门店 |
| `store_admin` | 门店管理员，可管理门店岗位 |
| `employer` | 雇主 |
| `worker` | 打工人 |

## 📝 常见操作

### 添加新菜单

1. 编辑 `conf/menus.yaml`
2. 添加新的菜单项配置
3. 重启服务

示例：
```yaml
- menu_id: "8"
  name: "new-module"
  label: "新功能模块"
  icon: "AppstoreOutlined"
  parent_id: null
  type: "group"
  sort_order: 8
  visible: true
  disabled: false
  permission: "new:view"
  roles: ["admin"]  # 配置可访问的角色
  children: []
```

### 修改菜单权限

只需修改菜单项的 `roles` 字段：

```yaml
# 从仅admin可访问
roles: ["admin"]

# 改为admin和brand_admin都可访问
roles: ["admin", "brand_admin"]
```

### 隐藏某个菜单

设置 `visible: false`:

```yaml
- menu_id: "7"
  name: "messages"
  visible: false  # 隐藏此菜单
```

## 🔍 调试技巧

### 查看用户所有角色

在业务逻辑中，用户的角色包括：
1. 基础角色（来自users.role）
2. 扩展角色（来自user_roles表）

### 查看菜单过滤结果

修改 `biz/logic/admin/get_menus.go`，添加日志：

```go
utils.Infof("用户 %d 的角色: %v", userID, userRoles)
utils.Infof("过滤后菜单数量: %d", len(filteredMenus))
```

## 📖 详细文档

- [完整功能文档](docs/menu_management.md)
- [实现总结](docs/menu_implementation_summary.md)
- [门店管理功能](docs/store_management.md)

## ⚙️ 高级配置

### 自定义图标

支持所有 Ant Design 图标：
- SettingOutlined
- UserOutlined
- ShopOutlined
- BankOutlined
- TeamOutlined
- HomeOutlined
- etc.

### 菜单类型

- `group`: 菜单组（父菜单）
- `menu`: 菜单项（可点击跳转）
- `button`: 按钮（用于页面内操作）

### 权限标识

建议使用 `资源:操作` 格式：
- `brands:view` - 查看权限
- `brands:list` - 列表权限
- `brands:create` - 创建权限
- `brands:edit` - 编辑权限
- `brands:delete` - 删除权限

## 🐛 常见问题

### Q: 修改配置后没有生效？
A: 需要重启服务，菜单配置在启动时加载

### Q: 某些角色看不到菜单？
A: 检查 `roles` 配置，确保包含了该角色

### Q: 用户角色变更后菜单没更新？
A: 用户需要重新登录以刷新JWT token

### Q: 如何让所有人都能访问某个菜单？
A: 将 `roles` 设置为空数组 `[]`

## 📞 技术支持

如有问题，请查看：
1. 应用日志：`logs/app.log`
2. 菜单配置：`conf/menus.yaml`
3. 文档目录：`docs/`

---

**提示**: 前端展示的菜单只是UI控制，后端API仍需要独立进行权限验证！

