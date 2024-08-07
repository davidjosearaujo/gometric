package main

import (
	"encoding/json"
	"net/http"

	"github.com/davidjosearaujo/gometric/metrics"
	"github.com/graphql-go/graphql"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        metrics.MetricsSchema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})

	server := &http.Server{
		Addr:    ":7000",
		Handler: mux,
	}

	server.ListenAndServe()
}
