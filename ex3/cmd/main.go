package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mfrw/gophercises/ex3"
)

func main() {
	file := flag.String("file", "gopher.json", "Ther JSON file with the story")
	port := flag.String("port", ":8080", "Port to start the webapp")
	flag.Parse()

	fmt.Println("Parsing", *file)

	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	story, err := ex3.JsonStory(r)
	if err != nil {
		log.Fatal(err)
	}
	h := ex3.NewHandler(story)
	log.Fatal(http.ListenAndServe(*port, h))
}
