package dao

import (
	"admin-go-api/api/entity"
	. "admin-go-api/pkg/db"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysOperationLog.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 20:46
 */

// CreateSysOperationLog 添加操作日志
func CreateSysOperationLog(dto entity.SysOperationLog) {
	Db.Create(&dto)
}

// GetSysOperationLogList 分页查询操作日志列表
func GetSysOperationLogList(Username, BeginTime, EndTime string, PageSize,
	PageNum int) (sysOperationLog []entity.SysOperationLog, count int64) {
	curDb := Db.Table("sys_operation_log")
	if Username != "" {
		curDb = curDb.Where("username =?", Username)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("`create_time` BETWEEN ? AND ?", BeginTime, EndTime)
	}
	curDb.Count(&count)
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&sysOperationLog)
	return sysOperationLog, count
}
