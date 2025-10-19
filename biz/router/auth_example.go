package router

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/xinjiyuan97/labor-clients/biz/handler/auth"
	"github.com/xinjiyuan97/labor-clients/middleware"
)

// RegisterAuthRoutes 注册认证相关路由
func RegisterAuthRoutes(h *server.Hertz) {
	// 不需要鉴权的路由
	authGroup := h.Group("/api/v1/auth")
	{
		// 注册和登录不需要鉴权
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/login", auth.Login)
	}

	// 需要鉴权的路由
	protectedGroup := h.Group("/api/v1/auth")
	protectedGroup.Use(middleware.JWTAuth()) // 使用JWT鉴权中间件
	{
		// 登出需要鉴权
		protectedGroup.POST("/logout", auth.Logout)

		// 刷新token需要鉴权
		protectedGroup.POST("/refresh", auth.RefreshToken)

		// 获取当前用户信息需要鉴权
		protectedGroup.GET("/me", auth.GetCurrentUserProfile)
	}

	// 需要特定角色权限的路由
	workerGroup := h.Group("/api/v1/worker")
	workerGroup.Use(middleware.JWTAuth())           // JWT鉴权
	workerGroup.Use(middleware.RequireWorkerRole()) // 要求worker角色
	{
		// 只有worker可以访问的路由
		workerGroup.PUT("/profile", auth.UpdateUserProfile)
		workerGroup.GET("/profile/:user_id", auth.GetUserProfile)
	}

	// 需要employer角色权限的路由
	employerGroup := h.Group("/api/v1/employer")
	employerGroup.Use(middleware.JWTAuth())             // JWT鉴权
	employerGroup.Use(middleware.RequireEmployerRole()) // 要求employer角色
	{
		// 只有employer可以访问的路由
		employerGroup.GET("/dashboard", func(ctx context.Context, c *app.RequestContext) {
			// employer dashboard logic
		})
	}

	// 可选鉴权的路由（有token就用，没有token也允许访问）
	publicGroup := h.Group("/api/v1/public")
	publicGroup.Use(middleware.OptionalJWTAuth()) // 可选JWT鉴权
	{
		// 这些路由有token时会显示个性化内容，没有token时显示通用内容
		publicGroup.GET("/jobs", func(ctx context.Context, c *app.RequestContext) {
			userID, exists := middleware.GetUserIDFromContext(c)
			if exists {
				// 显示个性化的工作推荐
				c.JSON(200, map[string]interface{}{
					"message": "个性化工作推荐",
					"user_id": userID,
				})
			} else {
				// 显示通用的工作列表
				c.JSON(200, map[string]interface{}{
					"message": "通用工作列表",
				})
			}
		})
	}
}
