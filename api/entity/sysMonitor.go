package entity

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  sysMonitor.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-05-13 10:39
 */

type SysResourceInfo struct {
	Hostname       string `json:"hostname"`
	Uptime         string `json:"uptime"`
	OS             string `json:"os"`
	Kernel         string `json:"kernel"`
	CPUCores       string `json:"cpuCores"`
	CPUVendorID    string `json:"cpuVendorId"`
	CPUModelName   string `json:"cpuModelName"`
	MemoryTotal    string `json:"memoryTotal"`
	MemoryAvail    string `json:"memoryAvailable"`
	MemoryUsed     string `json:"memoryUsed"`
	MemoryUsedPct  string `json:"memoryUsedPct"`
	DiskPath       string `json:"diskPath"`
	DiskFsType     string `json:"diskFsType"`
	DiskTotal      string `json:"diskTotal"`
	DiskFree       string `json:"diskFree"`
	DiskUsed       string `json:"diskUsed"`
	DiskUsedPct    string `json:"diskUsedPct"`
	ProcessMemPct  uint   `json:"pmem"`
	ProcessDiskPct uint   `json:"pdisk"`
}
