package entity

import "admin-go-api/common/util"

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysAdmin.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 11:13
 */

// SysAdmin 用户模型对象
type SysAdmin struct {
	ID     uint `gorm:"column:id;comment:'主键';primaryKey;NOT NULL"json:"id"` //ID
	PostId int  `gorm:"column:post_id;comment:'岗位id'" json:"postId"`
	// 岗位id
	DeptId int `gorm:"column:dept_id;comment:'部门id'" json:"deptId"`
	// 部门id
	Username   string     `gorm:"column:username;varchar(64);comment:'用户账号';NOT NULL" json:"username"`         // 用户账号
	Password   string     `gorm:"column:password;varchar(64);comment:'密码';NOTNULL" json:"password"`            // 密码
	Nickname   string     `gorm:"column:nickname;varchar(64);comment:'昵称'"json:"nickname"`                     // 昵称
	Status     int        `gorm:"column:status;default:1;comment:'帐号启用状态：1->启用,2->禁用';NOT NULL" json:"status"` // 帐号启用状态：1->启用,2->禁用
	Icon       string     `gorm:"column:icon;varchar(500);comment:'头像'"json:"icon"`                            // 头像
	Email      string     `gorm:"column:email;varchar(64);comment:'邮箱'"json:"email"`                           // 邮箱
	Phone      string     `gorm:"column:phone;varchar(64);comment:'电话'"json:"phone"`                           // 电话
	Note       string     `gorm:"column:note;varchar(500);comment:'备注'"json:"note"`                            // 备注
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL"json:"createTime"`                 // 创建时间
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// JwtAdmin 鉴权用户结构体
type JwtAdmin struct {
	ID       uint   `json :"id"`      //ID
	Username string `json:"username"` //用户名
	Nickname string `json:"nickname"` //昵称
	Icon     string `json:"icon"`     //头像
	Email    string `json:"email"`    //邮箱
	Phone    string `json:"phone"`    //电话
	Note     string `json:"note"`     //备注
}

// LoginDto 登录对象
type LoginDto struct {
	Username string `json:"username" validate:"required"`          //用户名
	Password string `json:"password" validate:"required"`          //密码
	Image    string `json:"image" validate:"required,min=4,max=6"` //验证码
	IdKey    string `json:"idKey" validate:"required"`             //uuid
}

// AddSysAdminDto 新增参数
type AddSysAdminDto struct {
	PostId   int    `validate:"required"` // 岗位id
	RoleId   uint   `validate:"required"` // 角色id
	DeptId   int    `validate:"required"` // 部门id
	Username string `validate:"required"` // 用户名
	Password string `validate:"required"` // 密码
	Nickname string `validate:"required"` // 昵称
	Phone    string `validate:"required"` // 手机号
	Email    string `validate:"required"` // 邮箱
	Note     string // 备注
	Status   int    `validate:"required"` // 状态：1->启用,2->禁用
}

// SysAdminInfo 详情视图
type SysAdminInfo struct {
	ID       uint   `json:"id"`       // ID
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Status   int    `json:"status"`   // 状态：1->启用,2->禁用
	PostId   int    `json:"postId"`   // 岗位id
	DeptId   int    `json:"deptId"`   // 部门id
	RoleId   uint   `json:"roleId" `  // 角色id
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Note     string `json:"note"`     // 备注
}

// UpdateSysAdminDto 修改参数
type UpdateSysAdminDto struct {
	Id       uint   // ID
	PostId   int    // 岗位id
	DeptId   int    // 部门id
	RoleId   uint   // 角色id
	Username string // 用户名
	Nickname string // 昵称
	Phone    string // 手机号
	Email    string // 邮箱
	Note     string // 备注
	Status   int    // 状态：1->启用,2->禁用
}

// SysAdminIdDto Id参数
type SysAdminIdDto struct {
	Id uint `json:"id"` // ID
}

// UpdateSysAdminStatusDto 修改状态参数
type UpdateSysAdminStatusDto struct {
	Id     uint `json:"id"`
	Status int  `json:"status"`
}

// ResetSysAdminPasswordDto 重置密码参数
type ResetSysAdminPasswordDto struct {
	Id       uint   // ID
	Password string //密码
}
