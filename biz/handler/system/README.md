# System Handlers 使用说明

本目录包含了系统相关的所有API处理器，实现了完整的系统管理功能。

## 已完成的功能

### 1. 获取系统配置 (`GetSystemConfig`)
- 路由: `GET /api/v1/system/config`
- 功能: 获取系统配置信息
- 返回: 系统配置键值对
- 包含: 应用名称、版本、最低支持版本、强制更新标志、维护状态等

### 2. 获取版本信息 (`GetVersion`)
- 路由: `GET /api/v1/system/version`
- 功能: 获取应用版本信息
- 返回: 版本号、构建时间、Git提交信息
- 用于客户端版本检查和更新提示

### 3. 提交反馈 (`SubmitFeedback`)
- 路由: `POST /api/v1/system/feedback`
- 功能: 用户提交反馈意见
- 参数: feedback_type, content, contact_info, attachments
- 返回: 反馈ID
- 支持多种反馈类型和附件上传

### 4. 获取通知列表 (`GetNoticeList`)
- 路由: `GET /api/v1/system/notices`
- 功能: 获取系统通知列表，支持分页
- 参数: page_req (分页参数)
- 返回: 通知列表和分页信息
- 只显示已发布的通知

### 5. 获取通知详情 (`GetNoticeDetail`)
- 路由: `GET /api/v1/system/notices/:notice_id`
- 功能: 获取单个通知的详细信息
- 参数: notice_id
- 返回: 通知详细信息

## 技术实现

### 架构设计
- **Handler层**: 处理HTTP请求和响应
- **Logic层**: 业务逻辑处理
- **MySQL层**: 数据库操作

### 数据库操作
- 使用GORM进行数据库操作
- 支持事务处理（使用闭包事务）
- 分页查询优化
- 状态过滤（只显示已发布的通知）

### 特性
- **分页支持**: 通知列表接口支持分页
- **状态管理**: 通知支持草稿、已发布、已归档状态
- **反馈管理**: 支持多种反馈类型和处理状态
- **版本管理**: 支持版本检查和更新提示
- **事务安全**: 使用GORM闭包事务确保数据一致性

## 使用示例

### 获取系统配置
```bash
GET /api/v1/system/config
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "获取系统配置成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "configs": {
    "app_name": "零工客户端",
    "app_version": "1.0.0",
    "min_version": "1.0.0",
    "force_update": "false",
    "maintenance": "false",
    "contact_phone": "400-123-4567"
  }
}
```

### 获取版本信息
```bash
GET /api/v1/system/version
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "获取版本信息成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "version": "1.0.0",
  "build_time": "2024-01-20T10:00:00Z",
  "git_commit": "abc123"
}
```

### 提交反馈
```bash
POST /api/v1/system/feedback
{
  "feedback_type": "bug",
  "content": "发现了一个bug",
  "contact_info": "user@example.com",
  "attachments": ["screenshot.jpg"]
}
```

### 获取通知列表
```bash
GET /api/v1/system/notices?page=1&limit=10
```

### 获取通知详情
```bash
GET /api/v1/system/notices/1
```

## 注意事项

1. 系统配置和版本信息目前使用硬编码数据，实际应用中应从数据库读取
2. 反馈功能需要用户登录（user_id需要从JWT token中获取）
3. 通知只显示已发布状态的内容
4. 分页参数有默认值：page=1, limit=10
5. 支持多种反馈类型：bug、feature、suggestion、other

## 数据库模型

### Feedback 反馈表
- `user_id`: 用户ID
- `type`: 反馈类型
- `content`: 反馈内容
- `contact`: 联系方式
- `status`: 处理状态（pending、processing、resolved、closed）
- `reply`: 回复内容
- `reply_time`: 回复时间

### Notice 通知表
- `title`: 通知标题
- `content`: 通知内容
- `type`: 通知类型
- `priority`: 优先级（low、normal、high、urgent）
- `status`: 状态（draft、published、archived）
- `start_time`: 开始时间
- `end_time`: 结束时间

### VersionInfo 版本信息表
- `version`: 版本号
- `build_number`: 构建号
- `min_version`: 最低支持版本
- `force_update`: 是否强制更新
- `update_url`: 更新下载地址
- `update_note`: 更新说明
