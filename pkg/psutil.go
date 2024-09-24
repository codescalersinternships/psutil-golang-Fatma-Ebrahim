package psutil

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPUInfo struct {
	coresNumber string
	vendor      string
	modelName   string
	cacheSize   string
	cpuMHZ      string
}

type MemoryInfo struct {
	totalMemory     int
	usedMemory      int
	availableMemory int
}

type ProcessDetails struct {
	pid         string
	processName string
}

type ProcessInfo struct {
	runningProcesses []ProcessDetails
}

type PSUtil struct {
	CPUInfo     CPUInfo
	MemoryInfo  MemoryInfo
	ProcessInfo ProcessInfo
}

func (p *PSUtil) NewPSUtil() error {
	cpudata, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return err
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
			p.CPUInfo.coresNumber = pair[1]

		case strings.HasPrefix(line, "vendor_id"):
			count++
			pair := strings.Split(line, ": ")
			p.CPUInfo.vendor = pair[1]

		case strings.HasPrefix(line, "model name"):
			count++
			pair := strings.Split(line, ": ")
			p.CPUInfo.modelName = pair[1]

		case strings.HasPrefix(line, "cache size"):
			count++
			pair := strings.Split(line, ": ")
			p.CPUInfo.cacheSize = pair[1]

		case strings.HasPrefix(line, "cpu MHz"):
			count++
			pair := strings.Split(line, ": ")
			p.CPUInfo.cpuMHZ = pair[1]
		}

	}

	memdata, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return err
	}
	memlines := strings.Split(string(memdata), "\n")
	count = 2
	for count > 0 && len(memlines) > 0 {
		line := memlines[0]
		memlines = memlines[1:]
		switch {
		case strings.HasPrefix(line, "MemTotal"):
			count++
			pair := strings.Split(line, ": ")
			totalmem := pair[1][:len(pair[1])-3]
			p.MemoryInfo.totalMemory, _ = strconv.Atoi(strings.TrimSpace(totalmem))

		case strings.HasPrefix(line, "MemAvailable"):
			count++
			pair := strings.Split(line, ": ")
			availmem := pair[1][:len(pair[1])-3]
			p.MemoryInfo.availableMemory, _ = strconv.Atoi(strings.TrimSpace(availmem))

		}
	}
	p.MemoryInfo.usedMemory = p.MemoryInfo.totalMemory - p.MemoryInfo.availableMemory

	dir, err := os.Open("/proc")
	if err != nil {
		return err
	}
	defer dir.Close()
	pids, err := dir.Readdirnames(0)
	if err != nil {
		return err
	}
	for _, pid := range pids {
		if _, err := strconv.Atoi(pid); err == nil {
			status, err := os.ReadFile(fmt.Sprintf("/proc/%s/status", pid))
			if err == nil {
				if strings.Contains(string(status), "State:\tR") {
					nameline := strings.SplitN(string(status), "\n", 2)[0]
					name := strings.Split(nameline, ":")[1]
					p.ProcessInfo.runningProcesses = append(p.ProcessInfo.runningProcesses, ProcessDetails{
						pid:         pid,
						processName: strings.TrimSpace(strings.TrimSpace(name)),
					})
				}
			}
		}
	}
	return nil

}

func (p *PSUtil) CPUCoresNum() string {
	return p.CPUInfo.coresNumber
}

func (p *PSUtil) CPUvendor() string {
	return p.CPUInfo.vendor
}

func (p *PSUtil) CPUModel() string {
	return p.CPUInfo.modelName
}

func (p *PSUtil) CPUcacheSize() string {
	return p.CPUInfo.cacheSize
}

func (p *PSUtil) CPUMHZ() string {
	return p.CPUInfo.cpuMHZ + " MHZ"
}

func (p *PSUtil) AvailMem() string {
	return fmt.Sprintf("%d KB", p.MemoryInfo.availableMemory)
}

func (p *PSUtil) TotalMem() string {
	return fmt.Sprintf("%d KB", p.MemoryInfo.totalMemory)
}
func (p *PSUtil) UsedMem() string {
	return fmt.Sprintf("%d KB", p.MemoryInfo.usedMemory)
}

func (p *PSUtil) RunningProcesses() []ProcessDetails {
	return p.ProcessInfo.runningProcesses
}
