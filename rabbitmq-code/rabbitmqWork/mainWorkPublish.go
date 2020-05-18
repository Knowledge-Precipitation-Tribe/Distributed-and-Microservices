package main

import (
	"fmt"
	"rabbitmq-code/rabbitmq"
	"strconv"
	"time"
)

//工作模式下发布多条消息
func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("testWork")

	for i:=0; i<=100; i++{
		rabbitmq.PublishSimple("hello world" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}