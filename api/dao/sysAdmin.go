package dao

import (
	"admin-go-api/api/entity"
	. "admin-go-api/pkg/db"
)

/**
 * @Author: 南宫乘风
 * @Description: 用户数据层
 * @File:  sysAdmin.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 11:41
 */

// SysAdminDetail 用户详情
func SysAdminDetail(dto entity.LoginDto) (sysAdmin entity.SysAdmin) {
	// 获取传入的username
	username := dto.Username
	// 使用username查询数据库
	Db.Where("username = ?", username).First(&sysAdmin)
	// 返回查询结果
	return sysAdmin
}
