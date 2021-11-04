package main

import (
	"context"
	"fmt"
	"log"

	pb "example.com/grpc/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.CreateUser(context.Background(), &pb.User{Id: 2, Name: "Pablo", Age: 21})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
