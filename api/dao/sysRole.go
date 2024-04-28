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

// DeleteSysRoleById 根据角色id删除角色
// 参数:
// dto - 包含要删除的角色id的数据传输对象 (SysRoleIdDto)
// 说明:
// 该函数首先从"sys_role"表中删除指定id的角色，然后从"sys_role_menu"表中删除所有与该角色id相关的角色菜单关联。
func DeleteSysRoleById(dto entity.SysRoleIdDto) {
	// 从"sys_role"表中删除指定id的角色
	Db.Table("sys_role").Delete(&entity.SysRole{}, dto.Id)

	// 从"sys_role_menu"表中删除所有与该角色id相关的角色菜单关联
	Db.Table("sys_role_menu").Where("role_id = ?", dto.Id).Delete(&entity.SysRoleMenu{})
}

// UpdateSysRoleStatus 更新角色状态。
// 该函数接收一个实体UpdateSysRoleStatusDto，其中包含需要更新的角色ID和新的状态，
func UpdateSysRoleStatus(dto entity.UpdateSysRoleStatusDto) bool {
	var sysRole entity.SysRole // 声明一个SysRole类型的变量，用于存储从数据库中查找到的角色

	// 根据dto中的Id查找并加载第一个匹配的SysRole角色
	Db.First(&sysRole, dto.Id)

	// 更新sysRole实体的状态为dto中的状态
	sysRole.Status = dto.Status

	// 使用事务保存更新后的sysRole实体
	tx := Db.Save(&sysRole)

	// 检查是否有行受到影響，若有则返回true，表示更新成功
	if tx.RowsAffected > 0 {
		return true
	}
	return false // 若没有行受到影響，则返回false，表示更新失败或未找到指定角色
}

// GetSysRoleList 获取系统角色列表。
// 参数：
// - PageNum: 请求的页码。
// - PageSize: 每页显示的数量。
// - RoleName: 角色名称，可选，用于筛选指定名称的角色。
// - Status: 状态，可选，用于筛选指定状态的角色。
// - BeginTime: 起始时间，可选，用于筛选创建时间在指定范围内的角色。
// - EndTime: 结束时间，可选，需与BeginTime一起使用，用于筛选创建时间在指定范围内的角色。
// 返回值：
// - sysRole: 符合条件的角色列表。
// - count: 符合条件的角色总数。
func GetSysRoleList(PageNum int, PageSize int, RoleName string, Status string, BeginTime string, EndTime string) (sysRole []*entity.SysRole, count int64) {
	curDb := Db.Table("sys_role") // 从数据库中获取"sys_role"表的引用

	// 根据提供的条件筛选角色
	if RoleName != "" {
		curDb = curDb.Where("role_name = ?", RoleName)
	}
	if BeginTime != "" && EndTime != "" {
		curDb = curDb.Where("create_time BETWEEN ? AND ?", BeginTime, EndTime)
	}
	if Status != "" {
		curDb = curDb.Where("status = ?", Status)
	}

	// 计算符合条件的角色总数
	curDb.Count(&count)

	// 分页查询，并按创建时间降序排列
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time DESC").Find(&sysRole)

	return sysRole, count
}

// QuerySysRoleVoList 角色下拉视图列表。
func QuerySysRoleVoList() (sysRoleVo []entity.SysRoleVo) {
	Db.Table("sys_role").Select("id,role_name").Scan(&sysRoleVo)
	return sysRoleVo
}

// QueryRoleMenuIdList 根据角色Id查询菜单权限
// 参数:
// Id int - 角色的ID
// 返回值:
// idVo []entity.IdVo - 菜单ID的集合
func QueryRoleMenuIdList(Id int) (idVo []entity.IdVo) {
	const menuType int = 3 // 定义菜单类型常量

	// 查询具有指定菜单类型且与给定角色ID相关的菜单ID
	Db.Table("sys_menu sm").
		Select("sm.id").
		Joins("LEFT JOIN sys_role_menu srm ON srm.menu_id = sm.id").
		Joins("LEFT JOIN sys_role sr ON sr.id = srm.role_id ").
		Where("sm.menu_type = ?", menuType).
		Where("sr.id = ?", Id).
		Scan(&idVo)
	return idVo
}

// AssignPermissions 分配权限
func AssignPermissions(menu entity.RoleMenu) (err error) {
	err = Db.Table("sys_role_menu").Where("role_id = ?", menu.Id).Delete(&entity.SysRoleMenu{}).Error
	if err != nil {
		return err
	}
	for _, v := range menu.MenuIds {
		var entity entity.SysRoleMenu
		entity.RoleId = menu.Id
		entity.MenuId = v
		Db.Create(&entity)
	}
	return err

}
