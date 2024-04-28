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

func CreateSysOperationLog(dto entity.SysOperationLog) {
	Db.Create(&dto)
}
