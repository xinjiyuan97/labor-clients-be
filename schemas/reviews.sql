-- 评价表
CREATE TABLE IF NOT EXISTS reviews (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '评价ID',
    job_id BIGINT NOT NULL COMMENT '岗位ID',
    employer_id BIGINT NOT NULL COMMENT '雇主ID',
    worker_id BIGINT NOT NULL COMMENT '零工ID',
    rating TINYINT NOT NULL CHECK (rating >= 1 AND rating <= 5) COMMENT '评分(1-5)',
    content TEXT COMMENT '评价内容',
    review_type ENUM('employer_to_worker', 'worker_to_employer') NOT NULL COMMENT '评价类型',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '评价时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_employer_id (employer_id),
    INDEX idx_worker_id (worker_id),
    INDEX idx_rating (rating)
) COMMENT '评价表';
