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
	coresNumber string
	vendor      string
	modelName   string
	cacheSize   string
	cpuMHZ      string
}

// MemoryInfo represents Memory information struct from /proc/meminfo
type MemoryInfo struct {
	totalMemory     int
	usedMemory      int
	availableMemory int
}

// ProcessDetails represents Process information struct
type ProcessDetails struct {
	pid         string
	processName string
}

// ProcessInfo represents Process information struct
type ProcessInfo struct {
	runningProcesses []ProcessDetails
}

func procinfo() (ProcessInfo, error) {
	procinfo := ProcessInfo{}
	dir, err := os.Open("/proc")
	if err != nil {
		return procinfo, err
	}
	defer dir.Close()
	pids, err := dir.Readdirnames(0)
	if err != nil {
		return procinfo, err
	}
	for _, pid := range pids {
		if _, err := strconv.Atoi(pid); err == nil {
			status, err := os.ReadFile(fmt.Sprintf("/proc/%s/status", pid))
			if err == nil {
				if strings.Contains(string(status), "State:\tR") {
					nameline := strings.SplitN(string(status), "\n", 2)[0]
					name := strings.Split(nameline, ":")[1]
					procinfo.runningProcesses = append(procinfo.runningProcesses, ProcessDetails{
						pid:         pid,
						processName: strings.TrimSpace(strings.TrimSpace(name)),
					})
				}
			}
		}
	}
	return procinfo, nil

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
			meminfo.totalMemory, _ = strconv.Atoi(strings.TrimSpace(totalmem))

		case strings.HasPrefix(line, "MemAvailable"):
			count++
			pair := strings.Split(line, ": ")
			availmem := pair[1][:len(pair[1])-3]
			meminfo.availableMemory, _ = strconv.Atoi(strings.TrimSpace(availmem))

		}
	}
	meminfo.usedMemory = meminfo.totalMemory - meminfo.availableMemory
	return meminfo, nil

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
			cpuinfo.coresNumber = pair[1]

		case strings.HasPrefix(line, "vendor_id"):
			count++
			pair := strings.Split(line, ": ")
			cpuinfo.vendor = pair[1]

		case strings.HasPrefix(line, "model name"):
			count++
			pair := strings.Split(line, ": ")
			cpuinfo.modelName = pair[1]

		case strings.HasPrefix(line, "cache size"):
			count++
			pair := strings.Split(line, ": ")
			cpuinfo.cacheSize = pair[1]

		case strings.HasPrefix(line, "cpu MHz"):
			count++
			pair := strings.Split(line, ": ")
			cpuinfo.cpuMHZ = pair[1]
		}

	}
	return cpuinfo, nil
}

// CPUCoresNum returns CPU cores number
func CPUCoresNum() (string, error) {
	cpuinfo, err := cpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.coresNumber, nil
}

// CPUvendor returns CPU vendor name
func CPUvendor() (string, error) {
	cpuinfo, err := cpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.vendor, nil
}

// CPUModel returns CPU model
func CPUModel() (string, error) {
	cpuinfo, err := cpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.modelName, nil
}

// CPUcacheSize returns CPU cache size in KB
func CPUcacheSize() (string, error) {
	cpuinfo, err := cpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.cacheSize, nil
}

// CPUMHZ returns CPU MHZ
func CPUMHZ() (string, error) {
	cpuinfo, err := cpuinfo()
	if err != nil {
		return "", err
	}
	return cpuinfo.cpuMHZ + " MHZ", nil
}

// AvailMem returns available memory in KB
func AvailMem() (string, error) {
	meminfo, err := meminfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d KB", meminfo.availableMemory), nil
}

// TotalMem returns total memory in KB
func TotalMem() (string, error) {
	meminfo, err := meminfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d KB", meminfo.totalMemory), nil
}

// UsedMem returns used memory in KB
func UsedMem() (string, error) {
	meminfo, err := meminfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d KB", meminfo.usedMemory), nil
}

// RunningProcesses returns running processes PIDs and names
func RunningProcesses() ([]ProcessDetails, error) {
	procinfo, err := procinfo()
	if err != nil {
		return procinfo.runningProcesses, err
	}
	return procinfo.runningProcesses, nil
}
