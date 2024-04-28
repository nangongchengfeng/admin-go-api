package dao

import (
	"admin-go-api/api/entity"
	"admin-go-api/common/util"
	. "admin-go-api/pkg/db"
	"time"
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

// GetSysAdminByUserName 根据用户名查询用户
func GetSysAdminByUserName(username string) (sysAdmin entity.SysAdmin) {
	Db.Where("username = ?", username).First(&sysAdmin)
	return sysAdmin
}

// CreateSysAdmin 创建一个系统管理员账号。
//
// 参数:
// dto - 添加系统管理员的数据传输对象，包含了新建管理员的全部信息。
//
// 返回值:
// 返回一个布尔值，true 表示管理员创建成功，false 表示创建失败（可能是用户名已存在）。
func CreateSysAdmin(dto entity.AddSysAdminDto) bool {
	// 通过用户名获取已存在的系统管理员，检查是否已存在。
	sysAdminByUSerName := GetSysAdminByUserName(dto.Username)
	if sysAdminByUSerName.ID != 0 {
		return false
	}

	// 构造一个新的系统管理员实例，并设置其属性。
	sysAdmin := entity.SysAdmin{
		PostId:     dto.PostId,
		DeptId:     dto.DeptId,
		Username:   dto.Username,
		Nickname:   dto.Nickname,
		Password:   util.EncryptionMd5(dto.Password), // 使用MD5加密密码。
		Phone:      dto.Phone,
		Email:      dto.Email,
		Note:       dto.Note,
		Status:     dto.Status,
		CreateTime: util.HTime{Time: time.Now()}, // 设置创建时间为当前时间。
	}

	// 在数据库中创建新的系统管理员。
	tx := Db.Create(&sysAdmin)

	// 再次通过用户名获取系统管理员，以确认是否创建成功。
	sysAdminExist := GetSysAdminByUserName(dto.Username)

	// 创建系统管理员角色关系。
	var entity entity.SysAdminRole
	entity.AdminId = sysAdminExist.ID
	entity.RoleId = dto.RoleId
	Db.Create(&entity)

	// 检查是否有行受到影响，即检查创建操作是否成功。
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// GetSysAdminInfo 根据提供的用户ID，查询并返回该用户的详细信息。
// 参数：
// Id - 用户的唯一标识符。
// 返回值：
// sysAdminInfo - 查询到的用户详情，包括用户基本信息和角色信息。
func GetSysAdminInfo(Id int) (sysAdminInfo entity.SysAdminInfo) {
	// 使用ORM查询语言，从数据库中选取特定ID的用户详情。
	// 通过LEFT JOIN关联了sys_admin_role和sys_role表，以获取用户的角色信息。
	Db.Table("sys_admin").
		Select("sys_admin.*, sys_admin_role.role_id").
		Joins("LEFT JOIN sys_admin_role ON sys_admin.id =sys_admin_role.admin_id").
		Joins("LEFT JOIN sys_role ON sys_admin_role.role_id = sys_role.id").
		First(&sysAdminInfo, Id)
	return sysAdminInfo
}

// UpdateSysAdmin 更新系统管理员信息
// 参数:
// dto - 包含要更新的系统管理员信息的数据传输对象
// 返回值:
// sysAdmin - 更新后的系统管理员实体
func UpdateSysAdmin(dto entity.UpdateSysAdminDto) (sysAdmin entity.SysAdmin) {
	// 根据ID从数据库中获取第一个系统管理员记录
	Db.First(&sysAdmin, dto.Id)
	// 如果提供了新用户名，则更新用户名
	if dto.Username != "" {
		sysAdmin.Username = dto.Username
	}
	// 更新职位ID、部门ID和状态
	sysAdmin.PostId = dto.PostId
	sysAdmin.DeptId = dto.DeptId
	sysAdmin.Status = dto.Status
	// 如果提供了新昵称，则更新昵称
	if dto.Nickname != "" {
		sysAdmin.Nickname = dto.Nickname
	}
	// 如果提供了新电话号码，则更新电话号码
	if dto.Phone != "" {
		sysAdmin.Phone = dto.Phone
	}
	// 如果提供了新电子邮件地址，则更新电子邮件地址
	if dto.Email != "" {
		sysAdmin.Email = dto.Email
	}
	// 如果提供了备注信息，则更新备注
	if dto.Note != "" {
		sysAdmin.Note = dto.Note
	}
	// 保存更新后的系统管理员信息到数据库
	Db.Save(&sysAdmin)

	// 删除之前的角色关联，为管理员重新分配角色
	var sysAdminRole entity.SysAdminRole
	// 根据管理员ID删除所有旧的角色关联
	Db.Where("admin_id = ?", dto.Id).Delete(&entity.SysAdminRole{})
	// 创建新的角色关联
	sysAdminRole.AdminId = dto.Id
	sysAdminRole.RoleId = dto.RoleId
	Db.Create(&sysAdminRole)

	return sysAdmin
}
