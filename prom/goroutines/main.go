package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	goroutineGuage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gr_count",
			Help: "Number of goroutines currently exist",
		},
		[]string{"hostname"},
	)
)

func init() {
	prometheus.MustRegister(goroutineGuage)
}

func observer() {
	hostname, _ := os.Hostname()

	for {
		value := runtime.NumGoroutine()
		goroutineGuage.With(prometheus.Labels{"hostname": hostname}).Set(float64(value))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go observer()
	http.Handle("/metrics", prometheus.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
