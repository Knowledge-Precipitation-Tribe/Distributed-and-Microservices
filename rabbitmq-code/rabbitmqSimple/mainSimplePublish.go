package rabbitmqSimple

import (
	"fmt"
	"rabbitmq-code/common"
)

func main() {
	rabbitmq := common.NewRabbitMQSimple("testSimple")
	rabbitmq.PublishSimple("hello world")
	fmt.Println("发送成功")
}