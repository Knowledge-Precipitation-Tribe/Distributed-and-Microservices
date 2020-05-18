package rabbitmq

// Simple模式

import "github.com/streadway/amqp"

// URL格式：amqp://账号:密码@rabbitmq服务器地址:端口号/vhost
const MOUEL = "amqp://guest:guest@127.0.0.1:5672/"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel

	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}

