package service

import (
	"admin-go-api/common/config"
	"admin-go-api/common/result"
	"admin-go-api/common/util"
	"fmt"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @Author: 南宫乘风
 * @Description: 图片上传服务层
 * @File:  Upload.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 15:17
 */

type IUploadService interface {
	Upload(c *gin.Context)
}
type UploadServiceImpl struct{}

// Upload 实现图片上传功能
// 参数:
// - c *gin.Context: Gin框架的上下文对象，用于处理HTTP请求和响应
// 返回值: 无
func (u UploadServiceImpl) Upload(c *gin.Context) {
	// 从HTTP请求中获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		// 如果获取文件失败，则返回上传错误的信息
		result.Failed(c, int(result.ApiCode.FILEUPLOADERROR),
			result.ApiCode.GetMessage(result.ApiCode.FILEUPLOADERROR))
	}

	// 获取当前时间，用于生成唯一的文件名
	now := time.Now()
	// 获取文件的扩展名
	ext := path.Ext(file.Filename)
	// 生成唯一的文件名，使用纳秒时间戳加上原始扩展名
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	// 构建文件保存路径，按照年、月、日分类保存
	filePath := fmt.Sprintf("%s%s%s%s",
		config.Config.ImageSettings.UploadDir,
		fmt.Sprintf("%04d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%04d", now.Day()))
	// 创建文件保存的目录
	util.CreateDir(filePath)
	// 拼接完整的文件路径
	fullPath := filePath + "/" + fileName
	// 保存上传的文件到指定路径
	c.SaveUploadedFile(file, fullPath)
	// 返回上传成功的消息，包括保存的文件路径
	result.Success(c, config.Config.ImageSettings.ImageHost+fullPath)
}

var uploadService = UploadServiceImpl{}

func UploadService() IUploadService {
	return &uploadService
}
