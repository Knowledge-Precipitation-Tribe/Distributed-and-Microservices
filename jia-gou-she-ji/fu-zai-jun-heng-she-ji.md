# 负载均衡设计

## 负载均衡系统

### 硬件

* F5
* A10
* Radware

### 软件

* LVS
* Nginx
* HAProxy

### 负载均衡算法

* Random
* RoundRobin
* ConsistenHash

## 负载均衡系统的意义

所谓的负载均衡不仅仅意味着我们要尽量保证请求分布在不同的服务器上，也包含完整的故障处理与恢复机制。因为服务的故障如果没有很好的处理机制很容易造成系统的雪崩。

如果我们业务模块的某个服务器故障，这时如何发现它的故障呢？

因为我们的服务都是会在注册中心例如ZK进行注册，而且服务模块会定期上报心跳，如果ZK发现某个服务模块没有按时上报心跳，那么就会从维护的列表中移除它的ip，此时网关通过订阅发现当前模块不可达，就会转移进行访问其他服务器。

一般情况下网关会与每个业务服务器建立多个长连接，如果发现服务不可达就将当前的长连接删除。

但是对于注册中心来说只要服务心跳存在，ZK 就认为服务是处在可用状态，但是服务如果处在假死的状态，ZK 是无从得知的。这个时候，业务逻辑服务是否真正可用只能够由网关知道。因为我们的请求都是有返回值的，所以我们可以根据每个服务器的返回值放入队列并进行统计，如果某个业务服务器的错误返回码比例过高就可以认定当前的业务服务器故障，此时通知报警，运维人员或注册中心进行业务服务器的重启。

在服务假死时，我们首先要保存业务的上下文信息 ，然后将业务服务器关机，并等待一会使得注册中心发现业务服务器没有心跳，进行业务服务器的摘除，此时再重启业务服务器，那么当前的服务器就会在注册中心进行重新注册了。

