package admin

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/admin"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/models"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// CreateBrandLogic 创建品牌业务逻辑
func CreateBrandLogic(req *admin.CreateBrandReq) (*admin.CreateBrandResp, error) {
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

	// 使用事务创建品牌
	var brandID int64
	err := mysql.Transaction(func(tx *gorm.DB) error {
		// 构建品牌对象
		brand := &models.Brand{
			Name:        req.CompanyShortName, // 使用公司简称作为品牌名称
			Logo:        req.Logo,
			Description: req.Description,
			AuthStatus:  "pending", // 默认待审核状态
		}

		// 如果没有提供简称，使用"新品牌"作为默认名称
		if brand.Name == "" {
			brand.Name = "新品牌"
		}

		// 生成唯一ID
		brand.ID = utils.GenerateID()

		// 创建品牌
		if err := mysql.CreateBrand(tx, brand); err != nil {
			return err
		}

		brandID = brand.ID
		return nil
	})

	if err != nil {
		utils.Errorf("创建品牌失败: %v", err)
		return &admin.CreateBrandResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "创建品牌失败",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	utils.Infof("品牌创建成功, brand_id: %d", brandID)
	return &admin.CreateBrandResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "品牌创建成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		BrandID: brandID,
	}, nil
}
