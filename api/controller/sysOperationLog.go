package controller

import (
	"admin-go-api/api/entity"
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

// DeleteSysOperationLogById 根据id删除操作日志
// @Summary 根据id删除操作日志
// @Produce json
// @Tags 操作日志管理
// @Description 根据id删除操作日志
// @Param data body entity.SysOperationLogIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/delete [delete]
// @Security ApiKeyAuth
func DeleteSysOperationLogById(c *gin.Context) {
	var dto entity.SysOperationLogIdDto
	_ = c.BindJSON(&dto)
	service.SysOperationLogService().DeleteSysOperationLogById(c, dto)
}

// BatchDeleteSysOperationLog 批量删除操作日志
// @Summary 批量删除操作日志接口
// @Produce json
// @Tags 操作日志管理
// @Description 批量删除操作日志接口
// @Param data body entity.BatchDeleteSysOperationLogDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/batch/delete [delete]
// @Security ApiKeyAuth
func BatchDeleteSysOperationLog(c *gin.Context) {
	var dto entity.BatchDeleteSysOperationLogDto
	_ = c.BindJSON(&dto)
	service.SysOperationLogService().BatchDeleteSysOperationLog(c, dto)
}

// CleanSysOperationLog 清空操作日志
// @Summary 清空操作日志接口
// @Produce json
// @Tags 操作日志管理
// @Description 清空操作日志接口
// @Success 200 {object} result.Result
// @router /api/sysOperationLog/clean [delete]
// @Security ApiKeyAuth
func CleanSysOperationLog(c *gin.Context) {
	service.SysOperationLogService().CleanSysOperationLog(c)
}
