package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// MenuItem 菜单项
type MenuItem struct {
	MenuID     string     `yaml:"menu_id" json:"menu_id"`
	Name       string     `yaml:"name" json:"name"`
	Label      string     `yaml:"label" json:"label"`
	Path       string     `yaml:"path,omitempty" json:"path,omitempty"`
	Icon       string     `yaml:"icon" json:"icon"`
	ParentID   *string    `yaml:"parent_id" json:"parent_id"`
	Type       string     `yaml:"type" json:"type"` // group, menu, button
	SortOrder  int        `yaml:"sort_order" json:"sort_order"`
	Visible    bool       `yaml:"visible" json:"visible"`
	Disabled   bool       `yaml:"disabled" json:"disabled"`
	Permission string     `yaml:"permission" json:"permission"`
	Roles      []string   `yaml:"roles" json:"roles"` // 允许访问的角色
	Children   []MenuItem `yaml:"children,omitempty" json:"children,omitempty"`
}

// MenuConfig 菜单配置
type MenuConfig struct {
	Menus []MenuItem `yaml:"menus"`
}

var menuConfig *MenuConfig

// LoadMenuConfig 加载菜单配置
func LoadMenuConfig(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	menuConfig = &MenuConfig{}
	if err := yaml.Unmarshal(data, menuConfig); err != nil {
		return err
	}

	log.Printf("菜单配置加载成功，共 %d 个顶级菜单", len(menuConfig.Menus))
	return nil
}

// GetMenuConfig 获取菜单配置
func GetMenuConfig() *MenuConfig {
	return menuConfig
}

// FilterMenusByRoles 根据角色过滤菜单
func (mc *MenuConfig) FilterMenusByRoles(roles []string) []MenuItem {
	if mc == nil {
		return []MenuItem{}
	}

	result := make([]MenuItem, 0)
	for _, menu := range mc.Menus {
		if filteredMenu := filterMenuItemByRoles(menu, roles); filteredMenu != nil {
			result = append(result, *filteredMenu)
		}
	}
	return result
}

// filterMenuItemByRoles 递归过滤菜单项
func filterMenuItemByRoles(item MenuItem, userRoles []string) *MenuItem {
	// 检查当前菜单项是否允许任一用户角色访问
	if !hasAnyRole(item.Roles, userRoles) {
		return nil
	}

	// 如果有子菜单，递归过滤
	if len(item.Children) > 0 {
		filteredChildren := make([]MenuItem, 0)
		for _, child := range item.Children {
			if filteredChild := filterMenuItemByRoles(child, userRoles); filteredChild != nil {
				filteredChildren = append(filteredChildren, *filteredChild)
			}
		}
		item.Children = filteredChildren

		// 如果是group类型，且过滤后没有子菜单，则不显示
		if item.Type == "group" && len(filteredChildren) == 0 {
			return nil
		}
	}

	return &item
}

// hasAnyRole 检查用户是否拥有任一所需角色
func hasAnyRole(requiredRoles []string, userRoles []string) bool {
	if len(requiredRoles) == 0 {
		return true // 没有角色限制，所有人可访问
	}

	roleMap := make(map[string]bool)
	for _, role := range userRoles {
		roleMap[role] = true
	}

	for _, required := range requiredRoles {
		if roleMap[required] {
			return true
		}
	}
	return false
}
