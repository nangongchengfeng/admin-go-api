package service

import (
	"admin-go-api/common/result"
	"admin-go-api/common/util"
	"fmt"
	"runtime"

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
	// 获取操作系统信息
	os := runtime.GOOS
	// 声明 os_info
	var osInfo util.SysResourceInfo

	// 判断操作系统类型
	switch os {
	case "windows":
		osInfo = util.GetResourceInfo()

		//osInfo = util.SerializeToJson(hostInfo)

		fmt.Println("当前操作系统为 Windows")
	case "linux":
		//osInfo = "当前操作系统为 Linux"
		fmt.Println("当前操作系统为 Linux")
	default:
		//osInfo = "未知操作系统"
		fmt.Println("未知操作系统")
	}
	result.Success(c, osInfo)
}

var sysMonitorService = SysMonitorServiceImpl{}

func SysMonitorService() ISysMonitorService {
	return &sysMonitorService
}
