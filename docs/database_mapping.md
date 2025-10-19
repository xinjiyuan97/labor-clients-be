# 接口与数据库字段映射表

## 1. 用户认证模块

### 1.1 用户注册/登录
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| phone | users.phone | 手机号 |
| password | users.password_hash | 密码哈希 |
| role | users.role | 用户角色 |
| username | users.username | 用户名 |
| user_id | users.id | 用户ID |

### 1.2 用户信息
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| user_id | users.id | 用户ID |
| username | users.username | 用户名 |
| phone | users.phone | 手机号 |
| avatar | users.avatar | 头像URL |
| role | users.role | 用户角色 |
| real_name | workers.real_name | 真实姓名 |
| gender | workers.gender | 性别 |
| age | workers.age | 年龄 |
| education | workers.education | 学历 |
| introduction | workers.introduction | 个人介绍 |

## 2. 日程管理模块

### 2.1 日程操作
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| schedule_id | schedules.id | 日程ID |
| title | schedules.title | 日程标题 |
| job_id | schedules.job_id | 关联岗位ID |
| start_time | schedules.start_time | 开始时间 |
| end_time | schedules.end_time | 结束时间 |
| location | schedules.location | 地点 |
| notes | schedules.notes | 备注 |
| status | schedules.status | 状态 |
| reminder_minutes | schedules.reminder_minutes | 提前提醒分钟数 |
| worker_id | schedules.worker_id | 零工ID |

## 3. 岗位推荐模块

### 3.1 岗位列表/详情
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| job_id | jobs.id | 岗位ID |
| title | jobs.title | 岗位标题 |
| job_type | jobs.job_type | 岗位类型 |
| description | jobs.description | 岗位描述 |
| salary | jobs.salary | 薪资 |
| salary_unit | jobs.salary_unit | 结算单位 |
| location | jobs.location | 工作地点 |
| latitude | jobs.latitude | 纬度 |
| longitude | jobs.longitude | 经度 |
| requirements | jobs.requirements | 工作要求 |
| benefits | jobs.benefits | 福利待遇 |
| start_time | jobs.start_time | 开始时间 |
| end_time | jobs.end_time | 结束时间 |
| status | jobs.status | 岗位状态 |
| applicant_count | jobs.applicant_count | 报名人数 |
| max_applicants | jobs.max_applicants | 最大报名人数 |
| employer_id | jobs.employer_id | 雇主ID |
| brand_id | jobs.brand_id | 所属品牌ID |
| category_id | jobs.category_id | 分类ID |

### 3.2 品牌信息
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| brand_id | brands.id | 品牌ID |
| name | brands.name | 品牌名称 |
| logo | brands.logo | 品牌Logo URL |
| description | brands.description | 品牌描述 |
| auth_status | brands.auth_status | 认证状态 |

### 3.3 雇主信息
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| user_id | employers.user_id | 关联用户ID |
| brand_id | employers.brand_id | 所属品牌ID |
| company_name | employers.company_name | 公司名称 |
| contact_person | employers.contact_person | 联系人姓名 |
| contact_phone | employers.contact_phone | 联系人手机 |
| business_license | employers.business_license | 营业执照号 |
| auth_status | employers.auth_status | 认证状态 |
| auth_time | employers.auth_time | 认证时间 |

### 3.4 岗位分类
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| category_id | job_categories.id | 分类ID |
| name | job_categories.name | 分类名称 |
| description | job_categories.description | 分类描述 |
| parent_id | job_categories.parent_id | 父级分类ID |
| sort_order | job_categories.sort_order | 排序 |

### 3.5 岗位标签
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| job_id | job_tags.job_id | 岗位ID |
| tag_name | job_tags.tag_name | 标签名称 |
| tag_type | job_tags.tag_type | 标签类型 |

## 4. 岗位申请模块

### 4.1 岗位申请
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| application_id | job_applications.id | 申请ID |
| job_id | job_applications.job_id | 岗位ID |
| worker_id | job_applications.worker_id | 零工ID |
| status | job_applications.status | 申请状态 |
| applied_at | job_applications.applied_at | 申请时间 |
| confirmed_at | job_applications.confirmed_at | 确认时间 |
| cancel_reason | job_applications.cancel_reason | 取消原因 |
| worker_rating | job_applications.worker_rating | 零工评分 |
| employer_rating | job_applications.employer_rating | 雇主评分 |
| review | job_applications.review | 评价内容 |

## 5. 消息模块

### 5.1 消息管理
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| message_id | messages.id | 消息ID |
| from_user | messages.from_user | 发送用户ID |
| to_user | messages.to_user | 接收用户ID |
| message_type | messages.message_type | 消息类型 |
| content | messages.content | 消息内容 |
| msg_category | messages.msg_category | 消息分类 |
| is_read | messages.is_read | 是否已读 |

## 6. 个人中心模块

### 6.1 零工信息
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| user_id | workers.user_id | 用户ID |
| real_name | workers.real_name | 真实姓名 |
| gender | workers.gender | 性别 |
| age | workers.age | 年龄 |
| id_card | workers.id_card | 身份证号 |
| health_cert | workers.health_cert | 健康证URL |
| education | workers.education | 学历 |
| height | workers.height | 身高 |
| introduction | workers.introduction | 个人介绍 |
| work_experience | workers.work_experience | 工作经历 |
| expected_salary | workers.expected_salary | 期望薪资 |

### 6.2 用户收藏
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| favorite_id | user_favorite_jobs.id | 收藏ID |
| user_id | user_favorite_jobs.user_id | 用户ID |
| job_id | user_favorite_jobs.job_id | 岗位ID |

## 7. 社区模块

### 7.1 社区帖子
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| post_id | community_posts.id | 帖子ID |
| author_id | community_posts.author_id | 作者用户ID |
| title | community_posts.title | 帖子标题 |
| content | community_posts.content | 帖子内容 |
| post_type | community_posts.post_type | 帖子类型 |
| view_count | community_posts.view_count | 浏览数 |
| like_count | community_posts.like_count | 点赞数 |
| status | community_posts.status | 状态 |

## 8. 考勤模块

### 8.1 考勤记录
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| record_id | attendance_records.id | 考勤ID |
| job_id | attendance_records.job_id | 岗位ID |
| worker_id | attendance_records.worker_id | 零工ID |
| check_in | attendance_records.check_in | 打卡时间 |
| check_out | attendance_records.check_out | 签退时间 |
| work_hours | attendance_records.work_hours | 工作时长 |
| check_in_location | attendance_records.check_in_location | 打卡位置 |
| check_out_location | attendance_records.check_out_location | 签退位置 |
| status | attendance_records.status | 考勤状态 |

## 9. 评价模块

### 9.1 评价管理
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| review_id | reviews.id | 评价ID |
| job_id | reviews.job_id | 岗位ID |
| employer_id | reviews.employer_id | 雇主ID |
| worker_id | reviews.worker_id | 零工ID |
| rating | reviews.rating | 评分 |
| content | reviews.content | 评价内容 |
| review_type | reviews.review_type | 评价类型 |

## 10. 支付模块

### 10.1 支付记录
| 接口字段 | 数据库表.字段 | 说明 |
|---------|--------------|------|
| payment_id | payments.id | 支付ID |
| job_id | payments.job_id | 岗位ID |
| worker_id | payments.worker_id | 零工ID |
| employer_id | payments.employer_id | 雇主ID |
| amount | payments.amount | 支付金额 |
| payment_method | payments.payment_method | 支付方式 |
| status | payments.status | 支付状态 |
| paid_at | payments.paid_at | 支付时间 |
| platform_fee | payments.platform_fee | 平台费用 |

## 11. 字段类型说明

### 枚举值定义

#### users.role
- `worker`: 零工
- `employer`: 雇主

#### jobs.job_type
- `standard`: 标准岗位
- `rush`: 抢班岗位
- `transfer`: 转让岗位

#### jobs.status
- `draft`: 草稿
- `published`: 已发布
- `filled`: 已招满
- `completed`: 已完成
- `cancelled`: 已取消

#### schedules.status
- `pending`: 待开始
- `in_progress`: 进行中
- `completed`: 已完成
- `cancelled`: 已取消

#### job_applications.status
- `applied`: 已申请
- `confirmed`: 已确认
- `rejected`: 已拒绝
- `cancelled`: 已取消
- `completed`: 已完成

#### messages.msg_category
- `system`: 系统消息
- `chat`: 聊天消息
- `community`: 社区消息

#### attendance_records.status
- `normal`: 正常
- `late`: 迟到
- `early_leave`: 早退
- `absent`: 缺勤
- `leave`: 请假

#### reviews.review_type
- `employer_to_worker`: 雇主评价零工
- `worker_to_employer`: 零工评价雇主

#### payments.status
- `pending`: 待支付
- `processing`: 处理中
- `completed`: 已完成
- `failed`: 支付失败

## 12. 索引说明

所有表都包含了适当的索引以提高查询性能：

- **主键索引**: 所有表的 `id` 字段
- **唯一索引**: `users.phone`, `job_applications(job_id, worker_id)`, `user_favorite_jobs(user_id, job_id)`
- **普通索引**: 外键字段、常用查询字段、时间字段等
- **软删除索引**: 所有表的 `deleted_at` 字段
