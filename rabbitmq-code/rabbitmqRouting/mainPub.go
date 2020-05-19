package main

import (
	"log"
	"rabbitmq-code/common"
	"strconv"
	"time"
)

func main() {
	testOne := common.NewRabbitMQRoutinh("testRouting", "testOne")
	testTwo := common.NewRabbitMQRoutinh("testRouting", "testTwo")

	for i := 0; i<100; i++{
		testOne.PublishRouting("Routing模式testOne第" + strconv.Itoa(i) +"条数据")
		testTwo.PublishRouting("Routing模式testTwo第" + strconv.Itoa(i) +"条数据")
		log.Println("Routing模式生产第" + strconv.Itoa(i) +"条数据")
		time.Sleep(1 * time.Second)
	}
}