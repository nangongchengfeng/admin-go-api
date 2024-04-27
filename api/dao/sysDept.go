package dao

import (
	"admin-go-api/api/entity"
	"admin-go-api/common/util"
	. "admin-go-api/pkg/db"
	"time"
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

// GetSysDeptByName 根据部门名称获取部门信息
func GetSysDeptByName(deptName string) (sysDept entity.SysDept) {
	Db.Where("dept_name = ?", deptName).First(&sysDept)
	return sysDept
}

// CreateSysDept 创建一个系统部门
// 参数:
// sysDept - 需要创建的部门实体
// 返回值:
// bool - 创建成功返回true，如果部门名称已存在则不创建并返回false
func CreateSysDept(sysDept entity.SysDept) bool {
	// 根据部门名称获取已存在的部门信息
	sysDeptName := GetSysDeptByName(sysDept.DeptName)
	// 如果已存在相同名称的部门，则不创建
	if sysDeptName.ID > 0 {
		return false
	}
	// 部门类型为1时，设置部门父ID为0，表示该部门为顶级部门
	if sysDept.DeptType == 1 {
		sysDept := entity.SysDept{
			DeptStatus: sysDept.DeptStatus,
			ParentId:   0,
			DeptType:   sysDept.DeptType,
			DeptName:   sysDept.DeptName,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysDept) // 创建部门
		return true
	} else {
		// 部门类型不为1时，使用提供的父ID
		sysDept := entity.SysDept{
			DeptStatus: sysDept.DeptStatus,
			ParentId:   sysDept.ParentId,
			DeptType:   sysDept.DeptType,
			DeptName:   sysDept.DeptName,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysDept) // 创建部门
		return true
	}
}

// GetSysDeptById 根据部门ID获取部门信息
func GetSysDeptById(Id int) (sysDept entity.SysDept) {
	Db.First(&sysDept, Id)
	return sysDept
}

// UpdateSysDept 更新部门信息
func UpdateSysDept(dept entity.SysDept) (sysDept entity.SysDept) {
	// 根据ID首先查询到对应的部门信息
	Db.First(&sysDept, dept.ID)
	// 更新部门状态、父部门ID、部门类型和部门名称
	sysDept.DeptStatus = dept.DeptStatus
	sysDept.ParentId = dept.ParentId
	sysDept.DeptType = dept.DeptType
	sysDept.DeptName = dept.DeptName
	// 保存更新后的部门信息
	Db.Save(&sysDept)
	return sysDept
}
