package main

import (
	"rabbitmq-code/common"
)

func main() {
	rabbitmq := common.NewRabbitMQSimple("testWork")
	rabbitmq.ConsumeSimple()
}