# Work工作模式

## Work Queue （工作队列）

生产者发布多条消息到消息队列中，可以有多个消费者监听队列，通过类似抢占的方式获取消息，一个消息只能被一个消费者获取，相当于简单模式下拓展多个消费者用于复杂均衡。

![](../../../../.gitbook/assets/image%20%2865%29.png)

## 创建RabbitMQ Work模式

![](../../../../.gitbook/assets/image%20%2841%29.png)

与Simple模式相比只是多了一个消费者而已

### 生产者代码

```go
//工作模式下发布多条消息
func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("testWork")

	for i:=0; i<=100; i++{
		rabbitmq.PublishSimple("hello world" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
```

### 消费者代码

```go
func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("testWork")
	rabbitmq.ConsumeSimple()
}
```

## 执行程序

### 生产者发布消息

![](../../../../.gitbook/assets/image%20%2828%29.png)

### 消费者获取消息

消费者1

![](../../../../.gitbook/assets/image%20%2815%29.png)

消费者2

![](../../../../.gitbook/assets/image%20%2820%29.png)

### 查看RabbitMQ管理界面

![](../../../../.gitbook/assets/image%20%2838%29.png)

## 代码位置

main函数：[rabbitmqWork](https://github.com/Knowledge-Precipitation-Tribe/Distributed-and-Microservices/tree/master/rabbitmq-code/rabbitmqWork)

