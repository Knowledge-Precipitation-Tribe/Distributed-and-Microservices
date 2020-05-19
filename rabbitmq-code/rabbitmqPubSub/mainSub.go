package main

import "rabbitmq-code/rabbitmq"

func main() {
	rabbitmq := rabbitmq.NewRabbitMQPubSub("testpubsub")
	rabbitmq.ReceiveSub()
}