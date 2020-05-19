package main

import (
	"log"
	"rabbitmq-code/common"
	"strconv"
	"time"
)

func main() {
	testOne := common.NewRabbitMQTopic("testTopic", "test.topic.one")
	testTwo := common.NewRabbitMQTopic("testTopic", "test.topic.two")

	for i := 0; i<100; i++{
		testOne.PublishTopic("Topic模式One第" + strconv.Itoa(i) +"条数据")
		testTwo.PublishTopic("Topic模式Two第" + strconv.Itoa(i) +"条数据")
		log.Println("Topic模式生产第" + strconv.Itoa(i) +"条数据")
		time.Sleep(1 * time.Second)
	}
}