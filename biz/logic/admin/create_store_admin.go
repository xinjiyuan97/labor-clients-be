package admin

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
	"gorm.io/gorm"
)

// CreateStoreAdmin 创建门店管理员
func CreateStoreAdmin(ctx context.Context, req *admin.CreateStoreAdminReq) (*admin.CreateStoreAdminResp, error) {
	var user *models.User
	var err error

	// 1. 根据手机号查询用户
	user, err = mysql.GetUserByPhone(ctx, req.Phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		return &admin.CreateStoreAdminResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "查询用户失败",
			},
		}, nil
	}

	// 2. 如果用户不存在，则创建用户
	if err == gorm.ErrRecordNotFound || user == nil {
		// 生成密码哈希
		passwordHash, err := utils.HashPassword(defaultPassword)
		if err != nil {
			return &admin.CreateStoreAdminResp{
				Base: &common.BaseResp{
					Code:    500,
					Message: "密码加密失败",
				},
			}, nil
		}

		// 设置用户名
		username := req.Phone
		if req.RealName != "" {
			username = req.RealName
		}

		// 创建新用户
		user = &models.User{
			Username:     username,
			Phone:        req.Phone,
			PasswordHash: passwordHash,
			Role:         "employer", // 默认角色为employer
			Status:       "active",
		}

		if err := mysql.CreateUser(ctx, user); err != nil {
			return &admin.CreateStoreAdminResp{
				Base: &common.BaseResp{
					Code:    500,
					Message: "创建用户失败",
				},
			}, nil
		}
	}

	// 3. 验证品牌是否存在
	brand, err := mysql.GetBrandByID(mysql.DB, req.BrandID)
	if err != nil || brand == nil {
		return &admin.CreateStoreAdminResp{
			Base: &common.BaseResp{
				Code:    404,
				Message: "品牌不存在",
			},
		}, nil
	}

	// 4. 验证门店是否存在且属于该品牌
	store, err := mysql.GetStoreByID(ctx, req.StoreID)
	if err != nil || store == nil {
		return &admin.CreateStoreAdminResp{
			Base: &common.BaseResp{
				Code:    404,
				Message: "门店不存在",
			},
		}, nil
	}

	if store.BrandID != req.BrandID {
		return &admin.CreateStoreAdminResp{
			Base: &common.BaseResp{
				Code:    400,
				Message: "门店不属于该品牌",
			},
		}, nil
	}

	// 5. 临时逻辑：检查该用户是否已经绑定了任何品牌
	existingRoles, err := mysql.GetUserRolesByUserID(ctx, user.ID)
	if err != nil {
		return &admin.CreateStoreAdminResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "检查角色失败",
			},
		}, err
	}

	// 检查是否已经有激活的品牌角色
	for _, role := range existingRoles {
		if role.Status == "active" && role.BrandID != nil {
			if *role.BrandID != req.BrandID {
				return &admin.CreateStoreAdminResp{
					Base: &common.BaseResp{
						Code:    400,
						Message: "该用户已绑定其他品牌，一个账号只能绑定一个品牌",
					},
				}, nil
			}
			// 如果已经是同一门店的管理员
			if role.RoleType == "store_admin" && role.StoreID != nil && *role.StoreID == req.StoreID {
				return &admin.CreateStoreAdminResp{
					Base: &common.BaseResp{
						Code:    400,
						Message: "该用户已是该门店的管理员",
					},
				}, nil
			}
		}
	}

	// 6. 创建用户角色关联
	userRole := &models.UserRole{
		UserID:   user.ID,
		RoleType: "store_admin",
		BrandID:  &req.BrandID,
		StoreID:  &req.StoreID,
		Status:   "active",
	}

	if err := mysql.CreateUserRole(ctx, userRole); err != nil {
		return &admin.CreateStoreAdminResp{
			Base: &common.BaseResp{
				Code:    500,
				Message: "创建门店管理员失败",
			},
		}, err
	}

	// 7. 返回成功结果
	return &admin.CreateStoreAdminResp{
		Base: &common.BaseResp{
			Code:    200,
			Message: "创建门店管理员成功",
		},
		RoleID: userRole.ID,
	}, nil
}
