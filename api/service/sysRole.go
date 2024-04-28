package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
	"admin-go-api/common/result"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 角色服务层
 * @File:  sysRole.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 09:20
 */

type ISysRoleService interface {
	CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto)
	GetSysRoleById(c *gin.Context, id int)
	UpdateSysRole(c *gin.Context, sysRole entity.UpdateSysRoleDto)
	DeleteSysRoleById(c *gin.Context, dto entity.SysRoleIdDto)
}

type SysRoleServiceImpl struct {
}

// DeleteSysRoleById 删除角色
func (s SysRoleServiceImpl) DeleteSysRoleById(c *gin.Context, dto entity.SysRoleIdDto) {
	dao.DeleteSysRoleById(dto)
	result.Success(c, true)
}

// UpdateSysRole 更新角色
func (s SysRoleServiceImpl) UpdateSysRole(c *gin.Context, sysRole entity.UpdateSysRoleDto) {
	result.Success(c, dao.UpdateSysRole(sysRole))
}

// GetSysRoleById 获取角色
func (s SysRoleServiceImpl) GetSysRoleById(c *gin.Context, id int) {
	sysRole := dao.GetSysRoleById(id)
	result.Success(c, sysRole)
}

// CreateSysRole 创建角色
func (s SysRoleServiceImpl) CreateSysRole(c *gin.Context, dto entity.AddSysRoleDto) {
	isCreate := dao.CreateSysRole(dto)
	if !isCreate {
		result.Failed(c, int(result.ApiCode.ROLENAMEALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.ROLENAMEALREADYEXISTS))
	} else {
		result.Success(c, dto)
	}
}

var sysRoleService = SysRoleServiceImpl{}

func SysRoleService() ISysRoleService {
	return &sysRoleService
}
