-- 岗位申请表
CREATE TABLE job_applications (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '申请ID',
    job_id BIGINT NOT NULL COMMENT '岗位ID',
    worker_id BIGINT NOT NULL COMMENT '零工ID',
    status ENUM('applied', 'confirmed', 'rejected', 'cancelled', 'completed') DEFAULT 'applied' COMMENT '申请状态',
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
    confirmed_at TIMESTAMP NULL COMMENT '确认时间',
    cancel_reason TEXT COMMENT '取消原因',
    worker_rating TINYINT COMMENT '零工评分',
    employer_rating TINYINT COMMENT '雇主评分',
    review TEXT COMMENT '评价内容',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    UNIQUE KEY uk_job_worker (job_id, worker_id),
    INDEX idx_worker_id (worker_id),
    INDEX idx_status (status),
    INDEX idx_applied_at (applied_at)
) COMMENT '岗位申请表';