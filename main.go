package main

import (
	"admin-go-api/common/config"
	_ "admin-go-api/docs"
	"admin-go-api/pkg/db"
	"admin-go-api/pkg/log"
	"admin-go-api/pkg/redis"
	"admin-go-api/router"
	"context"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 启动程序
 * @File:  main.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 17:03
 */

// initSwagger 函数用于调用 swag init 命令生成 Swagger 文档
func initSwagger() error {
	cmd := exec.Command("swag", "init")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// @title 通用后台管理系统
// @version 1.0
// @description 后台管理系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := initSwagger()
	if err != nil {
		return
	}
	// 加载日志
	log := log.Log()
	gin.SetMode(config.Config.Server.Model)
	router := router.InitRouter()
	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: router,
	}
	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen: %s \n", err)
		}
		log.Info("listen: %s \n", config.Config.Server.Address)
	}()
	quit := make(chan os.Signal)
	//监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

// 初始化连接
func init() {
	// mysql
	err := db.SetupDBLink()
	if err != nil {
		return
	}
	// redis
	err = redis.SetupRedisDb()
	if err != nil {
		return
	}
}
