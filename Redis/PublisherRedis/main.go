package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"pubsub/Entity"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

var ctx = context.Background()

func main() {

	app := gin.Default()
	router := app.Group("pubsub-redis")

	router.POST("/", func(c *gin.Context) {
		var user Entity.UserEntity

		if err := c.BindJSON(&user); err != nil {
			log.Fatal("failed to bind json:", err)
			return
		}

		payload, err := json.Marshal(user)
		if err != nil {
			log.Fatal("failed to marshal:", err)
			return
		}

		if err := redisClient.Publish(ctx, "send-user-data", payload).Err(); err != nil {
			log.Fatal("failed to publish message:", err)
			return
		}

		c.JSON(http.StatusOK, "ok")

	})

	app.Run("localhost:3000")
}
