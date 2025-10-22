# 签名URL使用示例

## 简介

签名URL（Presigned URL）是一种为对象存储文件生成临时访问链接的方式，可以在不公开文件的情况下，允许用户在限定时间内访问文件。

## 使用场景

### 1. 私密文件分享
```
场景：用户上传身份证照片，只允许特定用户在特定时间内查看
流程：
1. 用户上传证件照
2. 后端存储直接URL到数据库
3. 当需要展示时，生成临时签名URL
4. 签名URL在1小时后自动失效
```

### 2. 付费内容访问
```
场景：用户购买课程视频，生成24小时有效的观看链接
流程：
1. 用户支付成功
2. 系统生成86400秒（24小时）有效的签名URL
3. 用户在24小时内可以观看
4. 超时后需要重新生成
```

### 3. 临时文件下载
```
场景：系统生成导出文件，允许用户下载一次
流程：
1. 系统生成导出文件并上传到TOS
2. 生成1小时有效的签名URL
3. 发送下载链接给用户
4. 1小时后链接自动失效
```

## API调用示例

### 示例1：获取图片的临时访问链接（默认1小时）

**请求**:
```bash
GET /api/v1/upload/signed-url?file_url=https://ai-editor-dev.tos-cn-beijing.volces.com/uploads/avatar/20251022/photo.jpg
```

**响应**:
```json
{
  "base": {
    "code": 200,
    "message": "获取签名URL成功",
    "timestamp": "2025-10-22T10:30:00Z"
  },
  "signed_url": "https://ai-editor-dev.tos-cn-beijing.volces.com/uploads/avatar/20251022/photo.jpg?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=xxx&X-Tos-Date=20251022T103000Z&X-Tos-Expires=3600&X-Tos-SignedHeaders=host&X-Tos-Signature=xxx",
  "expire_seconds": 3600,
  "expire_time": "2025-10-22T11:30:00Z"
}
```

### 示例2：获取文档的24小时访问链接

**请求**:
```bash
GET /api/v1/upload/signed-url?file_url=https://ai-editor-dev.tos-cn-beijing.volces.com/uploads/document/20251022/contract.pdf&expire_seconds=86400
```

**响应**:
```json
{
  "base": {
    "code": 200,
    "message": "获取签名URL成功",
    "timestamp": "2025-10-22T10:30:00Z"
  },
  "signed_url": "https://ai-editor-dev.tos-cn-beijing.volces.com/uploads/document/20251022/contract.pdf?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=xxx&X-Tos-Date=20251022T103000Z&X-Tos-Expires=86400&X-Tos-SignedHeaders=host&X-Tos-Signature=xxx",
  "expire_seconds": 86400,
  "expire_time": "2025-10-23T10:30:00Z"
}
```

### 示例3：获取证书的短期访问链接（10分钟）

**请求**:
```bash
GET /api/v1/upload/signed-url?file_url=https://ai-editor-dev.tos-cn-beijing.volces.com/uploads/cert/id_card/20251022/idcard.jpg&expire_seconds=600
```

**响应**:
```json
{
  "base": {
    "code": 200,
    "message": "获取签名URL成功",
    "timestamp": "2025-10-22T10:30:00Z"
  },
  "signed_url": "https://ai-editor-dev.tos-cn-beijing.volces.com/uploads/cert/id_card/20251022/idcard.jpg?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=xxx&X-Tos-Date=20251022T103000Z&X-Tos-Expires=600&X-Tos-SignedHeaders=host&X-Tos-Signature=xxx",
  "expire_seconds": 600,
  "expire_time": "2025-10-22T10:40:00Z"
}
```

## 前端集成示例

### JavaScript/TypeScript

```javascript
// 上传文件并获取签名URL
async function uploadAndGetSignedUrl(file, uploadType = 'document') {
  // 1. 上传文件
  const formData = new FormData();
  formData.append('file', file);
  formData.append('upload_type', uploadType);
  
  const uploadResponse = await fetch('/api/v1/upload/file', {
    method: 'POST',
    body: formData
  });
  
  const uploadData = await uploadResponse.json();
  const fileUrl = uploadData.file_url;
  
  // 2. 获取签名URL（1小时有效）
  const signedUrlResponse = await fetch(
    `/api/v1/upload/signed-url?file_url=${encodeURIComponent(fileUrl)}&expire_seconds=3600`
  );
  
  const signedData = await signedUrlResponse.json();
  return signedData.signed_url;
}

// 使用示例
const file = document.getElementById('fileInput').files[0];
const signedUrl = await uploadAndGetSignedUrl(file);
// 使用signedUrl下载或显示文件
window.location.href = signedUrl;
```

### React示例

```jsx
import React, { useState } from 'react';

function FileUploader() {
  const [signedUrl, setSignedUrl] = useState('');
  const [loading, setLoading] = useState(false);

  const handleUpload = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    setLoading(true);
    try {
      // 上传文件
      const formData = new FormData();
      formData.append('file', file);
      
      const uploadRes = await fetch('/api/v1/upload/file', {
        method: 'POST',
        body: formData
      });
      const uploadData = await uploadRes.json();
      
      // 获取签名URL
      const signedRes = await fetch(
        `/api/v1/upload/signed-url?file_url=${encodeURIComponent(uploadData.file_url)}&expire_seconds=3600`
      );
      const signedData = await signedRes.json();
      
      setSignedUrl(signedData.signed_url);
      alert('文件上传成功！链接1小时内有效');
    } catch (error) {
      console.error('上传失败:', error);
      alert('上传失败');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <input 
        type="file" 
        onChange={handleUpload} 
        disabled={loading}
      />
      {loading && <p>上传中...</p>}
      {signedUrl && (
        <div>
          <p>文件已上传</p>
          <a href={signedUrl} target="_blank" rel="noopener noreferrer">
            下载文件（1小时内有效）
          </a>
        </div>
      )}
    </div>
  );
}

export default FileUploader;
```

## 后端集成示例（Go）

### 在业务逻辑中使用

```go
package example

import (
    "github.com/xinjiyuan97/labor-clients/config"
    "github.com/xinjiyuan97/labor-clients/utils"
)

// 示例：为用户生成文件下载链接
func GenerateDownloadLink(fileURL string, expireHours int) (string, error) {
    cfg := config.GetGlobalConfig()
    uploadService, err := utils.GetUploadService(&cfg.OSS)
    if err != nil {
        return "", err
    }
    
    // 转换为秒
    expireSeconds := int64(expireHours * 3600)
    
    signedURL, err := uploadService.GetSignedURL(fileURL, expireSeconds)
    if err != nil {
        return "", err
    }
    
    return signedURL, nil
}

// 示例：批量生成签名URL
func BatchGenerateSignedURLs(fileURLs []string, expireSeconds int64) ([]string, error) {
    cfg := config.GetGlobalConfig()
    uploadService, err := utils.GetUploadService(&cfg.OSS)
    if err != nil {
        return nil, err
    }
    
    signedURLs := make([]string, 0, len(fileURLs))
    for _, fileURL := range fileURLs {
        signedURL, err := uploadService.GetSignedURL(fileURL, expireSeconds)
        if err != nil {
            return nil, err
        }
        signedURLs = append(signedURLs, signedURL)
    }
    
    return signedURLs, nil
}
```

## 常见问题

### Q1: 签名URL过期后还能访问吗？
**A**: 不能。签名URL过期后会返回403错误，需要重新生成新的签名URL。

### Q2: 签名URL可以分享给其他人吗？
**A**: 可以。只要在有效期内，任何人都可以通过签名URL访问文件。如果需要更严格的权限控制，建议结合业务逻辑在应用层进行鉴权。

### Q3: 签名URL的最大有效期是多久？
**A**: 系统限制最大为7天（604800秒）。如果设置超过这个时间，会自动调整为7天。

### Q4: 如何在前端判断签名URL是否过期？
**A**: 可以保存`expire_time`字段，在前端判断当前时间是否超过这个时间。如果接近过期，可以提前重新请求新的签名URL。

```javascript
function isUrlExpired(expireTime) {
  const expireDate = new Date(expireTime);
  const now = new Date();
  return now >= expireDate;
}

// 使用示例
if (isUrlExpired(signedData.expire_time)) {
  // 重新获取签名URL
  const newSignedUrl = await getSignedUrl(fileUrl);
}
```

### Q5: 直接URL和签名URL可以混用吗？
**A**: 可以。可以在数据库中保存直接URL，需要时动态生成签名URL。这样既保证了灵活性，又不需要存储多个URL。

## 安全建议

1. **不要在客户端长期缓存签名URL**: 签名URL包含敏感信息，不建议在客户端存储
2. **根据场景设置合适的过期时间**: 
   - 临时预览：5-10分钟
   - 正常使用：1-2小时
   - 长期分享：1-7天
3. **敏感文件建议使用签名URL**: 如证件照、合同等
4. **定期轮换Access Key**: 增强安全性
5. **监控异常访问**: 关注TOS的访问日志，及时发现异常

## 性能优化

1. **缓存签名URL**: 在有效期内可以复用签名URL，减少API调用
2. **批量生成**: 如果需要展示多个文件，可以批量生成签名URL
3. **异步生成**: 对于非关键路径，可以异步生成签名URL
4. **CDN配合**: 签名URL也可以通过CDN加速访问

