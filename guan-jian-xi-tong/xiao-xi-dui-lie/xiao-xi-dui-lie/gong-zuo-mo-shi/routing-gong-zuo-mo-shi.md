# Routing工作模式

## Routing模式

生产者生产消息后发送给交换机，此时交换机根据其指定的key将消息发送给不同的queue，在创建消费者时也指定了其对应的key，所以保证该消费者只能获取对应queue上的消息，在这种模式下一个消息可以被多个消费者获取。并且消息的目标队列可被生产者指定。

![](../../../../.gitbook/assets/image%20%2866%29.png)

## 创建RabbitMQ实例

与pubsub模式不同的就是要指定routingkey

```go
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
```

## 创建生产者

与pubsub模式下不同的是这里要将交换机的类型修改为direct，而且在消息发送时要指定key

```go
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
```

## 创建消费者

与pubsub模式下不同的是这里要将交换机的类型修改为direct，而且在绑定交换机时要指定key

```go
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
```

## 执行程序

先启动消费者监听队列，后启动生产者生产消息

### 消费者

可以看到消费者已经获取到队列中的消息，而且根据路由规则，每个消费者只能获取当前队列的消息。

![](../../../../.gitbook/assets/image%20%2843%29.png)

![](../../../../.gitbook/assets/image%20%2848%29.png)

### 生产者

![](../../../../.gitbook/assets/image%20%2849%29.png)

### RabbitMQ管理界面

可以看到我们指定名称的交换机已经创建成功

![](../../../../.gitbook/assets/image%20%2844%29.png)

我们点击进入交换机后可以看到里面的正在监控消息的发送过程

![](../../../../.gitbook/assets/image%20%2861%29.png)

这时我们到队列中查看一下，可以看到RabbitMQ已经为我们创建了两个随机名称的队列

![](../../../../.gitbook/assets/image%20%2863%29.png)

点进其中一个队列查看一下具体信息，可以看到里面信息发送的监控

![](../../../../.gitbook/assets/image%20%2851%29.png)

## 代码位置

完整代码位置：[rabbitmq-Routing](https://github.com/Knowledge-Precipitation-Tribe/Distributed-and-Microservices/blob/master/rabbitmq-code/common/rabbitmq-Routing.go)

main函数：[rabbitmqRouting](https://github.com/Knowledge-Precipitation-Tribe/Distributed-and-Microservices/tree/master/rabbitmq-code/rabbitmqRouting)

