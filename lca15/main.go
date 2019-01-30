package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
)

func main() {

	// 1 Listen for connections
	if ln, err := net.Listen("tcp", ":8080"); err == nil {
		for {
			// 2 Accept a connection
			if conn, err := ln.Accept(); err == nil {
				reader := bufio.NewReader(conn)
				// 3 Read a request
				if req, err := http.ReadRequest(reader); err == nil {
					// 4 Dial to backend
					if be, err := net.Dial("tcp", "127.0.0.1:8081"); err == nil {
						be_reader := bufio.NewReader(be)
						// 5 Send request to backend
						if err := req.Write(be); err == nil {
							// 6 Read Response from backend
							if resp, err := http.ReadResponse(be_reader, req); err == nil {
								resp.Close = true
								// 7 Send Response from backend back to client
								if err := resp.Write(conn); err == nil {
									log.Printf("%s: %d", req.URL.Path, resp.StatusCode)
								}
								// 8 Close and be done!!
								conn.Close()
							}
						}
					}
				}
			}
		}
	}
}
