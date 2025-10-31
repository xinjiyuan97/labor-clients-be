# 微信云托管部署文档

本文档说明如何将系统部署到微信云托管，并测试获取微信用户信息的功能。

## 1. 概述

微信云托管会自动在容器内的HTTP请求头中注入微信用户信息，包括：
- `X-WX-OPENID`: 微信小程序用户的openid
- `X-WX-UNIONID`: 微信用户的unionid（可选，需要用户在微信开放平台绑定）
- `X-WX-APPID`: 微信小程序appid
- `X-WX-ENV`: 环境ID
- `X-WX-CLOUDBASE-ACCESS-TOKEN`: 访问令牌

## 2. Dockerfile构建

我们已经创建了优化的多阶段构建Dockerfile，位于项目根目录：

```dockerfile
# 构建阶段
FROM golang:1.24-alpine as builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY conf ./conf
COPY output/conf ./conf

EXPOSE 8888

CMD ["/app/main", "-mode", "server", "-env", "prod"]
```

### 构建镜像

在项目根目录执行：

```bash
docker build -t labor-clients:latest .
```

### 测试镜像

```bash
# 运行容器
docker run -d \
  --name labor-clients \
  -p 8888:8888 \
  -v $(pwd)/conf:/app/conf \
  labor-clients:latest

# 查看日志
docker logs -f labor-clients
```

## 3. 部署到微信云托管

### 3.1 准备工作

1. 登录[微信云托管控制台](https://console.cloud.tencent.com/tcb)
2. 创建云托管环境
3. 在环境中创建服务

### 3.2 配置服务

在服务配置中设置：
- **服务名称**: labor-clients
- **端口**: 8888
- **环境变量**: 
  - `ENV=prod` (可选，使用配置文件)

### 3.3 上传代码

```bash
# 方式1: 通过Git关联仓库自动构建
# 在控制台配置Git仓库地址

# 方式2: 本地构建镜像上传
docker build -t labor-clients:latest .
# 登录腾讯云容器镜像服务，然后推送镜像
```

### 3.4 发布服务

在控制台点击"发布新版本"，等待部署完成。

## 4. 测试微信用户信息接口

### 4.1 接口说明

**接口路径**: `GET /api/v1/system/wechat-user-info`

**功能**: 获取微信云托管容器内注入的微信用户信息

**返回示例**:
```json
{
  "base": {
    "code": 200,
    "message": "获取微信用户信息成功",
    "timestamp": "2025-01-23T10:00:00Z"
  },
  "openid": "oxxxxxxxxxxxxxxxxxxxxxx",
  "unionid": "oxxxxxxxxxxxxxxxxxxxxxx",
  "appid": "wx1234567890abcdef",
  "env": "production-123456",
  "cloudbase_access_token": "xxx"
}
```

### 4.2 从微信小程序调用

在微信小程序中使用 `wx.cloud.callContainer` 调用：

```javascript
const res = await wx.cloud.callContainer({
  config: {
    env: '你的云环境ID', // 云托管环境ID
  },
  path: '/api/v1/system/wechat-user-info',
  method: 'GET',
  header: {
    'X-WX-SERVICE': 'labor-clients', // 服务名称
  }
});

console.log('用户信息:', res.data);
```

### 4.3 从网页调用

在网页中使用云托管SDK调用：

```html
<script src="https://web-9gikcbug35bad3a8-1304825656.tcloudbaseapp.com/sdk/1.3.0/cloud.js"></script>
<script>
  window.onload = async function () {
    var c1 = new cloud.Cloud({
      identityless: true, // 普通网页开发
      resourceAppid: '你的小程序appid',
      resourceEnv: '你的云托管环境ID',
    });
    await c1.init();

    const res = await c1.callContainer({
      path: '/api/v1/system/wechat-user-info',
      method: 'GET',
      header: {
        'X-WX-SERVICE': 'labor-clients'
      }
    });

    console.log('用户信息:', res.data);
  }
</script>
```

### 4.4 从curl调用（本地测试）

```bash
# 需要手动添加微信云托管的请求头
curl -X GET \
  http://localhost:8888/api/v1/system/wechat-user-info \
  -H "X-WX-OPENID: test_openid_12345" \
  -H "X-WX-UNIONID: test_unionid_12345" \
  -H "X-WX-APPID: wx1234567890abcdef" \
  -H "X-WX-ENV: test-env-123" \
  -H "X-WX-CLOUDBASE-ACCESS-TOKEN: test_token"
```

## 5. 日志查看

### 5.1 本地日志

日志会输出到控制台和文件（根据配置）：
- 日志文件: `logs/app.log`
- 日志格式: JSON格式，便于分析

### 5.2 云托管日志

在微信云托管控制台查看：
1. 进入对应服务
2. 点击"服务日志"
3. 查看实时日志

查看微信用户信息的日志示例：
```json
{
  "level": "info",
  "msg": "获取微信用户信息",
  "openid": "oxxxxxxxxxxxxxxxxxxxxxx",
  "unionid": "oxxxxxxxxxxxxxxxxxxxxxx",
  "appid": "wx1234567890abcdef",
  "env": "production-123456",
  "cloudbase_access_token": "xxx",
  "has_openid": true,
  "has_unionid": true,
  "has_appid": true,
  "has_env": true,
  "has_cloudbase_access_token": true,
  "time": "2025-01-23T10:00:00Z"
}
```

## 6. 代码实现说明

### 6.1 接口定义

在 `idls/system.thrift` 中定义了接口：

```thrift
// 获取微信用户信息请求
struct GetWeChatUserInfoReq {
}

// 获取微信用户信息响应
struct GetWeChatUserInfoResp {
    1: common.BaseResp base (api.body="base");
    2: string openid (api.body="openid");
    3: string unionid (api.body="unionid");
    4: string appid (api.body="appid");
    5: string env (api.body="env");
    6: string cloudbase_access_token (api.body="cloudbase_access_token");
}

service SystemService {
    GetWeChatUserInfoResp GetWeChatUserInfo(1: GetWeChatUserInfoReq request) 
        (api.get="/api/v1/system/wechat-user-info");
}
```

### 6.2 Handler实现

在 `biz/handler/system/get_we_chat_user_info.go` 中：

```go
func GetWeChatUserInfo(ctx context.Context, c *app.RequestContext) {
    var req system.GetWeChatUserInfoReq
    err := c.BindAndValidate(&req)
    if err != nil {
        c.String(consts.StatusBadRequest, err.Error())
        return
    }

    resp, err := systemlogic.GetWeChatUserInfoLogic(ctx, c)
    if err != nil {
        c.String(consts.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(consts.StatusOK, resp)
}
```

### 6.3 Logic实现

在 `biz/logic/system/get_wechat_user_info.go` 中：

```go
func GetWeChatUserInfoLogic(ctx context.Context, c *app.RequestContext) (*system.GetWeChatUserInfoResp, error) {
    // 获取微信云托管容器内的用户信息
    openid := string(c.GetHeader("X-WX-OPENID"))
    unionid := string(c.GetHeader("X-WX-UNIONID"))
    appid := string(c.GetHeader("X-WX-APPID"))
    env := string(c.GetHeader("X-WX-ENV"))
    cloudbaseAccessToken := string(c.GetHeader("X-WX-CLOUDBASE-ACCESS-TOKEN"))

    // 记录所有微信用户信息到日志
    utils.LogWithFields(map[string]interface{}{
        "openid": openid,
        "unionid": unionid,
        "appid": appid,
        "env": env,
        "cloudbase_access_token": cloudbaseAccessToken,
        "has_openid": openid != "",
        "has_unionid": unionid != "",
        "has_appid": appid != "",
        "has_env": env != "",
        "has_cloudbase_access_token": cloudbaseAccessToken != "",
    }).Info("获取微信用户信息")

    return &system.GetWeChatUserInfoResp{
        Base: &common.BaseResp{
            Code:      200,
            Message:   "获取微信用户信息成功",
            Timestamp: time.Now().Format(time.RFC3339),
        },
        Openid:               openid,
        Unionid:              unionid,
        Appid:                appid,
        Env:                  env,
        CloudbaseAccessToken: cloudbaseAccessToken,
    }, nil
}
```

## 7. 注意事项

1. **UnionID获取条件**: 
   - 需要用户已授权小程序登录
   - 小程序需要在微信开放平台绑定
   - 用户需要关注同主体下的微信公众号或在开放平台绑定过其他应用

2. **本地开发**: 
   - 本地环境不会自动注入微信请求头
   - 可以通过mock请求头进行测试
   - 日志会显示 `has_*` 字段为 false

3. **安全性**: 
   - 不要将微信用户信息泄露到日志中（生产环境）
   - 建议对敏感信息进行脱敏处理
   - 在生产环境中可以通过日志级别控制输出

4. **性能**: 
   - 微信云托管会自动处理请求头注入
   - 无需额外调用微信API
   - 响应速度取决于服务本身的处理时间

## 8. 故障排查

### 8.1 无法获取到openid

**原因**: 
- 不是从微信云托管调用
- 服务名称配置错误

**解决**:
- 检查请求头中是否正确设置了 `X-WX-SERVICE`
- 确认调用方式是 `wx.cloud.callContainer` 或云托管SDK

### 8.2 UnionID为空

**原因**:
- 小程序未绑定微信开放平台
- 用户未授权或未满足条件

**解决**:
- 在微信公众平台绑定小程序到微信开放平台
- 检查用户授权状态

### 8.3 构建失败

**原因**:
- 依赖下载失败
- Go版本不匹配

**解决**:
```bash
# 清理并重新下载依赖
go clean -modcache
go mod download

# 确认Go版本
go version
```

## 9. 相关文档

- [微信云托管官方文档](https://developers.weixin.qq.com/miniprogram/dev/wxcloudservice/wxcloudrun/)
- [微信云托管快速入门-Golang](https://developers.weixin.qq.com/miniprogram/dev/wxcloudservice/wxcloudrun/src/quickstart/custom/golang.html)
- [小程序登录](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/login.html)
- [UnionID机制说明](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/union-id.html)

## 10. 更新代码

如果需要修改接口定义：

1. 修改 `idls/system.thrift`
2. 执行 `make generate_all` 重新生成代码
3. 重新构建镜像
4. 更新云托管服务

```bash
# 重新生成代码
make generate_all

# 构建镜像
docker build -t labor-clients:latest .

# 推送到云托管
# 在控制台执行发布操作
```

