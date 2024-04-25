package router

import (
	"admin-go-api/common/config"
	"admin-go-api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 访问接口路由配置
 * @File:  router.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 16:57
 */

func InitRouter() *gin.Engine {
	router := gin.New()
	// 迭机恢复
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	// 图片访问路径静态文件夹可直接访问
	router.StaticFS(config.Config.ImageSettings.UploadDir,
		http.Dir(config.Config.ImageSettings.UploadDir))
	// 日志log中间件
	router.Use(middleware.Logger())
	register(router)
	return router
}

// register 路由接口
func register(router *gin.Engine) {
	// todo 后续接口url

}
