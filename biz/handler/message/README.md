# Message Handlers 使用说明

本目录包含了消息相关的所有API处理器，实现了完整的消息系统功能。

## 已完成的功能

### 1. 发送消息 (`SendMessage`)
- 路由: `POST /api/v1/messages`
- 功能: 发送消息给指定用户
- 参数: to_user, message_type, content, msg_category
- 返回: 消息ID
- 支持多种消息类型和分类

### 2. 获取消息列表 (`GetMessageList`)
- 路由: `GET /api/v1/messages`
- 功能: 获取用户的消息列表，支持分页和分类过滤
- 参数: page_req (分页参数), msg_category (消息分类), is_read (已读状态)
- 返回: 消息列表和分页信息
- 支持按分类和已读状态过滤

### 3. 获取消息详情 (`GetMessageDetail`)
- 路由: `GET /api/v1/messages/:message_id`
- 功能: 获取单个消息的详细信息
- 参数: message_id
- 返回: 消息详细信息

### 4. 标记消息已读 (`MarkMessageRead`)
- 路由: `PUT /api/v1/messages/:message_id/read`
- 功能: 标记单个消息为已读
- 参数: message_id
- 返回: 操作结果

### 5. 批量标记已读 (`BatchMarkRead`)
- 路由: `PUT /api/v1/messages/batch-read`
- 功能: 批量标记多个消息为已读
- 参数: message_ids (消息ID数组)
- 返回: 操作结果

### 6. 获取未读消息数量 (`GetUnreadCount`)
- 路由: `GET /api/v1/messages/unread-count`
- 功能: 获取用户的未读消息数量
- 参数: msg_category (消息分类，可选)
- 返回: 未读消息数量

## 技术实现

### 架构设计
- **Handler层**: 处理HTTP请求和响应
- **Logic层**: 业务逻辑处理
- **MySQL层**: 数据库操作

### 数据库操作
- 使用GORM进行数据库操作
- 支持事务处理（使用闭包事务）
- 分页查询优化
- 分类和状态过滤

### 特性
- **分页支持**: 消息列表接口支持分页
- **分类管理**: 支持系统消息、聊天消息、社区消息等分类
- **已读状态**: 支持消息已读/未读状态管理
- **批量操作**: 支持批量标记已读
- **事务安全**: 使用GORM闭包事务确保数据一致性

## 使用示例

### 发送消息
```bash
POST /api/v1/messages
{
  "to_user": 123,
  "message_type": "text",
  "content": "你好，这是一条测试消息",
  "msg_category": "chat"
}
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "发送消息成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "message_id": 456
}
```

### 获取消息列表
```bash
GET /api/v1/messages?page=1&limit=10&msg_category=chat&is_read=false
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "获取消息列表成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "page_resp": {
    "page": 1,
    "limit": 10,
    "total": 25
  },
  "messages": [
    {
      "message_id": 456,
      "from_user": 789,
      "to_user": 123,
      "message_type": "text",
      "content": "你好，这是一条测试消息",
      "msg_category": "chat",
      "is_read": false,
      "created_at": "2024-01-20T10:00:00Z"
    }
  ]
}
```

### 获取消息详情
```bash
GET /api/v1/messages/456
```

### 标记消息已读
```bash
PUT /api/v1/messages/456/read
```

### 批量标记已读
```bash
PUT /api/v1/messages/batch-read
{
  "message_ids": [456, 457, 458]
}
```

### 获取未读消息数量
```bash
GET /api/v1/messages/unread-count?msg_category=chat
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "获取未读消息数量成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "unread_count": 5
}
```

## 注意事项

1. 所有消息相关操作需要用户登录（user_id需要从JWT token中获取）
2. 发送消息时，from_user需要从JWT token中获取
3. 消息分类支持：system（系统消息）、chat（聊天消息）、community（社区消息）
4. 分页参数有默认值：page=1, limit=10
5. 支持按消息分类和已读状态过滤
6. 批量操作使用事务确保数据一致性

## 数据库模型

### Message 消息表
- `from_user`: 发送用户ID
- `to_user`: 接收用户ID
- `message_type`: 消息类型（text、image、file等）
- `content`: 消息内容
- `msg_category`: 消息分类（system、chat、community）
- `is_read`: 是否已读
- `created_at`: 创建时间
- `updated_at`: 更新时间

## 消息分类说明

- **system**: 系统消息，如通知、公告等
- **chat**: 聊天消息，用户之间的私聊
- **community**: 社区消息，如评论、点赞等

## 扩展功能

未来可以考虑添加的功能：
- 消息推送（WebSocket/SSE）
- 消息撤回
- 消息转发
- 消息搜索
- 消息加密
- 消息状态（发送中、已送达、已读）
