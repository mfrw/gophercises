package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	cache "github.com/mfrw/gophercises/arccache"
)

var c *cache.Cache

func main() {
	var err error
	c, err = cache.New(8192)
	if err != nil {
		log.Fatal(err)
	}
	m := mux.NewRouter()
	m.HandleFunc("/add/{key}", addHandler).Methods("POST")
	m.HandleFunc("/contains/{key}", containsHandler).Methods("GET")
	m.HandleFunc("/containsoradd/{key}", containOrAddHandler).Methods("POST")
	m.HandleFunc("/get/{key}", getHandler).Methods("GET")
	m.HandleFunc("/keys", keysHandler).Methods("GET")
	m.HandleFunc("/peek/{key}", peekHandler).Methods("GET")
	m.HandleFunc("/purge", purgeHandler).Methods("GET")
	m.HandleFunc("/remove/{key}", removeHandler).Methods("GET")
	m.HandleFunc("/removeoldest", removeOldestHandler).Methods("GET")
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key, ok := mux.Vars(r)["key"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"result\": false")
		return
	}
	value := r.PostFormValue("value")
	c.Add(key, value)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"result\": true")
}

func containsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key, ok := mux.Vars(r)["key"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"result\": false")
		return
	}
	w.WriteHeader(http.StatusOK)
	ok = c.Contains(key)
	if !ok {
		fmt.Fprintf(w, "{\"result\": false")
		return
	}
	fmt.Fprintf(w, "{\"result\": true")
}

func containOrAddHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func keysHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func peekHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func purgeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func removeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func removeOldestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
