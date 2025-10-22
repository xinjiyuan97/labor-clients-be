# 火山引擎TOS文件上传集成指南

## 概述

本项目已经集成了火山引擎TOS（Tencent Object Storage）对象存储服务，用于上传图片、文件等存储在TOS上。

## 功能特性

- ✅ 支持图片上传（jpg, jpeg, png, gif, webp, bmp, svg）
- ✅ 支持文件上传（pdf, doc, docx, xls, xlsx, zip, rar等）
- ✅ 支持证书文件上传（id_card, passport, driver_license等）
- ✅ 文件大小限制检查
- ✅ 文件类型验证
- ✅ 自动生成唯一文件名（基于雪花算法）
- ✅ 按日期和类型组织文件目录结构

## 配置说明

### 1. 修改配置文件

编辑 `conf/config.yaml` 或 `conf/config.prod.yaml`，配置TOS相关参数：

```yaml
oss:
  provider: "volcengine"                                    # 存储提供商
  access_key: "your_access_key"                            # 火山引擎访问密钥
  secret_key: "your_secret_key"                            # 火山引擎秘密密钥
  region: "cn-beijing"                                      # TOS区域
  bucket: "your-bucket-name"                               # TOS存储桶名称
  endpoint: "tos-cn-beijing.volces.com"                    # TOS端点地址
  base_url: "https://your-bucket.tos-cn-beijing.volces.com/" # 文件访问基础URL
  upload_path: "uploads/"                                   # 上传路径前缀
  max_file_size: 10485760                                   # 最大文件大小(10MB)
  allowed_exts:                                             # 允许的文件扩展名
    - "jpg"
    - "jpeg"
    - "png"
    - "gif"
    - "webp"
    - "pdf"
    - "doc"
    - "docx"
    - "xls"
    - "xlsx"
    - "zip"
    - "rar"
```

### 2. 获取火山引擎凭证

1. 登录[火山引擎控制台](https://console.volcengine.com/)
2. 创建或选择TOS存储桶
3. 获取Access Key和Secret Key
4. 配置存储桶的公网访问权限（如需要公网访问）

## API接口

### 1. 上传图片（直接URL）

**接口地址**: `POST /api/v1/upload/image`

**请求参数**:
- `image_file` (file, 必填): 要上传的图片文件
- `upload_type` (string, 可选): 上传类型，用于文件分类（如：avatar, post, product）

**响应示例**:
```json
{
  "base": {
    "code": 200,
    "message": "图片上传成功",
    "timestamp": "2025-10-22T10:30:00Z"
  },
  "image_url": "https://your-bucket.tos-cn-beijing.volces.com/uploads/avatar/20251022/20251022103000_123456789.jpg",
  "file_name": "photo.jpg",
  "file_size": 102400
}
```

**注意**: 默认返回的是TOS的公网访问URL，如需要带签名的临时访问URL，请使用 **获取签名URL** 接口。

### 2. 上传文件（直接URL）

**接口地址**: `POST /api/v1/upload/file`

**请求参数**:
- `file` (file, 必填): 要上传的文件
- `upload_type` (string, 可选): 上传类型，用于文件分类

**响应示例**:
```json
{
  "base": {
    "code": 200,
    "message": "文件上传成功",
    "timestamp": "2025-10-22T10:30:00Z"
  },
  "file_url": "https://your-bucket.tos-cn-beijing.volces.com/uploads/document/20251022/20251022103000_123456789.pdf",
  "file_name": "document.pdf",
  "file_size": 512000,
  "file_type": "pdf"
}
```

**注意**: 默认返回的是TOS的公网访问URL，如需要带签名的临时访问URL，请使用 **获取签名URL** 接口。

### 3. 上传证书文件（直接URL）

**接口地址**: `POST /api/v1/upload/cert`

**请求参数**:
- `cert_file` (file, 必填): 要上传的证书文件（支持jpg, jpeg, png, pdf）
- `cert_type` (string, 必填): 证书类型
  - `id_card`: 身份证
  - `passport`: 护照
  - `driver_license`: 驾驶证
  - `business_license`: 营业执照
  - `qualification_cert`: 资格证书
  - `health_cert`: 健康证
  - `other`: 其他

**响应示例**:
```json
{
  "base": {
    "code": 200,
    "message": "证书文件上传成功",
    "timestamp": "2025-10-22T10:30:00Z"
  },
  "file_url": "https://your-bucket.tos-cn-beijing.volces.com/uploads/cert/id_card/20251022/20251022103000_123456789.jpg",
  "cert_type": "id_card",
  "file_name": "idcard.jpg",
  "file_size": 204800
}
```

**注意**: 默认返回的是TOS的公网访问URL，如需要带签名的临时访问URL，请使用 **获取签名URL** 接口。

### 4. 获取签名URL（临时访问链接）

**接口地址**: `GET /api/v1/upload/signed-url`

**请求参数**:
- `file_url` (string, 必填): 文件的原始URL
- `expire_seconds` (int64, 可选): 签名URL的过期时间（秒），默认3600秒（1小时），最大7天

**响应示例**:
```json
{
  "base": {
    "code": 200,
    "message": "获取签名URL成功",
    "timestamp": "2025-10-22T10:30:00Z"
  },
  "signed_url": "https://your-bucket.tos-cn-beijing.volces.com/uploads/avatar/20251022/20251022103000_123456789.jpg?X-Tos-Algorithm=...&X-Tos-Credential=...&X-Tos-Expires=...",
  "expire_seconds": 3600,
  "expire_time": "2025-10-22T11:30:00Z"
}
```

**使用场景**:
- 需要临时访问权限的文件
- 私有文件的安全访问
- 防止文件URL被长期盗用
- 需要控制文件访问时限

**示例请求**:
```bash
# 获取1小时有效期的签名URL
GET /api/v1/upload/signed-url?file_url=https://your-bucket.tos-cn-beijing.volces.com/uploads/avatar/20251022/20251022103000_123456789.jpg&expire_seconds=3600

# 获取24小时有效期的签名URL
GET /api/v1/upload/signed-url?file_url=https://your-bucket.tos-cn-beijing.volces.com/uploads/document/20251022/document.pdf&expire_seconds=86400
```

## 文件存储结构

上传的文件会按照以下目录结构存储：

```
uploads/
├── avatar/
│   └── 20251022/
│       └── 20251022103000_123456789.jpg
├── document/
│   └── 20251022/
│       └── 20251022103000_987654321.pdf
└── cert/
    ├── id_card/
    │   └── 20251022/
    │       └── 20251022103000_111111111.jpg
    └── passport/
        └── 20251022/
            └── 20251022103000_222222222.jpg
```

## 错误处理

常见错误码：

- `400`: 参数错误（文件类型不支持、文件大小超限等）
- `500`: 服务器错误（TOS连接失败、上传失败等）

错误响应示例：
```json
{
  "base": {
    "code": 400,
    "message": "不支持的图片格式，仅支持: jpg, jpeg, png, gif, webp, bmp, svg",
    "timestamp": "2025-10-22T10:30:00Z"
  }
}
```

## 本地开发

如果不想在本地开发时使用TOS，可以将配置中的 `provider` 设置为 `local`：

```yaml
oss:
  provider: "local"
  # ... 其他配置保持不变
```

这样文件会返回模拟URL，不会实际上传到TOS。

## 签名URL vs 直接URL

### 直接URL（公开访问）
- **优点**: 简单快速，无需额外请求
- **缺点**: URL长期有效，可能被盗用
- **适用场景**: 公开的图片、文件，如头像、商品图等

### 签名URL（临时访问）
- **优点**: 安全性高，可控制访问时限，防止盗用
- **缺点**: 需要额外请求生成，URL较长
- **适用场景**: 
  - 私密文件（证件照、合同文件）
  - 付费内容
  - 临时分享链接
  - 需要访问控制的场景

### 最佳实践

1. **公开内容**: 使用直接URL
   ```
   上传图片 → 直接使用返回的 image_url
   ```

2. **私密内容**: 使用签名URL
   ```
   上传文件 → 获取 file_url → 调用签名URL接口 → 使用 signed_url
   ```

3. **混合场景**: 
   - 存储时保存直接URL
   - 需要访问时动态生成签名URL
   - 根据用户权限决定是否使用签名

## 技术架构

### 核心组件

1. **TOS客户端管理器** (`dal/tos/tos.go`)
   - 初始化和管理TOS客户端连接
   - 提供客户端实例访问接口

2. **文件上传服务** (`utils/upload.go`)
   - `UploadService` 接口：定义文件上传服务规范
     - `UploadFile`: 上传文件
     - `DeleteFile`: 删除文件
     - `GetSignedURL`: 生成预签名URL
   - `TOSUploadService`: 火山引擎TOS实现
   - `LocalUploadService`: 本地开发实现

3. **业务逻辑层** (`biz/logic/upload/`)
   - `upload_image.go`: 图片上传逻辑
   - `upload_file.go`: 文件上传逻辑
   - `upload_cert_file.go`: 证书上传逻辑
   - `get_signed_url.go`: 签名URL生成逻辑

4. **接口处理层** (`biz/handler/upload/`)
   - 处理HTTP请求
   - 文件接收和参数验证
   - 调用业务逻辑层

## 安全建议

1. **访问控制**: 在TOS控制台配置适当的访问权限
2. **防盗链**: 配置Referer白名单防止资源盗用
3. **HTTPS**: 生产环境务必使用HTTPS传输
4. **文件扫描**: 考虑集成病毒扫描服务
5. **凭证管理**: 使用环境变量或密钥管理服务存储凭证

## 性能优化

1. **并发上传**: TOS SDK内部支持分片并发上传大文件
2. **CDN加速**: 可配置CDN加速文件访问
3. **缓存策略**: 设置合适的Cache-Control头

## 故障排查

### 1. 上传失败

检查：
- TOS凭证是否正确
- 存储桶是否存在
- 网络连接是否正常
- 文件大小是否超限

### 2. 文件无法访问

检查：
- 存储桶权限设置
- base_url配置是否正确
- 文件是否真正上传成功

### 3. 初始化失败

查看日志中的错误信息，常见原因：
- 配置项缺失
- 凭证无效
- 网络问题

## 相关链接

- [火山引擎TOS文档](https://www.volcengine.com/docs/6349)
- [TOS Go SDK文档](https://github.com/volcengine/ve-tos-golang-sdk)
- [TOS控制台](https://console.volcengine.com/tos)

## 更新日志

### 2025-10-22
- ✅ 集成火山引擎TOS SDK
- ✅ 实现图片、文件、证书上传功能
- ✅ 添加文件类型和大小验证
- ✅ 实现基于雪花算法的唯一文件名生成
- ✅ 支持按日期和类型组织文件目录
- ✅ 实现预签名URL生成功能
- ✅ 支持自定义签名URL过期时间
- ✅ 提供获取签名URL的API接口

