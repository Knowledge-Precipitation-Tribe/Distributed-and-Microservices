package rabbitmqSimple

import (
	"rabbitmq-code/common"
)

func main() {
	rabbitmq := common.NewRabbitMQSimple("testSimple")
	rabbitmq.ConsumeSimple()
}