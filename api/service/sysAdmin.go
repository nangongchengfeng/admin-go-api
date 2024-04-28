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
	CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto)
	GetSysAdminInfo(c *gin.Context, id int)
	UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto)
	DeleteSysAdminById(c *gin.Context, dto entity.SysAdminIdDto)
}

// SysAdminServiceImpl 实现ISysAdminService接口
type SysAdminServiceImpl struct{}

// DeleteSysAdminById 根据ID删除用户
func (s SysAdminServiceImpl) DeleteSysAdminById(c *gin.Context, dto entity.SysAdminIdDto) {
	dao.DeleteSysAdminById(dto)
	result.Success(c, true)
}

// UpdateSysAdmin 更新用户的信息
func (s SysAdminServiceImpl) UpdateSysAdmin(c *gin.Context, dto entity.UpdateSysAdminDto) {
	result.Success(c, dao.UpdateSysAdmin(dto))
}

// GetSysAdminInfo 根据ID获取用户信息
func (s SysAdminServiceImpl) GetSysAdminInfo(c *gin.Context, id int) {
	result.Success(c, dao.GetSysAdminInfo(id))
}

// CreateSysAdmin 创建一个系统管理员用户
// 参数:
// - c *gin.Context: Gin框架的上下文对象，用于处理HTTP请求和响应
// - dto entity.AddSysAdminDto: 添加管理员的数据传输对象，包含新管理员的信息
// 返回值: 无
func (s SysAdminServiceImpl) CreateSysAdmin(c *gin.Context, dto entity.AddSysAdminDto) {
	// 验证传入的DTO数据是否符合规范
	err := validator.New().Struct(dto)
	if err != nil {
		// 如果验证失败，返回缺少参数的错误信息
		result.Failed(c, int(result.ApiCode.MissingNewAdminParameter),
			result.ApiCode.GetMessage(result.ApiCode.MissingNewAdminParameter))
		return
	}

	// 尝试在数据库中创建系统管理员
	isCreate := dao.CreateSysAdmin(dto)
	if !isCreate {
		// 如果创建失败（如用户名已存在），返回相应的错误信息
		result.Failed(c, int(result.ApiCode.USERNAMEALREADYEXISTS),
			result.ApiCode.GetMessage(result.ApiCode.USERNAMEALREADYEXISTS))
		return
	}

	// 创建成功，返回成功的响应
	result.Success(c, true)
}

// Login 用户登录
// Login 登录函数
// :c *gin.Context 上下文
// :dto entity.LoginDto 登录参数
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
	// 校验 用户名和密码 查询数据库，获取用户信息
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

// 定义一个系统管理员服务实现类
var sysAdminService = SysAdminServiceImpl{}

// SysAdminService 定义一个系统管理员服务接口
func SysAdminService() ISysAdminService {
	// 返回系统管理员服务实现类的指针
	return &sysAdminService
}
