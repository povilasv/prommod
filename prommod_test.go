package prommod_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/povilasv/prommod"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ExampleNewCollector() {
	prometheus.MustRegister(prommod.NewCollector("test_app_name"))

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ExamplePrint() {
	fmt.Println(prommod.Print("test_app_name"))
	// Output: test_app_name
}
