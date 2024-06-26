package service

import (
	"admin-go-api/api/entity"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/yusufpapurcu/wmi"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  winHostInfo.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-05-13 10:55
 */

// GetStartTime 开机时间
func GetStartTime() string {
	GetTickCount := syscall.NewLazyDLL("kernel32.dll").NewProc("GetTickCount")
	r, _, _ := GetTickCount.Call()
	if r == 0 {
		return "获取时间失败"
	}
	ms := time.Duration(r * 1000 * 1000) // 转换为纳秒

	days := ms / (24 * time.Hour)
	ms -= days * 24 * time.Hour
	hours := ms / time.Hour
	ms -= hours * time.Hour
	minutes := ms / time.Minute
	ms -= minutes * time.Minute
	seconds := ms / time.Second

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

// GetSystemVersion 系统版本
func GetSystemVersion() string {
	version, err := syscall.GetVersion()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d.%d (%d)", byte(version), uint8(version>>8), version>>16)
}

// getHostname 获取主机名
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}

// GetMotherboardInfo 主板信息
func GetMotherboardInfo() string {
	var s []struct {
		Product string
	}
	err := wmi.Query("SELECT  Product  FROM Win32_BaseBoard WHERE (Product IS NOT NULL)", &s)
	if err != nil {
		return ""
	}
	return s[0].Product
}

// 获取 CPU 核心数
func GetCores() string {
	cmd := exec.Command("wmic", "cpu", "get", "NumberOfCores")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		// 第二行是核心数
		return strings.TrimSpace(lines[1]) + "核"
	}
	return ""
}

// 获取 CPU 厂商 ID
func GetVendorID() string {
	cmd := exec.Command("wmic", "cpu", "get", "Manufacturer")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		// 第二行是厂商 ID
		return strings.TrimSpace(lines[1])
	}
	return ""
}

// 获取 CPU 型号名称
func GetModelName() string {
	cmd := exec.Command("wmic", "cpu", "get", "Name")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		// 第二行是型号名称
		return strings.TrimSpace(lines[1])
	}
	return ""
}

// memoryStatusEx struct
type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

// 获取内存信息
func GetMemory() entity.MemoryInfo {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	GlobalMemoryStatusEx := kernel32.NewProc("GlobalMemoryStatusEx")

	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))

	mem, _, callErr := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		fmt.Println("Call to GlobalMemoryStatusEx failed:", callErr)
		return entity.MemoryInfo{}
	}

	total := fmt.Sprintf("%dGB", memInfo.ullTotalPhys/1024/1024/1024)
	available := fmt.Sprintf("%dGB", memInfo.ullAvailPhys/1024/1024/1024)
	used := fmt.Sprintf("%dGB", (memInfo.ullTotalPhys-memInfo.ullAvailPhys)/1024/1024/1024)
	usedPercent := fmt.Sprintf("%.2f%%", float64(memInfo.dwMemoryLoad))

	return entity.MemoryInfo{
		Total:       total,
		Available:   available,
		Used:        used,
		UsedPercent: usedPercent,
	}
}

func usage(getDiskFreeSpaceExW *syscall.LazyProc, path string) (entity.DiskInfo, error) {
	lpFreeBytesAvailable := int64(0)
	var totalBytes, freeBytes int64
	diskret, _, err := getDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&freeBytes)),
	)
	if diskret == 0 {
		return entity.DiskInfo{}, err
	}
	total := totalBytes / (1024 * 1024 * 1024) // convert to GB
	free := freeBytes / (1024 * 1024 * 1024)   // convert to GB
	used := total - free
	usedPercent := float64(used) / float64(total) * 100

	return entity.DiskInfo{
		Path:        path,
		FsType:      "", // Assuming no filesystem type information
		Total:       fmt.Sprintf("%dGB", total),
		Free:        fmt.Sprintf("%dGB", free),
		Used:        fmt.Sprintf("%dGB", used),
		UsedPercent: fmt.Sprintf("%.2f%%", usedPercent),
	}, nil
}

// GetDiskInfo 获取硬盘信息，只返回C盘的信息
func GetDiskInfo() entity.DiskInfo {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	GetDiskFreeSpaceExW := kernel32.NewProc("GetDiskFreeSpaceExW")

	// 获取C盘信息
	info, err := usage(GetDiskFreeSpaceExW, "C:\\")
	if err != nil {
		fmt.Println("Error retrieving disk usage:", err)
		return entity.DiskInfo{}
	}

	return info
}

func parsePercent(percent string) uint {
	percent = strings.TrimSuffix(percent, "%")
	value, err := strconv.ParseFloat(percent, 64)
	if err != nil {
		log.Fatalf("Failed to parse percent value: %v", err)
	}
	return uint(value) // 将小数转换为整数百分比
}

func GetResourceInfo() (resourceInfo entity.SysResourceInfo) {
	var wg sync.WaitGroup
	// 创建用于传递各种信息的 channels
	hostnameChan := make(chan string, 1)
	uptimeChan := make(chan string, 1)
	osChan := make(chan string, 1)
	kernelChan := make(chan string, 1)
	memChan := make(chan entity.MemoryInfo, 1)
	cpuChan := make(chan entity.CPUInfo, 1)
	diskChan := make(chan entity.DiskInfo, 1)

	wg.Add(7) // 七个并行任务

	// 获取主机名
	go func() {
		defer wg.Done()
		hostnameChan <- getHostname()
	}()

	// 获取启动时间
	go func() {
		defer wg.Done()
		uptimeChan <- GetStartTime()
	}()

	// 获取操作系统版本
	go func() {
		defer wg.Done()
		osChan <- GetSystemVersion()
	}()

	// 获取主板信息
	go func() {
		defer wg.Done()
		kernelChan <- GetMotherboardInfo()
	}()

	// 获取内存信息
	go func() {
		defer wg.Done()
		memInfo := GetMemory()
		memChan <- memInfo
	}()

	// 获取CPU信息
	go func() {
		defer wg.Done()
		cpuInfo := entity.CPUInfo{
			Cores:     GetCores(),
			VendorID:  GetVendorID(),
			ModelName: GetModelName(),
		}
		cpuChan <- cpuInfo
	}()

	// 获取磁盘信息
	go func() {
		defer wg.Done()
		diskInfo := GetDiskInfo()

		diskChan <- diskInfo
	}()

	wg.Wait() // 等待所有 goroutine 完成

	// 从 channels 读取数据
	hostname := <-hostnameChan
	uptime := <-uptimeChan
	os := <-osChan
	kernel := <-kernelChan
	memInfo := <-memChan
	cpuInfo := <-cpuChan
	diskInfo := <-diskChan

	// 汇总信息
	resourceInfo = entity.SysResourceInfo{
		Host: entity.HostInfo{
			Hostname: hostname,
			Uptime:   uptime,
			OS:       os,
			Kernel:   kernel,
		},
		CPU:         cpuInfo,
		Memory:      memInfo,
		Disk:        diskInfo,
		ProcessMem:  parsePercent(memInfo.UsedPercent),
		ProcessDisk: parsePercent(diskInfo.UsedPercent),
	}

	return resourceInfo
}
