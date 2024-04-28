package entity

import "admin-go-api/common/util"

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysLoginInfo.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-04-28 18:18
 */

// SysLoginInfo 登录日志
type SysLoginInfo struct {
	ID            uint       `gorm:"column:id;comment:'主键';primaryKey;NOT NULL"json:"id"`                    // ID
	Username      string     `gorm:"column:username;varchar(50);comment:'用户账号'"json:"username"`              // 用户账号
	IpAddress     string     `gorm:"column:ip_address;varchar(128);comment:'登录IP地址'" json:"ipAddress"`       // 登录IP地址
	LoginLocation string     `gorm:"column:login_location;varchar(255);comment:'登录地点'" json:"loginLocation"` // 登录地点
	Browser       string     `gorm:"column:browser;varchar(50);comment:'浏览器类型'" json:"browser"`              // 浏览器类型
	Os            string     `gorm:"column:os;varchar(50);comment:'操作系统'" json:"os"`                         // 操作系统
	LoginStatus   int        `gorm:"column:login_status;comment:'登录状态（1-成功 2-失败）'"json:"loginStatus"`        // 登录状态（1-成功 2-失败）
	Message       string     `gorm:"column:message;varchar(255);comment:'提示消息'" json:"message"`              // 提示消息
	LoginTime     util.HTime `gorm:"column:login_time;comment:'访问时间'" json:"loginTime"`                      // 访问时间
}

func (SysLoginInfo) TableName() string {
	return "sys_login_info"
}

// SysLoginInfoIdDto Id参数
type SysLoginInfoIdDto struct {
	Id uint `json:"id"` // ID
}
