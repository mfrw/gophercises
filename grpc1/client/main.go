package main

import (
	"context"
	"io"
	"log"

	pb "../customer"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatal(err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Customer : %w", customer)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewCustomerClient(conn)

	customer := &pb.CustomerRequest{
		Id:    101,
		Name:  "Muhammad Falak",
		Email: "some@ex.com",
		Phone: "1234",
		Address: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "1 Mission St",
				City:              "SF",
				State:             "CA",
				Zip:               "1234",
				IsShippingAddress: false,
			},
			&pb.CustomerRequest_Address{
				Street:            "GF",
				City:              "Kochi",
				State:             "KL",
				Zip:               "1233",
				IsShippingAddress: true,
			},
		},
	}
	createCustomer(client, customer)

}
