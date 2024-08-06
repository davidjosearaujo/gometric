package metrics

import (
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
					var cpu CPU
					
					host, _ := sysinfo.Host()
					
					if load, ok := host.(types.LoadAverage); ok {
						cpu.Load, _ = load.LoadAverage()
					}

					cpu.Time, _ = host.CPUTime()

					// TODO:
					//	- Add fields with core count
					//  - Add field with core number

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
			// "network": &graphql.Field{
			// 	Type: networkType,
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 	},
			// },
		},
	})

	MetricsSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
