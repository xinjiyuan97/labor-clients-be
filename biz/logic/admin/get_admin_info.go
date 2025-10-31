package admin

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/config"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/middleware"
	"github.com/xinjiyuan97/labor-clients/utils"
)

func CheckAdminRole(ctx context.Context, role string) *common.BaseResp {
	if utils.Contains([]string{"admin", "employer"}, role) {
		return &common.BaseResp{
			Code:    200,
			Message: "success",
		}
	}
	return &common.BaseResp{
		Code:    403,
		Message: "权限不足",
	}
}

// GetAdminInfo 获取管理员信息（包括用户信息、角色列表和菜单）
func GetAdminInfo(ctx context.Context, c *app.RequestContext, req *admin.GetAdminInfoReq) (*admin.GetAdminInfoResp, error) {
	// 1. 获取用户ID和基础角色
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		return &admin.GetAdminInfoResp{
			Base: &common.BaseResp{
				Code:    401,
				Message: "未登录",
			},
		}, nil
	}

	baseRole, _ := middleware.GetUserRoleFromContext(c)

	// 2. 获取用户基本信息
	user, err := mysql.GetUserByID(mysql.DB, userID)
	if err != nil || user == nil {
		return &admin.GetAdminInfoResp{
			Base: &common.BaseResp{
				Code:    404,
				Message: "用户不存在",
			},
		}, nil
	}

	// 3. 收集所有角色信息（临时逻辑：一个账号只能绑定一个品牌）
	roles := make([]*admin.AdminRoleInfo, 0)
	userRoles := []string{baseRole}
	var uniqueBrandID int64 = 0

	// 添加基础角色
	roles = append(roles, &admin.AdminRoleInfo{
		RoleType: baseRole,
	})

	// 4. 查询扩展角色（brand_admin, store_admin）
	// 临时逻辑：只获取第一个激活的品牌角色
	if baseRole != "admin" {
		roleRecords, err := mysql.GetUserRolesByUserID(ctx, userID)
		if err == nil && len(roleRecords) > 0 {
			for _, record := range roleRecords {
				if record.Status == "active" && record.BrandID != nil {
					// 临时逻辑：只保留第一个品牌的所有角色
					if uniqueBrandID == 0 {
						uniqueBrandID = *record.BrandID
					} else if uniqueBrandID != *record.BrandID {
						// 跳过其他品牌的角色
						continue
					}

					userRoles = append(userRoles, record.RoleType)

					// 构建角色信息
					roleInfo := &admin.AdminRoleInfo{
						RoleType: record.RoleType,
						RoleID:   record.ID,
					}

					// 获取品牌信息
					brand, err := mysql.GetBrandByID(mysql.DB, *record.BrandID)
					if err == nil && brand != nil {
						roleInfo.BrandID = *record.BrandID
						roleInfo.BrandName = brand.Name
					}

					// 获取门店信息
					if record.StoreID != nil {
						store, err := mysql.GetStoreByID(ctx, *record.StoreID)
						if err == nil && store != nil {
							roleInfo.StoreID = *record.StoreID
							roleInfo.StoreName = store.Name
						}
					}

					roles = append(roles, roleInfo)
				}
			}
		}
	}

	// 5. 获取菜单配置并根据角色过滤
	menuConfig := config.GetMenuConfig()
	if menuConfig == nil {
		return &admin.GetAdminInfoResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "菜单配置未加载",
			},
		}, nil
	}

	// 6. 返回完整的管理员信息
	return &admin.GetAdminInfoResp{
		Base: &common.BaseResp{
			Code:    200,
			Message: "success",
		},
		UserID:   user.ID,
		Username: user.Username,
		Phone:    user.Phone,
		Avatar:   user.Avatar,
		BaseRole: user.Role,
		Roles:    roles,
	}, nil
}
