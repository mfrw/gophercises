package main

import (
	"bytes"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	//fmt.Fprintf(w, "<h1 style='color: %s'>Welcome!</h1>You are visitor number %v!", r.FormValue("color"), vistNum)
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()
	buf.WriteString("<h1 style='color: ")
	buf.WriteString(r.FormValue("color"))
	buf.WriteString("'> Welcome!</h1> You are visitor number ")
	b := strconv.AppendInt(buf.Bytes(), int64(vistNum), 10)
	b = append(b, '!')
	w.Write(b)
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
