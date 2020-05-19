# Publish/Subscribe发布订阅模式

## **publish/subscribe发布订阅**

消息被路由投递到多个队列，一个消息被多个消费者获取，与之前不同的是这里在生产者后面又添加了一个深蓝色的X，这个X代表Exchange交换机。工作过程如下：

* 生产者将消息传递给交换机
* 交换机通过指定的规则将消息传递给不同的队列，所以同一条消息可以被多个消费者获取
* 消息由每个队列绑定的消费者进行消费

![](../../../.gitbook/assets/image%20%2836%29.png)

## 创建RabbitMQ实例

```go
//创建发布订阅模式下的rabbitmq实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	//创建RabbitMQ实例
	//在这个模式下不需要指定queuename，但是要指定交换机名称
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	//获取conn
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}
```

## 创建生产者

```go
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

	r.failOnErr(err, "failed to declare an exchange")

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
```

## 创建消费者

```go
//订阅模式下消费者
func (r *RabbitMQ) ReceiveSub() {
  //尝试创建交换机
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

	r.failOnErr(err, "failed to decalre an exchange")

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
	r.failOnErr(err, "failed to declare a queue")

  //将队列绑定在指定的交换机上
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
```

## 执行程序

先启动消费者监听队列，后启动生产者生产消息

### 消费者

可以看到消费者已经获取到队列中的消息

![](../../../.gitbook/assets/image%20%2846%29.png)

![](../../../.gitbook/assets/image%20%2856%29.png)

### 生产者

![](../../../.gitbook/assets/image%20%2860%29.png)

### RabbitMQ管理界面

可以看到我们指定名称的交换机已经创建成功

![](../../../.gitbook/assets/image%20%2857%29.png)

我们点击进入交换机后可以看到里面的正在监控消息的发送过程

![](../../../.gitbook/assets/image%20%2845%29.png)

这时我们到队列中查看一下，可以看到RabbitMQ已经为我们创建了两个随机名称的队列

![](../../../.gitbook/assets/image%20%2853%29.png)

点进其中一个队列查看一下具体信息，可以看到里面信息发送的监控

![](../../../.gitbook/assets/image%20%2862%29.png)

## 代码位置

完整代码位置：[rabbitmq-PubSub](https://github.com/Knowledge-Precipitation-Tribe/Distributed-and-Microservices/blob/master/rabbitmq-code/common/rabbitmq-PubSub.go)

main函数：[rabbitmqPubSub](https://github.com/Knowledge-Precipitation-Tribe/Distributed-and-Microservices/blob/master/rabbitmq-code/common/rabbitmq-PubSub.go)

