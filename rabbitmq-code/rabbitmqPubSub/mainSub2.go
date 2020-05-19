package main

import (
	"rabbitmq-code/common"
)

func main() {
	rabbitmq := common.NewRabbitMQPubSub("testpubsub")
	rabbitmq.ReceiveSub()
}