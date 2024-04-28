package util

import "os"

/**
 * @Author: 南宫乘风
 * @Description: 图片上传工具类
 * @File:  uploadTool.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 15:14
 */

// CreateDir 创建目录
func CreateDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// IsExist 判断是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
