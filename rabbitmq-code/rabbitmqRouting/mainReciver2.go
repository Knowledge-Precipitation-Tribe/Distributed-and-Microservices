package main

import "rabbitmq-code/common"

func main() {
	testTwo := common.NewRabbitMQRouting("testRouting", "testTwo")
	testTwo.ReceiveRouting()
}