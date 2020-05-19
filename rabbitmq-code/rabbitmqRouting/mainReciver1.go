package main

import "rabbitmq-code/common"

func main() {
	testOne := common.NewRabbitMQRoutinh("testRouting", "testOne")
	testOne.ReceiveRouting()
}