CREATE TABLE IF NOT EXISTS brands (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '品牌ID',
    name VARCHAR(100) NOT NULL COMMENT '品牌名称',
    logo VARCHAR(255) COMMENT '品牌Logo URL',
    description TEXT COMMENT '品牌描述',
    auth_status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending' COMMENT '认证状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_name (name)
) COMMENT '品牌信息表';