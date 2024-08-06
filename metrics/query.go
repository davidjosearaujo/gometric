package metrics

import (
	"runtime"

	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	"github.com/graphql-go/graphql"
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
					var net Network
					host, _ := sysinfo.Host()

					if n, ok := host.(types.NetworkCounters); ok {
						netcounter, _ := n.NetworkCounters()
						net.Network = *netcounter
					}

					return net, nil
				},
			},
		},
	})

	MetricsSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
