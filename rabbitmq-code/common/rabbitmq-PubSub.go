package common

import (
	"github.com/streadway/amqp"
	"log"
)

//创建发布订阅模式下的rabbitmq实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	//创建RabbitMQ实例
	//在这个模式下不需要指定queuename，但是要指定交换机名称
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	//获取conn
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnErr(err, "failed to connect rabbitmq")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailOnErr(err, "failed to open a channel")
	return rabbitmq
}

//订阅模式下生产者
func (r *RabbitMQ) PublishPub(message string) {
	//1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//设计交换机的类型为广播类型
		"fanout",
		//是否持久化
		true,
		//是否自动删除
		false,
		//如果设置为true，则这个exchange不可以被client用来推送消息，仅用来进行交换机之间的绑定
		false,
		//是否阻塞
		false,
		nil,
	)

	r.FailOnErr(err, "failed to declare an exchange")

	//发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

}

//订阅模式下消费者
func (r *RabbitMQ) ReceiveSub() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	r.FailOnErr(err, "failed to decalre an exchange")

	//试探性创建队列，注意不要写队列名称
	q, err := r.channel.QueueDeclare(
		//注意创建队列不要指定名称，rabbitmq会生成随机的队列名称
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.FailOnErr(err, "failed to declare a queue")

	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		"",
		r.Exchange,
		false,
		nil,
	)

	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range messges{
			log.Printf("Received a message : %s", d.Body)
		}
	}()

	<-forever
}