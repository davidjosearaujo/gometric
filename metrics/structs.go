package metrics

import (
	"github.com/elastic/go-sysinfo/types"
)

type CPU struct {
	Time		types.CPUTimes
	Load		*types.LoadAverageInfo
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
	InterfaceName string `json:"interfaceName"`
	BytesSent     uint64 `json:"bytesSent"`
	BytesRecv     uint64 `json:"bytesRecv"`
	PacketsSent   uint64 `json:"packetsSent"`
	PacketsRecv   uint64 `json:"packetsRecv"`
	ErrIn         uint64 `json:"errIn"`
	ErrOut        uint64 `json:"errOut"`
	DropIn        uint64 `json:"dropIn"`
	DropOut       uint64 `json:"dropOut"`
}

type Process struct {
	PID         int     `json:"pid"`
	Name        string  `json:"name"`
	CPUUsage    float64 `json:"cpuUsage"`
	MemoryUsage uint64  `json:"memoryUsage"`
	StartTime   int64   `json:"startTime"`
}