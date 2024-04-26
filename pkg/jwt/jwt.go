package jwt

import (
	"admin-go-api/api/entity"
	"admin-go-api/common/constant"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

/**
 * @Author: 南宫乘风
 * @Description: jwt 工具类
 * @File:  jwt.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 11:21
 */
// 生成token,解析token,获取当前登录的用户id及用户信息

type userStdClaims struct {
	entity.JwtAdmin
	jwt.StandardClaims
}

// TokenExpireDuration 过期时间
const (
	TokenExpireDuration = time.Hour * 24
)

// Secret token 密钥
var Secret = []byte("admin-api")

var (
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
)

// GenerateTokenByAdmin 根据管理员信息生成token
func GenerateTokenByAdmin(admin entity.JwtAdmin) (string, error) {
	// 复制管理员信息，避免修改原始数据
	var jwtAdmin = entity.JwtAdmin{
		ID:       admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Icon:     admin.Icon,
		Email:    admin.Email,
		Phone:    admin.Phone,
		Note:     admin.Note,
	}

	// 生成token
	c := userStdClaims{
		jwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "backstage",                                // 签发人
		},
	}
	// 使用指定的签名方法和声明创建一个新的token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// todo 设置 redis
	return token.SignedString(Secret)
}

// ValidateToken 解析JWT
func ValidateToken(tokenString string) (*entity.JwtAdmin, error) {
	// 如果token为空，则返回错误
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	// 解析token，无错误则继续，有错误则返回错误
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	// 解析token中的用户信息，无错误则继续，有错误则返回错误
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// 判断token的签名方法是否为HMAC，不是则返回错误
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v",
				token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	// 返回用户信息
	return &claims.JwtAdmin, err
}

// GetAdminId 返回id
func GetAdminId(c *gin.Context) (uint, error) {
	// 从上下文中获取用户对象
	u, exist := c.Get(constant.ContextKeyUserObj)
	// 如果不存在，则返回错误
	if !exist {
		return 0, errors.New("can't get user id")
	}
	// 将用户对象转换为JwtAdmin结构体
	admin, ok := u.(*entity.JwtAdmin)
	// 如果不转换成功，则返回错误
	if ok {
		// 转换成功，返回id
		return admin.ID, nil
	}
	// 如果转换失败，则返回错误
	return 0, errors.New("can't convert to id struct")
}

// GetAdminName 返回用户名
func GetAdminName(c *gin.Context) (string, error) {
	// 获取上下文中的用户对象
	u, exist := c.Get(constant.ContextKeyUserObj)
	// 如果不存在，则返回错误
	if !exist {
		return string(string(0)), errors.New("can't get user name")
	}
	// 将用户对象转换为JwtAdmin类型
	admin, ok := u.(*entity.JwtAdmin)
	// 如果转换成功，则返回用户名
	if ok {
		return admin.Username, nil
	}
	// 如果转换失败，则返回错误
	return string(string(0)), errors.New("can't convert to api name")
}

// GetAdmin 返回admin信息
func GetAdmin(c *gin.Context) (*entity.JwtAdmin, error) {
	// 从上下文中获取用户对象
	u, exist := c.Get(constant.ContextKeyUserObj)
	// 如果不存在，则返回错误
	if !exist {
		return nil, errors.New("can't get api")
	}
	// 将用户对象转换为admin结构体
	admin, ok := u.(*entity.JwtAdmin)
	// 如果转换成功，则返回admin结构体，否则返回错误
	if ok {
		return admin, nil
	}
	return nil, errors.New("can't convert to api struct")
}
