package main

import "rabbitmq-code/common"

func main() {
	rabbitmq := common.NewRabbitMQTopic("testTopic", "#")
	//接受所有的消息
	rabbitmq.ReceiveTopic()
}