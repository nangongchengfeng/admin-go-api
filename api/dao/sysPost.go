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

// GetSysPostList 分页查询岗位列表
// 参数：
// PageNum: 查询的当前页码
// PageSize: 每页显示的数量
// PostName: 岗位名称，可为空，为空时不过滤岗位名称
// PostStatus: 岗位状态，可为空，为空时不过滤岗位状态
// BeginTime: 查询开始时间，格式为YYYY-MM-DD，可为空，为空时不过滤创建时间下限
// EndTime: 查询结束时间，格式为YYYY-MM-DD，可为空，为空时不过滤创建时间上限
// 返回值：
// sysPost: 查询到的岗位实体列表
// count: 查询到的岗位总数
func GetSysPostList(PageNum, PageSize int, PostName, PostStatus, BeginTime, EndTime string) (sysPost []entity.SysPost, count int64) {
	curDb := Db.Table("sys_post") // 从数据库中获取"sys_post"表的实例
	// 根据提供的条件筛选查询
	if PostName != "" {
		curDb = curDb.Where("post_name = ?", PostName)
	}
	if PostStatus != "" {
		curDb = curDb.Where("post_status = ?", PostStatus)
	}
	if BeginTime != "" {
		curDb = curDb.Where("create_time >= ?", BeginTime)
	}
	if EndTime != "" {
		curDb = curDb.Where("create_time <= ?", EndTime)
	}
	// 计算符合条件的记录总数
	curDb.Count(&count)
	// 进行分页查询，并按创建时间降序排列
	curDb.Limit(PageSize).Offset((PageNum - 1) * PageSize).Order("create_time desc").Find(&sysPost)
	return sysPost, count
}
