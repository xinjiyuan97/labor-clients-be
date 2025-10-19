-- 考勤记录表
CREATE TABLE attendance_records (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '考勤ID',
    job_id BIGINT NOT NULL COMMENT '岗位ID',
    worker_id BIGINT NOT NULL COMMENT '零工ID',
    check_in DATETIME NULL COMMENT '打卡时间',
    check_out DATETIME NULL COMMENT '签退时间',
    work_hours DECIMAL(4,2) COMMENT '工作时长',
    check_in_location VARCHAR(255) COMMENT '打卡位置',
    check_out_location VARCHAR(255) COMMENT '签退位置',
    status ENUM('normal', 'late', 'early_leave', 'absent', 'leave') DEFAULT 'normal' COMMENT '考勤状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_job_worker (job_id, worker_id),
    INDEX idx_check_in (check_in)
) COMMENT '考勤记录表';