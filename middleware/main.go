package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/jessicalins/instrumentation-practices-examples/middleware/httpmiddleware"
)

func main() {
	// Create non-global registry
	registry := prometheus.NewRegistry()

	// Add go runtime metrics and process collectors
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	// Expose /metrics HTTP endpoint using the created custom registry.
	http.Handle(
		"/metrics",
		httpmiddleware.New(
			registry, nil).
			WrapHandler("/metrics", promhttp.HandlerFor(
				registry,
				promhttp.HandlerOpts{}),
			))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
