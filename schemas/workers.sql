-- 零工信息表
CREATE TABLE workers (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '零工ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    real_name VARCHAR(50) COMMENT '真实姓名',
    gender ENUM('male', 'female') COMMENT '性别',
    age TINYINT UNSIGNED COMMENT '年龄',
    id_card VARCHAR(20) COMMENT '身份证号',
    health_cert VARCHAR(255) COMMENT '健康证URL',
    education VARCHAR(50) COMMENT '学历',
    height DECIMAL(4,1) COMMENT '身高(cm)',
    introduction TEXT COMMENT '个人介绍',
    work_experience TEXT COMMENT '工作经历',
    expected_salary DECIMAL(10,2) COMMENT '期望薪资',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_user_id (user_id),
    INDEX idx_real_name (real_name)
) COMMENT '零工详细信息表';