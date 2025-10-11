package middleware

import (
	"github.com/gin-gonic/gin"
	"mini_shop/jwt"
	"net/http"
	"strings"
)

// JWTAuthMiddleware 检查 token 是否有效 & 是否有权限访问
func JWTAuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": -1, "msg": "未提供 Token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": -1, "msg": "Token 无效或已过期"})
			c.Abort()
			return
		}

		// 权限判断
		hasPermission := false
		for _, role := range allowedRoles {
			if claims.Role == role {
				hasPermission = true
				break
			}
		}
		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{"code": -1, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 保存用户信息到上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
