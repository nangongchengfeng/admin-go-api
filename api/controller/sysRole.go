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
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
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
// @Param pageSize query int false "PageSize"
// @Param pageNum query int false "PageNum"
// @Param roleName query string false "RoleName"
// @Param status query string false "Status"
// @Param beginTime query string false "BeginTime"
// @Param endTime query string false "EndTime"
// @Success 200 {object} result.Result
// @router /api/role/list [get]
// @Security ApiKeyAuth
func GetSysRoleList(c *gin.Context) {
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	RoleName := c.Query("roleName")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	service.SysRoleService().GetSysRoleList(c, PageNum, PageSize, RoleName, Status, BeginTime, EndTime)
}

// QuerySysRoleVoList 获取角色列表
// @Summary 获取角色列表
// @Description: 获取角色列表
// @Tags 角色管理
// @Produce json
// @Success 200 {object} result.Result
// @router /api/role/vo/list [get]
// @Security ApiKeyAuth
func QuerySysRoleVoList(c *gin.Context) {
	service.SysRoleService().QuerySysRoleVoList(c)
}

// QueryRoleMenuIdList 根据角色id查询菜单数据
// @Summary 根据角色id查询菜单数据接口
// @Produce json
// @Tags 角色管理
// @Description 根据角色id查询菜单数据接口
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/role/vo/idList [get]
// @Security ApiKeyAuth
func QueryRoleMenuIdList(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysRoleService().QueryRoleMenuIdList(c, Id)
}

// AssignPermissions 分配权限
// @Summary 分配权限
// @Description: 分配权限
// @Tags 角色管理
// @Produce json
// @Param data body entity.RoleMenu true "data"
// @Success 200 {object} result.Result
// @router /api/role/assignPermissions [PUT]
// @Security ApiKeyAuth
func AssignPermissions(c *gin.Context) {
	var RoleMenu entity.RoleMenu
	_ = c.ShouldBind(&RoleMenu)
	service.SysRoleService().AssignPermissions(c, RoleMenu)
}
