package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 角色控制层
 * @File:  sysRole.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 09:19
 */

var sysRole entity.AddSysRoleDto

// CreateSysRole 创建角色
// @Summary 创建角色
// @Description: 创建角色
// @Tags 角色管理
// @Produce json
// @Param data body entity.AddSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/add [post]
func CreateSysRole(c *gin.Context) {
	_ = c.ShouldBind(&sysRole)
	service.SysRoleService().CreateSysRole(c, sysRole)
}

// GetSysRoleById 获取角色详情
// @Summary 获取角色详情
// @Description: 获取角色详情
// @Tags 角色管理
// @Produce json
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/role/info [get]
func GetSysRoleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().GetSysRoleById(c, id)
}

// UpdateSysRole 修改角色
// @Summary 修改角色
// @Description: 修改角色
// @Tags 角色管理
// @Produce json
// @Param data body entity.UpdateSysRoleDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/update [PUT]
func UpdateSysRole(c *gin.Context) {
	var dto entity.UpdateSysRoleDto
	_ = c.ShouldBind(&dto)
	service.SysRoleService().UpdateSysRole(c, dto)
}

// DeleteSysRoleById 删除角色
// @Summary 删除角色
// @Description: 删除角色
// @Tags 角色管理
// @Produce json
// @Param data body entity.SysRoleIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/delete [DELETE]
func DeleteSysRoleById(c *gin.Context) {
	var dto entity.SysRoleIdDto
	_ = c.ShouldBind(&dto)
	service.SysRoleService().DeleteSysRoleById(c, dto)
}

// UpdateSysRoleStatus 修改角色状态
// @Summary 修改角色状态
// @Description: 修改角色状态
// @Tags 角色管理
// @Produce json
// @Param data body entity.UpdateSysRoleStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/role/updateStatus [PUT]
func UpdateSysRoleStatus(c *gin.Context) {
	var dto entity.UpdateSysRoleStatusDto
	_ = c.ShouldBind(&dto)
	service.SysRoleService().UpdateSysRoleStatus(c, dto)
}

// GetSysRoleList 获取角色列表
// @Summary 获取角色列表
// @Description: 获取角色列表
// @Tags 角色管理
// @Produce json
// @Param PageSize query int false "PageSize"
// @Param PageNum query int false "PageNum"
// @Param RoleName query string false "RoleName"
// @Param Status query string false "Status"
// @Param BeginTime query string false "BeginTime"
// @Param EndTime query string false "EndTime"
// @Success 200 {object} result.Result
// @router /api/role/list [get]
func GetSysRoleList(c *gin.Context) {
	PageSize, _ := strconv.Atoi(c.Query("PageSize"))
	PageNum, _ := strconv.Atoi(c.Query("PageNum"))
	RoleName := c.Query("RoleName")
	Status := c.Query("Status")
	BeginTime := c.Query("BeginTime")
	EndTime := c.Query("EndTime")
	service.SysRoleService().GetSysRoleList(c, PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
}
