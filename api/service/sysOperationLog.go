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
 * @File:  sysOperationLog.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 20:46
 */

// ISysOperationLogService 接口定义
type ISysOperationLogService interface {
	GetSysOperationLogList(c *gin.Context, Username, BeginTime, EndTime string, PageSize, PageNum int)
	DeleteSysOperationLogById(c *gin.Context, dto entity.SysOperationLogIdDto)
	BatchDeleteSysOperationLog(c *gin.Context, dto entity.BatchDeleteSysOperationLogDto)
	CleanSysOperationLog(c *gin.Context)
}

type SysOperationLogServiceImpl struct {
}

// CleanSysOperationLog 清空操作日志
func (s SysOperationLogServiceImpl) CleanSysOperationLog(c *gin.Context) {
	dao.CleanSysOperationLog()
	result.Success(c, true)
}

// BatchDeleteSysOperationLog 批量删除操作日志
func (s SysOperationLogServiceImpl) BatchDeleteSysOperationLog(c *gin.Context,
	dto entity.BatchDeleteSysOperationLogDto) {
	dao.BatchDeleteSysOperationLog(dto)
	result.Success(c, true)
}

// DeleteSysOperationLogById 根据id删除操作日志
func (s SysOperationLogServiceImpl) DeleteSysOperationLogById(c *gin.Context, dto entity.SysOperationLogIdDto) {
	dao.DeleteSysOperationLogById(dto)
	result.Success(c, true)
}

// GetSysOperationLogList 分页查询
func (s SysOperationLogServiceImpl) GetSysOperationLogList(c *gin.Context, Username, BeginTime, EndTime string, PageSize, PageNum int) {
	// 分页查询操作日志列表
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysOperationLog, count := dao.GetSysOperationLogList(Username, BeginTime,
		EndTime, PageSize, PageNum)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize,
		"pageNum": PageNum, "list": sysOperationLog})
}

var sysOperationLogService = SysOperationLogServiceImpl{}

func SysOperationLogService() ISysOperationLogService {
	return &sysOperationLogService
}
