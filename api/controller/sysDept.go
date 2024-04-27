package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"

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
func CreateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().CreateSysDept(c, sysDept)
}
