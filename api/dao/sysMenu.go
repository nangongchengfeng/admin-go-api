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
 * @File:  sysMenu.go.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-27 10:33
 */

// GetSysMenuByName 根据菜单名称查询菜单
func GetSysMenuByName(menuName string) (sysMenu entity.SysMenu) {
	Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}

// CreateSysMenu 创建系统菜单。
// 根据传入的菜单信息，在数据库中创建新的菜单项。
// 如果菜单名称已存在，则不创建并返回false；否则根据菜单类型创建相应的菜单记录，并返回true。
// 参数：
//
//	addSysMenu - 要添加的菜单实体，包含菜单的各种属性。
//
// 返回值：
//
//	bool - 创建成功返回true，否则返回false。
func CreateSysMenu(addSysMenu entity.SysMenu) bool {
	// 通过菜单名检查是否已存在同名菜单
	sysMenuName := GetSysMenuByName(addSysMenu.MenuName)
	if sysMenuName.ID != 0 {
		return false
	}

	// 创建目录类型的菜单
	if addSysMenu.MenuType == 1 {
		sysMenu := entity.SysMenu{
			MenuName:   addSysMenu.MenuName,
			MenuType:   addSysMenu.MenuType,
			Icon:       addSysMenu.Icon,
			Url:        addSysMenu.Url,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
			ParentId:   0,
		}
		Db.Create(&sysMenu)
		return true
	} else if addSysMenu.MenuType == 2 {
		// 创建菜单类型的菜单
		sysMenu := entity.SysMenu{
			ParentId:   addSysMenu.ParentId,
			MenuName:   addSysMenu.MenuName,
			MenuType:   addSysMenu.MenuType,
			Icon:       addSysMenu.Icon,
			Url:        addSysMenu.Url,
			Sort:       addSysMenu.Sort,
			MenuStatus: addSysMenu.MenuStatus,
			Value:      addSysMenu.Value,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	} else if addSysMenu.MenuType == 3 {
		// 创建按钮类型的菜单
		sysMenu := entity.SysMenu{
			ParentId:   addSysMenu.ParentId,
			MenuName:   addSysMenu.MenuName,
			MenuType:   addSysMenu.MenuType,
			Sort:       addSysMenu.Sort,
			MenuStatus: addSysMenu.MenuStatus,
			Value:      addSysMenu.Value,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	}
	return false
}

// QuerySysMenuVoList 查询菜单列表
func QuerySysMenuVoList() (sysMenuVo []entity.SysMenuVo) {
	Db.Table("sys_menu").Select("id,menu_name as label,parent_id").Scan(&sysMenuVo)
	return sysMenuVo
}

// GetSysMenu 根据菜单ID获取菜单信息
func GetSysMenu(Id int) (SysMenu entity.SysMenu) {
	Db.First(&SysMenu, Id)
	return SysMenu
}

// UpdateSysMenu 更新菜单信息
func UpdateSysMenu(menu entity.SysMenu) (sysMenu entity.SysMenu) {
	Db.First(&sysMenu, menu.ID)
	sysMenu.ParentId = menu.ParentId
	sysMenu.MenuName = menu.MenuName
	sysMenu.Icon = menu.Icon
	sysMenu.Value = menu.Value
	sysMenu.MenuType = menu.MenuType
	sysMenu.Url = menu.Url
	sysMenu.MenuStatus = menu.MenuStatus
	sysMenu.Sort = menu.Sort
	Db.Save(&sysMenu)
	return sysMenu
}

// GetSysRoleMenu 根据菜单ID获取菜单权限信息
//
// 参数:
// dto - 包含菜单ID的数据传输对象
//
// 返回值:
// 返回一个实体SysRoleMenu，包含根据菜单ID查询到的菜单权限信息
func GetSysRoleMenu(dto entity.SysMenuIdDto) (sysRoleMenu entity.SysRoleMenu) {
	// 使用菜单ID查询数据库，获取第一条匹配的数据
	Db.Where("menu_id = ?", dto.Id).First(&sysRoleMenu)
	return sysRoleMenu
}

// DeleteSysMenu 删除指定菜单
// 参数 dto: 包含菜单ID的数据传输对象
// 返回值 bool: 删除操作是否成功。成功返回true，失败返回false。
func DeleteSysMenu(dto entity.SysMenuIdDto) bool {
	// 根据菜单ID获取关联的角色菜单信息
	sysRoleMenu := GetSysRoleMenu(dto)
	// 如果该菜单ID有关联的角色菜单，则不删除，返回false
	if sysRoleMenu.MenuId > 0 {
		return false
	}
	// 在数据库中根据ID删除菜单记录
	Db.Where("parent_id = ?", dto.Id).Delete(&entity.SysMenu{})
	// 删除关联的子菜单
	Db.Delete(&entity.SysMenu{}, dto.Id)
	return true
}

// GetSysMenuList 获取系统菜单列表。
// 根据提供的菜单名称和状态过滤菜单项，并返回符合条件的菜单列表。
//
// 参数:
//
//	MenuName string - 菜单名称，留空则不过滤菜单名称。
//	MenuStatus string - 菜单状态，留空则不过滤菜单状态。
//
// 返回值:
//
//	[]*entity.SysMenu - 符合条件的菜单实体列表。
func GetSysMenuList(MenuName string, MenuStatus string) (sysMenu []*entity.SysMenu) {
	// 使用当前数据库连接，查询sys_menu表，按sort字段排序
	curDb := Db.Table("sys_menu").Order("sort")
	// 如果提供了菜单名称，则添加名称过滤条件
	if MenuName != "" {
		curDb.Where("menu_name = ?", MenuName)
	}
	// 如果提供了菜单状态，则添加状态过滤条件
	if MenuStatus != "" {
		curDb.Where("menu_status = ?", MenuStatus)
	}
	// 执行查询，并将结果填充到sysMenu变量中
	curDb.Find(&sysMenu)
	return
}
