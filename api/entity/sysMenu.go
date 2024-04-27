package entity

import "admin-go-api/common/util"

/**
 * @Author: 南宫乘风
 * @Description: 菜单相关模型
 * @File:  sysMenu.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-27 10:29
 */

// SysMenu 菜单模型
type SysMenu struct {
	ID         uint   `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`              // ID
	ParentId   uint   `gorm:"column:parent_id;comment:'父菜单id'" json:"parentId"`                  // 父菜单id
	MenuName   string `gorm:"column:menu_name;varchar(100);comment:'菜单名称'" json:"menuName"`      // 菜单名称
	Icon       string `gorm:"column:icon;varchar(100);comment:'菜单图标'" json:"icon"`               // 菜单图标
	Value      string `gorm:"column:value;varchar(100);comment:'权限值'" json:"value"`              // 权限值
	MenuType   uint   `gorm:"column:menu_type;comment:'菜单类型：1->目录；2->菜单；3->按钮'" json:"menuType"` // 菜单类型：1->目录；2->菜单；3->按钮
	Url        string `gorm:"column:url;varchar(100);comment:'菜单url'" json:"url"`                // 菜单url
	MenuStatus uint   `gorm:"column:menu_status;comment:'启用状态；1->禁用；2-> 启用'" json:"menuStatus"`  // 启用状态；1->禁用；2->启用
	Sort       uint   `gorm:"column:sort;comment:'排序'" json:"sort"`
	// 排序
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间'" json:"createTime"` // 创建时间
	Children   []SysMenu  `json:"children" gorm:"-"`
	// 子集
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
