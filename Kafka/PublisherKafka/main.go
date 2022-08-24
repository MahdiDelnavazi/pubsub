package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"pubsub/Entity"
)

var ctx = context.Background()

var ConnectionKafka *kafka.Conn

func main() {

	app := gin.Default()
	router := app.Group("pubsub-kafka")

	topic := "my-topic"
	partition := 0

	var kafkaErr error
	ConnectionKafka, kafkaErr = kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if kafkaErr != nil {
		log.Fatal("failed to dial leader:", kafkaErr)
	}

	router.POST("/", func(ginContext *gin.Context) {

		//conn.SetWriteDeadline(time.Now().Add(50 * time.Second))

		var user Entity.UserEntity
		ginContext.ShouldBindJSON(&user)
		userJson, err := json.Marshal(user)

		_, err = ConnectionKafka.WriteMessages(
			kafka.Message{Value: userJson},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		ginContext.JSON(http.StatusOK, "ok")

	})

	app.Run("localhost:3000")

	if err := ConnectionKafka.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
