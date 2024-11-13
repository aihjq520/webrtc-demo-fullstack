# WebRtc Demo Server端

## 1. 前言

用go语言框架编写，用来当作webrtc的信令服务器，并且提供了用户模块功能。

以下是相关的介绍webrtc文章:

1. https://juejin.cn/post/7323087699479838730

2. https://juejin.cn/post/7267892210304565303

## 2.准备工作

在`config/application.yml`配置好MySql数据库

```yml
server:
  port: 8080  // 启动端口
datasource:
  driverName: mysql
  host:  // 数据库host
  port: 3306
  database:   // 数据库名称
  username:      // 数据库账号
  password:  // 密码
  charset: utf8mb4
```

然后运行命令启动

```SHELL
go run routes.go main.go
```


如果看到如下命令则启动成功

```SHELL
GIN-debug] POST   /api/auth/register        --> webrtcdemo/api/controller.Register (5 handlers)
[GIN-debug] POST   /api/auth/login           --> webrtcdemo/api/controller.Login (5 handlers)
[GIN-debug] GET    /api/auth/info            --> webrtcdemo/api/controller.Info (6 handlers)
[GIN-debug] GET    /rtc                      --> webrtcdemo/api/controller.WebRtcHandler (5 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```

## 3.功能说明

这个项目目前主要是用来充当信令交换服务器，并且有用户模块，记录p2p双方通信。

后续可以考虑加入


- [ ] 房间创建
- [ ] 一对一，一对多 
- [ ] 弹幕
- [ ] 流加工


## 4.最后

欢迎各位提出宝贵的意见和讨论，如果有帮助到你非常开心~
