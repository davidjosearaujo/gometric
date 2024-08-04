package metrics

import (
	"github.com/graphql-go/graphql"
)

var (
	cpuType           *graphql.Object
	memoryType        *graphql.Object
	diskType          *graphql.Object
	networkType       *graphql.Object
	processType       *graphql.Object
	systemMetricsType *graphql.Object
)

func initTypes() {
	cpuType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "CPU",
		Description: "CPU metrics",
		Fields: graphql.Fields{
			"usage": &graphql.Field{
				Type:        graphql.Float,
				Description: "The total CPU usage percentage",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if cpu, ok := p.Source.(CPU); ok {
						return cpu.Usage, nil
					}
					return nil, nil
				},
			},
			"perCoreUsage": &graphql.Field{
				Type:        graphql.NewList(graphql.Float),
				Description: "CPU usage percentage per core",
				Args: graphql.FieldConfigArgument{
					"number": &graphql.ArgumentConfig{
						Description: "Number of core",
						Type:        graphql.NewList(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if cpu, ok := p.Source.(CPU); ok {
						var result []float64

						for _, index := range p.Args["number"].([]interface{}) {
							result = append(result, cpu.PerCoreUsage[index.(int)])
						}

						return result, nil
					}
					return nil, nil
				},
			},
		},
	})

	memoryType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Memory",
		Description: "Memory metrics",
		Fields: graphql.Fields{
			"total": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(Memory); ok {
						return mem.Total, nil
					}
					return nil, nil
				},
			},
			"used": &graphql.Field{
				Type:        graphql.Float,
				Description: "Used memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(Memory); ok {
						return mem.Used, nil
					}
					return nil, nil
				},
			},
			"free": &graphql.Field{
				Type:        graphql.Float,
				Description: "Free memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(Memory); ok {
						return mem.Free, nil
					}
					return nil, nil
				},
			},
			"cached": &graphql.Field{
				Type:        graphql.Float,
				Description: "Cached memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(Memory); ok {
						return mem.Cached, nil
					}
					return nil, nil
				},
			},
			"swapTotal": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total swap memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(Memory); ok {
						return mem.SwapTotal, nil
					}
					return nil, nil
				},
			},
			"swapUsed": &graphql.Field{
				Type:        graphql.Float,
				Description: "Used swap memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(Memory); ok {
						return mem.SwapUsed, nil
					}
					return nil, nil
				},
			},
			"swapFree": &graphql.Field{
				Type:        graphql.Float,
				Description: "Free swap memory in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if mem, ok := p.Source.(Memory); ok {
						return mem.SwapFree, nil
					}
					return nil, nil
				},
			},
		},
	})

	diskType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Disk",
		Description: "Disk metrics",
		Fields: graphql.Fields{
			// TODO: Fields bellow need argument for disk name
			"total": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total disk space in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.Total, nil
					}
					return nil, nil
				},
			},
			"used": &graphql.Field{
				Type:        graphql.Float,
				Description: "Used disk space in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.Used, nil
					}
					return nil, nil
				},
			},
			"free": &graphql.Field{
				Type:        graphql.Float,
				Description: "Free disk space in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.Free, nil
					}
					return nil, nil
				},
			},
			"inodesTotal": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total inodes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.InodesTotal, nil
					}
					return nil, nil
				},
			},
			"inodesUsed": &graphql.Field{
				Type:        graphql.Float,
				Description: "Used inodes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.InodesUsed, nil
					}
					return nil, nil
				},
			},
			"inodesFree": &graphql.Field{
				Type:        graphql.Float,
				Description: "Free inodes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if disk, ok := p.Source.(Disk); ok {
						return disk.InodesFree, nil
					}
					return nil, nil
				},
			},
		},
	})

	networkType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Network",
		Description: "Network metrics",
		Fields: graphql.Fields{
			"interfaceName": &graphql.Field{
				Type:        graphql.String,
				Description: "Network interface name",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						return net.InterfaceName, nil
					}
					return nil, nil
				},
			},
			// TODO: Fields bellow need argument for interface name
			"bytesSent": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total bytes sent",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						return net.BytesSent, nil
					}
					return nil, nil
				},
			},
			"bytesRecv": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total bytes received",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						return net.BytesRecv, nil
					}
					return nil, nil
				},
			},
			"packetsSent": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total packets sent",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						return net.PacketsSent, nil
					}
					return nil, nil
				},
			},
			"packetsRecv": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total packets received",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						return net.PacketsRecv, nil
					}
					return nil, nil
				},
			},
			"errIn": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total input errors",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						return net.ErrIn, nil
					}
					return nil, nil
				},
			},
			"errOut": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total output errors",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						return net.ErrOut, nil
					}
					return nil, nil
				},
			},
			"dropIn": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total input drops",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						return net.DropIn, nil
					}
					return nil, nil
				},
			},
			"dropOut": &graphql.Field{
				Type:        graphql.Float,
				Description: "Total output drops",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if net, ok := p.Source.(Network); ok {
						return net.DropOut, nil
					}
					return nil, nil
				},
			},
		},
	})

	processType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Process",
		Description: "Process metrics",
		Fields: graphql.Fields{
			"pid": &graphql.Field{
				Type:        graphql.Int,
				Description: "Process ID",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if proc, ok := p.Source.(Process); ok {
						return proc.PID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Process name",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if proc, ok := p.Source.(Process); ok {
						return proc.Name, nil
					}
					return nil, nil
				},
			},
			// TODO: Fields bellow need argument for PID number
			"cpuUsage": &graphql.Field{
				Type:        graphql.Float,
				Description: "CPU usage percentage",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if proc, ok := p.Source.(Process); ok {
						return proc.CPUUsage, nil
					}
					return nil, nil
				},
			},
			"memoryUsage": &graphql.Field{
				Type:        graphql.Float,
				Description: "Memory usage in bytes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if proc, ok := p.Source.(Process); ok {
						return proc.MemoryUsage, nil
					}
					return nil, nil
				},
			},
			"startTime": &graphql.Field{
				Type:        graphql.Float,
				Description: "Process start time (Unix timestamp)",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if proc, ok := p.Source.(Process); ok {
						return proc.StartTime, nil
					}
					return nil, nil
				},
			},
		},
	})

	systemMetricsType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "SystemMetrics",
		Description: "Overall system metrics",
		Fields: graphql.Fields{
			"cpu": &graphql.Field{
				Type:        cpuType,
				Description: "CPU metrics",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if sysMetrics, ok := p.Source.(SystemMetrics); ok {
						return sysMetrics.CPU, nil
					}
					return nil, nil
				},
			},
			"memory": &graphql.Field{
				Type:        memoryType,
				Description: "Memory metrics",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if sysMetrics, ok := p.Source.(SystemMetrics); ok {
						return sysMetrics.Memory, nil
					}
					return nil, nil
				},
			},
			"disk": &graphql.Field{
				Type:        graphql.NewList(diskType),
				Description: "Disk metrics",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if sysMetrics, ok := p.Source.(SystemMetrics); ok {
						return sysMetrics.Disk, nil
					}
					return nil, nil
				},
			},
			"network": &graphql.Field{
				Type:        graphql.NewList(networkType),
				Description: "Network metrics",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if sysMetrics, ok := p.Source.(SystemMetrics); ok {
						return sysMetrics.Network, nil
					}
					return nil, nil
				},
			},
			"processes": &graphql.Field{
				Type:        graphql.NewList(processType),
				Description: "Processes metrics",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if sysMetrics, ok := p.Source.(SystemMetrics); ok {
						return sysMetrics.Processes, nil
					}
					return nil, nil
				},
			},
		},
	})

}
