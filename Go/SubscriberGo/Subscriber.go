package SubscriberGo

import (
	"fmt"
	"pubsub/Entity"
	"sync"
)

type Subscriber struct {
	subChanel *chan Entity.UserEntity
}

func NewSubscriber(subChan *chan Entity.UserEntity) *Subscriber {
	count = 0
	return &Subscriber{subChanel: subChan}
}

var mutex sync.Mutex
var count int

func (subscriber *Subscriber) Receiver() {
	msg := <-*subscriber.subChanel
	mutex.Lock()
	count = count + 1
	mutex.Unlock()
	fmt.Println("this is user :", msg, count)
}
