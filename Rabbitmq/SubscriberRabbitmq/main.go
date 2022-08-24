package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var ctx = context.Background()

type UserSubEntity struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// connect to rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErrorRabbitmq(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErrorRabbitmq(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnErrorRabbitmq(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErrorRabbitmq(err, "Failed to register a Consumer")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var user UserSubEntity
			fmt.Println("this is user :", msg.Body)
			if err := json.Unmarshal(msg.Body, &user); err != nil {
				log.Fatal("failed to unmarshal:", err)
			}

			fmt.Println("Received message from main channel.")
			fmt.Printf("%+v\n", user)
		}
	}()

	<-forever

}

// failOnErrorRabbitmq is a func for handle rabbitmq errors
func failOnErrorRabbitmq(err error, msg string) {
	if err != nil {
		fmt.Errorf("error %s: %s", msg, err)
	}
}
