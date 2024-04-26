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
 * @File:  sysPost.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 14:06
 */

type ISysPostService interface {
	CreateSysPost(c *gin.Context, sysPost entity.SysPost)
}

type SysPostSeerviceImpl struct{}

// CreateSysPost 新增岗位
func (s SysPostSeerviceImpl) CreateSysPost(c *gin.Context, sysPost entity.SysPost) {
	// 尝试在数据库中创建系统岗位
	isCreated := dao.CreateSysPost(sysPost)

	// 如果岗位创建失败（已存在），则返回失败响应
	if !isCreated {
		code := result.ApiCode.POSTALREADYEXISTS
		result.Failed(c, int(code), result.ApiCode.GetMessage(code))
		return
	}
	// 岗位创建成功，返回成功响应
	result.Success(c, true)
}

var sysPostService = SysPostSeerviceImpl{}

func SysPostService() ISysPostService {
	return &sysPostService
}
