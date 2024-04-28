package entity

/**
 * @Author: 南宫乘风
 * @Description: 用户与角色相关模型
 * @File:  sysAdminRole.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 12:10
 */

type SysAdminRole struct {
	RoleId  uint `gorm:"column:role_id;comment:'角色id';NOT NULL" json:"roleId"`
	AdminId uint `gorm:"column:admin_id;comment:'用户id';NOT NULL" json:"menuId"`
}

func (SysAdminRole) TableName() string {
	return "sys_admin_role"
}
