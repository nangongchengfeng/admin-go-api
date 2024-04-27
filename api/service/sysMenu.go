package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
	"admin-go-api/common/result"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysMenu.go.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-27 10:33
 */

type ISysMenuService interface {
	CreateSysMenu(c *gin.Context, sysMenu entity.SysMenu)
	QuerySysMenuVoList(c *gin.Context)
	UpdateSysMenu(c *gin.Context, sysMenu entity.SysMenu)
	DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto)
	GetSysMenuList(c *gin.Context, MenuName string, MenuStatus string)
}
type SysMenuServiceImpl struct {
}

func (s SysMenuServiceImpl) GetSysMenuList(c *gin.Context, MenuName string, MenuStatus string) {
	result.Success(c, dao.GetSysMenuList(MenuName, MenuStatus))
}

// DeleteSysMenu 删除菜单
func (s SysMenuServiceImpl) DeleteSysMenu(c *gin.Context, dto entity.SysMenuIdDto) {
	isDelete := dao.DeleteSysMenu(dto)
	if !isDelete {
		result.Failed(c, int(result.ApiCode.DELSYSMENUFAILED), result.ApiCode.GetMessage(result.ApiCode.DELSYSMENUFAILED))
	} else {
		result.Success(c, true)
	}
}

// UpdateSysMenu 更新菜单
func (s SysMenuServiceImpl) UpdateSysMenu(c *gin.Context, sysMenu entity.SysMenu) {
	menu := dao.UpdateSysMenu(sysMenu)
	result.Success(c, menu)
}

// QuerySysMenuVoList 查询菜单
func (s SysMenuServiceImpl) QuerySysMenuVoList(c *gin.Context) {
	sysMenuVo := dao.QuerySysMenuVoList()
	result.Success(c, sysMenuVo)
}

// CreateSysMenu 创建菜单
func (s SysMenuServiceImpl) CreateSysMenu(c *gin.Context, sysMenu entity.SysMenu) {
	isCreate := dao.CreateSysMenu(sysMenu)
	if !isCreate {
		result.Failed(c, int(result.ApiCode.MENUISEXIST), result.ApiCode.GetMessage(result.ApiCode.MENUISEXIST))
	} else {
		result.Success(c, true)
	}
}

var sysMenuService = SysMenuServiceImpl{}

func SysMenuService() ISysMenuService {
	return &sysMenuService
}
