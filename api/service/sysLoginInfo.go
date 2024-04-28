package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
	"admin-go-api/common/result"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysLoginInfo.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 18:18
 */

type ISysLoginInfoService interface {
	GetSysLoginInfoList(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageSize, pageNum int)
	BatchDeleteSysLoginInfo(c *gin.Context, dto entity.DelSysLoginInfoDto)
	DeleteSysLoginInfo(c *gin.Context, dto entity.SysLoginInfoIdDto)
	CleanSysLoginInfo(c *gin.Context)
}

type SysLoginInfoServiceImpl struct{}

func (s SysLoginInfoServiceImpl) GetSysLoginInfoList(c *gin.Context, Username, LoginStatus, BeginTime, EndTime string, PageSize, pageNum int) {
	if PageSize < 1 {
		PageSize = 10
	}
	if pageNum < 1 {
		pageNum = 1
	}
	sysLoginInfo, count := dao.GetSysLoginInfoList(Username, LoginStatus, BeginTime, EndTime, PageSize, pageNum)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize,
		"pageNum": pageNum, "list": sysLoginInfo})
}

// BatchDeleteSysLoginInfo 批量删除登录日志
func (s SysLoginInfoServiceImpl) BatchDeleteSysLoginInfo(c *gin.Context, dto entity.DelSysLoginInfoDto) {
	dao.BatchDeleteSysLoginInfo(dto)
	result.Success(c, true)
}

// DeleteSysLoginInfo 根据id删除登录日志
func (s SysLoginInfoServiceImpl) DeleteSysLoginInfo(c *gin.Context, dto entity.SysLoginInfoIdDto) {
	dao.DeleteSysLoginInfoById(dto)
	result.Success(c, true)
}

// CleanSysLoginInfo 清空登录日志
func (s SysLoginInfoServiceImpl) CleanSysLoginInfo(c *gin.Context) {
	dao.CleanSysLoginInfo()
	result.Success(c, true)
}

var sysLoginInfoService = SysLoginInfoServiceImpl{}

func SysLoginInfoService() ISysLoginInfoService {
	return &sysLoginInfoService
}
