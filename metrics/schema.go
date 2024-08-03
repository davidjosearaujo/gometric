package metrics

import (
	"github.com/graphql-go/graphql"
)

var (
	MetricsSchema graphql.Schema

	cpuType           *graphql.Object
	memoryType        *graphql.Object
	diskType          *graphql.Object
	networkType       *graphql.Object
	processType       *graphql.Object
	systemMetricsType *graphql.Object
)

type CPU struct {
	Usage        float64   `json:"usage"`
	PerCoreUsage []float64 `json:"perCoreUsage"`
	LoadAverage  []float64 `json:"loadAverage"`
}

type Memory struct {
	Total     uint64 `json:"total"`
	Used      uint64 `json:"used"`
	Free      uint64 `json:"free"`
	Cached    uint64 `json:"cached"`
	SwapTotal uint64 `json:"swapTotal"`
	SwapUsed  uint64 `json:"swapUsed"`
	SwapFree  uint64 `json:"swapFree"`
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

type SystemMetrics struct {
	CPU       CPU       `json:"cpu"`
	Memory    Memory    `json:"memory"`
	Disk      []Disk    `json:"disk"`
	Network   []Network `json:"network"`
	Processes []Process `json:"processes"`
}

func init() {
	cpuType = graphql.NewObject(graphql.ObjectConfig{
		Name:			"CPU",
		Description:	"CPU metrics",
		Fields: graphql.Fields{
			"usage": &graphql.Field{
				Type: 			&graphql.Scalar,
				Description:	"The total CPU usage percentage",
				Resolve: 		func(p graphql.ResolveParams) (interface{}, error) {
					// TODO
				}
			}
		}
	})
}
