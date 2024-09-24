package psutil

import (

)

type CPUInfo struct {
	CoresNumber int
	Vendor      string
	ModelName   string
	CacheSize   int
	CPUMHZ      int
}

type MemoryInfo struct {
	TotalMemory     int
	UsedMemory      int
	AvailableMemory int
}
type ProcessDetails struct {
	PID         int
	ProcessName string
}

type ProcessInfo struct {
	RunningProcesses []ProcessDetails
}

type PSUtil struct {
	CPUInfo     CPUInfo
	MemoryInfo  MemoryInfo
	ProcessInfo ProcessInfo
}

func (p *PSUtil) GetCPUInfo() *CPUInfo {
	cpuinfo:=&CPUInfo{}
	return cpuinfo
}

func (p *PSUtil) GetMemoryInfo() *MemoryInfo {
	meminfo:=&MemoryInfo{}
	return meminfo
}

func (p *PSUtil) GetProcessInfo() *ProcessInfo {
	processinfo:=&ProcessInfo{}
	return processinfo
}