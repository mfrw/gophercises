package main

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	pb "../rpc"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func main() {
	perRPC := oauth.NewOauthAccess(fetchToken())
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}),
		),
	}
	conn, err := grpc.Dial(":8080", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "authenticated-client"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greeting: %s\n", r.Message)
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
