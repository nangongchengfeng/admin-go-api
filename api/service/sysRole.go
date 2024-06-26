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
	UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto)
	GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, RoleStatus, BeginTime, EndTime string)
	QuerySysRoleVoList(c *gin.Context)
	QueryRoleMenuIdList(c *gin.Context, id int)
	AssignPermissions(c *gin.Context, menu entity.RoleMenu)
}

type SysRoleServiceImpl struct {
}

// AssignPermissions 分配权限
func (s SysRoleServiceImpl) AssignPermissions(c *gin.Context, menu entity.RoleMenu) {
	result.Success(c, dao.AssignPermissions(menu))
}

// QueryRoleMenuIdList 获取角色菜单列表
func (s SysRoleServiceImpl) QueryRoleMenuIdList(c *gin.Context, Id int) {
	voList := dao.QueryRoleMenuIdList(Id)
	var idList = make([]int, 0)
	for _, vo := range voList {
		idList = append(idList, vo.Id)
	}
	result.Success(c, idList)
}

// QuerySysRoleVoList 获取角色列表
func (s SysRoleServiceImpl) QuerySysRoleVoList(c *gin.Context) {
	result.Success(c, dao.QuerySysRoleVoList())
}

// GetSysRoleList 获取角色列表
func (s SysRoleServiceImpl) GetSysRoleList(c *gin.Context, PageNum, PageSize int, RoleName, RoleStatus, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}

	sysRoleList, count := dao.GetSysRoleList(PageNum, PageSize, RoleName, RoleStatus, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{
		"total":    count,
		"pageSize": PageSize,
		"pageNum":  PageNum,
		"list":     sysRoleList,
	})
}

// UpdateSysRoleStatus 更新角色状态
func (s SysRoleServiceImpl) UpdateSysRoleStatus(c *gin.Context, dto entity.UpdateSysRoleStatusDto) {
	isUpdate := dao.UpdateSysRoleStatus(dto)
	if !isUpdate {
		return
	}
	result.Success(c, true)
}

// DeleteSysRoleById 删除角色
func (s SysRoleServiceImpl) DeleteSysRoleById(c *gin.Context, dto entity.SysRoleIdDto) {
	deletes := dao.DeleteSysRoleById(dto)
	if !deletes {
		result.Failed(c, int(result.ApiCode.DELSYSMUSERAILED), result.ApiCode.GetMessage(result.ApiCode.DELSYSMUSERAILED))
	} else {
		result.Success(c, true)
	}

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
