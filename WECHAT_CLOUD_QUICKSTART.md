# 微信云托管快速开始指南

本文档提供快速部署和测试微信云托管功能的步骤。

## 已完成的工作

✅ 创建了优化后的 Dockerfile（多阶段构建）  
✅ 定义了微信用户信息测试接口 `/api/v1/system/wechat-user-info`  
✅ 实现了接口handler和logic逻辑  
✅ 添加了详细的日志输出  
✅ 创建了完整的部署文档  

## 快速测试

### 1. 本地构建测试

```bash
# 确保在项目根目录
cd /Users/jiyuanxin/work/src/github.com/xinjiyuan97/labor-clients-be

# 构建项目
go build -o output/labor-clients

# 运行服务器
./output/labor-clients -mode server -env example

# 在另一个终端测试接口（需要手动添加微信请求头）
curl -X GET \
  http://localhost:8888/api/v1/system/wechat-user-info \
  -H "X-WX-OPENID: test_openid_12345" \
  -H "X-WX-UNIONID: test_unionid_12345" \
  -H "X-WX-APPID: wx1234567890abcdef" \
  -H "X-WX-ENV: test-env-123" \
  -H "X-WX-CLOUDBASE-ACCESS-TOKEN: test_token"
```

### 2. Docker构建测试

```bash
# 构建Docker镜像
docker build -t labor-clients:latest .

# 运行容器
docker run -d \
  --name labor-clients-test \
  -p 8888:8888 \
  -v $(pwd)/conf:/app/conf \
  labor-clients:latest

# 查看日志
docker logs -f labor-clients-test

# 测试接口
curl -X GET \
  http://localhost:8888/api/v1/system/wechat-user-info \
  -H "X-WX-OPENID: test_openid_12345" \
  -H "X-WX-UNIONID: test_unionid_12345" \
  -H "X-WX-APPID: wx1234567890abcdef" \
  -H "X-WX-ENV: test-env-123" \
  -H "X-WX-CLOUDBASE-ACCESS-TOKEN: test_token"

# 清理
docker stop labor-clients-test
docker rm labor-clients-test
```

## 文件清单

### 新增/修改的文件

1. **Dockerfile** - 多阶段构建的Docker配置
2. **idls/system.thrift** - 添加了 `GetWeChatUserInfo` 接口定义
3. **biz/handler/system/get_we_chat_user_info.go** - 接口handler
4. **biz/logic/system/get_wechat_user_info.go** - 业务逻辑实现
5. **biz/router/system/system.go** - 自动生成的路由注册（已包含新接口）
6. **docs/wechat_cloud_deployment.md** - 完整的部署和使用文档

### 自动生成的文件

以下文件由 `make generate_all` 自动生成，不需要手动修改：

- `biz/model/system/*.go` - Model结构体
- `biz/router/system/*.go` - 路由配置

## 接口说明

### 接口路径
`GET /api/v1/system/wechat-user-info`

### 功能
从微信云托管请求头中提取并返回微信用户信息（openid、unionid等），同时记录到日志中。

### 返回格式
```json
{
  "base": {
    "code": 200,
    "message": "获取微信用户信息成功",
    "timestamp": "2025-01-23T10:00:00Z"
  },
  "openid": "微信openid",
  "unionid": "微信unionid（可选）",
  "appid": "小程序appid",
  "env": "环境ID",
  "cloudbase_access_token": "访问令牌"
}
```

### 日志输出示例
```json
{
  "level": "info",
  "msg": "获取微信用户信息",
  "openid": "test_openid_12345",
  "unionid": "test_unionid_12345",
  "appid": "wx1234567890abcdef",
  "env": "test-env-123",
  "cloudbase_access_token": "test_token",
  "has_openid": true,
  "has_unionid": true,
  "has_appid": true,
  "has_env": true,
  "has_cloudbase_access_token": true,
  "time": "2025-01-23T10:00:00Z"
}
```

## 部署到微信云托管

### 步骤1: 准备代码

确保代码已提交到Git仓库或准备上传镜像。

### 步骤2: 创建云托管服务

1. 登录 [微信云托管控制台](https://console.cloud.tencent.com/tcb)
2. 创建环境（如已存在则跳过）
3. 创建服务，服务名称为 `labor-clients`

### 步骤3: 配置服务

- **端口**: 8888
- **环境变量**: `ENV=prod`（可选）
- **Dockerfile路径**: `/Dockerfile`

### 步骤4: 发布部署

1. 关联Git仓库或上传代码
2. 点击"发布新版本"
3. 等待构建和部署完成

### 步骤5: 测试接口

从微信小程序或网页调用接口：

```javascript
// 小程序调用示例
const res = await wx.cloud.callContainer({
  config: {
    env: '你的云环境ID',
  },
  path: '/api/v1/system/wechat-user-info',
  method: 'GET',
  header: {
    'X-WX-SERVICE': 'labor-clients',
  }
});

console.log('用户信息:', res.data);
```

## 监控和调试

### 查看日志

1. 进入云托管控制台
2. 选择对应服务
3. 点击"服务日志"查看实时日志

### 日志字段说明

- `has_openid`: 是否获取到openid
- `has_unionid`: 是否获取到unionid
- 其他字段为空则表示未从微信云托管调用

## 常见问题

### Q: 为什么本地测试时所有字段都为空？

A: 微信云托管只在线上环境自动注入请求头。本地测试需要使用mock请求头（参考上面的curl示例）。

### Q: 如何获取UnionID？

A: 需要满足以下条件：
1. 小程序已绑定到微信开放平台
2. 用户已授权登录
3. 用户满足微信UnionID获取条件

### Q: 如何修改接口定义？

A: 修改 `idls/system.thrift` 后执行 `make generate_all` 重新生成代码。

## 下一步

1. 根据实际需求修改接口逻辑
2. 添加更多的微信用户信息处理
3. 集成到现有的用户系统中
4. 添加缓存机制提高性能

## 参考文档

详细文档请参考：[wechat_cloud_deployment.md](./docs/wechat_cloud_deployment.md)

