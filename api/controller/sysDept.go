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
 * @File:  sysDept.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 18:30
 */
var sysDept entity.SysDept

// GetSysDeptList 获取部门列表
// @Summary 获取部门列表
// @Tags 部门管理
// @Produce json
// @Param DeptName query string false "部门名称"
// @Param DeptStatus query string false "部门状态： 1->启动 2->停用"
// @Success 200 {object} result.Result
// @router /api/dept/list [get]
// @Security ApiKeyAuth
func GetSysDeptList(c *gin.Context) {
	DeptName := c.Query("DeptName")
	DeptStatus := c.Query("DeptStatus")
	service.SysDeptService().GetSysDeptList(c, DeptName, DeptStatus)
}

// CreateSysDept 创建部门
// @Summary 创建部门
// @Tags 部门管理
// @Produce json
// @Param sysDept body entity.SysDept true "部门信息"
// @Success 200 {object} result.Result
// @router /api/dept/add  [post]
// @Security ApiKeyAuth
func CreateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().CreateSysDept(c, sysDept)
}

// GetSysDeptById 根据id获取部门信息
// @Summary 根据id获取部门信息
// @Tags 部门管理
// @Produce json
// @Param id query int true "部门id"
// @Success 200 {object} result.Result
// @router /api/dept/info [get]
// @Security ApiKeyAuth
func GetSysDeptById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysDeptService().GetSysDeptById(c, Id)
}

// UpdateSysDept 修改部门信息
// @Summary 修改部门信息
// @Tags 部门管理
// @Produce json
// @Param sysDept body entity.SysDept true "部门信息"
// @Success 200 {object} result.Result
// @router /api/dept/update [put]
// @Security ApiKeyAuth
func UpdateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().UpdateSysDept(c, sysDept)
}

// DeleteSysDeptById 根据id删除部门
// @Summary 根据id删除部门
// @Tags 部门管理
// @Produce json
// @Param dto body entity.SysDeptIdDto true "部门id"
// @Success 200 {object} result.Result
// @router /api/dept/delete [delete]
// @Security ApiKeyAuth
func DeleteSysDeptById(c *gin.Context) {
	var dto entity.SysDeptIdDto
	_ = c.BindJSON(&dto)
	service.SysDeptService().DeleteSysDeptById(c, dto)
}

// QuerySysDeptVoList 查询部门树状结构
// @Summary 查询部门树状结构
// @Tags 部门管理
// @Produce json
// @Success 200 {object} result.Result
// @router /api/dept/vo/list [get]
// @Security ApiKeyAuth
func QuerySysDeptVoList(c *gin.Context) {
	service.SysDeptService().QuerySysDeptVoList(c)
}
