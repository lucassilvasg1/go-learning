package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/lucassilvasg1/go-learning/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(cliente pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Lucas",
		Email: "lucas@gmail.com",
	}

	res, err := cliente.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(cliente pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Lucas",
		Email: "lucas@gmail.com",
	}

	responseStream, err := cliente.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receivethe msg: %v", err)
		}

		fmt.Println("Status:", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "l1",
			Name:  "Lucas1",
			Email: "lucas1gmail.com",
		},
		&pb.User{
			Id:    "l2",
			Name:  "Lucas2",
			Email: "lucas2gmail.com",
		},
		&pb.User{
			Id:    "l3",
			Name:  "Lucas3",
			Email: "lucas3gmail.com",
		},
		&pb.User{
			Id:    "l4",
			Name:  "Lucas4",
			Email: "lucas4gmail.com",
		},
		&pb.User{
			Id:    "l5",
			Name:  "Lucas5",
			Email: "lucas5gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating quest: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
