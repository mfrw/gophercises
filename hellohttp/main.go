package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

var httpString = `
<html>
<head>
<title>Test</title>
</head>
<body>Test Body</body>
</html>
`

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, httpString)
}
