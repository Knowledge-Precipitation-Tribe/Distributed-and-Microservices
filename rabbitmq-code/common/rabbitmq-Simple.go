package common

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//简单模式：step1 创建简单模式下的RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	//simple模式下使用默认的exchange
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式：step2 创建简单模式下的生产者
func (r *RabbitMQ) PublishSimple(message string) {
	//1. 申请队列，如果队列不存在则自动创建，如果存在则获取存在的队列
	//保证队列存在，使消息发送到队列中
	_, err := r.channel.QueueDeclare(r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性，独占队列
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//2. 发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		// 如果为true， 则根据exchange类型和routkey规则，如果无法找到符合条件的队列
		// 那么会把发送的消息回退给publish
		false,
		//如果为true，当exchange发送消息到队列后发现没有consume，则会把发送的消息返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Println(err)
	}
}

// 简单模式：step3 创建简单模式下的消费者
func (r *RabbitMQ) ConsumeSimple() {
	//1. 申请队列，如果队列不存在则自动创建，如果存在则获取存在的队列
	//保证队列存在，使消息发送到队列中
	_, err := r.channel.QueueDeclare(r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性，独占队列
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//接受消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		//用于区分多个不同的消费者
		"",
		//是否自动应答，也就是消费者消费一个队列后是否主动告知rabbitmq当前的消息我已经消费完
		//rabbitmq会根据这个判断是否可以删除该消息
		//为false的话要手动实现
		true,
		//是否具有排他性
		false,
		//如果为true不能在同一个connection中发送消息传递给当前conn的消费者
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
			//实现其他的逻辑函数
		}
	}()
	log.Println("waiting for message")
	<-forever
}