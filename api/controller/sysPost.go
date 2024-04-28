package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"
	"fmt"
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
// @Security ApiKeyAuth
func CreateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().CreateSysPost(c, sysPost)
}

// GetSysPostList 获取岗位列表
// @Summary 获取岗位列表接口
// @Tags 岗位管理
// @Produce json
// @Description 获取岗位列表接口
// @Param PageNum query string false "分页数"
// @Param PageSize query string false "每页数量"
// @Param PostName query string false "岗位名称"
// @Param PostStatus query string false "状态： 1->启动 2->停用"
// @Param BeginTime query string false "开始时间"
// @Param EndTime query string false "结束时间"
// @Success 200 {object} result.Result
// @router /api/post/list [get]
// @Security ApiKeyAuth
func GetSysPostList(c *gin.Context) {
	var PageNum, PageSize int
	var PostName, PostStatus, BeginTime, EndTime string
	PageNum, _ = strconv.Atoi(c.Query("PageNum"))
	PageSize, _ = strconv.Atoi(c.Query("PageSize"))
	PostName = c.Query("PostName")
	PostStatus = c.Query("PostStatus")
	BeginTime = c.Query("BeginTime")
	EndTime = c.Query("EndTime")
	fmt.Println("PageNum:", PageNum, "PageSize:", PageSize, "PostName:", PostName, "PostStatus:", PostStatus, "BeginTime:", BeginTime, "EndTime:", EndTime)
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
// @Security ApiKeyAuth
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
// @router /api/post/update [put]
// @Security ApiKeyAuth
func UpdateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().UpdateSysPost(c, sysPost)
}

// DeleteSysPostById 删除岗位
// @Summary 删除岗位接口
// @Tags 岗位管理
// @Produce json
// @Description 删除岗位接口
// @Param data body entity.SysPostIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/delete [delete]
// @Security ApiKeyAuth
func DeleteSysPostById(c *gin.Context) {
	var dto entity.SysPostIdDto
	_ = c.BindJSON(&dto)
	service.SysPostService().DeleteSysPostById(c, dto)
}

// BatchDeleteSysPost 批量删除岗位
// @Summary 批量删除岗位接口
// @Tags 岗位管理
// @Produce json
// @Description 批量删除岗位接口
// @Param data body entity.DelSysPostDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/batch/delete [delete]
// @Security ApiKeyAuth
func BatchDeleteSysPost(c *gin.Context) {
	var dto entity.DelSysPostDto
	_ = c.BindJSON(&dto)
	service.SysPostService().BatchDeleteSysPost(c, dto)
}

// UpdateSysPostStatus 修改岗位状态
// @Summary 角色状态启用/停用接口
// @Tags 岗位管理
// @Produce json
// @Description 修改岗位状态接口
// @Param data body entity.UpdateSysPostStatusDto true "data"
// @Success 200 {object} result.Result
// @router /api/post/updateStatus [put]
// @Security ApiKeyAuth
func UpdateSysPostStatus(c *gin.Context) {
	var dto entity.UpdateSysPostStatusDto
	_ = c.BindJSON(&dto)
	service.SysPostService().UpdateSysPostStatus(c, dto)
}

// QuerySysPostVoList 获取岗位列表
// @Summary 获取岗位列表接口
// @Tags 岗位管理
// @Produce json
// @Description 获取岗位列表接口
// @Success 200 {object} result.Result
// @router /api/post/vo/list [get]
// @Security ApiKeyAuth
func QuerySysPostVoList(c *gin.Context) {
	service.SysPostService().QuerySysPostVoList(c)
}
