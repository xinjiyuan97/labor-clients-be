# JWT鉴权中间件使用说明

## 概述

本中间件提供了完整的JWT鉴权功能，包括token验证、用户信息提取、角色权限控制等功能。

## 中间件类型

### 1. JWTAuth() - 强制JWT鉴权
- **功能**: 强制要求有效的JWT token
- **使用场景**: 需要登录才能访问的API
- **行为**: 如果没有token或token无效，返回401错误

### 2. OptionalJWTAuth() - 可选JWT鉴权
- **功能**: 可选的JWT token验证
- **使用场景**: 有token时显示个性化内容，没有token时显示通用内容
- **行为**: 有token就验证并设置用户信息，没有token也允许访问

### 3. RequireAuth() - 要求已认证
- **功能**: 确保用户已通过JWT鉴权
- **使用场景**: 在OptionalJWTAuth之后使用，确保用户已登录
- **行为**: 检查上下文中是否有用户ID

### 4. RequireRole(role) - 要求特定角色
- **功能**: 要求用户具有特定角色
- **使用场景**: 需要特定权限的API
- **行为**: 检查用户角色是否匹配

### 5. RequireWorkerRole() - 要求worker角色
- **功能**: 要求用户是worker角色
- **使用场景**: 只有worker可以访问的功能

### 6. RequireEmployerRole() - 要求employer角色
- **功能**: 要求用户是employer角色
- **使用场景**: 只有employer可以访问的功能

## 辅助函数

### GetUserIDFromContext(c) - 获取用户ID
```go
userID, exists := middleware.GetUserIDFromContext(c)
if exists {
    // 使用userID
}
```

### GetUserRoleFromContext(c) - 获取用户角色
```go
userRole, exists := middleware.GetUserRoleFromContext(c)
if exists {
    // 使用userRole
}
```

### GetJWTClaimsFromContext(c) - 获取JWT claims
```go
claims, exists := middleware.GetJWTClaimsFromContext(c)
if exists {
    // 使用完整的claims信息
}
```

## 使用示例

### 1. 基础路由配置
```go
// 不需要鉴权的路由
authGroup := h.Group("/api/v1/auth")
{
    authGroup.POST("/register", auth.Register)
    authGroup.POST("/login", auth.Login)
}

// 需要鉴权的路由
protectedGroup := h.Group("/api/v1/auth")
protectedGroup.Use(middleware.JWTAuth())
{
    protectedGroup.POST("/logout", auth.Logout)
    protectedGroup.GET("/me", auth.GetCurrentUserProfile)
}
```

### 2. 角色权限控制
```go
// 只有worker可以访问
workerGroup := h.Group("/api/v1/worker")
workerGroup.Use(middleware.JWTAuth())
workerGroup.Use(middleware.RequireWorkerRole())
{
    workerGroup.PUT("/profile", auth.UpdateUserProfile)
}

// 只有employer可以访问
employerGroup := h.Group("/api/v1/employer")
employerGroup.Use(middleware.JWTAuth())
employerGroup.Use(middleware.RequireEmployerRole())
{
    employerGroup.GET("/dashboard", employerDashboard)
}
```

### 3. 可选鉴权
```go
// 可选鉴权的路由
publicGroup := h.Group("/api/v1/public")
publicGroup.Use(middleware.OptionalJWTAuth())
{
    publicGroup.GET("/jobs", func(ctx context.Context, c *app.RequestContext) {
        userID, exists := middleware.GetUserIDFromContext(c)
        if exists {
            // 显示个性化内容
            c.JSON(200, map[string]interface{}{
                "message": "个性化工作推荐",
                "user_id": userID,
            })
        } else {
            // 显示通用内容
            c.JSON(200, map[string]interface{}{
                "message": "通用工作列表",
            })
        }
    })
}
```

### 4. Handler中使用用户信息
```go
func GetCurrentUserProfile(ctx context.Context, c *app.RequestContext) {
    // 获取用户ID
    userID, exists := middleware.GetUserIDFromContext(c)
    if !exists {
        c.JSON(401, &common.BaseResp{
            Code:      401,
            Message:   "未登录",
            Timestamp: middleware.GetCurrentTime(),
        })
        return
    }

    // 获取用户角色
    userRole, _ := middleware.GetUserRoleFromContext(c)

    // 使用用户信息进行业务逻辑处理
    // ...
}
```

## 中间件链式使用

中间件可以链式使用，按顺序执行：

```go
// 先验证JWT，再检查角色权限
group.Use(middleware.JWTAuth())
group.Use(middleware.RequireWorkerRole())
group.GET("/worker-only", handler)
```

## 错误响应格式

所有中间件都返回统一的错误响应格式：

```json
{
    "code": 401,
    "message": "缺少Authorization header",
    "timestamp": "2024-01-01T00:00:00Z"
}
```

## 注意事项

1. **JWT密钥**: 确保JWT密钥配置正确
2. **Token格式**: 客户端需要在Authorization header中发送 `Bearer <token>` 格式
3. **中间件顺序**: 中间件按添加顺序执行，确保顺序正确
4. **错误处理**: 中间件会自动中断请求并返回错误响应
5. **上下文存储**: 用户信息存储在Hertz的上下文中，可以在后续handler中获取
