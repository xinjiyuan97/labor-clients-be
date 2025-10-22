-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
    username VARCHAR(50) NOT NULL COMMENT '用户名',
    phone VARCHAR(20) UNIQUE NOT NULL COMMENT '手机号',
    avatar VARCHAR(255) COMMENT '头像URL',
    password_hash VARCHAR(255) NOT NULL COMMENT '密码哈希',
    role ENUM('worker', 'employer', 'admin') NOT NULL COMMENT '用户角色',
    status ENUM('active', 'disabled', 'pending') NOT NULL DEFAULT 'active' COMMENT '账号状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_phone (phone),
    INDEX idx_role (role),
    INDEX idx_status (status)
) COMMENT '用户基础信息表';