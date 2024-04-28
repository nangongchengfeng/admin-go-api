package controller

import (
	"admin-go-api/api/service"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  Upload.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 15:17
 */

// Upload 单图片上传
// @Summary 单图片上传接口
// @Description 单图片上传接口
// @Tags 工具接口
// @Produce json
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} result.Result
// @Router /api/upload [post]
func Upload(c *gin.Context) {
	service.UploadService().Upload(c)
}
