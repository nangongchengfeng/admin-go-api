package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"

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
