package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var ok = true

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if ok {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")
}

func main() {
	go func() {
		c := time.NewTicker(200 * time.Millisecond)
		for {
			<-c.C
			ok = !ok
		}
	}()

	router := mux.NewRouter()
	router.Use(handlers.CompressHandler)

	router.HandleFunc("/", index).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
