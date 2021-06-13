# Table of contents

* [Distributed and Cloud Native](README.md)
* [集群,分布式，微服务概念和区别理解](ji-qun-fen-bu-shi-wei-fu-wu-gai-nian-he-qu-bie-li-jie.md)
* [IaaS,PaaS,SaaS](iaas-paas-saas.md)
* [到底什么是“云原生”？](yun-yuan-sheng.md)
* [架构师图谱](jia-gou-shi-tu-pu/README.md)
  * [架构师图谱（上）](jia-gou-shi-tu-pu/jia-gou-shi-tu-pu-shang.md)

## 架构设计

* [架构的演进](jia-gou-she-ji/untitled/README.md)
  * [查询分离模式](jia-gou-she-ji/untitled/cha-xun-fen-li-mo-shi.md)
  * [内容分发模式](jia-gou-she-ji/untitled/nei-rong-fen-fa-mo-shi.md)
  * [多级缓存模式](jia-gou-she-ji/untitled/duo-ji-huan-cun-mo-shi.md)
  * [分库分表模式](jia-gou-she-ji/untitled/fen-ku-fen-biao-mo-shi.md)
  * [弹性伸缩模式](jia-gou-she-ji/untitled/tan-xing-shen-suo-mo-shi.md)
  * [多机房模式](jia-gou-she-ji/untitled/duo-ji-fang-mo-shi.md)
  * [高并发秒杀系统](jia-gou-she-ji/untitled/gao-bing-fa-miao-sha-xi-tong.md)
* [架构设计哲学](jia-gou-she-ji/jia-gou-she-ji-zhe-xue.md)
* [高可用设计](jia-gou-she-ji/gao-ke-yong-she-ji.md)
* [高并发设计](jia-gou-she-ji/gao-bing-fa-she-ji/README.md)
  * [高并发系统的分析和设计](jia-gou-she-ji/gao-bing-fa-she-ji/untitled-2.md)
* [服务无状态设计](jia-gou-she-ji/fu-wu-wu-zhuang-tai-she-ji/README.md)
  * [一致性Hash](jia-gou-she-ji/fu-wu-wu-zhuang-tai-she-ji/yi-zhi-xing-hash.md)
* [负载均衡设计](jia-gou-she-ji/fu-zai-jun-heng-she-ji.md)
* [服务密等设计](jia-gou-she-ji/fu-wu-mi-deng-she-ji.md)
* [分布式锁设计](jia-gou-she-ji/fen-bu-shi-suo-she-ji.md)
* [分布式事务设计](jia-gou-she-ji/fen-bu-shi-shi-wu-she-ji.md)
* [服务降级设计](jia-gou-she-ji/fu-wu-jiang-ji-she-ji.md)
* [服务熔断设计](jia-gou-she-ji/fu-wu-xian-liu-rong-duan-she-ji.md)
* [服务灰度发布\(金丝雀发布\)设计](jia-gou-she-ji/fu-wu-hui-du-fa-bu-she-ji.md)
* [服务全链路压测设计](jia-gou-she-ji/fu-wu-quan-lian-lu-ya-ce-she-ji/README.md)
  * [如何做压测](jia-gou-she-ji/fu-wu-quan-lian-lu-ya-ce-she-ji/ru-he-zuo-ya-ce.md)

## 云原生 <a id="yun-yuan-sheng-1"></a>

* [推荐资源](yun-yuan-sheng-1/tui-jian-zi-yuan.md)
* [技术文章](yun-yuan-sheng-1/ji-shu-wen-zhang.md)

## 关键系统

* [注册/控制中心](guan-jian-xi-tong/zhu-ce-zhong-xin/README.md)
  * [Zookeeper](guan-jian-xi-tong/zhu-ce-zhong-xin/zookeeper/README.md)
    * [ZooKeeper服务器搭建](guan-jian-xi-tong/zhu-ce-zhong-xin/zookeeper/zookeeper-fu-wu-qi-da-jian.md)
  * [ETCD](guan-jian-xi-tong/zhu-ce-zhong-xin/etcd.md)
  * [Eureka](guan-jian-xi-tong/zhu-ce-zhong-xin/eureka.md)
  * [Consul](guan-jian-xi-tong/zhu-ce-zhong-xin/consul.md)
* [配置中心](guan-jian-xi-tong/pei-zhi-zhong-xin/README.md)
  * [Apollo](guan-jian-xi-tong/pei-zhi-zhong-xin/apollo.md)
* [消息队列](guan-jian-xi-tong/xiao-xi-dui-lie/README.md)
  * [使用MQ可能会遇到的问题](guan-jian-xi-tong/xiao-xi-dui-lie/shi-yong-mq-ke-neng-hui-yu-dao-de-wen-ti.md)
  * [kafka](guan-jian-xi-tong/xiao-xi-dui-lie/kafka.md)
  * [RocketMQ](guan-jian-xi-tong/xiao-xi-dui-lie/rocketmq.md)
  * [Redis](guan-jian-xi-tong/xiao-xi-dui-lie/redis.md)
  * [RabbitMQ](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/README.md)
    * [如何保证消息的发送与消费](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/ru-he-bao-zheng-xiao-xi-de-fa-song-yu-xiao-fei.md)
    * [如何保证消息的可靠传输](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/ru-he-bao-zheng-xiao-xi-de-ke-kao-chuan-shu.md)
    * [如何保证RabbitMQ高可用](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/ru-he-bao-zheng-rabbitmq-gao-ke-yong.md)
    * [核心概念](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/he-xin-gai-nian/README.md)
      * [Virtual hosts](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/he-xin-gai-nian/virtual-hosts.md)
      * [Exchanges](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/he-xin-gai-nian/exchanges.md)
      * [Queue](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/he-xin-gai-nian/queue.md)
      * [Binding](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/he-xin-gai-nian/binding.md)
    * [工作模式](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/gong-zuo-mo-shi/README.md)
      * [基础代码](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/gong-zuo-mo-shi/simple-mo-shi.md)
      * [Simple工作模式](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/gong-zuo-mo-shi/simple-gong-zuo-mo-shi.md)
      * [Work工作模式](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/gong-zuo-mo-shi/work-gong-zuo-mo-shi.md)
      * [Publish/Subscribe发布订阅模式](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/gong-zuo-mo-shi/publishsubscribe-fa-bu-ding-yue-mo-shi.md)
      * [Routing工作模式](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/gong-zuo-mo-shi/routing-gong-zuo-mo-shi.md)
      * [Topic工作模式](guan-jian-xi-tong/xiao-xi-dui-lie/xiao-xi-dui-lie/gong-zuo-mo-shi/topic-gong-zuo-mo-shi.md)
* [链路追踪](guan-jian-xi-tong/lian-lu-zhui-zong/README.md)
  * [开源方案](guan-jian-xi-tong/lian-lu-zhui-zong/kai-yuan-fang-an.md)
* [服务治理平台](guan-jian-xi-tong/fu-wu-zhi-li-ping-tai/README.md)
  * [系列教程](guan-jian-xi-tong/fu-wu-zhi-li-ping-tai/xi-lie-jiao-cheng.md)
* [缓存系统](guan-jian-xi-tong/huan-cun-xi-tong/README.md)
  * [缓存系统稳定性](guan-jian-xi-tong/huan-cun-xi-tong/huan-cun-xi-tong-wen-ding-xing.md)
  * [缓存数据一致性](guan-jian-xi-tong/huan-cun-xi-tong/huan-cun-shu-ju-yi-zhi-xing.md)
  * [缓存系统的取舍](guan-jian-xi-tong/huan-cun-xi-tong/huan-cun-xi-tong-de-qu-she.md)
* [个性化推荐](guan-jian-xi-tong/ge-xing-hua-tui-jian.md)
* [智能搜索引擎](guan-jian-xi-tong/zhi-neng-sou-suo-yin-qing.md)

## 微服务

* [微服务架构](wei-fu-wu/wei-fu-wu-jia-gou.md)
* [服务发布和引用](wei-fu-wu/fu-wu-fa-bu-he-yin-yong.md)
* [注册和发现服务](wei-fu-wu/zhu-ce-he-fa-xian-fu-wu.md)
* [RPC远程服务调用](wei-fu-wu/rpc-yuan-cheng-fu-wu-tiao-yong/README.md)
  * [GPRC](wei-fu-wu/rpc-yuan-cheng-fu-wu-tiao-yong/gprc.md)
  * [zRPC](wei-fu-wu/rpc-yuan-cheng-fu-wu-tiao-yong/zrpc.md)
* [微服务监控](wei-fu-wu/wei-fu-wu-jian-kong.md)
* [微服务追踪](wei-fu-wu/wei-fu-wu-zhui-zong.md)
* [微服务治理](wei-fu-wu/wei-fu-wu-zhi-li.md)
* [推荐文章](wei-fu-wu/tui-jian-wen-zhang.md)

## 网关

* [什么是网关](wang-guan/shen-me-shi-wang-guan.md)
* [代表产品](wang-guan/dai-biao-chan-pin/README.md)
  * [Easegress](wang-guan/dai-biao-chan-pin/easegress.md)

## serverless <a id="serverless-1"></a>

* [serverless](serverless-1/serverless.md)
* [Serverless 的价值](serverless-1/serverless-de-jia-zhi.md)
* [Serverless 架构模式](serverless-1/serverless-jia-gou-mo-shi.md)
* [Serverless 技术选型](serverless-1/serverless-ji-shu-xuan-xing.md)
* [技术文章](serverless-1/ji-shu-wen-zhang.md)

## service

* [RPC](service/rpc.md)
* [Dubbo](service/dubbo/README.md)
  * [Dubbo-go](service/dubbo/dubbo-go.md)

## 安全

* [TLS](an-quan/tls.md)

## 负载均衡

* [http/https](fu-zai-jun-heng/http-https.md)
* [负载均衡](fu-zai-jun-heng/fu-zai-jun-heng.md)
* [正向代理](fu-zai-jun-heng/zheng-xiang-dai-li.md)
* [反向代理](fu-zai-jun-heng/fan-xiang-dai-li.md)
* [Nginx](fu-zai-jun-heng/nginx.md)
* [HAProxy](fu-zai-jun-heng/haproxy.md)
* [HAProxy与Nginx对比](fu-zai-jun-heng/haproxy-yu-nginx-dui-bi.md)
* [Traefik](fu-zai-jun-heng/traefik.md)

## DB

* [数据库中间件](db/shu-ju-ku-zhong-jian-jian.md)
* [分布式数据库](db/fen-bu-shi-shu-ju-ku.md)
* [cockroach](db/cockroach.md)

## 缓存

* [Untitled](huan-cun/untitled.md)

## CI/CD

* [CI/CD](ci-cd/ci-cd.md)
* [搭建Gitlab服务器](ci-cd/da-jian-gitlab-fu-wu-qi.md)
* [Gitlab CI](ci-cd/gitlab-ci.md)
* [Github Actions](ci-cd/github-actions.md)
* [一次push到多个仓库](ci-cd/yi-ci-push-dao-duo-ge-cang-ku.md)

## 常见概念汇总

* [CAP理论](chang-jian-gai-nian-hui-zong/cap-li-lun.md)
* [Base理论](chang-jian-gai-nian-hui-zong/base-li-lun.md)
* [TCP长连接与短连接](chang-jian-gai-nian-hui-zong/tcp-chang-lian-jie-yu-duan-lian-jie.md)
* [长轮询与短轮询](chang-jian-gai-nian-hui-zong/chang-lun-xun-yu-duan-lun-xun.md)
* [谈谈HTTP协议中的短轮询、长轮询、长连接和短连接](chang-jian-gai-nian-hui-zong/tan-tan-http-xie-yi-zhong-de-duan-lun-xun-chang-lun-xun-chang-lian-jie-he-duan-lian-jie.md)
* [Websocket](chang-jian-gai-nian-hui-zong/websocket.md)
* [中心化与去中心化](chang-jian-gai-nian-hui-zong/zhong-xin-hua-yu-qu-zhong-xin-hua.md)
* [QPS、TPS、PV、UV、IP、GMV、RPS](chang-jian-gai-nian-hui-zong/qps-tps-pv-uv-ip-gmv-rps.md)
* [MMAP](chang-jian-gai-nian-hui-zong/mmap.md)

---

* [参考文献](can-kao-wen-xian.md)
* [技术文章](ji-shu-wen-zhang.md)

