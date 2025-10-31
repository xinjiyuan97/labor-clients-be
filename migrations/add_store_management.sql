-- 添加门店管理相关的数据库变更

-- 1. 创建门店表
CREATE TABLE IF NOT EXISTS stores (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '门店ID',
    brand_id BIGINT NOT NULL COMMENT '所属品牌ID',
    name VARCHAR(100) NOT NULL COMMENT '门店名称',
    address VARCHAR(255) NOT NULL COMMENT '门店地址',
    latitude DECIMAL(10,8) COMMENT '纬度',
    longitude DECIMAL(11,8) COMMENT '经度',
    contact_phone VARCHAR(20) COMMENT '联系电话',
    contact_person VARCHAR(50) COMMENT '联系人',
    description TEXT COMMENT '门店描述',
    status ENUM('active', 'disabled') NOT NULL DEFAULT 'active' COMMENT '门店状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_brand_id (brand_id),
    INDEX idx_status (status)
) COMMENT '门店信息表';

-- 2. 创建用户角色关联表
CREATE TABLE IF NOT EXISTS user_roles (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '角色关联ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    role_type ENUM('brand_admin', 'store_admin') NOT NULL COMMENT '角色类型',
    brand_id BIGINT NULL COMMENT '关联品牌ID',
    store_id BIGINT NULL COMMENT '关联门店ID',
    status ENUM('active', 'disabled') NOT NULL DEFAULT 'active' COMMENT '角色状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_user_id (user_id),
    INDEX idx_role_type (role_type),
    INDEX idx_brand_id (brand_id),
    INDEX idx_store_id (store_id),
    INDEX idx_status (status),
    UNIQUE KEY unique_user_brand_store (user_id, brand_id, store_id, deleted_at)
) COMMENT '用户角色关联表';

-- 3. 修改岗位表，添加门店ID字段
ALTER TABLE jobs 
    ADD COLUMN store_id BIGINT NULL COMMENT '所属门店ID' AFTER brand_id,
    ADD INDEX idx_store_id (store_id);

