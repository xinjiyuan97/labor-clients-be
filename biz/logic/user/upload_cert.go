package user

import (
	"time"

	"gorm.io/gorm"

	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/biz/model/user"
	"github.com/xinjiyuan97/labor-clients/dal/mysql"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// UploadCertLogic 上传证书业务逻辑
func UploadCertLogic(req *user.UploadCertReq, userID int64) (*user.UploadCertResp, error) {
	// 这里应该实现文件上传逻辑，将文件保存到文件服务器或云存储
	// 目前简化处理，直接返回一个模拟的URL
	fileURL := "https://example.com/certs/" + req.CertFile

	// 使用事务保存证书信息
	var certID int64
	err := mysql.Transaction(func(tx *gorm.DB) error {
		// 创建证书记录（这里假设有一个证书表，实际项目中需要根据业务需求设计）
		// 由于没有证书表模型，这里简化处理
		certID = time.Now().Unix() // 使用时间戳作为证书ID

		// 这里可以添加证书信息到数据库的逻辑
		// 例如：创建证书记录，关联用户ID等

		return nil
	})

	if err != nil {
		utils.Errorf("上传证书失败: %v", err)
		return &user.UploadCertResp{
			Base: &common.BaseResp{
				Code:      500,
				Message:   "系统错误",
				Timestamp: time.Now().Format(time.RFC3339),
			},
		}, nil
	}

	return &user.UploadCertResp{
		Base: &common.BaseResp{
			Code:      200,
			Message:   "证书上传成功",
			Timestamp: time.Now().Format(time.RFC3339),
		},
		CertType: req.CertType,
		FileURL:  fileURL,
		CertID:   certID,
	}, nil
}
