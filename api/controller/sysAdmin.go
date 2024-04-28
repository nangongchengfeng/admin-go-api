package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysAdmin.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 11:45
 */

// Login
// @Summary 用户登录接口
// @Tags 用户管理
// @Produce json
// @Description 用户登录接口
// @Param data body entity.LoginDto true "data"
// @Success 200 {object} result.Result
// @router /api/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().Login(c, dto)
}

// CreateSysAdmin 创建用户
// @Summary 创建用户
// @Tags 用户管理
// @Produce json
// @Description 创建用户
// @Param data body entity.AddSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/add [post]
func CreateSysAdmin(c *gin.Context) {
	var dto entity.AddSysAdminDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().CreateSysAdmin(c, dto)
}

// GetSysAdminInfo 获取用户详情
// @Summary 根据id查询用户详情
// @Tags 用户管理
// @Produce json
// @Description 根据id查询用户详情
// @Param id query int true "id"
// @Success 200 {object} result.Result
// @router /api/admin/info [get]
func GetSysAdminInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	service.SysAdminService().GetSysAdminInfo(c, id)
}

// UpdateSysAdmin 修改用户信息
// @Summary 修改用户信息
// @Tags 用户管理
// @Produce json
// @Description 修改用户信息
// @Param data body entity.UpdateSysAdminDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/update [put]
func UpdateSysAdmin(c *gin.Context) {
	var dto entity.UpdateSysAdminDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdmin(c, dto)
}

// DeleteSysAdminById 根据ID删除用户
// @Summary 根据ID删除用户
// @Tags 用户管理
// @Produce json
// @Description 根据ID删除用户
// @Param data body entity.SysAdminIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/delete [delete]
func DeleteSysAdminById(c *gin.Context) {
	var dto entity.SysAdminIdDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().DeleteSysAdminById(c, dto)
}

// UpdateSysAdminStatus 修改用户状态
// @Summary 修改用户状态
// @Tags 用户管理
// @Produce json
// @Description 修改用户状态
// @Param data body entity.UpdateSysAdminStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updateStatus [put]
func UpdateSysAdminStatus(c *gin.Context) {
	var dto entity.UpdateSysAdminStatusDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdminStatus(c, dto)
}
