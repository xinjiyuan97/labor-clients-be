# Attendance Handlers 使用说明

本目录包含了考勤相关的所有API处理器，实现了完整的考勤管理功能。

## 已完成的功能

### 1. 打卡签到 (`CheckIn`)
- 路由: `POST /api/v1/attendance/checkin`
- 功能: 员工打卡签到
- 参数: job_id, check_in_location, latitude, longitude
- 返回: 考勤记录ID、打卡时间
- 防止重复打卡（每天每个工作只能打卡一次）

### 2. 签退 (`CheckOut`)
- 路由: `POST /api/v1/attendance/checkout`
- 功能: 员工签退
- 参数: job_id, check_out_location, latitude, longitude
- 返回: 考勤记录ID、签退时间、工作时长
- 自动计算工作时长
- 需要先打卡才能签退

### 3. 获取考勤记录列表 (`GetAttendanceRecords`)
- 路由: `GET /api/v1/attendance/records`
- 功能: 获取考勤记录列表，支持分页和筛选
- 参数: page_req (分页参数), job_id (可选), start_date, end_date
- 返回: 考勤记录列表和分页信息
- 支持按工作和日期范围筛选

### 4. 获取考勤记录详情 (`GetAttendanceDetail`)
- 路由: `GET /api/v1/attendance/records/:record_id`
- 功能: 获取单个考勤记录的详细信息
- 参数: record_id
- 返回: 考勤记录详细信息
- 包含打卡/签退时间、位置、工作时长、状态等

### 5. 申请请假 (`ApplyLeave`)
- 路由: `POST /api/v1/attendance/apply-leave`
- 功能: 申请请假
- 参数: job_id, leave_date, leave_reason
- 返回: 请假记录ID
- 创建状态为"leave"的考勤记录
- 防止同一天重复请假

### 6. 申请补签 (`ApplyMakeup`)
- 路由: `POST /api/v1/attendance/apply-makeup`
- 功能: 申请补签
- 参数: job_id, makeup_date, makeup_reason, makeup_time
- 返回: 补签记录ID
- 创建状态为"normal"的考勤记录
- 防止同一天重复补签

## 技术实现

### 架构设计
- **Handler层**: 处理HTTP请求和响应
- **Logic层**: 业务逻辑处理
- **MySQL层**: 数据库操作

### 数据库操作
- 使用GORM进行数据库操作
- 支持事务处理（使用闭包事务）
- 分页查询优化
- 自动计算工作时长

### 特性
- **防重复**: 防止同一天重复打卡、请假或补签
- **自动计算**: 自动计算工作时长（签退时间 - 打卡时间）
- **位置记录**: 记录打卡和签退的地理位置
- **状态管理**: 支持多种考勤状态（normal, late, early_leave, absent, leave）
- **分页支持**: 考勤记录列表支持分页
- **日期筛选**: 支持按日期范围筛选考勤记录
- **事务安全**: 使用GORM闭包事务确保数据一致性

## 使用示例

### 打卡签到
```bash
POST /api/v1/attendance/checkin
{
  "job_id": 123,
  "check_in_location": "办公室A座",
  "latitude": 39.9042,
  "longitude": 116.4074
}
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "打卡成功",
    "timestamp": "2024-01-20T08:30:00Z"
  },
  "record_id": 456,
  "check_in_time": "2024-01-20T08:30:00Z"
}
```

### 签退
```bash
POST /api/v1/attendance/checkout
{
  "job_id": 123,
  "check_out_location": "办公室A座",
  "latitude": 39.9042,
  "longitude": 116.4074
}
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "签退成功",
    "timestamp": "2024-01-20T18:00:00Z"
  },
  "record_id": 456,
  "check_out_time": "2024-01-20T18:00:00Z",
  "work_hours": 9.5
}
```

### 获取考勤记录列表
```bash
GET /api/v1/attendance/records?page=1&limit=10&start_date=2024-01-01&end_date=2024-01-31&job_id=123
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "获取考勤记录列表成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "page_resp": {
    "page": 1,
    "limit": 10,
    "total": 25
  },
  "records": [
    {
      "record_id": 456,
      "job_id": 123,
      "worker_id": 789,
      "check_in": "2024-01-20T08:30:00Z",
      "check_out": "2024-01-20T18:00:00Z",
      "work_hours": 9.5,
      "check_in_location": "办公室A座",
      "check_out_location": "办公室A座",
      "status": "normal"
    }
  ]
}
```

### 获取考勤记录详情
```bash
GET /api/v1/attendance/records/456
```

### 申请请假
```bash
POST /api/v1/attendance/apply-leave
{
  "job_id": 123,
  "leave_date": "2024-01-25",
  "leave_reason": "个人事务"
}
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "申请请假成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "leave_id": 457
}
```

### 申请补签
```bash
POST /api/v1/attendance/apply-makeup
{
  "job_id": 123,
  "makeup_date": "2024-01-18",
  "makeup_reason": "忘记打卡",
  "makeup_time": "2024-01-18 08:30:00"
}
```

响应示例：
```json
{
  "base": {
    "code": 200,
    "message": "申请补签成功",
    "timestamp": "2024-01-20T10:00:00Z"
  },
  "makeup_id": 458
}
```

## 注意事项

1. 所有操作需要用户登录（worker_id需要从JWT token中获取）
2. 每天每个工作只能打卡一次
3. 必须先打卡才能签退
4. 工作时长自动计算（签退时间 - 打卡时间）
5. 考勤状态类型：
   - `normal`: 正常
   - `late`: 迟到
   - `early_leave`: 早退
   - `absent`: 缺勤
   - `leave`: 请假
6. 请假和补签会创建考勤记录，状态分别为"leave"和"normal"
7. 所有写操作使用事务确保数据一致性
8. 分页参数有默认值：page=1, limit=10
9. 地理位置使用经纬度记录

## 数据库模型

### AttendanceRecord 考勤记录表
- `job_id`: 岗位ID
- `worker_id`: 零工ID
- `check_in`: 打卡时间
- `check_out`: 签退时间
- `work_hours`: 工作时长（小时）
- `check_in_location`: 打卡位置
- `check_out_location`: 签退位置
- `status`: 考勤状态
- `created_at`: 创建时间
- `updated_at`: 更新时间

## 工作时长计算

工作时长 = 签退时间 - 打卡时间

- 使用decimal类型确保精度
- 自动保留2位小数
- 单位：小时

## API路由总览

- `POST /api/v1/attendance/checkin` - 打卡签到
- `POST /api/v1/attendance/checkout` - 签退
- `GET /api/v1/attendance/records` - 获取考勤记录列表
- `GET /api/v1/attendance/records/:record_id` - 获取考勤记录详情
- `POST /api/v1/attendance/apply-leave` - 申请请假
- `POST /api/v1/attendance/apply-makeup` - 申请补签

## 扩展功能

未来可以考虑添加的功能：
- 自动判断迟到/早退
- 考勤统计报表
- 考勤异常提醒
- 考勤审批流程
- 加班记录
- 请假审批
- 补签审批
- 导出考勤记录
- 考勤打卡地点限制（地理围栏）
- 人脸识别打卡

