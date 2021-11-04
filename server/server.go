package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "example.com/grpc/gen/proto"
	"google.golang.org/grpc"
)

type TestApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *TestApiServer) CreateUser(ctx context.Context, req *pb.User) (*pb.Response, error) {
	msg := pb.Response{Msg: "Usuario Creado"}
	fmt.Println(req)
	return &msg, nil
}

func main() {
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &TestApiServer{})

	log.Println("Server on port 8080")
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal(err)
	}
}
