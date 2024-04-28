package controller

import (
	"admin-go-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysLoginInfo.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 18:18
 */

// GetSysLoginInfoList 分页获取登录日志列表
// @Summary 分页获取登录日志列表接口
// @Produce json
// @Tags 日志管理
// @Description 分页获取登录日志列表接口
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param username query string false "用户名"
// @Param loginStatus query string false "登录状态（1-成功 2-失败）"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysLoginInfo/list [get]
// @Security ApiKeyAuth
func GetSysLoginInfoList(c *gin.Context) {
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	Username := c.Query("username")
	LoginStatus := c.Query("loginStatus")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysLoginInfoService().GetSysLoginInfoList(c, Username, LoginStatus,
		BeginTime, EndTime, PageSize, PageNum)
}
