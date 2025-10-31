-- 门店表
CREATE TABLE IF NOT EXISTS stores (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '门店ID',
    brand_id BIGINT NOT NULL COMMENT '所属品牌ID',
    name VARCHAR(100) NOT NULL COMMENT '门店名称',
    address VARCHAR(255) NOT NULL COMMENT '门店地址',
    latitude DECIMAL(10,8) COMMENT '纬度',
    longitude DECIMAL(11,8) COMMENT '经度',
    contact_phone VARCHAR(20) COMMENT '联系电话',
    contact_person VARCHAR(50) COMMENT '联系人',
    description TEXT COMMENT '门店描述',
    status ENUM('active', 'disabled') NOT NULL DEFAULT 'active' COMMENT '门店状态',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX idx_brand_id (brand_id),
    INDEX idx_status (status),
    FOREIGN KEY (brand_id) REFERENCES brands(id)
) COMMENT '门店信息表';

