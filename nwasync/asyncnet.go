package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go copyToStderr(conn)
	}
}

func copyToStderr(conn net.Conn) {
	n, err := io.Copy(os.Stderr, conn)
	log.Printf("Copied %d bytes; finished with err = %v\n", n, err)
}
