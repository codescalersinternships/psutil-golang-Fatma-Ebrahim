package test

import (
	"reflect"
	"testing"

	psutil "github.com/codescalersinternships/psutil-golang-Fatma-Ebrahim/pkg"
)

func (m *MockInfo) Getcpuinfo() (psutil.CPUInfo, error) {
	cpuinfo := psutil.CPUInfo{CoresNumber: "2", Vendor: "vendor name", ModelName: "model name", CacheSize: "512 KB", CPUMHZ: "1000"}
	return cpuinfo, nil
}
func (m *MockInfo) Getmeminfo() (psutil.MemoryInfo, error) {
	meminfo := psutil.MemoryInfo{TotalMemory: 1024, AvailableMemory: 512, UsedMemory: 512}
	return meminfo, nil
}
func (m *MockInfo) Getprocinfo() (psutil.ProcessInfo, error) {
	procinfo := psutil.ProcessInfo{RunningProcesses: []psutil.ProcessDetails{{PID: "1", ProcessName: "main"}}}
	return procinfo, nil
}

type MockInfo struct{}

func TestCPUInfo(t *testing.T) {
	mockinfo := &MockInfo{}
	t.Run("test cpu cores number", func(t *testing.T) {
		coresnumber, err := psutil.CPUCoresNum(mockinfo)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if coresnumber != "2" {
			t.Errorf("expected core number: 2 got %s", coresnumber)
		}
	})
	t.Run("test cpu vendor", func(t *testing.T) {
		vendor, err := psutil.CPUVendor(mockinfo)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if vendor != "vendor name" {
			t.Errorf("expected vendor: vendor name got %s", vendor)
		}
	})

	t.Run("test cpu model name", func(t *testing.T) {
		model, err := psutil.CPUModel(mockinfo)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if model != "model name" {
			t.Errorf("expected model: model name got %s", model)
		}
	})

	t.Run("test cpu cache size", func(t *testing.T) {
		cachesize, err := psutil.CPUcacheSize(mockinfo)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if cachesize != "512 KB" {
			t.Errorf("expected cache size: 512 KB got %s", cachesize)
		}
	})

	t.Run("test cpu MHZ", func(t *testing.T) {
		cpumhz, err := psutil.CPUMHZ(mockinfo)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if cpumhz != "1000 MHZ" {
			t.Errorf("expected CPU MHZ: 1000 MHZ got %s", cpumhz)
		}
	})

}

func TestMemInfo(t *testing.T) {
	mockinfo := &MockInfo{}
	t.Run("test total memory", func(t *testing.T) {
		totalmem, err := psutil.TotalMem(mockinfo)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if totalmem != "1024 KB" {
			t.Errorf("expected total memory: 1024 KB got %s", totalmem)
		}
	})
	t.Run("test available memory", func(t *testing.T) {
		availmem, err := psutil.AvailMem(mockinfo)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if availmem != "512 KB" {
			t.Errorf("expected available memory: 512 KB got %s", availmem)
		}
	})

	t.Run("test used memory", func(t *testing.T) {
		usedmem, err := psutil.UsedMem(mockinfo)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		if usedmem != "512 KB" {
			t.Errorf("expected used memory: 512 KB got %s", usedmem)
		}
	})
}

func TestProcInfo(t *testing.T) {
	mockinfo := &MockInfo{}
	t.Run("test running processes", func(t *testing.T) {
		runproc, err := psutil.RunningProcesses(mockinfo)
		if err != nil {
			t.Errorf("expected nil got %v", err)
		}
		want := []psutil.ProcessDetails{{PID: "1", ProcessName: "main"}}
		if !reflect.DeepEqual(runproc, want) {
			t.Errorf("expected running processes: %v got %v", want, runproc)
		}
	})

}
