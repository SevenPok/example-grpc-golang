package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"github.com/segmentio/kafka-go"

	pb "example.com/grpc/gen/proto"
	"google.golang.org/grpc"
)
const (
	topic         = "user-log"
	brokerAddress = "localhost:9092"
)
type TestApiServer struct {
	pb.UnimplementedTestApiServer
}

func produce(ctx context.Context, req *pb.User) {
	l := log.New(os.Stdout, "kafka escribiendo: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger: l,
	})
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(strconv.Itoa(int(req.Id))),
		Value: []byte("{Id:"+strconv.Itoa(int(req.Id))+", Name: "+req.Name+", Age: "+strconv.Itoa(int(req.Age))+"}"),
	})
	if err != nil {
		panic("No se pudo enviar mensaje: " + err.Error())
	}
}


func (s *TestApiServer) CreateUser(ctx context.Context, req *pb.User) (*pb.Response, error) {
	msg := pb.Response{Msg: "Usuario Creado"}
	fmt.Println(req)
	produce(ctx,req)
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
