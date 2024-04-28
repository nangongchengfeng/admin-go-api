package util

import (
	"github.com/gin-gonic/gin"
	useragent "github.com/wenlng/go-user-agent"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  osAndbrowser.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 18:16
 */

func GetOs(c *gin.Context) string {
	userAgent := c.Request.Header.Get("User-Agent")
	os := useragent.GetOsName(userAgent)
	return os
}

// GetBrowser 获取browser
func GetBrowser(c *gin.Context) string {
	userAgent := c.Request.Header.Get("User-Agent")
	browser := useragent.GetBrowserName(userAgent)
	return browser
}
