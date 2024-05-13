package controller

import (
	"admin-go-api/api/service"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 主机信息接口
 * @File:  sysMonitor.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-05-13 10:34
 */

// GetHostInfo 获取主机信息
// @Summary 获取主机信息
// @Description 获取主机信息
// @Tags 主机信息
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"cpu":{"cpuNum":4,"cpuNumUsed":0,"cpuRate":0},"disk":{"diskNum":1,"diskNumUsed":0,"diskRate":0},"mem":{"memNum":0,"memNumUsed":0,"memRate":0},"net":{"netNum":0,"netNumUsed":0,"netRate":0}}}"
// @Router /api/hostInfo [get]
func GetHostInfo(c *gin.Context) {
	service.SysMonitorService().GetHostInfo(c)
}
