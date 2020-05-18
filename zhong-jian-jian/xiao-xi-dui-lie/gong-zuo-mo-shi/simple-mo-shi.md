# 基础代码

## Go语言操作

在go语言中我们使用[amqp](https://github.com/streadway/amqp)这个项目来操作rabbitmq：

```go
import "github.com/streadway/amqp"
```

![](../../../.gitbook/assets/image%20%2812%29.png)

## 创建URL

首先创建一下RabbitMQ对应的URL信息

```go
// URL格式：amqp://账号:密码@rabbitmq服务器地址:端口号/vhost
const MQURL = "amqp://guest:guest@127.0.0.1:5671/"
```

## 创建RabbitMQ结构体

结构体中用于存放一些RabbitMQ的信息

```go
type RabbitMQ struct {
	conn    *amqp.Connection
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
```

## 通用方法

```go
// 创建RabbitMQ实例
func NewRabbitMQ(queuqName string,
	exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queuqName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	//创建rabbitmq连接
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}

// 断开channel和connection的连接释放资源
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//自定义错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}
```

