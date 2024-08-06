package metrics

import (
	"github.com/elastic/go-sysinfo/types"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
)

type CPU struct {
	Info      []cpu.InfoStat
	Time      types.CPUTimes
	Load      *types.LoadAverageInfo
	CoreCount int16
}

// TODO
type Disk struct {
	Partitions     []disk.PartitionStat
	IOCountersStat map[string]disk.IOCountersStat
	UsageStat      map[string]disk.UsageStat
}

type Network struct {
	Network types.NetworkCountersInfo
}

type Process struct {
	PID         int     `json:"pid"`
	Name        string  `json:"name"`
	CPUUsage    float64 `json:"cpuUsage"`
	MemoryUsage uint64  `json:"memoryUsage"`
	StartTime   int64   `json:"startTime"`
}
