package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xinjiyuan97/labor-clients/biz/model/common"
	"github.com/xinjiyuan97/labor-clients/utils"
)

// JWTAuth JWT鉴权中间件
func JWTAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取Authorization header
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(401, &common.BaseResp{
				Code:      401,
				Message:   "缺少Authorization header",
				Timestamp: GetCurrentTime(),
			})
			c.Abort()
			return
		}

		// 提取token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(401, &common.BaseResp{
				Code:      401,
				Message:   "无效的token格式",
				Timestamp: GetCurrentTime(),
			})
			c.Abort()
			return
		}

		// 验证token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(401, &common.BaseResp{
				Code:      401,
				Message:   "无效的token",
				Timestamp: GetCurrentTime(),
			})
			c.Abort()
			return
		}

		// 将用户信息存储到context中
		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Set("jwt_claims", claims)

		c.Next(ctx)
	}
}

// OptionalJWTAuth 可选的JWT鉴权中间件（不强制要求token）
func OptionalJWTAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取Authorization header
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			// 没有token，继续执行但不设置用户信息
			c.Next(ctx)
			return
		}

		// 提取token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			// token格式错误，继续执行但不设置用户信息
			c.Next(ctx)
			return
		}

		// 验证token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			// token无效，继续执行但不设置用户信息
			c.Next(ctx)
			return
		}

		// 将用户信息存储到context中
		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Set("jwt_claims", claims)

		c.Next(ctx)
	}
}

// GetCurrentTime 获取当前时间字符串
func GetCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

// GetUserIDFromContext 从上下文中获取用户ID
func GetUserIDFromContext(c *app.RequestContext) (int64, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	if id, ok := userID.(int64); ok {
		return id, true
	}

	return 0, false
}

// GetUserRoleFromContext 从上下文中获取用户角色
func GetUserRoleFromContext(c *app.RequestContext) (string, bool) {
	userRole, exists := c.Get("user_role")
	if !exists {
		return "", false
	}

	if role, ok := userRole.(string); ok {
		return role, true
	}

	return "", false
}

// GetJWTClaimsFromContext 从上下文中获取JWT claims
func GetJWTClaimsFromContext(c *app.RequestContext) (*utils.Claims, bool) {
	claims, exists := c.Get("jwt_claims")
	if !exists {
		return nil, false
	}

	if jwtClaims, ok := claims.(*utils.Claims); ok {
		return jwtClaims, true
	}

	return nil, false
}

// RequireAuth 要求用户已认证的中间件
func RequireAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userID, exists := GetUserIDFromContext(c)
		if !exists || userID == 0 {
			c.JSON(401, &common.BaseResp{
				Code:      401,
				Message:   "需要登录",
				Timestamp: GetCurrentTime(),
			})
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}

// RequireRole 要求特定角色的中间件
func RequireRole(requiredRole string) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		userRole, exists := GetUserRoleFromContext(c)
		if !exists || userRole != requiredRole {
			c.JSON(403, &common.BaseResp{
				Code:      403,
				Message:   "权限不足",
				Timestamp: GetCurrentTime(),
			})
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}

// RequireWorkerRole 要求worker角色的中间件
func RequireWorkerRole() app.HandlerFunc {
	return RequireRole("worker")
}

// RequireEmployerRole 要求employer角色的中间件
func RequireEmployerRole() app.HandlerFunc {
	return RequireRole("employer")
}

// RequireAdminRole 要求admin角色的中间件
func RequireAdminRole() app.HandlerFunc {
	return RequireRole("admin")
}
