package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

var ctx = context.Background()

func main() {

	// to consume messages
	config := kafka.ReaderConfig{Brokers: []string{"localhost:9092"},
		Topic:    "my-topic",
		GroupID:  "g1",
		MaxBytes: 20}

	reader := kafka.NewReader(config)

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("some error :", err)
			continue
		}
		fmt.Println("message :", string(message.Value))
	}

}
