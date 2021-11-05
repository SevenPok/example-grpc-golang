package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)
const (
	topic         = "user-log"
	brokerAddress = "localhost:9092"
)
func main() {
	ctx := context.Background()
	consume(ctx)
}
func consume(ctx context.Context) {
	l := log.New(os.Stdout, "kafka leyendo: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		Logger: l,
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("No se pudo leer mensaje: " + err.Error())
		}
		fmt.Println("mensaje recibido: ", string(msg.Value))
	}
}