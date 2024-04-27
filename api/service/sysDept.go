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
}

type SysDeptServiceImpl struct{}

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
