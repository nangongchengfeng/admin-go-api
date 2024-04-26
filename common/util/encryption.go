package util

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * @Author: 南宫乘风
 * @Description: 加密工具类
 * @File:  encryption.go.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 11:11
 */

// 加密
func EncryptionMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
