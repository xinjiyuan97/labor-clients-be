-- 岗位标签表
CREATE TABLE job_tags (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '标签ID',
    job_id BIGINT NOT NULL COMMENT '岗位ID',
    tag_name VARCHAR(50) NOT NULL COMMENT '标签名称',
    tag_type VARCHAR(20) NOT NULL COMMENT '标签类型',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_job_id (job_id),
    INDEX idx_tag_type (tag_type)
) COMMENT '岗位标签表';
