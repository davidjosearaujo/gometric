package metrics

import (
	"runtime"

	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	"github.com/graphql-go/graphql"
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
					var cpu CPU

					// CPU Load
					if load, ok := host.(types.LoadAverage); ok {
						cpu.Load, _ = load.LoadAverage()
					}

					// CPU timers
					cpu.Time, _ = host.CPUTime()

					// CPU Number of cores
					cpu.CoreCount = int16(runtime.NumCPU())

					return cpu, nil
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
			// "disk": &graphql.Field{
			// 	Type: diskType,
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					
			// 		
			// 	},
			// },
		},
	})

	MetricsSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
