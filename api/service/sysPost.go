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
	GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string)
	GetSysPostById(c *gin.Context, id int)
	UpdateSysPost(c *gin.Context, sysPost entity.SysPost)
	DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto)
}

type SysPostSeerviceImpl struct{}

// DeleteSysPostById 删除岗位
func (s SysPostSeerviceImpl) DeleteSysPostById(c *gin.Context, dto entity.SysPostIdDto) {
	dao.DeleteSysPostById(dto)
	result.Success(c, true)
}

// UpdateSysPost 修改岗位
func (s SysPostSeerviceImpl) UpdateSysPost(c *gin.Context, sysPost entity.SysPost) {
	sysPost = dao.UpdateSysPost(sysPost)
	result.Success(c, sysPost)
}

// GetSysPostById 通过ID获取系统岗位信息
func (s SysPostSeerviceImpl) GetSysPostById(c *gin.Context, id int) {
	// 调用dao层方法，根据id获取系统岗位信息，并通过result封装返回给前端
	result.Success(c, dao.GetSysPostById(id))
}

// GetSysPostList 分页查询岗位列表
func (s SysPostSeerviceImpl) GetSysPostList(c *gin.Context, PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	sysPost, count := dao.GetSysPostList(PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{
		"total":    count,
		"pageSize": PageSize,
		"pageNum":  PageNum,
		"list":     sysPost,
	})
}

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
