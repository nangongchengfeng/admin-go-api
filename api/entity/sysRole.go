package entity

import "admin-go-api/common/util"

/**
 * @Author: 南宫乘风
 * @Description: 角色相关模型
 * @File:  sysRole.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 09:19
 */

// SysRole 角色模型
type SysRole struct {
	ID          uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL"json:"id"`                          // ID
	RoleName    string     `gorm:"column:role_name;varchar(64);comment:'角色名称';NOT NULL" json:"roleName"`         // 角色名称
	RoleKey     string     `gorm:"column:role_key;varchar(64);comment:'权限字符串';NOT NULL" json:"roleKey"`          // 权限字符串
	Status      int        `gorm:"column:status;default:1;comment:'帐号启用状态：1-> 启用,2->禁用';NOT NULL" json:"status"` // 帐号启用状态：1->启用,2->禁用
	Description string     `gorm:"column:description;varchar(500);comment:'描述'" json:"description"`              // 描述
	CreateTime  util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                 // 创建时间
}

func (SysRole) TableName() string {
	return "sys_role"
}

// AddSysRoleDto 新增参数
type AddSysRoleDto struct {
	RoleName    string // 角色名称
	RoleKey     string // 角色key
	Status      int    // 状态：1->启用,2->禁用
	Description string // 描述
}

// UpdateSysRoleDto 更新参数
type UpdateSysRoleDto struct {
	ID          uint
	RoleName    string
	RoleKey     string // 角色key
	Status      int    // 状态：1->启用,2->禁用
	Description string
}

// SysRoleIdDto 删除参数
type SysRoleIdDto struct {
	Id uint `json:"id"`
}

// UpdateSysRoleStatusDto 更新状态参数
type UpdateSysRoleStatusDto struct {
	Id     uint
	Status int // 状态：1->启用,2->禁用
}

// SysRoleVo 角色下拉列表
type SysRoleVo struct {
	ID       int    `json:"id"`
	RoleName string `json:"roleName"`
}

// IdVo 当前角色的菜单权限id
type IdVo struct {
	Id int `json:"id"` // ID
}

// RoleMenu 角色id,菜单id视图
type RoleMenu struct {
	Id      uint   `json:"id" binding:"required"`      // ID
	MenuIds []uint `json:"menuIds" binding:"required"` // 菜单id列表
}
