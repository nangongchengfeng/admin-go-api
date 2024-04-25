package result

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  code.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-25 15:52
 */

// Codes Code 定义装置错误码
type Codes struct {
	SUCCESS         uint
	FAILED          uint
	Message         map[uint]string
	NOAUTH          uint
	AUTHFORMATERROR uint
}

// ApiCode 状态码

var ApiCode = &Codes{
	SUCCESS:         200,
	FAILED:          501,
	NOAUTH:          401,
	AUTHFORMATERROR: 405,
}

// init 初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:         "操作成功",
		ApiCode.FAILED:          "操作失败",
		ApiCode.NOAUTH:          "请求头的auth为空",
		ApiCode.AUTHFORMATERROR: "请求头的auth格式错误",
	}
}

// GetMessage 供外部调用
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if ok {
		return message
	}
	return "未知错误"
}
