package metrics

import (
	"runtime"

	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	"github.com/graphql-go/graphql"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
)

var (
	MetricsSchema graphql.Schema
)

func init() {
	initTypes()
	initQuery()
}

func initQuery() {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"host": &graphql.Field{
				Type: hostType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					host, _ := sysinfo.Host()
					return host.Info(), nil
				},
			},
			"os": &graphql.Field{
				Type: osType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					host, _ := sysinfo.Host()
					hostinfo := host.Info()
					os := *hostinfo.OS
					return os, nil
				},
			},
			"cpu": &graphql.Field{
				Type: cpuType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					host, _ := sysinfo.Host()
					var cpuObj CPU

					// CPU Load
					if load, ok := host.(types.LoadAverage); ok {
						cpuObj.Load, _ = load.LoadAverage()
					}

					// CPU timers
					cpuObj.Time, _ = host.CPUTime()

					cpuObj.Info, _ = cpu.Info()

					// CPU Number of cores
					cpuObj.CoreCount = int16(runtime.NumCPU())

					return cpuObj, nil
				},
			},
			"memory": &graphql.Field{
				Type: memoryType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					host, _ := sysinfo.Host()
					memory, _ := host.Memory()
					return *memory, nil
				},
			},
			"network": &graphql.Field{
				Type: networkType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var network Network
					host, _ := sysinfo.Host()

					if n, ok := host.(types.NetworkCounters); ok {
						netcounter, _ := n.NetworkCounters()
						network.Network = *netcounter
					}

					return network, nil
				},
			},
			// TODO
			"disk": &graphql.Field{
				Type: diskType,
				Args: graphql.FieldConfigArgument{
					"device": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Name of the device",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var diskObj Disk
					diskObj.Partitions, _ = disk.Partitions(false)

					if p.Args["device"] != nil {
						diskObj.Partitions, _ = disk.Partitions(false)
						for _, partition := range diskObj.Partitions {
							if partition.Device == p.Args["device"] {
								diskObj.Partitions = []disk.PartitionStat{partition}
								break
							}
						}

						tempUsageStat, _ := disk.Usage(diskObj.Partitions[0].Mountpoint)
						diskObj.UsageStat = *tempUsageStat
					}

					return diskObj, nil
				},
			},
		},
	})

	MetricsSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
