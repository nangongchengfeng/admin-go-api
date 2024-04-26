package dao

import (
	"admin-go-api/api/entity"
	. "admin-go-api/pkg/db"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysDept.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 18:30
 */

// GetSysDeptList 查询部门列表
//
// 参数:
// DeptName string - 部门名称，如果提供此参数，则查询条件将包含部门名称的匹配。
// DeptStatus string - 部门状态，如果提供此参数，则查询条件将包含部门状态的匹配。
//
// 返回值:
// sysDept []entity.SysDept - 查询结果，为SysDept类型的切片，包含所有匹配的部门记录。
func GetSysDeptList(DeptName string, DeptStatus string) (sysDept []entity.SysDept) {
	curDb := Db.Table("sys_dept") // 从全局Db变量中获取当前数据库连接，并指定查询的表为"sys_dept"
	if DeptName != "" {
		// 如果提供了部门名称，则添加对应的查询条件
		curDb = curDb.Where("dept_name = ?", DeptName)
	}
	if DeptStatus != "" {
		// 如果提供了部门状态，则添加对应的查询条件
		curDb = curDb.Where("dept_status = ?", DeptStatus)
	}
	// 执行查询，并将结果存储到sysDept变量中
	curDb.Find(&sysDept)
	return sysDept // 返回查询结果
}
