package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pubsub/Entity"
	"pubsub/Go/SubscriberGo"
)

var subscriber SubscriberGo.Subscriber

func main() {

	app := gin.Default()
	router := app.Group("pubsub-golang")

	subChanel := make(chan Entity.UserEntity, 10)
	Subscriber := SubscriberGo.NewSubscriber(&subChanel)

	router.POST("/", func(ginContext *gin.Context) {

		var user Entity.UserEntity
		ginContext.BindJSON(&user)
		subChanel <- user
		Subscriber.Receiver()

		ginContext.JSON(http.StatusOK, "ok")

	})

	app.Run("localhost:3000")

}
