package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mfrw/gophercises/ex18/primitive"
)

func trackTime(s time.Time, msg string) {
	e := time.Since(s)
	fmt.Println(msg, ":", e)
}

func main() {
	b, err := primitive.Primitive("test.png", "out.png", 10, ModeTriangle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
