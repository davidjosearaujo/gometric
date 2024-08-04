package metrics

import (
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
			"cpu": &graphql.Field{
				Type: cpuType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return CPU{
						Usage:        5.6,
						PerCoreUsage: []float64{5.3, 6.4}}, nil
				},
			},
		},
	})

	MetricsSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
