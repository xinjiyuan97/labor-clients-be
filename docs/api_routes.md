# API 路由汇总表

## 基础路径
所有API的基础路径为：`/api/v1`

## 1. 用户认证模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| POST | `/auth/register` | 用户注册 | ❌ |
| POST | `/auth/login` | 用户登录 | ❌ |
| POST | `/auth/logout` | 用户登出 | ✅ |
| POST | `/auth/refresh` | 刷新Token | ❌ |
| GET | `/auth/profile` | 获取用户信息 | ✅ |

## 2. 日程管理模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/schedules` | 获取日程列表 | ✅ |
| GET | `/schedules/today` | 获取今日排班 | ✅ |
| GET | `/schedules/{id}` | 获取日程详情 | ✅ |
| POST | `/schedules` | 创建日程 | ✅ |
| PUT | `/schedules/{id}` | 更新日程 | ✅ |
| PUT | `/schedules/{id}/status` | 更新日程状态 | ✅ |
| DELETE | `/schedules/{id}` | 删除日程 | ✅ |

## 3. 岗位推荐模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/jobs` | 获取岗位列表 | ✅ |
| GET | `/jobs/{id}` | 获取岗位详情 | ✅ |
| GET | `/jobs/recommend` | 获取推荐岗位 | ✅ |
| GET | `/jobs/search` | 搜索岗位 | ✅ |
| GET | `/job-categories` | 获取岗位分类 | ❌ |
| GET | `/brands` | 获取品牌列表 | ❌ |

## 4. 岗位申请模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/job-applications` | 获取申请列表 | ✅ |
| GET | `/job-applications/my` | 获取我的申请 | ✅ |
| GET | `/job-applications/{id}` | 获取申请详情 | ✅ |
| POST | `/job-applications` | 申请岗位 | ✅ |
| PUT | `/job-applications/{id}/cancel` | 取消申请 | ✅ |
| PUT | `/job-applications/{id}/confirm` | 确认申请 | ✅ |

## 5. 消息模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/messages` | 获取消息列表 | ✅ |
| GET | `/messages/{id}` | 获取消息详情 | ✅ |
| POST | `/messages` | 发送消息 | ✅ |
| PUT | `/messages/{id}/read` | 标记消息已读 | ✅ |
| PUT | `/messages/batch-read` | 批量标记已读 | ✅ |
| GET | `/messages/unread-count` | 获取未读消息数 | ✅ |

## 6. 个人中心模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/user/profile` | 获取个人信息 | ✅ |
| PUT | `/user/profile` | 更新个人信息 | ✅ |
| POST | `/user/upload-avatar` | 上传头像 | ✅ |
| POST | `/user/upload-cert` | 上传认证文件 | ✅ |
| GET | `/user/favorites` | 获取我的收藏 | ✅ |
| POST | `/user/favorites` | 收藏岗位 | ✅ |
| DELETE | `/user/favorites/{job_id}` | 取消收藏 | ✅ |
| GET | `/user/income` | 获取收入统计 | ✅ |
| GET | `/user/income/detail` | 获取收入详情 | ✅ |
| GET | `/user/performance` | 获取个人表现 | ✅ |

## 7. 社区模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/community/posts` | 获取帖子列表 | ❌ |
| GET | `/community/posts/{id}` | 获取帖子详情 | ❌ |
| POST | `/community/posts` | 发布帖子 | ✅ |
| PUT | `/community/posts/{id}` | 更新帖子 | ✅ |
| DELETE | `/community/posts/{id}` | 删除帖子 | ✅ |
| POST | `/community/posts/{id}/like` | 点赞帖子 | ✅ |
| DELETE | `/community/posts/{id}/like` | 取消点赞 | ✅ |
| POST | `/community/posts/{id}/comment` | 评论帖子 | ✅ |
| GET | `/community/posts/{id}/comments` | 获取评论列表 | ❌ |
| GET | `/community/courses` | 获取入门课程 | ❌ |

## 8. 考勤模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/attendance/records` | 获取考勤记录 | ✅ |
| GET | `/attendance/records/{id}` | 获取考勤详情 | ✅ |
| POST | `/attendance/checkin` | 打卡 | ✅ |
| POST | `/attendance/checkout` | 签退 | ✅ |
| POST | `/attendance/apply-leave` | 申请请假 | ✅ |
| POST | `/attendance/apply-makeup` | 申请补卡 | ✅ |

## 9. 评价模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/reviews/received` | 获取收到的评价 | ✅ |
| GET | `/reviews/given` | 获取给出的评价 | ✅ |
| GET | `/reviews/{id}` | 获取评价详情 | ✅ |
| POST | `/reviews` | 发布评价 | ✅ |
| PUT | `/reviews/{id}` | 更新评价 | ✅ |

## 10. 支付模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/payments` | 获取支付记录 | ✅ |
| GET | `/payments/{id}` | 获取支付详情 | ✅ |
| POST | `/payments/withdraw` | 申请提现 | ✅ |
| GET | `/payments/withdraw-records` | 获取提现记录 | ✅ |

## 11. 系统模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| GET | `/system/config` | 获取系统配置 | ❌ |
| GET | `/system/version` | 获取版本信息 | ❌ |
| POST | `/system/feedback` | 提交反馈 | ✅ |
| GET | `/system/notices` | 获取公告列表 | ❌ |
| GET | `/system/notices/{id}` | 获取公告详情 | ❌ |

## 12. 文件上传模块

| 方法 | 路径 | 功能 | 需要认证 |
|------|------|------|----------|
| POST | `/upload/image` | 上传图片 | ✅ |
| POST | `/upload/file` | 上传文件 | ✅ |
| POST | `/upload/cert` | 上传认证文件 | ✅ |

## 路由分组说明

### 认证要求
- ❌ 不需要认证：公开接口，如注册、登录、获取公开信息等
- ✅ 需要认证：需要JWT token的接口，如用户操作、个人数据等

### 权限级别
1. **公开接口**: 任何人都可以访问
2. **用户接口**: 需要登录的用户
3. **管理员接口**: 需要管理员权限（如果需要）

### 响应格式
所有接口统一返回格式：
```json
{
  "code": 200,
  "message": "success",
  "data": {},
  "timestamp": "2024-01-01T10:00:00Z"
}
```

### 分页参数
列表接口统一支持分页：
- `page`: 页码，从1开始
- `limit`: 每页数量，默认20，最大100

### 排序参数
支持排序的接口统一参数：
- `sort`: 排序字段
- `order`: 排序方向，asc/desc

### 搜索参数
支持搜索的接口统一参数：
- `keyword`: 搜索关键词
- `filters`: 筛选条件（JSON格式）

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 409 | 资源冲突 |
| 422 | 数据验证失败 |
| 429 | 请求过于频繁 |
| 500 | 服务器内部错误 |

## 接口版本管理

- 当前版本：v1
- 版本路径：`/api/v1`
- 向后兼容：新版本保持向后兼容
- 废弃通知：通过响应头 `Deprecated` 通知废弃接口
