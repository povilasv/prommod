// +build go1.12
package promverauto

import (
	"github.com/povilasv/promver"
	"github.com/prometheus/client_golang/prometheus"
)

// NewCollector works like the function of the same name in the promver package
// but it automatically registers the Exporter with the
// prometheus.DefaultRegisterer. If the registration fails, NewCounter panics.
func NewCollector(program string) prometheus.Counter {
	c := promver.NewExporter(program)
	prometheus.MustRegister(c)
	return c
}
