package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
	"admin-go-api/common/result"
	"admin-go-api/common/util"
	"admin-go-api/pkg/jwt"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 用户服务层
 * @File:  sysAdmin.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 11:45
 */

// ISysAdminService 定义接口
type ISysAdminService interface {
	Login(c *gin.Context, dto entity.LoginDto)
}

// SysAdminServiceImpl 实现ISysAdminService接口
type SysAdminServiceImpl struct{}

// Login 用户登录
func (s SysAdminServiceImpl) Login(c *gin.Context, dto entity.LoginDto) {
	// 登录参数校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.MissingLoginParameter),
			result.ApiCode.GetMessage(result.ApiCode.MissingLoginParameter))
		return
	}
	// 验证码是否过期
	code := util.RedisStore{}.Get(dto.IdKey, true)
	if len(code) == 0 {
		result.Failed(c, int(result.ApiCode.VerificationCodeHasExpired),
			result.ApiCode.GetMessage(result.ApiCode.VerificationCodeHasExpired))
		return
	}
	// 校验验证码
	verifyRes := CaptVerify(dto.IdKey, dto.Image)
	if !verifyRes {
		result.Failed(c, int(result.ApiCode.CAPTCHANOTTRUE),
			result.ApiCode.GetMessage(result.ApiCode.CAPTCHANOTTRUE))
		return
	}
	// 校验
	sysAdmin := dao.SysAdminDetail(dto)
	if sysAdmin.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.PASSWORDNOTTRUE),
			result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}
	const status int = 2
	if sysAdmin.Status == status {
		result.Failed(c, int(result.ApiCode.STATUSISENABLE),
			result.ApiCode.GetMessage(result.ApiCode.STATUSISENABLE))
		return
	}
	jwtAdmin := entity.JwtAdmin{
		ID:       sysAdmin.ID,
		Username: sysAdmin.Username,
		Nickname: sysAdmin.Nickname,
		Icon:     sysAdmin.Icon,
		Email:    sysAdmin.Email,
		Phone:    sysAdmin.Phone,
		Note:     sysAdmin.Note,
	}
	// 生成token
	tokenString, _ := jwt.GenerateTokenByAdmin(jwtAdmin)
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdmin})
}

var sysAdminService = SysAdminServiceImpl{}

func SysAdminService() ISysAdminService {
	return &sysAdminService
}
