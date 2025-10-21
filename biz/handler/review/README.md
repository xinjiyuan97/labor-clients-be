# Review Handlers 使用说明

本目录包含了评价相关的所有API处理器，实现了完整的评价管理功能。

## 已完成的功能

### 1. 创建评价 (`CreateReview`)
- 路由: `POST /api/v1/reviews`
- 功能: 创建新评价
- 参数: job_id, employer_id, rating, content, review_type
- 返回: 评价ID
- 防止重复评价同一个工作

### 2. 获取收到的评价列表 (`GetReceivedReviews`)
- 路由: `GET /api/v1/reviews/received`
- 功能: 获取用户收到的评价列表，支持分页
- 参数: page_req (分页参数), start_date, end_date
- 返回: 评价列表和分页信息
- 显示他人对当前用户的评价

### 3. 获取发出的评价列表 (`GetGivenReviews`)
- 路由: `GET /api/v1/reviews/given`
- 功能: 获取用户发出的评价列表，支持分页
- 参数: page_req (分页参数), start_date, end_date
- 返回: 评价列表和分页信息
- 显示当前用户对他人的评价

### 4. 获取评价详情 (`GetReviewDetail`)
- 路由: `GET /api/v1/reviews/:review_id`
- 功能: 获取单个评价的详细信息
- 参数: review_id
- 返回: 评价详细信息

### 5. 更新评价 (`UpdateReview`)
- 路由: `PUT /api/v1/reviews/:review_id`
- 功能: 更新评价内容
- 参数: review_id, rating, content
- 返回: 操作结果
- 支持修改评分和评价内容

## 技术实现

### 架构设计
- **Handler层**: 处理HTTP请求和响应
- **Logic层**: 业务逻辑处理
- **MySQL层**: 数据库操作

### 数据库操作
- 使用GORM进行数据库操作
- 支持事务处理（使用闭包事务）
- 分页查询优化
- 防重复评价检查

### 特性
- **分页支持**: 评价列表接口支持分页
- **双向评价**: 支持雇主评价零工、零工评价雇主
- **评分范围**: 评分1-5分
- **重复检查**: 防止对同一工作重复评价
- **事务安全**: 使用GORM闭包事务确保数据一致性
- **评价分类**: 按收到/发出分类查询

## 使用示例

### 创建评价
```bash
POST /api/v1/reviews
{
  "job_id": 123,
  "employer_id": 456,
  "rating": 5,
  "content": "工作认真负责，值得推荐！",
  "review_type": "employer_to_worker"
}
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "创建评价成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "review_id": 789
}
```

### 获取收到的评价列表
```bash
GET /api/v1/reviews/received?page=1&limit=10
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "获取收到的评价列表成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "page_resp": {
    "page": 1,
    "limit": 10,
    "total": 25
  },
  "reviews": [
    {
      "review_id": 789,
      "job_id": 123,
      "employer_id": 456,
      "worker_id": 0,
      "rating": 5,
      "content": "工作认真负责，值得推荐！",
      "review_type": "employer_to_worker",
      "created_at": "2024-01-20T10:00:00Z"
    }
  ]
}
```

### 获取发出的评价列表
```bash
GET /api/v1/reviews/given?page=1&limit=10
```

### 获取评价详情
```bash
GET /api/v1/reviews/789
```

### 更新评价
```bash
PUT /api/v1/reviews/789
{
  "rating": 4,
  "content": "更新后的评价内容"
}
```

## 注意事项

1. 所有操作需要用户登录（user_id需要从JWT token中获取）
2. 创建评价时，worker_id需要从JWT token中获取
3. 评价类型支持：
   - `employer_to_worker`: 雇主评价零工
   - `worker_to_employer`: 零工评价雇主
4. 评分范围：1-5分
5. 防止对同一工作重复评价
6. 所有写操作使用事务确保数据一致性
7. 分页参数有默认值：page=1, limit=10

## 数据库模型

### Review 评价表
- `job_id`: 岗位ID
- `employer_id`: 雇主ID
- `worker_id`: 零工ID
- `rating`: 评分(1-5)
- `content`: 评价内容
- `review_type`: 评价类型
- `created_at`: 创建时间
- `updated_at`: 更新时间

## 评价类型说明

- **employer_to_worker**: 雇主评价零工
  - 雇主对完成工作的零工进行评价
  - 评价内容包括工作态度、专业技能、准时性等

- **worker_to_employer**: 零工评价雇主
  - 零工对雇主进行评价
  - 评价内容包括工作环境、薪资待遇、沟通情况等

## API路由总览

- `POST /api/v1/reviews` - 创建评价
- `GET /api/v1/reviews/received` - 获取收到的评价列表
- `GET /api/v1/reviews/given` - 获取发出的评价列表
- `GET /api/v1/reviews/:review_id` - 获取评价详情
- `PUT /api/v1/reviews/:review_id` - 更新评价

## 扩展功能

未来可以考虑添加的功能：
- 评价回复
- 评价举报
- 评价统计（平均分、评价数量）
- 评价标签
- 图片上传
- 评价排序（按时间、按评分）
- 评价过滤（按评分范围）

