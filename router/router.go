package router

import (
	"admin-go-api/api/controller"
	"admin-go-api/common/config"
	"admin-go-api/middleware"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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
	//Swag
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 用户接口
	router.GET("/api/captcha", controller.Captcha)
	router.POST("/api/login", controller.Login)

	// 岗位接口
	router.POST("/api/post/add", controller.CreateSysPost)
	router.GET("/api/post/list", controller.GetSysPostList)
	router.GET("/api/post/info", controller.GetSysPostById)
	router.PUT("/api/post/update", controller.UpdateSysPost)
	router.DELETE("/api/post/delete", controller.DeleteSysPostById)
	router.DELETE("/api/post/batch/delete", controller.BatchDeleteSysPost)
	router.PUT("/api/post/updateStatus", controller.UpdateSysPostStatus)
	router.GET("/api/post/vo/list", controller.QuerySysPostVoList)

	// 部门接口
	router.GET("/api/dept/list", controller.GetSysDeptList)
	router.POST("/api/dept/add", controller.CreateSysDept)
	router.GET("/api/dept/info", controller.GetSysDeptById)
	router.PUT("/api/dept/update", controller.UpdateSysDept)
	router.DELETE("/api/dept/delete", controller.DeleteSysDeptById)
	router.GET("/api/dept/vo/list", controller.QuerySysDeptVoList)

	// 菜单接口
	router.POST("/api/menu/add", controller.CreateSysMenu)
}
