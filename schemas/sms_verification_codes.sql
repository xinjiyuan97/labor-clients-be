-- 短信验证码表
CREATE TABLE IF NOT EXISTS sms_verification_codes (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    phone VARCHAR(20) NOT NULL COMMENT '手机号',
    code VARCHAR(10) NOT NULL COMMENT '验证码',
    status ENUM('unused', 'used', 'expired') NOT NULL DEFAULT 'unused' COMMENT '状态',
    expires_at DATETIME NOT NULL COMMENT '过期时间',
    used_at DATETIME DEFAULT NULL COMMENT '使用时间',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_phone (phone),
    INDEX idx_phone_code (phone, code),
    INDEX idx_status (status),
    INDEX idx_expires_at (expires_at)
) COMMENT '短信验证码表';

