# RabbitMQ

RabbitMQ官网：[https://www.rabbitmq.com/](https://www.rabbitmq.com/)

![](../../../.gitbook/assets/image%20%2868%29.png)

## 安装

关于RabbitMQ安装启动请查看：[Docker-Compose方式启动RabbitMQ服务](https://docs.docker.knowledge-precipitation.site/chang-yong-fu-wu-pei-zhi/da-jian-rabbitmq-fu-wu)

## 什么是RabbitMQ

首先介绍一下什么是MQ，消息队列（Message Queue，简称MQ），从字面意思上看，本质是个队列，FIFO先入先出，只不过队列中存放的内容是message而已。 其主要用途：不同进程Process/线程Thread之间通信。

RabbitMQ是基于AMQP协议的面向消息的中间件，使用Erlang编写，用于组件之间的解耦，主要体现在消息的生产者和消费者之间无强依赖关系，代码整体是开源的。

支持AMQP/XMPP/SMTP/STOMP多种协议。适合企业级开发

由于 TCP 连接的创建和销毁开销较大，且并发数受系统资源限制，会造成性能瓶颈。**RabbitMQ 使用信道的方式来传输数据**。信道是建立在真实的 TCP 连接内的虚拟连接，且每条 TCP 连接上的信道数量没有限制。

## 特点

* 高可用
* 拓展性
* 多语言客户端
* 界面管理

## 使用场景

* 流量削峰
* 异步处理
* 应用接耦

## 推荐阅读

\[1\] RabbitMQ教程：[https://blog.csdn.net/hellozpc/article/details/81436980](https://blog.csdn.net/hellozpc/article/details/81436980)

