package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  test.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-05-13 15:56
 */

// SystemInfo 包含系统信息的结构体
type HostInfo struct {
	Hostname      string
	Uptime        string
	OS            string
	KernelVersion string
}

// CPUInfo 包含 CPU 信息的结构体
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

// GetStartTime 获取开机时间
func GetStartTime() string {
	uptimeFile := "/proc/uptime"
	content, err := ioutil.ReadFile(uptimeFile)
	if err != nil {
		return "获取时间失败"
	}

	fields := strings.Fields(string(content))
	if len(fields) < 1 {
		return "获取时间失败"
	}

	uptimeSeconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return "获取时间失败"
	}

	uptimeDuration := time.Duration(int(uptimeSeconds)) * time.Second

	days := uptimeDuration / (24 * time.Hour)
	uptimeDuration -= days * 24 * time.Hour
	hours := uptimeDuration / time.Hour
	uptimeDuration -= hours * time.Hour
	minutes := uptimeDuration / time.Minute
	uptimeDuration -= minutes * time.Minute
	seconds := uptimeDuration / time.Second

	result := ""
	if days > 0 {
		result += strconv.FormatInt(int64(days), 10) + "天"
	}
	if hours > 0 {
		result += strconv.FormatInt(int64(hours), 10) + "小时"
	}
	if minutes > 0 {
		result += strconv.FormatInt(int64(minutes), 10) + "分钟"
	}
	if seconds > 0 {
		result += strconv.FormatInt(int64(seconds), 10) + "秒"
	}

	return result
}

// GetSystemInfo 获取系统信息
func GetSystemInfo() (HostInfo, error) {
	var info HostInfo

	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		return info, err
	}
	info.Hostname = hostname
	info.Uptime = GetStartTime()
	// 获取系统版本信息
	contents, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return info, err
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "VERSION=") {
			version := strings.Trim(line[len("VERSION="):], "\"")
			info.OS = version
			break
		}
	}

	// 获取内核版本信息
	output, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return info, err
	}
	info.KernelVersion = strings.TrimSpace(string(output))

	return info, nil
}

// GetCPUInfo 获取 CPU 信息
func GetCPUInfo() (CPUInfo, error) {

	// 打开 /proc/cpuinfo 文件
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error opening /proc/cpuinfo:", err)
		return CPUInfo{}, nil
	}
	defer file.Close()

	// 初始化变量
	var processorCount int
	var model, vendor string
	seenModels := make(map[string]bool)
	seenVendors := make(map[string]bool)

	// 读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		// 确保行是键值对
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			// 根据键名解析信息
			switch key {
			case "processor":
				processorCount++
			case "model name":
				if !seenModels[value] {
					seenModels[value] = true
					model = value
				}
			case "vendor_id":
				if !seenVendors[value] {
					seenVendors[value] = true
					vendor = value
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading /proc/cpuinfo:", err)
		return CPUInfo{}, nil
	}

	// 创建结构体实例
	cpuInfo := CPUInfo{
		Cores:     fmt.Sprintf("%d", processorCount),
		VendorID:  vendor,
		ModelName: model,
	}
	return cpuInfo, nil
}

func GetMemoryInfo() MemoryInfo {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println("Error opening /proc/meminfo:", err)
		return MemoryInfo{}
	}
	defer file.Close()

	memStats := make(map[string]int64)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		// Convert kB to GB
		value, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			continue
		}
		value /= 1024 * 1024 // Convert from kB to GB

		memStats[parts[0]] = value
	}

	total := memStats["MemTotal:"]
	available := memStats["MemAvailable:"]
	used := total - available
	usedPercent := 100 * used / total

	memoryInfo := MemoryInfo{
		Total:       fmt.Sprintf("%dGB", total),
		Available:   fmt.Sprintf("%dGB", available),
		Used:        fmt.Sprintf("%dGB", used),
		UsedPercent: fmt.Sprintf("%d%%", usedPercent),
	}
	return memoryInfo
}

// GetDiskInfo 获取磁盘信息并以 GB 为单位返回
func GetDiskInfo() (DiskInfo, error) {
	var info DiskInfo

	// 获取磁盘信息
	dfCmd := "df /"
	dfOutput, err := exec.Command("sh", "-c", dfCmd).Output()
	if err != nil {
		return info, err
	}

	lines := strings.Split(string(dfOutput), "\n")
	if len(lines) < 2 {
		return info, fmt.Errorf("unexpected output format")
	}

	fields := strings.Fields(lines[1])
	if len(fields) < 6 {
		return info, fmt.Errorf("unexpected output format")
	}

	// 将大小转换为 GB
	totalSize, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return info, fmt.Errorf("error parsing total size: %v", err)
	}
	info.Total = fmt.Sprintf("%dGB", totalSize/(1024*1024))

	usedSize, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return info, fmt.Errorf("error parsing used size: %v", err)
	}
	info.Used = fmt.Sprintf("%dGB", usedSize/(1024*1024))

	freeSize, err := strconv.ParseUint(fields[3], 10, 64)
	if err != nil {
		return info, fmt.Errorf("error parsing free size: %v", err)
	}
	info.Free = fmt.Sprintf("%dGB", freeSize/(1024*1024))

	// 其他字段直接复制
	info.Path = fields[5]
	info.FsType = fields[0]
	info.UsedPercent = fields[4]

	return info, nil
}

type SysResourceInfo struct {
	Host        HostInfo   `json:"host"`
	CPU         CPUInfo    `json:"cpu"`
	Memory      MemoryInfo `json:"memory"`
	Disk        DiskInfo   `json:"disk"`
	ProcessMem  uint       `json:"pmem"`
	ProcessDisk uint       `json:"pdisk"`
}

func parsePercent(percent string) uint {
	percent = strings.TrimSuffix(percent, "%")
	value, err := strconv.ParseFloat(percent, 64)
	if err != nil {
		log.Fatalf("Failed to parse percent value: %v", err)
	}
	return uint(value) // 将小数转换为整数百分比
}

func main() {
	hostInfo, _ := GetSystemInfo()
	cpuInfo, _ := GetCPUInfo()
	memInfo := GetMemoryInfo()
	diskInfo, _ := GetDiskInfo()

	resourceInfo := SysResourceInfo{
		Host:   hostInfo,
		CPU:    cpuInfo,
		Memory: memInfo,
		Disk:   diskInfo,
		// 这里的进程内存和磁盘使用数据是示例，需要实际获取
		ProcessMem:  parsePercent(memInfo.UsedPercent),
		ProcessDisk: parsePercent(diskInfo.UsedPercent),
	}

	fmt.Println("Host Info:", hostInfo)
	fmt.Println("CPU Info:", cpuInfo)
	fmt.Println("Memory Info:", memInfo)
	fmt.Println("Disk Info:", diskInfo)
	fmt.Println("Resource Info:", resourceInfo)
}
