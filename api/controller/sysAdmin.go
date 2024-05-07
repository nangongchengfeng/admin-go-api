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
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
func UpdateSysAdminStatus(c *gin.Context) {
	var dto entity.UpdateSysAdminStatusDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdateSysAdminStatus(c, dto)
}

// ResetSysAdminPassword 重置用户密码
// @Summary 重置用户密码
// @Tags 用户管理
// @Produce json
// @Description 重置用户密码
// @Param data body entity.ResetSysAdminPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updatePassword [put]
// @Security ApiKeyAuth
func ResetSysAdminPassword(c *gin.Context) {
	var dto entity.ResetSysAdminPasswordDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().ResetSysAdminPassword(c, dto)
}

// GetSysAdminList 分页获取列表
// @Summary 分页获取列表接口
// @Tags 用户管理
// @Produce json
// @Description 分页获取列表接口
// @Param pageNum query int false "分页数"
// @Param pageSize query int false "每页数"
// @Param username query string false "用户名"
// @Param status query string false "帐号启用状态：1->启用,2->禁用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/admin/list [get]
// @Security ApiKeyAuth
func GetSysAdminList(c *gin.Context) {
	PageNum, _ := strconv.Atoi(c.Query("pageNum"))
	PageSize, _ := strconv.Atoi(c.Query("pageSize"))
	Username := c.Query("username")
	Status := c.Query("status")
	BeginTime := c.Query("beginTime")
	EndTime := c.Query("endTime")
	//fmt.Println("行数", PageSize, "页数", PageNum)
	service.SysAdminService().GetSysAdminList(c, PageSize, PageNum, Username,
		Status, BeginTime, EndTime)
}

// UpdatePersonal 修改个人信息
// @Summary 修改个人信息
// @Tags 用户管理
// @Produce json
// @Description 修改个人信息
// @Param data body entity.UpdatePersonalDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updatePersonal [put]
// @Security ApiKeyAuth
func UpdatePersonal(c *gin.Context) {
	var dto entity.UpdatePersonalDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdatePersonal(c, dto)
}

// UpdatePersonalPassword 修改密码
// @Summary 修改密码接口
// @Tags 用户管理
// @Produce json
// @Description 修改密码接口
// @Param data body entity.UpdatePersonalPasswordDto true "data"
// @Success 200 {object} result.Result
// @router /api/admin/updatePersonalPassword [put]
// @Security ApiKeyAuth
func UpdatePersonalPassword(c *gin.Context) {
	var dto entity.UpdatePersonalPasswordDto
	_ = c.BindJSON(&dto)
	service.SysAdminService().UpdatePersonalPassword(c, dto)
}
