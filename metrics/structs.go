package metrics

import (
	"github.com/elastic/go-sysinfo/types"
)

type CPU struct {
	Time		types.CPUTimes
	Load		*types.LoadAverageInfo
	CoreCount	int16
}

type Disk struct {
	Total       uint64 `json:"total"`
	Used        uint64 `json:"used"`
	Free        uint64 `json:"free"`
	InodesTotal uint64 `json:"inodesTotal"`
	InodesUsed  uint64 `json:"inodesUsed"`
	InodesFree  uint64 `json:"inodesFree"`
}

type Network struct {
	Network		types.NetworkCountersInfo
}

type Process struct {
	PID         int     `json:"pid"`
	Name        string  `json:"name"`
	CPUUsage    float64 `json:"cpuUsage"`
	MemoryUsage uint64  `json:"memoryUsage"`
	StartTime   int64   `json:"startTime"`
}