package main

import "rabbitmq-code/rabbitmq"

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("testSimple")
	rabbitmq.ConsumeSimple()
}