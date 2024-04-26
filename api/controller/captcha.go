package controller

import (
	"admin-go-api/api/service"
	"admin-go-api/common/result"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  captcha.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 17:47
 */

// 验证码
// @Summary 验证码接口
// @Tags 工具接口
// @Produce json
// @Description 验证码接口
// @Success 200 {object} result.Result
// @router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, base64Image := service.CaptMake()
	result.Success(c, map[string]interface{}{"idKey": id, "image": base64Image})
}
