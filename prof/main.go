package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
)

var visitors struct {
	sync.Mutex
	n int
}

var colorRegexp = regexp.MustCompile(`^\w*$`)

func main() {
	log.Printf("Starting on port 8080")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	if match := colorRegexp.MatchString(r.FormValue("color")); !match {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	visitors.Lock()
	visitors.n++
	vistNum := visitors.n
	visitors.Unlock()
	w.Header().Set("Content-Type", "text/html; charset-utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome!</h1>You are visitor number " + fmt.Sprint(vistNum) + "!"))
}
