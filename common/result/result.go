package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description:通用结构
 * @File:  result.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 15:52
 */

type Result struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 返回的数据
}

func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Msg = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data
	c.JSON(http.StatusOK, res)
}
func Failed(c *gin.Context, code int, message string) {
	res := Result{}
	res.Code = code
	res.Msg = message
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}
