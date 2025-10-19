-- 消息表
CREATE TABLE messages (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '消息ID',
    from_user BIGINT NOT NULL COMMENT '发送用户ID',
    to_user BIGINT NOT NULL COMMENT '接收用户ID',
    message_type VARCHAR(20) NOT NULL COMMENT '消息类型',
    content TEXT NOT NULL COMMENT '消息内容',
    msg_category ENUM('system', 'chat', 'community') DEFAULT 'chat' COMMENT '消息分类',
    is_read BOOLEAN DEFAULT FALSE COMMENT '是否已读',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '发送时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_from_user (from_user),
    INDEX idx_to_user (to_user),
    INDEX idx_created_at (created_at),
    INDEX idx_msg_category (msg_category)
) COMMENT '消息表';