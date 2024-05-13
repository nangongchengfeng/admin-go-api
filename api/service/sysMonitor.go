package service

import (
	"admin-go-api/common/result"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 主机信息服务层
 * @File:  sysMonitor.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-05-13 10:35
 */

type ISysMonitorService interface {
	GetHostInfo(c *gin.Context)
}

type SysMonitorServiceImpl struct {
}

func (s SysMonitorServiceImpl) GetHostInfo(c *gin.Context) {
	result.Success(c, true)

}

var sysMonitorService = SysMonitorServiceImpl{}

func SysMonitorService() ISysMonitorService {
	return &sysMonitorService
}
