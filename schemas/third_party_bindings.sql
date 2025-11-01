-- 第三方账号绑定表
CREATE TABLE IF NOT EXISTS third_party_bindings (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    platform VARCHAR(50) NOT NULL COMMENT '第三方平台',
    openid VARCHAR(255) NOT NULL COMMENT '第三方平台OpenID',
    unionid VARCHAR(255) COMMENT '第三方平台UnionID',
    appid VARCHAR(100) NOT NULL COMMENT '应用AppID',
    nickname VARCHAR(100) COMMENT '第三方平台昵称',
    avatar VARCHAR(500) COMMENT '第三方平台头像',
    status ENUM('active', 'disabled') NOT NULL DEFAULT 'active' COMMENT '绑定状态',
    last_login_at DATETIME COMMENT '最后登录时间',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_user_id (user_id),
    INDEX idx_platform (platform),
    UNIQUE KEY idx_platform_openid (platform, openid),
    INDEX idx_unionid (unionid),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) COMMENT '第三方账号绑定表';

