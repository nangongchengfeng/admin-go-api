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

// GetSysPostById 根据提供的id查询系统中的岗位信息。
// 参数:
// id int - 岗位的唯一标识符
// 返回值:
//
//	entity.SysPost - 查询到的岗位信息。
func GetSysPostById(id int) (sysPost entity.SysPost) {
	// 使用ORM框架的Find方法，根据id查询岗位信息并存储到sysPost变量中
	Db.Find(&sysPost, id)
	return sysPost
}

// UpdateSysPost 用于修改岗位信息
// 参数:
// post - 需要修改的岗位实体，包含岗位的全部或部分信息
// 返回值:
// sysPost - 修改后的岗位实体，包含完整的岗位信息
func UpdateSysPost(post entity.SysPost) (sysPost entity.SysPost) {
	// 根据ID查询岗位信息
	Db.Find(&sysPost, post.ID)
	// 更新岗位信息
	sysPost.PostCode = post.PostCode
	sysPost.PostName = post.PostName
	sysPost.PostStatus = post.PostStatus
	// 如果有提供备注信息，则更新备注
	if post.Remark != "" {
		sysPost.Remark = post.Remark
	}
	// 保存更新后的岗位信息
	Db.Save(&sysPost)
	return sysPost
}

// DeleteSysPostById 根据Id删除岗位信息
func DeleteSysPostById(dto entity.SysPostIdDto) {
	Db.Delete(&entity.SysPost{}, dto.Id)
}

// BatchDeleteSysPost 批量删除岗位信息
func BatchDeleteSysPost(dto entity.DelSysPostDto) {
	Db.Where("id in ?", dto.Ids).Delete(&entity.SysPost{})
}

// UpdateSysPostStatus 用于更新岗位的状态。
// 参数 dto: 包含需要更新的岗位ID和新的岗位状态。
func UpdateSysPostStatus(dto entity.UpdateSysPostStatusDto) {
	// 根据ID查询岗位信息
	var sysPost entity.SysPost
	Db.Find(&sysPost, dto.Id)

	// 更新岗位状态
	sysPost.PostStatus = dto.PostStatus

	// 保存更新后的岗位信息
	Db.Save(&sysPost)
}

func QuerySysPostVoList() (sysPostVo []entity.SysPostVo) {
	Db.Table("sys_post").Select("id,post_name").Scan(&sysPostVo)
	return sysPostVo
}
