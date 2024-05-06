package middleware

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
	"admin-go-api/common/util"
	"admin-go-api/pkg/jwt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 操作日志中间件
 * @File:  LogMiddleware.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 20:42
 */

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := strings.ToLower(c.Request.Method)
		sysAdmin, _ := jwt.GetAdmin(c)
		if method != "get" {
			log := entity.SysOperationLog{
				AdminId:    sysAdmin.ID,
				Username:   sysAdmin.Username,
				Method:     method,
				Ip:         c.ClientIP(),
				Url:        c.Request.URL.Path,
				CreateTime: util.HTime{Time: time.Now()},
			}
			dao.CreateSysOperationLog(log)
		}
	}
}
