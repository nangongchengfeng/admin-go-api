package entity

import "syscall"

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysMonitor.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-05-13 10:39
 */

type SysResourceInfo struct {
	Host        HostInfo   `json:"host"`
	CPU         CPUInfo    `json:"cpu"`
	Memory      MemoryInfo `json:"memory"`
	Disk        DiskInfo   `json:"disk"`
	ProcessMem  uint       `json:"pmem"`
	ProcessDisk uint       `json:"pdisk"`
}

type HostInfo struct {
	Hostname string `json:"hostname"`
	Uptime   string `json:"uptime"`
	OS       string `json:"os"`
	Kernel   string `json:"kernel"`
}

type CPUInfo struct {
	Cores     string `json:"cores"`
	VendorID  string `json:"vendorId"`
	ModelName string `json:"modelName"`
}

type MemoryInfo struct {
	Total       string `json:"total"`
	Available   string `json:"available"`
	Used        string `json:"used"`
	UsedPercent string `json:"usedPercent"`
}

type DiskInfo struct {
	Path        string `json:"path"`
	FsType      string `json:"fstype"`
	Total       string `json:"total"`
	Free        string `json:"free"`
	Used        string `json:"used"`
	UsedPercent string `json:"usedPercent"`
}

var (
	advapi = syscall.NewLazyDLL("Advapi32.dll")
	kernel = syscall.NewLazyDLL("Kernel32.dll")
)
