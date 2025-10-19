-- 支付表
CREATE TABLE payments (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '支付ID',
    job_id BIGINT NOT NULL COMMENT '岗位ID',
    worker_id BIGINT NOT NULL COMMENT '零工ID',
    employer_id BIGINT NOT NULL COMMENT '雇主ID',
    amount DECIMAL(10,2) NOT NULL COMMENT '支付金额',
    payment_method VARCHAR(20) NOT NULL COMMENT '支付方式',
    status ENUM('pending', 'processing', 'completed', 'failed') DEFAULT 'pending' COMMENT '支付状态',
    paid_at DATETIME NULL COMMENT '支付时间',
    platform_fee DECIMAL(10,2) DEFAULT 0 COMMENT '平台费用',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_worker_id (worker_id),
    INDEX idx_employer_id (employer_id),
    INDEX idx_status (status)
) COMMENT '支付记录表';