
serverRedis:
	go run Redis/PublisherRedis/main.go

serverSubscriberRedis:
	go run Redis/SubscriberRedis/main.go

serverKafka:
	go run Kafka/PublisherKafka/main.go

serverSubscriberKafka:
	go run Kafka/SubscriberKafka/main.go

serverRabbitmq:
	go run Rabbitmq/PublisherRabbitmq/main.go

serverSubscriberRabbitmq:
	go run Rabbitmq/SubscriberRabbitmq/main.go

serverGo:
	go run Go/PublisherGo/main.go


.PHONY: server