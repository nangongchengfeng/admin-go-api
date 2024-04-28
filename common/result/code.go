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
	SUCCESS                                 uint
	FAILED                                  uint
	Message                                 map[uint]string
	NOAUTH                                  uint
	AUTHFORMATERROR                         uint
	INVALIDTOKEN                            uint
	MissingLoginParameter                   uint
	VerificationCodeHasExpired              uint
	CAPTCHANOTTRUE                          uint
	PASSWORDNOTTRUE                         uint
	STATUSISENABLE                          uint
	ROLENAMEALREADYEXISTS                   uint
	MENUISEXIST                             uint
	DELSYSMENUFAILED                        uint
	DEPTISEXIST                             uint
	DEPTISDISTRIBUTE                        uint
	POSTALREADYEXISTS                       uint
	MissingNewAdminParameter                uint
	USERNAMEALREADYEXISTS                   uint
	MissingModificationOfPersonalParameters uint
	MissingChangePasswordParameter          uint
	RESETPASSWORD                           uint
	FILEUPLOADERROR                         uint
}

// ApiCode 状态码

var ApiCode = &Codes{
	SUCCESS:                                 200,
	FAILED:                                  501,
	NOAUTH:                                  401,
	AUTHFORMATERROR:                         405,
	INVALIDTOKEN:                            406,
	MissingLoginParameter:                   407,
	VerificationCodeHasExpired:              408,
	CAPTCHANOTTRUE:                          409,
	PASSWORDNOTTRUE:                         410,
	STATUSISENABLE:                          411,
	ROLENAMEALREADYEXISTS:                   412,
	MENUISEXIST:                             413,
	DELSYSMENUFAILED:                        414,
	DEPTISEXIST:                             415,
	DEPTISDISTRIBUTE:                        416,
	POSTALREADYEXISTS:                       417,
	MissingNewAdminParameter:                418,
	USERNAMEALREADYEXISTS:                   419,
	MissingModificationOfPersonalParameters: 420,
	MissingChangePasswordParameter:          421,
	RESETPASSWORD:                           422,
	FILEUPLOADERROR:                         427,
}

// init 初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:                                 "操作成功",
		ApiCode.FAILED:                                  "操作失败",
		ApiCode.NOAUTH:                                  "请求头的auth为空",
		ApiCode.AUTHFORMATERROR:                         "请求头的auth格式错误",
		ApiCode.INVALIDTOKEN:                            "无效的Token或者登录过期,请重新登录！",
		ApiCode.MissingLoginParameter:                   "缺少登录参数",
		ApiCode.VerificationCodeHasExpired:              "验证码已失效",
		ApiCode.CAPTCHANOTTRUE:                          "验证码不正确，请重新输入",
		ApiCode.PASSWORDNOTTRUE:                         "密码不正确",
		ApiCode.STATUSISENABLE:                          "您的账号已被停用,请联系管理员",
		ApiCode.ROLENAMEALREADYEXISTS:                   "角色名称或权限字符已存在，请重新输入",
		ApiCode.MENUISEXIST:                             "菜单名称已存在，请重新输入",
		ApiCode.DELSYSMENUFAILED:                        "菜单已分配，不能删除",
		ApiCode.DEPTISEXIST:                             "部门名称已存在，请重新输入",
		ApiCode.DEPTISDISTRIBUTE:                        "部门已分配，不能删除",
		ApiCode.POSTALREADYEXISTS:                       "岗位名称或岗位编码已存在，请重新输入",
		ApiCode.MissingNewAdminParameter:                "缺少新增管理员参数",
		ApiCode.USERNAMEALREADYEXISTS:                   "用户名称已存在，请重新输入",
		ApiCode.MissingModificationOfPersonalParameters: "缺少修改个人参数",
		ApiCode.MissingChangePasswordParameter:          "缺少修改密码参数",
		ApiCode.RESETPASSWORD:                           "两次输入的密码不一致，请重新输入",
		ApiCode.FILEUPLOADERROR:                         "文件上传错误",
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
