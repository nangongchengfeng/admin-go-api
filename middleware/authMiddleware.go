package middleware

import (
	"admin-go-api/common/constant"
	"admin-go-api/common/result"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 鉴权中间件
 * @File:  authMiddleware.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 16:19
 */

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 鉴权逻辑
		// 验证token
		authHeader := c.Request.Header.Get("Authorization")
		// 如果验证失败，返回401 Unauthorized错误
		if authHeader == "" {
			result.Failed(c, int(result.ApiCode.NOAUTH),
				result.ApiCode.GetMessage(result.ApiCode.NOAUTH))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.Failed(c, int(result.ApiCode.AUTHFORMATERROR),
				result.ApiCode.GetMessage(result.ApiCode.AUTHFORMATERROR))
			c.Abort()
			return
		}

		// 如果验证成功，继续处理请求
		token := "token"
		c.Set(constant.ContextKeyUserObj, token)
		c.Next()

	}
}
