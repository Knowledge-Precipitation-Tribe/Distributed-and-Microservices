package common

import (
	"github.com/streadway/amqp"
	"log"
)

//创建路由模式下RabbitMQ实例
func NewRabbitMQRoutinh(exchangeName string, routingkey string) *RabbitMQ {
	//创建RabbitMQ实例
	//传入指定的routingkey
	rabbitmq := NewRabbitMQ("", exchangeName, routingkey)
	var err error
	//获取conn
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnErr(err, "failed to connect rabbitmq")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailOnErr(err, "failed to open a channel")
	return rabbitmq
}

//路由模式下生产者
func (r *RabbitMQ) PublishRouting(message string) {
	//1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//设计交换机的类型为direct类型
		"direct",
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
		//指定routingkey
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

}

//路由模式下消费者
func (r *RabbitMQ) ReceiveRouting() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//修改为direct
		"direct",
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
		//在Routing下要指定routingkey
		r.Key,
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