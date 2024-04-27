package entity

/**
 * @Author: 南宫乘风
 * @Description: 角色菜单相关模型
 * @File:  sysRoleMenu.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-27 10:25
 */

// SysRoleMenu 角色与菜单关系模型
type SysRoleMenu struct {
	RoleId uint `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"roleId"`
	MenuId uint `gorm:"column:menu_id;comment:'用户id';NOT NULL" json:"menuId"`
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
