package admin

import (
	"context"

	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/errno"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
	"gorm.io/gorm"
)

const defaultPassword = "123456"

// CreateBrandAdmin 创建品牌管理员
func CreateBrandAdmin(ctx context.Context, phone, name string, brandID int64, roleType string) (roleID int64, err error) {
	var user *models.User

	// 1. 根据手机号查询用户
	user, err = mysql.GetUserByPhone(ctx, phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	// 2. 如果用户不存在，则创建用户
	if err == gorm.ErrRecordNotFound || user == nil {
		// 生成密码哈希
		passwordHash, err := utils.HashPassword(defaultPassword)
		if err != nil {
			return 0, errno.NewError(500, "密码加密失败")
		}

		// 设置用户名
		username := phone
		if name != "" {
			username = name
		}

		// 创建新用户
		user = &models.User{
			Username:     username,
			Phone:        phone,
			PasswordHash: passwordHash,
			Role:         "employer", // 默认角色为employer
			Status:       "active",
		}

		if err := mysql.CreateUser(ctx, user); err != nil {
			return 0, errno.NewError(500, "创建用户失败")
		}
	}

	// 4. 临时逻辑：检查该用户是否已经绑定了任何品牌
	existingRoles, err := mysql.GetUserRolesByUserID(ctx, user.ID)
	if err != nil {
		return 0, errno.NewError(500, "检查角色失败")
	}

	// 检查是否已经有激活的品牌角色
	for _, role := range existingRoles {
		if role.Status == "active" && role.BrandID != nil {
			if *role.BrandID != brandID {
				return 0, errno.NewError(400, "该用户已绑定其他品牌，一个账号只能绑定一个品牌")
			}
			// 如果已经是同一品牌的管理员
			if role.RoleType == "brand_admin" {
				return 0, errno.NewError(400, "该用户已是该品牌的管理员")
			}
		}
	}

	// 5. 创建用户角色关联
	userRole := &models.UserRole{
		UserID:   user.ID,
		RoleType: "brand_admin",
		BrandID:  &brandID,
		Status:   "active",
	}

	if err := mysql.CreateUserRole(ctx, userRole); err != nil {
		return 0, errno.NewError(500, "创建品牌管理员失败")
	}

	// 6. 返回成功结果
	return userRole.ID, nil
}
