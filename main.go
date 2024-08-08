package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davidjosearaujo/gometric/metrics"
	"github.com/graphql-go/graphql"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/gometric", func(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("Gometric server running at: http://localhost:7000/gometric")

	server.ListenAndServe()
}
