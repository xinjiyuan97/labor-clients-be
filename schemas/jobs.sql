-- 岗位表
CREATE TABLE jobs (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '岗位ID',
    employer_id BIGINT NOT NULL COMMENT '雇主ID',
    brand_id BIGINT NOT NULL COMMENT '所属品牌ID',
    category_id BIGINT NOT NULL COMMENT '分类ID',
    title VARCHAR(100) NOT NULL COMMENT '岗位标题',
    job_type ENUM('standard', 'rush', 'transfer') NOT NULL DEFAULT 'standard' COMMENT '岗位类型',
    description TEXT COMMENT '岗位描述',
    salary DECIMAL(10,2) NOT NULL COMMENT '薪资',
    salary_unit VARCHAR(20) DEFAULT '天' COMMENT '结算单位',
    location VARCHAR(255) NOT NULL COMMENT '工作地点',
    latitude DECIMAL(10,8) COMMENT '纬度',
    longitude DECIMAL(11,8) COMMENT '经度',
    requirements TEXT COMMENT '工作要求',
    benefits TEXT COMMENT '福利待遇',
    start_time DATETIME NOT NULL COMMENT '开始时间',
    end_time DATETIME NOT NULL COMMENT '结束时间',
    status ENUM('draft', 'published', 'filled', 'completed', 'cancelled') DEFAULT 'draft' COMMENT '岗位状态',
    max_applicants INT NOT NULL DEFAULT 1 COMMENT '该岗位最大招募人数',  -- 新增字段
    applicant_count INT DEFAULT 0 COMMENT '报名人数',
    max_applicants INT DEFAULT 1 COMMENT '最大报名人数',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_employer_id (employer_id),
    INDEX idx_brand_id (brand_id),  -- 新增索引
    INDEX idx_category_id (category_id),
    INDEX idx_job_type (job_type),
    INDEX idx_status (status),
    INDEX idx_start_time (start_time)
) COMMENT '岗位信息表';