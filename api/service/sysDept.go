package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
	"admin-go-api/common/result"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysDept.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 18:30
 */

type ISysDeptService interface {
	GetSysDeptList(c *gin.Context, DeptName string, DeptStatus string)
	CreateSysDept(c *gin.Context, sysDept entity.SysDept)
	GetSysDeptById(c *gin.Context, id int)
	UpdateSysDept(c *gin.Context, sysDept entity.SysDept)
	DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto)
	QuerySysDeptVoList(c *gin.Context)
}

type SysDeptServiceImpl struct{}

// QuerySysDeptVoList 查询部门列表
func (s SysDeptServiceImpl) QuerySysDeptVoList(c *gin.Context) {
	dao.QuerySysDeptVoList()
	result.Success(c, dao.QuerySysDeptVoList())
}

// DeleteSysDeptById 删除部门
func (s SysDeptServiceImpl) DeleteSysDeptById(c *gin.Context, dto entity.SysDeptIdDto) {
	isCreate := dao.DeleteSysDeptById(dto)
	if !isCreate {
		result.Failed(c, int(result.ApiCode.DEPTISDISTRIBUTE), result.ApiCode.GetMessage(result.ApiCode.DEPTISDISTRIBUTE))
	} else {
		result.Success(c, true)
	}
}

// UpdateSysDept 更新部门
func (s SysDeptServiceImpl) UpdateSysDept(c *gin.Context, dept entity.SysDept) {
	sysDept := dao.UpdateSysDept(dept)
	result.Success(c, sysDept)
}

// GetSysDeptById 获取部门详情
func (s SysDeptServiceImpl) GetSysDeptById(c *gin.Context, id int) {
	dto := dao.GetSysDeptById(id)
	result.Success(c, dto)
}

// CreateSysDept 创建部门
func (s SysDeptServiceImpl) CreateSysDept(c *gin.Context, sysDept entity.SysDept) {
	isCreate := dao.CreateSysDept(sysDept)
	if !isCreate {
		result.Failed(c, int(result.ApiCode.DEPTISEXIST), result.ApiCode.GetMessage(result.ApiCode.DEPTISEXIST))
	} else {
		result.Success(c, true)
	}
}

// GetSysDeptList 获取部门列表
func (s SysDeptServiceImpl) GetSysDeptList(c *gin.Context, DeptName string, DeptStatus string) {
	sysDept := dao.GetSysDeptList(DeptName, DeptStatus)
	result.Success(c, sysDept)
}

var sysDeptService = SysDeptServiceImpl{}

func SysDeptService() ISysDeptService {
	return &sysDeptService
}
