package main

import (
	"bytes"
	"context"
	"log"
	"net"

	"github.com/mfrw/gophercises/tracing/grpc/rpc"
	grpc "google.golang.org/grpc"
)

type fetchIt int

var _ rpc.FetchServer = (*fetchIt)(nil)

func (f *fetchIt) Captilize(ctx context.Context, in *rpc.Request) (*rpc.Payload, error) {
	out := &rpc.Payload{
		Data: bytes.ToUpper(in.Data),
	}
	return out, nil
}

func main() {
	addr := ":8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("gRPC server: failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	rpc.RegisterFetchServer(srv, new(fetchIt))
	log.Printf("fetchIt gRPC server serving at %q", addr)
	if err := srv.Serve(ln); err != nil {
		log.Fatalf("gRPC server: error serving: %v", err)
	}
}
