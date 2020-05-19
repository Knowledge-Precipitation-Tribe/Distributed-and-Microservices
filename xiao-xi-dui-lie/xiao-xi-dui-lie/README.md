# RabbitMQ

RabbitMQ官网：[https://www.rabbitmq.com/](https://www.rabbitmq.com/)

![](../../.gitbook/assets/image%20%2868%29.png)

## 安装

关于RabbitMQ安装启动请查看：[Docker-Compose方式启动RabbitMQ服务](https://docs.docker.knowledge-precipitation.site/chang-yong-fu-wu-pei-zhi/da-jian-rabbitmq-fu-wu)

## 什么是RabbitMQ

首先介绍一下什么是MQ，消息队列（Message Queue，简称MQ），从字面意思上看，本质是个队列，FIFO先入先出，只不过队列中存放的内容是message而已。 其主要用途：不同进程Process/线程Thread之间通信。

RabbitMQ是面向消息的中间件，用于组件之间的解耦，主要体现在消息的生产者和消费者之间无强依赖关系。

其是在[AMQP](https://baike.baidu.com/item/AMQP)基础上实现的完整的消息系统。

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

