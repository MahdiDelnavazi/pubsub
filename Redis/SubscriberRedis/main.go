package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"pubsub/Entity"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {

	subscriber := redisClient.Subscribe(ctx, "send-user-data")

	var user Entity.UserEntity

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			log.Fatal("failed to receive:", err)
			return
		}
		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			log.Fatal("failed to unmarshal:", err)
			return
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", user)
	}
}
