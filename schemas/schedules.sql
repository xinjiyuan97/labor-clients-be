-- 日程表
CREATE TABLE schedules (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '日程ID',
    worker_id BIGINT NOT NULL COMMENT '零工ID',
    job_id BIGINT NULL COMMENT '关联岗位ID',
    title VARCHAR(100) NOT NULL COMMENT '日程标题',
    start_time DATETIME NOT NULL COMMENT '开始时间',
    end_time DATETIME NOT NULL COMMENT '结束时间',
    location VARCHAR(255) COMMENT '地点',
    notes TEXT COMMENT '备注',
    status ENUM('pending', 'in_progress', 'completed', 'cancelled') DEFAULT 'pending' COMMENT '状态',
    reminder_minutes INT DEFAULT 15 COMMENT '提前提醒分钟数',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_worker_id (worker_id),
    INDEX idx_start_time (start_time),
    INDEX idx_status (status)
) COMMENT '个人日程表';