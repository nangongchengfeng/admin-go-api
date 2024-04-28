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

	router.GET("/api/captcha", controller.Captcha)
	router.POST("/api/login", controller.Login)
	// JWT 鉴权接口
	jwt := router.Group("/api", middleware.AuthMiddleware(), middleware.LogMiddleware())
	{
		// 用户接口

		jwt.POST("/upload", controller.Upload)

		jwt.POST("/admin/add", controller.CreateSysAdmin)
		jwt.GET("/admin/info", controller.GetSysAdminInfo)
		jwt.PUT("/admin/update", controller.UpdateSysAdmin)
		jwt.DELETE("/admin/delete", controller.DeleteSysAdminById)
		jwt.PUT("/admin/updateStatus", controller.UpdateSysAdminStatus)
		jwt.PUT("/admin/updatePassword", controller.ResetSysAdminPassword)
		jwt.GET("/admin/list", controller.GetSysAdminList)
		jwt.PUT("/admin/updatePersonal", controller.UpdatePersonal)
		jwt.PUT("/admin/updatePersonalPassword", controller.UpdatePersonalPassword)

		// 岗位接口
		jwt.POST("/post/add", controller.CreateSysPost)
		jwt.GET("/post/list", controller.GetSysPostList)
		jwt.GET("/post/info", controller.GetSysPostById)
		jwt.PUT("/post/update", controller.UpdateSysPost)
		jwt.DELETE("/post/delete", controller.DeleteSysPostById)
		jwt.DELETE("/post/batch/delete", controller.BatchDeleteSysPost)
		jwt.PUT("/post/updateStatus", controller.UpdateSysPostStatus)
		jwt.GET("/post/vo/list", controller.QuerySysPostVoList)

		// 部门接口
		jwt.GET("/dept/list", controller.GetSysDeptList)
		jwt.POST("/dept/add", controller.CreateSysDept)
		jwt.GET("/dept/info", controller.GetSysDeptById)
		jwt.PUT("/dept/update", controller.UpdateSysDept)
		jwt.DELETE("/dept/delete", controller.DeleteSysDeptById)
		jwt.GET("/dept/vo/list", controller.QuerySysDeptVoList)

		// 菜单接口
		jwt.POST("/menu/add", controller.CreateSysMenu)
		jwt.GET("/menu/vo/list", controller.QuerySysMenuVoList)
		jwt.PUT("/menu/update", controller.UpdateSysMenu)
		jwt.DELETE("/menu/delete", controller.DeleteSysMenu)
		jwt.GET("/menu/list", controller.GetSysMenuList)

		// 角色管理
		jwt.POST("/role/add", controller.CreateSysRole)
		jwt.GET("/role/info", controller.GetSysRoleById)
		jwt.PUT("/role/update", controller.UpdateSysRole)
		jwt.DELETE("/role/delete", controller.DeleteSysRoleById)
		jwt.PUT("/role/updateStatus", controller.UpdateSysRoleStatus)
		jwt.GET("/role/list", controller.GetSysRoleList)
		jwt.GET("/role/vo/list", controller.QuerySysRoleVoList)
		jwt.GET("/role/vo/idList", controller.QueryRoleMenuIdList)
		jwt.PUT("/role/assignPermissions", controller.AssignPermissions)

		// 日志管理
		jwt.GET("/sysLoginInfo/list", controller.GetSysLoginInfoList)
		jwt.DELETE("/sysLoginInfo/batch/delete", controller.BatchDeleteSysLoginInfo)
		jwt.DELETE("/sysLoginInfo/delete", controller.DeleteSysLoginInfoById)
		jwt.DELETE("/sysLoginInfo/clean", controller.CleanSysLoginInfo)
	}
}
