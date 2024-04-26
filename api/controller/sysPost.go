package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 岗位控制层
 * @File:  sysPost.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 14:06
 */

var sysPost entity.SysPost

// CreateSysPost 新增岗位
// @Summary 新增岗位接口
// @Tags 岗位管理
// @Produce json
// @Description 新增岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/add [post]
func CreateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().CreateSysPost(c, sysPost)
}

// GetSysPostList 获取岗位列表
// @Summary 获取岗位列表接口
// @Tags 岗位管理
// @Produce json
// @Description 获取岗位列表接口
// @Param pageNum query string false "分页数"
// @Param pageSize query string false "每页数量"
// @Param postName query string false "岗位名称"
// @Param postStatus query string false "状态： 1->启动 2->停用"
// @Param beginTime query string false "开始时间"
// @Param endTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/post/list [get]
func GetSysPostList(c *gin.Context) {
	var PageNum, PageSize int
	var PostName, PostStatus, BeginTime, EndTime string
	PageNum, _ = strconv.Atoi(c.DefaultQuery("PageNum", "1"))
	PageSize, _ = strconv.Atoi(c.DefaultQuery("PageSize", "10"))
	PostName = c.DefaultQuery("PostName", "")
	PostStatus = c.DefaultQuery("PostStatus", "")
	BeginTime = c.DefaultQuery("BeginTime", "")
	EndTime = c.DefaultQuery("EndTime", "")
	service.SysPostService().GetSysPostList(c, PageNum, PageSize, PostName, PostStatus, BeginTime, EndTime)
}

// GetSysPostById 获取岗位详情
// @Summary 根据ID获取岗位详情接口
// @Tags 岗位管理
// @Produce json
// @Description 根据ID获取岗位详情接口
// @Param id query string true "ID"
// @Success 200 {object} result.Result
// @router /api/post/info [get]
func GetSysPostById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysPostService().GetSysPostById(c, Id)
}

// UpdateSysPost 修改岗位
// @Summary 修改岗位接口
// @Tags 岗位管理
// @Produce json
// @Description 修改岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @router /api/post/update [post]
func UpdateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().UpdateSysPost(c, sysPost)
}
