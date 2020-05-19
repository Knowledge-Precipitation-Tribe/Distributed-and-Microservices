package main

import "rabbitmq-code/common"

func main() {
	testOne := common.NewRabbitMQRouting("testRouting", "testOne")
	testOne.ReceiveRouting()
}