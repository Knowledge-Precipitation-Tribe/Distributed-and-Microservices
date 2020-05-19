package main

import (
	"log"
	"rabbitmq-code/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQPubSub("testpubsub")
	for i := 0; i<100; i++{
		rabbitmq.PublishPub("订阅模式第" + strconv.Itoa(i) +"条数据")
		log.Println("订阅模式生产第" + strconv.Itoa(i) +"条数据")
		time.Sleep(1 * time.Second)
	}
}