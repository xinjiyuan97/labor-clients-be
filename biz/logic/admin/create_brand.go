package admin

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/constants"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/errno"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CreateBrandLogic 创建品牌业务逻辑
func CreateBrandLogic(ctx context.Context, req *admin.CreateBrandReq) (*admin.CreateBrandResp, error) {
	// 验证必填字段
	if req.ContactPerson == "" {
		utils.Warn("联系人不能为空")
		return &admin.CreateBrandResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "联系人不能为空",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if req.ContactPhone == "" {
		utils.Warn("联系电话不能为空")
		return &admin.CreateBrandResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "联系电话不能为空",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	if req.ContactEmail == "" {
		utils.Warn("联系邮箱不能为空")
		return &admin.CreateBrandResp{
			Base: &common.BaseResp{
				Code:      400,
				Message:   "联系邮箱不能为空",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	brand := &models.Brand{
		BaseModel: models.BaseModel{
			ID:        utils.GenerateID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:             req.CompanyName,
		CompanyShortName: req.CompanyShortName,
		Logo:             req.Logo,
		Description:      req.Description,
		Website:          req.Website,
		Industry:         req.Industry,
		CompanySize:      req.CompanySize,
		CreditCode:       req.CreditCode,
		CompanyAddress:   req.CompanyAddress,
		BusinessScope:    req.BusinessScope,
		BusinessLicense:  req.BusinessLicense,
		BankAccount:      req.BankAccount,
		AuthStatus:       constants.BrandAuthStatusApproved, // 默认已审核状态
		AccountStatus:    constants.BrandAccountStatusActive,
	}

	// 如果没有提供公司名称，使用简称
	if brand.Name == "" {
		brand.Name = req.CompanyShortName
	}
	// 如果简称也没有，使用默认名称
	if brand.Name == "" {
		brand.Name = "新品牌"
	}
	// 使用事务创建品牌
	err := mysql.Transaction(ctx, func(tx *gorm.DB) error {
		// 构建品牌对象
		ctx = mysql.GetContextWithDB(ctx)
		contactUser, err := createBrandAdminForContact(ctx, req.ContactPhone, req.ContactPerson)
		if err != nil {
			utils.Errorf("创建联系人用户失败: %v", err)
			return err
		}

		brand.ContactUserID = &contactUser.ID
		// 创建品牌
		if err := mysql.CreateBrand(tx, brand); err != nil {
			return err
		}

		_, err = CreateBrandAdmin(ctx, req.ContactPhone, req.ContactPerson, brand.ID, "brand_admin")
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		utils.Errorf("创建品牌失败: %v", err)
		return &admin.CreateBrandResp{
			Base: errno.GetBaseResp(err),
		}, nil
	}

	return &admin.CreateBrandResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "品牌创建成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		BrandID: brand.ID,
	}, err
}

// createBrandAdminForContact 为联系人创建或获取用户
// 返回用户对象和错误
func createBrandAdminForContact(ctx context.Context, phone, realName string) (*models.User, error) {
	// 1. 根据手机号查询用户
	user, err := mysql.GetUserByPhone(ctx, phone)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// 2. 如果用户不存在，则创建用户
	if err == gorm.ErrRecordNotFound || user == nil {
		// 生成密码哈希
		passwordHash, err := utils.HashPassword(defaultPassword)
		if err != nil {
			return nil, err
		}

		// 设置用户名
		username := phone
		if realName != "" {
			username = realName
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
			return nil, err
		}
		utils.Infof("为品牌联系人创建用户成功, phone: %s, user_id: %d", phone, user.ID)
	}

	return user, nil
}

// assignBrandAdminRole 为用户分配品牌管理员角色
// Deprecated: 使用 CreateBrandAdmin 代替
func assignBrandAdminRole(ctx context.Context, userID, brandID int64) error {
	// 检查该用户是否已经是该品牌的管理员
	exists, err := mysql.CheckUserRoleExists(ctx, userID, brandID, "brand_admin", nil)
	if err != nil {
		return err
	}

	if exists {
		utils.Infof("用户已是该品牌的管理员, user_id: %d, brand_id: %d", userID, brandID)
		return nil
	}

	// 创建品牌管理员角色
	userRole := &models.UserRole{
		UserID:   userID,
		RoleType: "brand_admin",
		BrandID:  &brandID,
		Status:   "active",
	}

	if err := mysql.CreateUserRole(ctx, userRole); err != nil {
		return err
	}

	utils.Infof("创建品牌管理员角色成功, user_id: %d, brand_id: %d, role_id: %d", userID, brandID, userRole.ID)
	return nil
}
