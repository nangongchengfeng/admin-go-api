package entity

import "admin-go-api/common/util"

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysOperationLog.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 20:46
 */

// SysOperationLog 操作日志
type SysOperationLog struct {
	ID       uint   `gorm:"column:id;comment:'主键';primaryKey;NOT NULL" json:"id"`                 // ID
	AdminId  uint   `gorm:"column:admin_id;comment:'管理员id';NOT NULL" json:"adminId"`              // 管理员id
	Username string `gorm:"column:username;varchar(64);comment:'管理员账号';NOT NULL" json:"username"` // 管理员账号
	Method   string `gorm:"column:method;varchar(64);comment:'请求方式';NOT NULL" json:"method"`      // 请求方式
	Ip       string `gorm:"column:ip;varchar(64);comment:'IP'; json:"ip"`
	// IP
	Url        string     `gorm:"column:url;varchar(500);comment:'URL'; json:"url"`             // URL
	CreateTime util.HTime `gorm:"column:create_time;comment:'创建时间';NOT NULL" json:"createTime"` // 创建时间
}

func (SysOperationLog) TableName() string {
	return "sys_operation_log"
}

// SysOperationLogIdDto Id参数
type SysOperationLogIdDto struct {
	Id uint `json:"id"` // ID
}

// BatchDeleteSysOperationLogDto 批量删除id参数
type BatchDeleteSysOperationLogDto struct {
	Ids []uint //id列表
}
