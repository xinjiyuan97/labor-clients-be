-- 用户角色关联表
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
    UNIQUE KEY unique_user_brand_store (user_id, brand_id, store_id, deleted_at),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (brand_id) REFERENCES brands(id),
    FOREIGN KEY (store_id) REFERENCES stores(id)
) COMMENT '用户角色关联表';

