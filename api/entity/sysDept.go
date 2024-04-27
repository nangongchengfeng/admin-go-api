package entity

import "admin-go-api/common/util"

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysDept.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 18:30
 */

// SysDept 部门模型
type SysDept struct {
	ID         uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                         // ID
	ParentId   uint       `gorm:"column:parent_id;comment:'父id';NOT NULL" json:"parentId"`                      // 父id
	DeptType   uint       `gorm:"column:dept_type;comment:'部门类型（1->公司, 2->中心，3->部门）';NOT NULL" json:"deptType"` // 部门类型（1->公司, 2->中心，3->部门）
	DeptName   string     `gorm:"column:dept_name;varchar(30);comment:'部门名称';NOT NULL" json:"deptName"`         // 部门名称
	DeptStatus int        `gorm:"column:dept_status;default:1;comment:'部门状态（1->正常 2->停用）'" json:"deptStatus"`   // 部门状态（1->正常 2->停用）
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"`                 // 创建时间
	Children   []SysDept  `json:"children" gorm:"-"`
	// 子集
}

func (SysDept) TableName() string {
	return "sys_dept"
}

type SysDeptIdDto struct {
	Id int `json:"id"`
}

type SysDeptVo struct {
	Id       int    `json:"id"`
	ParentId uint   `json:"parentId"`
	Label    string `json:"label"`
}
