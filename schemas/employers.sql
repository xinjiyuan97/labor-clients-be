CREATE TABLE employers (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '雇主ID',
    user_id BIGINT NOT NULL COMMENT '关联用户ID',
    brand_id BIGINT NOT NULL COMMENT '所属品牌ID',  -- 新增字段
    company_name VARCHAR(100) COMMENT '公司名称（可保留，用于未加盟品牌的独立雇主）',
    contact_person VARCHAR(50) COMMENT '联系人姓名',
    contact_phone VARCHAR(20) COMMENT '联系人手机',
    business_license VARCHAR(100) COMMENT '营业执照号',
    auth_status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending' COMMENT '认证状态',
    auth_time DATETIME NULL COMMENT '认证时间',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_user_id (user_id),
    INDEX idx_brand_id (brand_id),  -- 新增索引
    INDEX idx_auth_status (auth_status)
) COMMENT '雇主信息表';