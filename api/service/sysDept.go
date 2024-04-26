package service

import (
	"admin-go-api/api/dao"
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
}
type SysDeptServiceImpl struct {
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
