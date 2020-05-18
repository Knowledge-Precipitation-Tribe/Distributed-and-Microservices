package main

import (
	"fmt"
	"rabbitmq-code/rabbitmq"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("testSimple")
	rabbitmq.PublishSimple("hello world")
	fmt.Println("发送成功")
}