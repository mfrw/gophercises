package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mfrw/gophercises/ex18/primitive"
)

func trackTime(s time.Time, msg string) {
	e := time.Since(s)
	fmt.Println(msg, ":", e)
}

type mw func(w http.ResponseWriter, r *http.Request)

func auth(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("auth")
}

func auth(next mw) mw {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://google.com/", http.StatusFound)
}

func main() {
	b, err := primitive.Primitive("test.png", "out.png", 10, primitive.ModeTriangle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
