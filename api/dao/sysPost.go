package dao

import (
	"admin-go-api/api/entity"
	"admin-go-api/common/util"
	. "admin-go-api/pkg/db"
	"time"
)

/**
 * @Author: 南宫乘风
 * @Description: 岗位数据层
 * @File:  sysPost.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-26 14:06
 */

// GetSysPostByCode 根据编码查询
func GetSysPostByCode(postCode string) (sysPost entity.SysPost) {
	Db.Where("post_code = ?", postCode).First(&sysPost)
	return sysPost
}

// GetSysPostByName 根据名称查询
func GetSysPostByName(postName string) (sysPost entity.SysPost) {
	Db.Where("post_name = ?", postName).First(&sysPost)
	return sysPost
}

// CreateSysPost 新增岗位
func CreateSysPost(sysPost entity.SysPost) bool {
	// 查询岗位代码是否已经存在
	sysPostByCode := GetSysPostByCode(sysPost.PostCode)
	// 如果存在，返回false
	if sysPostByCode.ID > 0 {
		return false
	}
	// 查询岗位名称是否已经存在
	sysPostByName := GetSysPostByName(sysPost.PostName)

	// 如果存在，返回false
	if sysPostByName.ID > 0 {
		return false
	}
	// 将要添加的岗位赋值给addSysPost
	addSysPost := entity.SysPost{
		PostCode:   sysPost.PostCode,
		PostName:   sysPost.PostName,
		PostStatus: sysPost.PostStatus,
		CreateTime: util.HTime{Time: time.Now()},
		Remark:     sysPost.Remark,
	}
	// 保存addSysPost
	tx := Db.Save(&addSysPost)
	// 如果受影响的行数大于0，返回true
	if tx.RowsAffected > 0 {
		return true
	}
	// 否则返回false
	return false
}
