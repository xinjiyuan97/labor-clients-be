-- ============================================
-- 用户表添加 status 字段迁移脚本
-- 创建时间: 2024-10-22
-- 说明: 为 users 表添加账号状态字段
-- ============================================

-- 第一步：添加 status 字段
-- 说明：添加账号状态字段，支持 active(活跃)、disabled(停用)、pending(待审核) 三种状态
ALTER TABLE users 
ADD COLUMN status ENUM('active', 'disabled', 'pending') NOT NULL DEFAULT 'active' 
COMMENT '账号状态' 
AFTER role;

-- 第二步：为 status 字段添加索引
-- 说明：添加索引以提升按状态查询的性能
ALTER TABLE users 
ADD INDEX idx_status (status);

-- 第三步：将现有用户的状态设置为 active（如果需要）
-- 说明：确保所有现有用户都有明确的状态
UPDATE users 
SET status = 'active' 
WHERE status IS NULL OR status = '';

-- ============================================
-- 验证脚本
-- ============================================

-- 查看表结构，确认字段已添加
DESC users;

-- 查看索引，确认索引已创建
SHOW INDEX FROM users;

-- 统计各状态的用户数量
SELECT status, COUNT(*) as count 
FROM users 
GROUP BY status;

-- ============================================
-- 回滚脚本（如果需要删除该字段）
-- ============================================

-- 警告：执行回滚操作会永久删除 status 字段的数据
-- 请在执行前确保已备份数据

-- 删除索引
-- ALTER TABLE users DROP INDEX idx_status;

-- 删除字段
-- ALTER TABLE users DROP COLUMN status;

