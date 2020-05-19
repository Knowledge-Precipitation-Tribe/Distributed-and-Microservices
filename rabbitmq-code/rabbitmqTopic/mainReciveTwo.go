package main

import "rabbitmq-code/common"

func main() {
	rabbitmq := common.NewRabbitMQTopic("testTopic", "test.*.two")
	rabbitmq.ReceiveTopic()
}