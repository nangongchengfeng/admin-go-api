package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"

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

// 新增岗位
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
