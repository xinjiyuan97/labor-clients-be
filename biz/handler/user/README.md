# User Handlers 实现说明

## 概述

本目录下的所有user handlers已经完成实现，使用JWT进行鉴权，数据库操作在mysql中实现，事务操作使用gorm的闭包transaction实现。

## 实现的功能

### 1. 更新用户信息 (`update_profile.go`)
- **路由**: `PUT /api/v1/user/profile`
- **功能**: 更新用户基础信息和零工详细信息
- **请求参数**:
  - `username`: 用户名
  - `avatar`: 头像URL
  - `worker_info`: 零工详细信息（可选）
- **响应**: 返回更新后的用户信息和零工信息

### 2. 上传头像 (`upload_avatar.go`)
- **路由**: `POST /api/v1/user/upload-avatar`
- **功能**: 上传用户头像
- **请求参数**:
  - `avatar_file`: 头像文件
- **响应**: 返回头像URL

### 3. 上传证书 (`upload_cert.go`)
- **路由**: `POST /api/v1/user/upload-cert`
- **功能**: 上传用户证书文件
- **请求参数**:
  - `cert_file`: 证书文件
  - `cert_type`: 证书类型
- **响应**: 返回证书URL和证书ID

### 4. 收藏工作 (`favorite_job.go`)
- **路由**: `POST /api/v1/user/favorites`
- **功能**: 收藏工作
- **请求参数**:
  - `job_id`: 工作ID
- **响应**: 返回收藏ID

### 5. 取消收藏 (`unfavorite_job.go`)
- **路由**: `DELETE /api/v1/user/favorites/:job_id`
- **功能**: 取消收藏工作
- **请求参数**:
  - `job_id`: 工作ID (path参数)
- **响应**: 取消收藏成功消息

### 6. 获取收藏列表 (`get_my_favorites.go`)
- **路由**: `GET /api/v1/user/favorites`
- **功能**: 获取我的收藏列表
- **请求参数**:
  - `page_req`: 分页参数
- **响应**: 返回收藏列表和分页信息

### 7. 获取收入统计 (`get_income.go`)
- **路由**: `GET /api/v1/user/income`
- **功能**: 获取收入统计信息
- **请求参数**:
  - `period`: 时间段 (week/month/year/custom)
  - `year`: 年份 (自定义时间段时)
  - `month`: 月份 (自定义时间段时)
- **响应**: 返回总收入、待支付收入、已支付收入和最近支付记录

### 8. 获取收入详情 (`get_income_detail.go`)
- **路由**: `GET /api/v1/user/income/detail`
- **功能**: 获取收入详情列表
- **请求参数**:
  - `page_req`: 分页参数
  - `start_date`: 开始日期
  - `end_date`: 结束日期
  - `status`: 支付状态
- **响应**: 返回支付记录列表和分页信息

### 9. 获取绩效统计 (`get_performance.go`)
- **路由**: `GET /api/v1/user/performance`
- **功能**: 获取用户绩效统计
- **请求参数**:
  - `period`: 时间段
- **响应**: 返回完成工作数、总申请数、成功率和平均评分

## 技术实现

### 数据库操作层 (`dal/mysql/`)
- **user_profile.go**: 用户基础信息和零工信息相关操作
- **user_favorite.go**: 用户收藏工作相关操作
- **user_income.go**: 用户收入相关操作
- **user_performance.go**: 用户绩效相关操作

### 业务逻辑层 (`biz/logic/user/`)
- 分离业务逻辑和HTTP处理
- 使用事务确保数据一致性
- 统一的错误处理和响应格式

### JWT鉴权
- 所有API都需要JWT鉴权
- 从JWT token中获取用户ID
- 使用middleware进行统一的鉴权处理

## 使用示例

### 更新用户信息
```bash
curl -X PUT http://localhost:8888/api/v1/user/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newname",
    "avatar": "https://example.com/avatar.jpg",
    "worker_info": {
      "real_name": "张三",
      "gender": "male",
      "age": 25,
      "education": "本科",
      "height": 175.5,
      "introduction": "有丰富的工作经验",
      "work_experience": "3年工作经验",
      "expected_salary": 5000.0
    }
  }'
```

### 收藏工作
```bash
curl -X POST http://localhost:8888/api/v1/user/favorites \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "job_id": 123
  }'
```

### 获取收入统计
```bash
curl -X GET "http://localhost:8888/api/v1/user/income?period=month" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 获取收藏列表
```bash
curl -X GET "http://localhost:8888/api/v1/user/favorites?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 注意事项

1. **JWT鉴权**: 所有API都需要在Authorization header中提供有效的JWT token
2. **事务支持**: 涉及多表操作的业务逻辑都使用事务确保数据一致性
3. **分页支持**: 列表类API都支持分页查询
4. **错误处理**: 统一的错误响应格式，包含code、message和timestamp字段
5. **文件上传**: 头像和证书上传功能需要实现具体的文件存储逻辑
6. **数据验证**: 所有输入参数都进行了验证
7. **权限控制**: 只有worker角色可以更新零工详细信息
