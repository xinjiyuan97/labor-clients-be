CREATE TABLE IF NOT EXISTS brands (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '品牌ID',
    
    -- 基本信息
    name VARCHAR(100) NOT NULL COMMENT '品牌名称（公司全称）',
    company_short_name VARCHAR(50) COMMENT '公司简称',
    logo VARCHAR(255) COMMENT '品牌Logo URL',
    description TEXT COMMENT '品牌描述',
    website VARCHAR(255) COMMENT '公司网站',
    
    -- 公司信息
    industry VARCHAR(50) COMMENT '所属行业',
    company_size VARCHAR(20) COMMENT '公司规模',
    credit_code VARCHAR(50) COMMENT '统一社会信用代码',
    company_address VARCHAR(255) COMMENT '公司地址',
    business_scope TEXT COMMENT '经营范围',
    established_date DATE COMMENT '成立日期',
    registered_capital DECIMAL(15,2) COMMENT '注册资本',
    
    -- 联系人信息（存储user_id）
    contact_user_id BIGINT COMMENT '联系人用户ID',
    contact_position VARCHAR(50) COMMENT '联系人职位',
    
    -- 证件信息
    id_card_number VARCHAR(50) COMMENT '身份证号',
    id_card_front VARCHAR(255) COMMENT '身份证正面照URL',
    id_card_back VARCHAR(255) COMMENT '身份证反面照URL',
    business_license VARCHAR(255) COMMENT '营业执照URL',
    tax_certificate VARCHAR(255) COMMENT '税务登记证URL',
    org_code_certificate VARCHAR(255) COMMENT '组织机构代码证URL',
    bank_license VARCHAR(255) COMMENT '开户许可证URL',
    other_certificates TEXT COMMENT '其他证件URL（JSON数组）',
    
    -- 财务信息
    bank_account VARCHAR(50) COMMENT '银行账号',
    settlement_cycle VARCHAR(20) COMMENT '结算周期',
    deposit_amount DECIMAL(15,2) COMMENT '保证金金额',
    
    -- 状态信息
    auth_status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending' COMMENT '认证状态',
    account_status ENUM('active', 'disabled', 'frozen') DEFAULT 'active' COMMENT '账号状态',
    
    -- 时间戳
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    
    -- 索引
    INDEX idx_name (name),
    INDEX idx_credit_code (credit_code),
    INDEX idx_contact_user_id (contact_user_id),
    INDEX idx_auth_status (auth_status),
    INDEX idx_account_status (account_status)
) COMMENT '品牌信息表';
