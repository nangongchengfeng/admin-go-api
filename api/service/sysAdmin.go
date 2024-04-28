package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
	"admin-go-api/common/result"
	"admin-go-api/common/util"
	"admin-go-api/pkg/jwt"
	"fmt"

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
	UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto)
	ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto)
	GetSysAdminList(c *gin.Context, PageNum, PageSize int, UserName, Status, BeginTime, EndTime string)
	UpdatePersonal(c *gin.Context, dto entity.UpdatePersonalDto)
	UpdatePersonalPassword(c *gin.Context, dto entity.UpdatePersonalPasswordDto)
}

// SysAdminServiceImpl 实现ISysAdminService接口
type SysAdminServiceImpl struct{}

// UpdatePersonalPassword 更新个人密码
// 参数:
// - c *gin.Context: Gin框架的上下文对象，用于处理HTTP请求和响应
// - dto entity.UpdatePersonalPasswordDto: 包含更新密码所需数据的传输对象
// 返回值: 无
func (s SysAdminServiceImpl) UpdatePersonalPassword(c *gin.Context, dto entity.UpdatePersonalPasswordDto) {
	// 验证更新密码所需的DTO参数
	err := validator.New().Struct(dto)
	if err != nil {
		// 如果参数验证失败，返回错误信息
		result.Failed(c, int(result.ApiCode.MissingChangePasswordParameter),
			result.ApiCode.GetMessage(result.ApiCode.MissingChangePasswordParameter))
		return
	}

	// 从JWT中获取当前管理员信息
	sysAdmin, _ := jwt.GetAdmin(c)
	dto.Id = sysAdmin.ID
	//dto.Id = 89
	//var Username string = "admin"
	// 根据用户名获取管理员信息，以验证当前密码是否正确
	sysAdminExist := dao.GetSysAdminByUserName(sysAdmin.Username)
	if sysAdminExist.Password != util.EncryptionMd5(dto.Password) {
		// 如果当前密码不匹配，返回错误信息
		result.Failed(c, int(result.ApiCode.PASSWORDNOTTRUE),
			result.ApiCode.GetMessage(result.ApiCode.PASSWORDNOTTRUE))
		return
	}

	// 验证新密码和重置密码是否一致
	if dto.NewPassword != dto.ResetPassword {
		// 如果不一致，返回错误信息
		result.Failed(c, int(result.ApiCode.RESETPASSWORD),
			result.ApiCode.GetMessage(result.ApiCode.RESETPASSWORD))
		return
	}

	// 对新密码进行加密处理
	dto.NewPassword = util.EncryptionMd5(dto.NewPassword)
	// 更新密码
	sysAdminUpdatePwd := dao.UpdatePersonalPassword(dto)
	// 生成新的JWT token 的结构体
	var jwtAdmin = entity.JwtAdmin{
		ID:       sysAdminUpdatePwd.ID,
		Username: sysAdminUpdatePwd.Username,
		Nickname: sysAdminUpdatePwd.Nickname,
		Icon:     sysAdminUpdatePwd.Icon,
		Email:    sysAdminUpdatePwd.Email,
		Phone:    sysAdminUpdatePwd.Phone,
		Note:     sysAdminUpdatePwd.Note,
	}
	// 生成新的JWT token
	tokenString, _ := jwt.GenerateTokenByAdmin(jwtAdmin)
	// 更新成功，返回成功信息及新的token和管理员信息
	result.Success(c, map[string]interface{}{"token": tokenString, "sysAdmin": sysAdminUpdatePwd})
	return

}

// UpdatePersonal 更新个人信息
func (s SysAdminServiceImpl) UpdatePersonal(c *gin.Context, dto entity.UpdatePersonalDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c,
			int(result.ApiCode.MissingModificationOfPersonalParameters),
			result.ApiCode.GetMessage(result.ApiCode.MissingModificationOfPersonalParameters))
	}
	id, _ := jwt.GetAdminId(c)
	dto.Id = id
	result.Success(c, dao.UpdatePersonal(dto))
}

// GetSysAdminList 获取用户列表
func (s SysAdminServiceImpl) GetSysAdminList(c *gin.Context, PageNum, PageSize int, UserName, Status, BeginTime, EndTime string) {
	if PageSize < 1 {
		PageSize = 10
	}
	if PageNum < 1 {
		PageNum = 1
	}
	fmt.Println(PageSize, PageNum, UserName, Status, BeginTime, EndTime)
	sysAdmin, count := dao.GetSysAdminList(PageSize, PageNum, UserName, Status, BeginTime, EndTime)
	result.Success(c, map[string]interface{}{"total": count, "pageSize": PageSize,
		"pageNum": PageNum, "list": sysAdmin})
	return
}

// ResetSysAdminPassword 重置密码
func (s SysAdminServiceImpl) ResetSysAdminPassword(c *gin.Context, dto entity.ResetSysAdminPasswordDto) {
	dao.ResetSysAdminPassword(dto)
	result.Success(c, true)
}

// UpdateSysAdminStatus 更新用户状态
func (s SysAdminServiceImpl) UpdateSysAdminStatus(c *gin.Context, dto entity.UpdateSysAdminStatusDto) {
	dao.UpdateSysAdminStatus(dto)
	result.Success(c, true)
}

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
