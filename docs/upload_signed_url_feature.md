# 上传接口返回签名URL功能说明

## 功能概述

上传接口现在会同时返回两个URL地址：
1. **原始地址（image_url/file_url）**：用于存储在数据库中
2. **签名后的展示地址（display_url）**：用于前端展示和访问

## 背景说明

在使用对象存储（如火山引擎TOS）时，文件访问通常需要签名URL才能访问。为了简化前端使用，上传接口在返回时会自动生成签名URL。

## 接口变更

### 1. 上传图片接口

**接口**: `POST /api/v1/upload/image`

**响应示例**:
```json
{
  "base": {
    "code": 200,
    "message": "图片上传成功",
    "timestamp": "2025-10-24T12:00:00Z"
  },
  "image_url": "https://bucket.tos.region.com/uploads/images/20251024/123456.jpg",
  "display_url": "https://bucket.tos.region.com/uploads/images/20251024/123456.jpg?X-Tos-Signature=xxx&X-Tos-Expires=1234567890",
  "file_name": "example.jpg",
  "file_size": "102400"
}
```

**字段说明**:
| 字段 | 类型 | 说明 |
|------|------|------|
| image_url | string | 原始地址，存储在数据库中 |
| display_url | string | 签名后的展示地址，有效期7天 |
| file_name | string | 文件名 |
| file_size | int64 | 文件大小（字节） |

### 2. 上传文件接口

**接口**: `POST /api/v1/upload/file`

**响应示例**:
```json
{
  "base": {
    "code": 200,
    "message": "文件上传成功",
    "timestamp": "2025-10-24T12:00:00Z"
  },
  "file_url": "https://bucket.tos.region.com/uploads/files/20251024/document.pdf",
  "display_url": "https://bucket.tos.region.com/uploads/files/20251024/document.pdf?X-Tos-Signature=xxx&X-Tos-Expires=1234567890",
  "file_name": "document.pdf",
  "file_size": "512000",
  "file_type": "pdf"
}
```

**字段说明**:
| 字段 | 类型 | 说明 |
|------|------|------|
| file_url | string | 原始地址，存储在数据库中 |
| display_url | string | 签名后的展示地址，有效期7天 |
| file_name | string | 文件名 |
| file_size | int64 | 文件大小（字节） |
| file_type | string | 文件类型（image/pdf/word等） |

### 3. 上传认证文件接口

**接口**: `POST /api/v1/upload/cert`

**响应示例**:
```json
{
  "base": {
    "code": 200,
    "message": "证书文件上传成功",
    "timestamp": "2025-10-24T12:00:00Z"
  },
  "file_url": "https://bucket.tos.region.com/uploads/cert/id_card/20251024/cert.jpg",
  "display_url": "https://bucket.tos.region.com/uploads/cert/id_card/20251024/cert.jpg?X-Tos-Signature=xxx&X-Tos-Expires=1234567890",
  "cert_type": "id_card",
  "file_name": "cert.jpg",
  "file_size": "204800"
}
```

**字段说明**:
| 字段 | 类型 | 说明 |
|------|------|------|
| file_url | string | 原始地址，存储在数据库中 |
| display_url | string | 签名后的展示地址，有效期7天 |
| cert_type | string | 证书类型 |
| file_name | string | 文件名 |
| file_size | int64 | 文件大小（字节） |

## 使用说明

### 前端使用

#### 上传时

```javascript
// 上传图片
const formData = new FormData();
formData.append('image_file', file);
formData.append('upload_type', 'avatar');

const response = await fetch('/api/v1/upload/image', {
  method: 'POST',
  body: formData
});

const data = await response.json();

// 保存到数据库（使用原始地址）
saveToDatabase({
  avatar: data.image_url
});

// 用于页面展示（使用签名地址）
setAvatarPreview(data.display_url);
```

#### 展示时

```javascript
// 从数据库获取的数据
const user = {
  avatar: "https://bucket.tos.region.com/uploads/images/20251024/123456.jpg"
};

// 如果需要展示，调用签名URL接口
const signedUrl = await getSignedURL(user.avatar);

// 显示图片
<img src={signedUrl} alt="avatar" />
```

### 后端使用

#### 存储原始地址

```go
// 上传成功后，将 image_url 存储到数据库
brand := &models.Brand{
  Logo: resp.ImageURL,  // 存储原始地址
}
db.Create(brand)
```

#### 返回时获取签名URL

```go
// 查询数据库
brand, _ := GetBrandByID(brandID)

// 生成签名URL用于前端展示
uploadService, _ := utils.GetUploadService(cfg)
displayURL, _ := uploadService.GetSignedURL(brand.Logo, 7*24*3600)

// 返回给前端
return &BrandDetail{
  Logo: brand.Logo,          // 原始地址
  LogoDisplayURL: displayURL, // 签名后的展示地址
}
```

## 签名URL有效期

- **默认有效期**: 7天（604800秒）
- **适用场景**: 
  - 用户头像
  - 品牌Logo
  - 产品图片
  - 证书文件
  - 其他需要展示的文件

## 核心逻辑

### 生成签名URL

```go
// 生成签名URL用于展示（默认7天有效期）
displayURL, err := uploadService.GetSignedURL(fileURL, 7*24*3600)
if err != nil {
  utils.Warnf("生成签名URL失败: %v, 使用原始URL", err)
  displayURL = fileURL
}
```

### 容错处理

- 如果签名URL生成失败，会自动降级使用原始URL
- 不影响上传成功的结果
- 记录警告日志便于排查问题

## 存储提供商支持

### 火山引擎TOS

- ✅ 完整支持签名URL
- 使用 TOS SDK 的 `PreSignedURL` 方法
- 自动包含认证信息和过期时间

### 本地存储

- ⚠️ 直接返回原始URL
- 无需签名即可访问
- 适用于开发测试环境

## 注意事项

### 1. URL存储

**推荐做法**:
```sql
-- 只存储原始URL
CREATE TABLE brands (
  id BIGINT PRIMARY KEY,
  logo VARCHAR(255),  -- 存储原始URL
  ...
);
```

**不推荐做法**:
```sql
-- 不要存储签名URL（会过期）
CREATE TABLE brands (
  id BIGINT PRIMARY KEY,
  logo VARCHAR(500),  -- 存储签名URL会在7天后失效
  ...
);
```

### 2. 前端缓存

- 签名URL可以在前端缓存
- 建议缓存时间不超过签名URL的有效期
- 可以在URL即将过期时重新获取

### 3. 性能优化

**批量获取签名URL**:
```go
// 批量处理品牌列表
brands := GetBrands()
for i, brand := range brands {
  displayURL, _ := uploadService.GetSignedURL(brand.Logo, 7*24*3600)
  brands[i].LogoDisplayURL = displayURL
}
```

### 4. 安全考虑

- 原始URL通常无法直接访问（需要签名）
- 签名URL在有效期内可以被任何人访问
- 敏感文件可以使用较短的有效期

## 迁移指南

### 现有接口兼容性

✅ **向后兼容**: 
- 新增了 `display_url` 字段
- 保留了原有的 `image_url`/`file_url` 字段
- 不影响现有前端代码的使用

### 前端迁移步骤

1. **第一步**: 更新上传逻辑，使用 `display_url` 进行预览
```javascript
// 原来
setPreview(data.image_url);

// 现在
setPreview(data.display_url);  // 使用签名URL预览
saveToDatabase(data.image_url); // 仍然存储原始URL
```

2. **第二步**: 更新显示逻辑，调用签名URL接口
```javascript
// 从数据库获取的原始URL
const imageUrl = brand.logo;

// 获取签名URL用于显示
const signedUrl = await getSignedURL(imageUrl);
<img src={signedUrl} />
```

3. **第三步**: 优化批量查询，后端直接返回签名URL
```javascript
// 后端返回时已包含签名URL
const brands = await getBrandList();
brands.forEach(brand => {
  <img src={brand.logo_display_url} />
});
```

## 获取签名URL接口

如果需要为已存储的原始URL生成签名URL，可以使用独立的签名URL接口：

**接口**: `GET /api/v1/upload/signed-url`

**请求参数**:
```
file_url: https://bucket.tos.region.com/uploads/images/123.jpg
expire_seconds: 604800  // 7天
```

**响应**:
```json
{
  "base": {
    "code": 200,
    "message": "success"
  },
  "signed_url": "https://bucket.tos.region.com/uploads/images/123.jpg?X-Tos-Signature=xxx&X-Tos-Expires=1234567890",
  "expire_seconds": "604800",
  "expire_time": "2025-10-31T12:00:00Z"
}
```

## 总结

### 优势

✅ **简化前端使用**: 上传后即可直接用于预览
✅ **提高安全性**: 文件访问需要签名验证
✅ **向后兼容**: 不影响现有代码
✅ **灵活存储**: 原始URL便于长期存储

### 最佳实践

1. **数据库**: 存储原始URL（`image_url`/`file_url`）
2. **上传后预览**: 使用返回的 `display_url`
3. **长期展示**: 调用签名URL接口获取新的签名URL
4. **批量查询**: 后端统一处理签名URL生成

