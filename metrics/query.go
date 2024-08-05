package metrics

import (
	"fmt"

	"github.com/elastic/go-sysinfo"
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
			// "memory":	&graphql.Field{
			// 	Type: memoryType,
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					
			// 	},
			// }
		},
	})

	MetricsSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
