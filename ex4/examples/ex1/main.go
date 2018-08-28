package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mfrw/gophercises/ex4"
)

var exampleHtml = `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page">A link to other page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := ex4.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(links)
}
