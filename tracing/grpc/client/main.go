package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mfrw/gophercises/tracing/grpc/rpc"
	grpc "google.golang.org/grpc"
)

func main() {
	serverAddr := ":8080"
	cc, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fetchIt gRPC client failed to dial to server: %v", err)
	}
	fc := rpc.NewFetchClient(cc)

	fIn := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _, err := fIn.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()

		out, err := fc.Captilize(ctx, &rpc.Request{Data: line})
		if err != nil {
			log.Printf("fetchIt gRPC client got error from server: %v", err)
			continue
		}
		fmt.Printf("< %s\n\n", out.Data)
	}
}
