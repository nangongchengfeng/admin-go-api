package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysMenu.go.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-27 10:33
 */
var sysMenu entity.SysMenu

// CreateSysMenu 创建菜单
// @Summary CreateSysMenu 创建菜单
// @Description 创建一个新的系统菜单项
// @Tags 菜单管理
// @Produce json
// @Param sysDept body entity.SysMenu true "部门信息"
// @Success 200 {object} result.Result
// @router /api/menu/add [post]
func CreateSysMenu(c *gin.Context) {

	_ = c.ShouldBindJSON(&sysMenu)
	service.SysMenuService().CreateSysMenu(c, sysMenu)
}

// QuerySysMenuVoList 查询菜单列表
// @Summary QuerySysMenuVoList 查询菜单列表
// @Description 查询菜单列表
// @Tags 菜单管理
// @Produce json
// @Success 200 {object} entity.SysMenuVo
// @router /api/menu/vo/list [get]
func QuerySysMenuVoList(c *gin.Context) {
	service.SysMenuService().QuerySysMenuVoList(c)
}
