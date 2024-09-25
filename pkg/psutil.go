// psutil-go is a lightweight library in Go that provides essential system information
package psutil

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CPUInfo represents CPU information struct from /proc/cpuinfo
type CPUInfo struct {
	CoresNumber string
	Vendor      string
	ModelName   string
	CacheSize   string
	CPUMHZ      string
}

// MemoryInfo represents Memory information struct from /proc/meminfo
type MemoryInfo struct {
	TotalMemory     int
	UsedMemory      int
	AvailableMemory int
}

// ProcessDetails represents Process information struct
type ProcessDetails struct {
	PID         string
	ProcessName string
}

// ProcessInfo represents Process information struct
type ProcessInfo struct {
	RunningProcesses []ProcessDetails
}

func cpuinfo() (CPUInfo, error) {
	cpuinfo := CPUInfo{}
	cpudata, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return cpuinfo, err
	}
	cpulines := strings.Split(string(cpudata), "\n")
	count := 5
	for count > 0 && len(cpulines) > 0 {
		line := cpulines[0]
		cpulines = cpulines[1:]
		switch {
		case strings.HasPrefix(line, "cpu cores"):
			count++
			pair := strings.Split(line, ": ")
			cpuinfo.CoresNumber = pair[1]

		case strings.HasPrefix(line, "Vendor_id"):
			count++
			pair := strings.Split(line, ": ")
			cpuinfo.Vendor = pair[1]

		case strings.HasPrefix(line, "model name"):
			count++
			pair := strings.Split(line, ": ")
			cpuinfo.ModelName = pair[1]

		case strings.HasPrefix(line, "cache size"):
			count++
			pair := strings.Split(line, ": ")
			cpuinfo.CacheSize = pair[1]

		case strings.HasPrefix(line, "cpu MHz"):
			count++
			pair := strings.Split(line, ": ")
			cpuinfo.CPUMHZ = pair[1]
		}

	}
	return cpuinfo, nil
}
func meminfo() (MemoryInfo, error) {
	meminfo := MemoryInfo{}
	memdata, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return meminfo, err
	}
	memlines := strings.Split(string(memdata), "\n")
	count := 2
	for count > 0 && len(memlines) > 0 {
		line := memlines[0]
		memlines = memlines[1:]
		switch {
		case strings.HasPrefix(line, "MemTotal"):
			count++
			pair := strings.Split(line, ": ")
			totalmem := pair[1][:len(pair[1])-3]
			meminfo.TotalMemory, _ = strconv.Atoi(strings.TrimSpace(totalmem))

		case strings.HasPrefix(line, "MemAvailable"):
			count++
			pair := strings.Split(line, ": ")
			availmem := pair[1][:len(pair[1])-3]
			meminfo.AvailableMemory, _ = strconv.Atoi(strings.TrimSpace(availmem))

		}
	}
	meminfo.UsedMemory = meminfo.TotalMemory - meminfo.AvailableMemory
	return meminfo, nil

}
func procinfo() (ProcessInfo, error) {
	procinfo := ProcessInfo{}
	dir, err := os.Open("/proc")
	if err != nil {
		return procinfo, err
	}
	defer dir.Close()
	PIDs, err := dir.Readdirnames(0)
	if err != nil {
		return procinfo, err
	}
	for _, PID := range PIDs {
		if _, err := strconv.Atoi(PID); err == nil {
			status, err := os.ReadFile(fmt.Sprintf("/proc/%s/status", PID))
			if err == nil {
				if strings.Contains(string(status), "State:\tR") {
					nameline := strings.SplitN(string(status), "\n", 2)[0]
					name := strings.Split(nameline, ":")[1]
					procinfo.RunningProcesses = append(procinfo.RunningProcesses, ProcessDetails{
						PID:         PID,
						ProcessName: strings.TrimSpace(strings.TrimSpace(name)),
					})
				}
			}
		}
	}
	return procinfo, nil

}

type FilesInfo struct{}

func (f *FilesInfo) Getcpuinfo() (CPUInfo, error) {
	return cpuinfo()
}
func (f *FilesInfo) Getmeminfo() (MemoryInfo, error) {
	return meminfo()
}
func (f *FilesInfo) Getprocinfo() (ProcessInfo, error) {
	return procinfo()
}

type Info interface {
	Getcpuinfo() (CPUInfo, error)
	Getmeminfo() (MemoryInfo, error)
	Getprocinfo() (ProcessInfo, error)
}

// CPUCoresNum returns CPU cores number
func CPUCoresNum(info Info) (string, error) {
	cpuinfo, err := info.Getcpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.CoresNumber, nil
}

// CPUVendor returns CPU Vendor name
func CPUVendor(info Info) (string, error) {
	cpuinfo, err := info.Getcpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.Vendor, nil
}

// CPUModel returns CPU model
func CPUModel(info Info) (string, error) {
	cpuinfo, err := info.Getcpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.ModelName, nil
}

// CPUcacheSize returns CPU cache size in KB
func CPUcacheSize(info Info) (string, error) {
	cpuinfo, err := info.Getcpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.CacheSize, nil
}

// CPUMHZ returns CPU MHZ
func CPUMHZ(info Info) (string, error) {
	cpuinfo, err := info.Getcpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.CPUMHZ + " MHZ", nil
}

// AvailMem returns available memory in KB
func AvailMem(info Info) (string, error) {
	meminfo, err := info.Getmeminfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d KB", meminfo.AvailableMemory), nil
}

// TotalMem returns total memory in KB
func TotalMem(info Info) (string, error) {
	meminfo, err := info.Getmeminfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d KB", meminfo.TotalMemory), nil
}

// UsedMem returns used memory in KB
func UsedMem(info Info) (string, error) {
	meminfo, err := info.Getmeminfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d KB", meminfo.UsedMemory), nil
}

// RunningProcesses returns running processes PIDs and names
func RunningProcesses(info Info) ([]ProcessDetails, error) {
	procinfo, err := info.Getprocinfo()
	if err != nil {
		return procinfo.RunningProcesses, err
	}
	return procinfo.RunningProcesses, nil
}
