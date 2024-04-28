package dao

import (
	"admin-go-api/api/entity"
	"admin-go-api/common/util"
	. "admin-go-api/pkg/db"
	"time"
)

/**
 * @Author: 南宫乘风
 * @Description: 角色数据层
 * @File:  sysRole.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 09:19
 */

// GetSysRoleByName 根据角色名称查询角色
func GetSysRoleByName(roleName string) (sysRole entity.SysRole) {
	Db.Where("role_name = ?", roleName).First(&sysRole)
	return sysRole
}

// GetSysRoleByKey 根据角色名称查询角色
func GetSysRoleByKey(roleKey string) (sysRole entity.SysRole) {
	Db.Where("role_key = ?", roleKey).First(&sysRole)
	return sysRole
}

// CreateSysRole 创建一个系统角色。
// 参数 dto: 包含新角色信息的 AddSysRoleDto 结构体。
// 返回值 bool: 表示角色是否成功创建。成功返回 true，失败返回 false。
func CreateSysRole(dto entity.AddSysRoleDto) bool {
	// 通过角色名检查该角色是否已存在
	sysRoleByName := GetSysRoleByName(dto.RoleName)
	if sysRoleByName.ID > 0 {
		return false
	}
	// 通过角色键检查该角色是否已存在
	sysRoleByKey := GetSysRoleByKey(dto.RoleKey)
	if sysRoleByKey.ID > 0 {
		return false
	}
	// 创建一个新的系统角色实例
	sysRole := entity.SysRole{
		RoleName:    dto.RoleName,
		RoleKey:     dto.RoleKey,
		Status:      dto.Status,
		Description: dto.Description,
		CreateTime:  util.HTime{Time: time.Now()},
	}
	// 尝试在数据库中创建该角色
	tx := Db.Create(&sysRole)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// GetSysRoleById 根据角色id查询角色
func GetSysRoleById(Id int) (sysRole entity.SysRole) {
	Db.Where("id = ?", Id).First(&sysRole)
	return sysRole
}

// UpdateSysRole  用于更新系统角色的信息。
// 参数 dto 为更新角色所需的数据传输对象，包含要更新的角色ID、状态、角色名、角色键以及可选的角色描述。
// 返回值为更新后的角色实体。
func UpdateSysRole(dto entity.UpdateSysRoleDto) (sysRole entity.SysRole) {
	// 根据ID从数据库中找到对应的系统角色
	Db.First(&sysRole, dto.ID)
	// 更新角色的状态、角色名、角色键
	sysRole.Status = dto.Status
	sysRole.RoleName = dto.RoleName
	sysRole.RoleKey = dto.RoleKey
	// 如果描述不为空，则更新角色的描述
	if dto.Description != "" {
		sysRole.Description = dto.Description
	}
	// 保存更新后的角色到数据库
	Db.Save(&sysRole)
	return sysRole
}