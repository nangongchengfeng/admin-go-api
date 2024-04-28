package controller

import (
	"admin-go-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysOperationLog.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 20:46
 */

// GetSysOperationLogList 分页获取操作日志列表
// @Summary 分页获取操作日志列表接口
// @Produce json
// @Tags 操作日志管理
// @Description 分页获取操作日志列表接口
// @Param PageSize query int false "每页数"
// @Param PageNum query int false "分页数"
// @Param Username query string false "用户名"
// @Param BeginTime query string false "开始时间"
// @Param EndTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/list [get]
// @Security ApiKeyAuth
func GetSysOperationLogList(c *gin.Context) {
	Username := c.Query("username")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	service.SysOperationLogService().GetSysOperationLogList(c, Username,
		BeginTime, EndTime, PageSize, PageNum)
}
