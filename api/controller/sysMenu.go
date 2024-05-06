package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"
	"fmt"
	"strconv"

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
// @Security ApiKeyAuth
func CreateSysMenu(c *gin.Context) {

	_ = c.ShouldBindJSON(&sysMenu)
	// 使用 fmt.Printf 打印结构体的每个字段
	fmt.Printf("%+v\n", sysMenu)
	service.SysMenuService().CreateSysMenu(c, sysMenu)
}

// QuerySysMenuVoList 查询菜单列表
// @Summary QuerySysMenuVoList 查询菜单列表
// @Description 查询菜单列表
// @Tags 菜单管理
// @Produce json
// @Success 200 {object} entity.SysMenuVo
// @router /api/menu/vo/list [get]
// @Security ApiKeyAuth
func QuerySysMenuVoList(c *gin.Context) {
	service.SysMenuService().QuerySysMenuVoList(c)
}

// UpdateSysMenu 更新菜单
// @Summary UpdateSysMenu 更新菜单
// @Description 更新菜单
// @Tags 菜单管理
// @Produce json
// @Param sysDept body entity.SysMenu true "部门信息"
// @Success 200 {object} entity.SysMenu
// @router /api/menu/update [put]
// @Security ApiKeyAuth
func UpdateSysMenu(c *gin.Context) {
	_ = c.ShouldBindJSON(&sysMenu)
	service.SysMenuService().UpdateSysMenu(c, sysMenu)
}

// DeleteSysMenu 删除菜单
// @Summary DeleteSysMenu 删除菜单
// @Description 删除菜单
// @Tags 菜单管理
// @Produce json
// @Param sysDept body entity.SysMenuIdDto true "部门信息"
// @Success 200 {object} result.Result
// @router /api/menu/delete [delete]
// @Security ApiKeyAuth
func DeleteSysMenu(c *gin.Context) {
	var dto entity.SysMenuIdDto
	_ = c.ShouldBindJSON(&dto)
	service.SysMenuService().DeleteSysMenu(c, dto)
}

// GetSysMenuList 获取菜单列表
// @Summary GetSysMenuList 获取菜单列表
// @Description 获取菜单列表
// @Tags 菜单管理
// @Produce json
// @Param menuName query string false "菜单名称"
// @Param menuStatus query string false "菜单状态"
// @Success 200 {object} entity.SysMenu
// @router /api/menu/list [get]
// @Security ApiKeyAuth
func GetSysMenuList(c *gin.Context) {
	menuName := c.Query("menuName")
	menuStatus := c.Query("menuStatus")
	service.SysMenuService().GetSysMenuList(c, menuName, menuStatus)
}

// GetSysMenu 获取菜单详情
// @Tags 菜单管理
// @Summary 获取菜单详情
// @Produce json
// @Param id query int true "菜单id"
// @Success 200 {object} entity.SysMenu
// @router /api/menu/info [get]
// @Security ApiKeyAuth
func GetSysMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	service.SysMenuService().GetSysMenu(c, id)
}
