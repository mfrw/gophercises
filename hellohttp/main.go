package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer r.Body.Close()
	code := http.StatusBadRequest

	defer func() {
		httpDuration := time.Since(start)
		hist.WithLabelValues(fmt.Sprintf("%d", code)).Observe(httpDuration.Seconds())
		ctr.WithLabelValues(fmt.Sprintf("%d", code)).Inc()
	}()

	w.Header().Set("Content-Type", "text/plain")
	if r.Method == "GET" {
		code = http.StatusOK
		vars := r.URL.Query()
		name := vars["name"]
		greet := fmt.Sprintf("Hello %s \n", strings.Join(name, ","))
		w.WriteHeader(code)
		w.Write([]byte(greet))
	} else {
		w.WriteHeader(code)
	}
}

type promStruct struct{}

func (h *promStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

var ctr = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "greetings_count",
	Help: "Nr time greeting handler was called",
}, []string{"code"})

var hist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "greeting_seconds",
	Help:    "Time Taken to greet",
	Buckets: []float64{1, 2, 3, 4, 5},
}, []string{"code"})

func init() {
	prometheus.Register(hist)
	prometheus.Register(ctr)
}

var defHandler = new(handler)

func main() {
	http.Handle("/", defHandler)
	http.Handle("/metrics", prometheus.Handler())
	log.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
