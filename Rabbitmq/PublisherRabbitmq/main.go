package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"pubsub/Entity"
)

var ctx = context.Background()

var rabbitmqChanel *amqp.Channel

func main() {

	app := gin.Default()
	router := app.Group("pubsub-rabbitmq")

	// connect to rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErrorRabbitmq(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// open channel
	rabbitmqChanel, err = conn.Channel()
	failOnErrorRabbitmq(err, "Failed to open a channel")
	defer rabbitmqChanel.Close()

	q, err := rabbitmqChanel.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnErrorRabbitmq(err, "Failed to declare a queue")

	router.POST("/", func(ginContext *gin.Context) {

		var user Entity.UserEntity

		if err := ginContext.BindJSON(&user); err != nil {
			log.Fatal("failed to bind json:", err)
			return
		}

		payload, err := json.Marshal(user)
		if err != nil {
			log.Fatal("failed to marshal:", err)
			return
		}

		err = rabbitmqChanel.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        payload,
			})
		failOnErrorRabbitmqGin(err, "Failed to publish a message", ginContext)

		ginContext.JSON(http.StatusOK, "ok")

	})

	app.Run("localhost:3000")
}

// failOnErrorRabbitmq is a func for handle rabbitmq errors
func failOnErrorRabbitmqGin(err error, msg string, context *gin.Context) {
	if err != nil {
		fmt.Errorf("error %s: %s", msg, err)
		context.JSON(http.StatusInternalServerError, err)
		context.Done()
	}
}

// failOnErrorRabbitmq is a func for handle rabbitmq errors
func failOnErrorRabbitmq(err error, msg string) {
	if err != nil {
		fmt.Errorf("error %s: %s", msg, err)
	}
}
