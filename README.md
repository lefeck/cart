# 微服务熔断

## hystrix-go的作用
* 阻止故障的连锁反应
* 快速失败并迅速恢复
* 回退并优雅降级
* 提供近实时的监控和告警

什么是服务雪崩效应

## hystrix-go的观测面板安装

hystrix并没有自带一个仪表盘，无法直观的查看接口的健康状况。所以，我们采用GitHub的一个开源实现hystrix-dashboard。
```
docker pull hystrix-dashboard
docker run --name hystrix-dashboard -d -p 9002:9002 mlabouardy/hystrix-dashboard:latest
```
访问http://ip:9002


# 限流作用
* 限制流量， 在服务端生效
* 保护后端服务
* 与熔断互补

## uber/limit漏桶算法原理
漏桶算法其实非常形象，如下图所示可以理解为一个漏水的桶，当有突发流量来临的时候，会先到桶里面，桶下有一个洞，可以以固定的速率向外流水，如果水从桶中外溢了出来，那么这个请求就会被拒绝掉。具体的表现就会向下图右侧的图表一样，突发流量就被整形成了一个平滑的流量。

![漏桶算法](https://mohuishou-blog-sz.oss-cn-shenzhen.aliyuncs.com/image/1617718978961-0a125409-6fbb-4ca9-b335-5bef61cd44a8.png)


# 负载均衡作用
* 提高系统横向扩展性
* 支持： http，https，tpc，udp
* 主要算法：轮训算法和随机算法，默认是随机算法。
  
  ![LOADBALANCE](https://github.com/asveg/picture/blob/master/loadbalance.png)
  
将来自客户端的请求均匀的分配到后端多台服务器上，实现负载均衡。


# 微服务api网关

## api网关的总体架构

* 第一层micro api网关层
* 第二层聚合业务层BFF层
  提高可扩展性
* 第三层基础服务层

## API网关路径说明
* 通过网关请求/greeter/say/hello， 这个路径，网关会将请求转发到go.micro.api.greeter服务的say.hello方法处理。
* go.micro.api 是网关的默认服务名的前缀。
* 路径中/cartApi/cartAPi/findAll 可以写成/cartAPi/findAll

![API](https://github.com/asveg/picture/blob/master/api.png)

# 微服务cart领域代码实现
* cart需求分析和项目目录创建
* cart代码开发
* cart 加入熔断，限流， 负载均衡。
* api网关使用及最终效果展示。

# Cart Service

This is the Cart service

Generated with

```
micro new --namespace=go.micro --type=service cart
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.cart
- Type: service
- Alias: cart

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./cart-service
```

Build a docker image
```
make docker
```
