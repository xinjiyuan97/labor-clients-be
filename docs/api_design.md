# 零工APP接口设计文档

基于PRD文档和数据库schema，设计完整的API接口体系。

## 1. 用户认证模块

### 1.1 用户注册
- **接口**: `POST /api/v1/auth/register`
- **功能**: 用户注册
- **输入参数**:
```json
{
  "phone": "13800138000",           // 手机号，对应 users.phone
  "password": "123456",             // 密码
  "role": "worker",                 // 角色，对应 users.role
  "username": "张三"                // 用户名，对应 users.username
}
```
- **输出参数**:
```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user_id": 123456789,           // 对应 users.id
    "token": "jwt_token_string",
    "expires_at": "2024-01-01T12:00:00Z"
  }
}
```

### 1.2 用户登录
- **接口**: `POST /api/v1/auth/login`
- **输入参数**:
```json
{
  "phone": "13800138000",           // 对应 users.phone
  "password": "123456"
}
```
- **输出参数**: 同注册接口

### 1.3 获取用户信息
- **接口**: `GET /api/v1/user/profile`
- **输出参数**:
```json
{
  "code": 200,
  "data": {
    "user_id": 123456789,           // users.id
    "username": "张三",              // users.username
    "phone": "138****8000",         // users.phone (脱敏)
    "avatar": "http://...",         // users.avatar
    "role": "worker",               // users.role
    "worker_info": {                // workers表信息
      "real_name": "张三",           // workers.real_name
      "gender": "male",             // workers.gender
      "age": 25,                    // workers.age
      "education": "本科",           // workers.education
      "introduction": "..."         // workers.introduction
    }
  }
}
```

## 2. 日程管理模块

### 2.1 创建日程
- **接口**: `POST /api/v1/schedules`
- **功能**: 创建个人日程
- **输入参数**:
```json
{
  "title": "早班快递",              // schedules.title
  "job_id": 123456,                // schedules.job_id (可选)
  "start_time": "2024-01-01T09:00:00Z",  // schedules.start_time
  "end_time": "2024-01-01T17:00:00Z",    // schedules.end_time
  "location": "北京市朝阳区",        // schedules.location
  "notes": "需要自带工具",           // schedules.notes
  "reminder_minutes": 15           // schedules.reminder_minutes
}
```
- **输出参数**:
```json
{
  "code": 200,
  "data": {
    "schedule_id": 789012,         // schedules.id
    "title": "早班快递",
    "status": "pending"            // schedules.status
  }
}
```

### 2.2 获取今日排班
- **接口**: `GET /api/v1/schedules/today`
- **输出参数**:
```json
{
  "code": 200,
  "data": [
    {
      "schedule_id": 789012,       // schedules.id
      "title": "早班快递",          // schedules.title
      "start_time": "2024-01-01T09:00:00Z",  // schedules.start_time
      "end_time": "2024-01-01T17:00:00Z",    // schedules.end_time
      "location": "北京市朝阳区",    // schedules.location
      "status": "pending",         // schedules.status
      "job_info": {                // 关联岗位信息
        "job_id": 123456,          // jobs.id
        "employer_name": "顺丰快递", // employers.company_name
        "contact_phone": "138..."   // employers.contact_phone
      }
    }
  ]
}
```

### 2.3 更新日程状态
- **接口**: `PUT /api/v1/schedules/{schedule_id}/status`
- **输入参数**:
```json
{
  "status": "in_progress"          // schedules.status
}
```

## 3. 岗位推荐模块

### 3.1 获取岗位列表
- **接口**: `GET /api/v1/jobs`
- **查询参数**:
```
category_id=1&job_type=standard&distance=5&salary_min=100&salary_max=500&page=1&limit=20
```
- **输出参数**:
```json
{
  "code": 200,
  "data": {
    "total": 100,
    "page": 1,
    "limit": 20,
    "jobs": [
      {
        "job_id": 123456,          // jobs.id
        "title": "快递分拣员",       // jobs.title
        "job_type": "standard",    // jobs.job_type
        "salary": 200.00,          // jobs.salary
        "salary_unit": "天",        // jobs.salary_unit
        "location": "北京市朝阳区",  // jobs.location
        "distance": 2.5,           // 计算得出
        "start_time": "2024-01-01T09:00:00Z",  // jobs.start_time
        "end_time": "2024-01-01T17:00:00Z",    // jobs.end_time
        "applicant_count": 5,      // jobs.applicant_count
        "max_applicants": 10,      // jobs.max_applicants
        "brand_info": {            // brands表信息
          "brand_id": 789,         // brands.id
          "name": "顺丰快递",       // brands.name
          "logo": "http://...",    // brands.logo
          "auth_status": "approved" // brands.auth_status
        },
        "employer_info": {         // employers表信息
          "company_name": "顺丰快递北京分公司", // employers.company_name
          "contact_person": "李经理", // employers.contact_person
          "auth_status": "approved" // employers.auth_status
        },
        "tags": [                  // job_tags表信息
          {
            "tag_name": "日结",     // job_tags.tag_name
            "tag_type": "benefit"  // job_tags.tag_type
          }
        ]
      }
    ]
  }
}
```

### 3.2 获取岗位详情
- **接口**: `GET /api/v1/jobs/{job_id}`
- **输出参数**:
```json
{
  "code": 200,
  "data": {
    "job_id": 123456,              // jobs.id
    "title": "快递分拣员",          // jobs.title
    "description": "负责快递分拣工作", // jobs.description
    "job_type": "standard",        // jobs.job_type
    "salary": 200.00,              // jobs.salary
    "salary_unit": "天",            // jobs.salary_unit
    "location": "北京市朝阳区",      // jobs.location
    "latitude": 39.9042,           // jobs.latitude
    "longitude": 116.4074,         // jobs.longitude
    "requirements": "身体健康，能吃苦耐劳", // jobs.requirements
    "benefits": "包餐，交通补贴",    // jobs.benefits
    "start_time": "2024-01-01T09:00:00Z",  // jobs.start_time
    "end_time": "2024-01-01T17:00:00Z",    // jobs.end_time
    "status": "published",         // jobs.status
    "applicant_count": 5,          // jobs.applicant_count
    "max_applicants": 10,          // jobs.max_applicants
    "brand_info": {                // brands表信息
      "brand_id": 789,             // brands.id
      "name": "顺丰快递",           // brands.name
      "logo": "http://...",        // brands.logo
      "description": "..."         // brands.description
    },
    "employer_info": {             // employers表信息
      "company_name": "顺丰快递北京分公司", // employers.company_name
      "contact_person": "李经理",   // employers.contact_person
      "contact_phone": "138..."    // employers.contact_phone
    },
    "category_info": {             // job_categories表信息
      "category_id": 1,            // job_categories.id
      "name": "快递物流"           // job_categories.name
    },
    "tags": [                      // job_tags表信息
      {
        "tag_name": "日结",
        "tag_type": "benefit"
      }
    ]
  }
}
```

### 3.3 获取岗位分类
- **接口**: `GET /api/v1/job-categories`
- **输出参数**:
```json
{
  "code": 200,
  "data": [
    {
      "category_id": 1,            // job_categories.id
      "name": "快递物流",           // job_categories.name
      "description": "快递分拣、配送等工作", // job_categories.description
      "parent_id": 0,              // job_categories.parent_id
      "sort_order": 1              // job_categories.sort_order
    }
  ]
}
```

## 4. 岗位申请模块

### 4.1 申请岗位
- **接口**: `POST /api/v1/job-applications`
- **输入参数**:
```json
{
  "job_id": 123456                 // job_applications.job_id
}
```
- **输出参数**:
```json
{
  "code": 200,
  "data": {
    "application_id": 456789,      // job_applications.id
    "job_id": 123456,              // job_applications.job_id
    "status": "applied",           // job_applications.status
    "applied_at": "2024-01-01T10:00:00Z"  // job_applications.applied_at
  }
}
```

### 4.2 获取我的申请
- **接口**: `GET /api/v1/job-applications/my`
- **查询参数**: `status=applied&page=1&limit=20`
- **输出参数**:
```json
{
  "code": 200,
  "data": {
    "total": 50,
    "applications": [
      {
        "application_id": 456789,  // job_applications.id
        "job_id": 123456,          // job_applications.job_id
        "status": "applied",       // job_applications.status
        "applied_at": "2024-01-01T10:00:00Z",  // job_applications.applied_at
        "confirmed_at": null,      // job_applications.confirmed_at
        "job_info": {              // 关联岗位信息
          "title": "快递分拣员",
          "salary": 200.00,
          "location": "北京市朝阳区",
          "start_time": "2024-01-01T09:00:00Z"
        }
      }
    ]
  }
}
```

### 4.3 取消申请
- **接口**: `PUT /api/v1/job-applications/{application_id}/cancel`
- **输入参数**:
```json
{
  "cancel_reason": "时间冲突"      // job_applications.cancel_reason
}
```

## 5. 消息模块

### 5.1 获取消息列表
- **接口**: `GET /api/v1/messages`
- **查询参数**: `category=system&page=1&limit=20`
- **输出参数**:
```json
{
  "code": 200,
  "data": {
    "total": 100,
    "messages": [
      {
        "message_id": 789012,      // messages.id
        "from_user": 123456,       // messages.from_user
        "to_user": 654321,         // messages.to_user
        "message_type": "text",    // messages.message_type
        "content": "您的申请已通过", // messages.content
        "msg_category": "system",  // messages.msg_category
        "is_read": false,          // messages.is_read
        "created_at": "2024-01-01T10:00:00Z"  // messages.created_at
      }
    ]
  }
}
```

### 5.2 发送消息
- **接口**: `POST /api/v1/messages`
- **输入参数**:
```json
{
  "to_user": 654321,               // messages.to_user
  "message_type": "text",          // messages.message_type
  "content": "您好，我想了解一下这个岗位的详情", // messages.content
  "msg_category": "chat"           // messages.msg_category
}
```

### 5.3 标记消息已读
- **接口**: `PUT /api/v1/messages/{message_id}/read`

## 6. 个人中心模块

### 6.1 更新个人信息
- **接口**: `PUT /api/v1/user/profile`
- **输入参数**:
```json
{
  "username": "张三",              // users.username
  "avatar": "http://...",          // users.avatar
  "worker_info": {                 // workers表信息
    "real_name": "张三",           // workers.real_name
    "gender": "male",              // workers.gender
    "age": 25,                     // workers.age
    "education": "本科",            // workers.education
    "height": 175.5,               // workers.height
    "introduction": "有丰富的工作经验", // workers.introduction
    "work_experience": "曾在顺丰工作2年", // workers.work_experience
    "expected_salary": 200.00      // workers.expected_salary
  }
}
```

### 6.2 上传认证文件
- **接口**: `POST /api/v1/user/upload-cert`
- **输入参数**: 文件上传
- **输出参数**:
```json
{
  "code": 200,
  "data": {
    "cert_type": "health_cert",    // 认证类型
    "file_url": "http://...",      // 文件URL
    "cert_id": 789012              // 认证记录ID
  }
}
```

### 6.3 获取我的收藏
- **接口**: `GET /api/v1/user/favorites`
- **输出参数**:
```json
{
  "code": 200,
  "data": [
    {
      "favorite_id": 123456,       // user_favorite_jobs.id
      "job_id": 789012,            // user_favorite_jobs.job_id
      "created_at": "2024-01-01T10:00:00Z", // user_favorite_jobs.created_at
      "job_info": {                // 关联岗位信息
        "title": "快递分拣员",
        "salary": 200.00,
        "location": "北京市朝阳区"
      }
    }
  ]
}
```

### 6.4 收藏/取消收藏岗位
- **接口**: `POST /api/v1/user/favorites`
- **输入参数**:
```json
{
  "job_id": 789012                 // user_favorite_jobs.job_id
}
```

### 6.5 获取收入统计
- **接口**: `GET /api/v1/user/income`
- **查询参数**: `period=month&year=2024&month=1`
- **输出参数**:
```json
{
  "code": 200,
  "data": {
    "total_income": 5000.00,       // 总收入
    "pending_income": 500.00,      // 待结算
    "paid_income": 4500.00,        // 已到账
    "payments": [                  // payments表信息
      {
        "payment_id": 123456,      // payments.id
        "job_id": 789012,          // payments.job_id
        "amount": 200.00,          // payments.amount
        "status": "completed",     // payments.status
        "paid_at": "2024-01-01T17:00:00Z", // payments.paid_at
        "job_info": {
          "title": "快递分拣员"
        }
      }
    ]
  }
}
```

## 7. 社区模块

### 7.1 获取社区帖子
- **接口**: `GET /api/v1/community/posts`
- **输出参数**:
```json
{
  "code": 200,
  "data": [
    {
      "post_id": 123456,           // community_posts.id
      "author_id": 789012,         // community_posts.author_id
      "title": "新手必看：如何快速接单", // community_posts.title
      "content": "分享一些经验...", // community_posts.content
      "post_type": "discussion",   // community_posts.post_type
      "view_count": 100,           // community_posts.view_count
      "like_count": 20,            // community_posts.like_count
      "status": "published",       // community_posts.status
      "created_at": "2024-01-01T10:00:00Z", // community_posts.created_at
      "author_info": {             // 作者信息
        "username": "张三",
        "avatar": "http://..."
      }
    }
  ]
}
```

### 7.2 发布帖子
- **接口**: `POST /api/v1/community/posts`
- **输入参数**:
```json
{
  "title": "新手必看：如何快速接单", // community_posts.title
  "content": "分享一些经验...",    // community_posts.content
  "post_type": "discussion"       // community_posts.post_type
}
```

## 8. 考勤模块

### 8.1 打卡
- **接口**: `POST /api/v1/attendance/checkin`
- **输入参数**:
```json
{
  "job_id": 123456,               // attendance_records.job_id
  "check_in_location": "北京市朝阳区", // attendance_records.check_in_location
  "latitude": 39.9042,            // 当前纬度
  "longitude": 116.4074           // 当前经度
}
```

### 8.2 签退
- **接口**: `POST /api/v1/attendance/checkout`
- **输入参数**:
```json
{
  "job_id": 123456,               // attendance_records.job_id
  "check_out_location": "北京市朝阳区", // attendance_records.check_out_location
  "latitude": 39.9042,            // 当前纬度
  "longitude": 116.4074           // 当前经度
}
```

### 8.3 获取考勤记录
- **接口**: `GET /api/v1/attendance/records`
- **查询参数**: `job_id=123456&page=1&limit=20`
- **输出参数**:
```json
{
  "code": 200,
  "data": [
    {
      "record_id": 789012,         // attendance_records.id
      "job_id": 123456,            // attendance_records.job_id
      "check_in": "2024-01-01T09:00:00Z", // attendance_records.check_in
      "check_out": "2024-01-01T17:00:00Z", // attendance_records.check_out
      "work_hours": 8.0,           // attendance_records.work_hours
      "status": "normal",          // attendance_records.status
      "job_info": {
        "title": "快递分拣员"
      }
    }
  ]
}
```

## 9. 评价模块

### 9.1 获取收到的评价
- **接口**: `GET /api/v1/reviews/received`
- **输出参数**:
```json
{
  "code": 200,
  "data": [
    {
      "review_id": 123456,         // reviews.id
      "job_id": 789012,            // reviews.job_id
      "rating": 5,                 // reviews.rating
      "content": "工作认真负责",    // reviews.content
      "review_type": "employer_to_worker", // reviews.review_type
      "created_at": "2024-01-01T18:00:00Z", // reviews.created_at
      "job_info": {
        "title": "快递分拣员"
      },
      "employer_info": {
        "company_name": "顺丰快递"
      }
    }
  ]
}
```

### 9.2 评价雇主
- **接口**: `POST /api/v1/reviews`
- **输入参数**:
```json
{
  "job_id": 789012,               // reviews.job_id
  "employer_id": 456789,          // reviews.employer_id
  "rating": 5,                    // reviews.rating
  "content": "雇主很好，工作环境不错", // reviews.content
  "review_type": "worker_to_employer" // reviews.review_type
}
```

## 10. 通用响应格式

所有接口统一使用以下响应格式：

```json
{
  "code": 200,                    // 状态码：200成功，其他失败
  "message": "操作成功",           // 提示信息
  "data": {},                     // 数据内容
  "timestamp": "2024-01-01T10:00:00Z" // 响应时间戳
}
```

## 11. 错误码定义

- `200`: 成功
- `400`: 请求参数错误
- `401`: 未授权
- `403`: 权限不足
- `404`: 资源不存在
- `500`: 服务器内部错误

## 12. 认证方式

所有需要认证的接口都需要在请求头中携带JWT token：

```
Authorization: Bearer <jwt_token>
```
