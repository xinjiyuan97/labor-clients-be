-- 更新brands表，添加完整的品牌信息字段
-- 执行时间: 2025-10-24

-- 添加基本信息字段
ALTER TABLE brands
ADD COLUMN company_short_name VARCHAR(50) COMMENT '公司简称' AFTER name,
ADD COLUMN website VARCHAR(255) COMMENT '公司网站' AFTER description;

-- 添加公司信息字段
ALTER TABLE brands
ADD COLUMN industry VARCHAR(50) COMMENT '所属行业' AFTER website,
ADD COLUMN company_size VARCHAR(20) COMMENT '公司规模' AFTER industry,
ADD COLUMN credit_code VARCHAR(50) COMMENT '统一社会信用代码' AFTER company_size,
ADD COLUMN company_address VARCHAR(255) COMMENT '公司地址' AFTER credit_code,
ADD COLUMN business_scope TEXT COMMENT '经营范围' AFTER company_address,
ADD COLUMN established_date DATE COMMENT '成立日期' AFTER business_scope,
ADD COLUMN registered_capital DECIMAL(15,2) COMMENT '注册资本' AFTER established_date;

-- 添加联系人信息字段（存储user_id）
ALTER TABLE brands
ADD COLUMN contact_user_id BIGINT COMMENT '联系人用户ID' AFTER registered_capital,
ADD COLUMN contact_position VARCHAR(50) COMMENT '联系人职位' AFTER contact_user_id;

-- 添加证件信息字段
ALTER TABLE brands
ADD COLUMN id_card_number VARCHAR(50) COMMENT '身份证号' AFTER contact_position,
ADD COLUMN id_card_front VARCHAR(255) COMMENT '身份证正面照URL' AFTER id_card_number,
ADD COLUMN id_card_back VARCHAR(255) COMMENT '身份证反面照URL' AFTER id_card_front,
ADD COLUMN business_license VARCHAR(255) COMMENT '营业执照URL' AFTER id_card_back,
ADD COLUMN tax_certificate VARCHAR(255) COMMENT '税务登记证URL' AFTER business_license,
ADD COLUMN org_code_certificate VARCHAR(255) COMMENT '组织机构代码证URL' AFTER tax_certificate,
ADD COLUMN bank_license VARCHAR(255) COMMENT '开户许可证URL' AFTER org_code_certificate,
ADD COLUMN other_certificates TEXT COMMENT '其他证件URL（JSON数组）' AFTER bank_license;

-- 添加财务信息字段
ALTER TABLE brands
ADD COLUMN bank_account VARCHAR(50) COMMENT '银行账号' AFTER other_certificates,
ADD COLUMN settlement_cycle VARCHAR(20) COMMENT '结算周期' AFTER bank_account,
ADD COLUMN deposit_amount DECIMAL(15,2) COMMENT '保证金金额' AFTER settlement_cycle;

-- 修改状态字段
ALTER TABLE brands
MODIFY COLUMN auth_status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending' COMMENT '认证状态';

-- 添加账号状态字段
ALTER TABLE brands
ADD COLUMN account_status ENUM('active', 'disabled', 'frozen') DEFAULT 'active' COMMENT '账号状态' AFTER auth_status;

-- 添加索引
ALTER TABLE brands
ADD INDEX idx_credit_code (credit_code),
ADD INDEX idx_contact_user_id (contact_user_id),
ADD INDEX idx_account_status (account_status);

-- 更新name字段注释
ALTER TABLE brands
MODIFY COLUMN name VARCHAR(100) NOT NULL COMMENT '品牌名称（公司全称）';

-- 说明：
-- 1. contact_user_id 关联到 users 表的 id 字段
-- 2. 联系人的姓名、电话、邮箱等信息从 users 表获取
-- 3. 原有的联系人直接存储字段（contact_person, contact_phone, contact_email）已被废弃，改为通过user_id关联

