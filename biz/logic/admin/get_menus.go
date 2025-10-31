package admin

import (
	"context"
	"errors"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/config"
)

// GetMenus 获取用户可访问的菜单列表
func GetMenus(ctx context.Context, role string) ([]*admin.MenuItem, error) {
	// 获取菜单配置并根据角色过滤
	menuConfig := config.GetMenuConfig()
	if menuConfig == nil {
		return nil, errors.New("menu config not found")
	}

	filteredMenus := menuConfig.FilterMenusByRoles([]string{role})

	// 转换为Thrift结构
	thriftMenus := make([]*admin.MenuItem, 0, len(filteredMenus))
	for _, menu := range filteredMenus {
		thriftMenus = append(thriftMenus, convertToThriftMenuItem(menu))
	}

	return thriftMenus, nil
}

// convertToThriftMenuItem 将配置中的菜单项转换为Thrift结构
func convertToThriftMenuItem(item config.MenuItem) *admin.MenuItem {
	thriftItem := &admin.MenuItem{
		MenuID:     item.MenuID,
		Name:       item.Name,
		Label:      item.Label,
		Path:       item.Path,
		Icon:       item.Icon,
		Type:       item.Type,
		SortOrder:  int32(item.SortOrder),
		Visible:    item.Visible,
		Disabled:   item.Disabled,
		Permission: item.Permission,
	}

	if item.ParentID != nil {
		thriftItem.ParentID = *item.ParentID
	}

	// 递归转换子菜单
	if len(item.Children) > 0 {
		thriftItem.Children = make([]*admin.MenuItem, 0, len(item.Children))
		for _, child := range item.Children {
			thriftItem.Children = append(thriftItem.Children, convertToThriftMenuItem(child))
		}
	}

	return thriftItem
}
