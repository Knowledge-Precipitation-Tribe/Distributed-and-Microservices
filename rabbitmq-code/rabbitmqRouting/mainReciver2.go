package main

import "rabbitmq-code/common"

func main() {
	testTwo := common.NewRabbitMQRoutinh("testRouting", "testTwo")
	testTwo.ReceiveRouting()
}