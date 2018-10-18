package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mfrw/gophercises/webkv/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := flag.String("port", "8080", "Port to listen on")
	addrsStr := flag.String("addrs", "", "(Requred) Redis addrs")

	ttl := flag.Duration("ttl", time.Second*15, "service TTL check duration")
	flag.Parse()

	addrs := strings.Split(*addrsStr, ";")

	s, err := service.New(addrs, *ttl)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", prometheus.InstrumentHandler("webkv", s))
	http.Handle("/metrics", promhttp.Handler())
	l := fmt.Sprintf(":%s", *port)
	log.Print("Listening on ", l)
	log.Fatal(http.ListenAndServe(l, nil))
}
