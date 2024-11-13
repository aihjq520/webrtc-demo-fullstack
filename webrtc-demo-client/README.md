# WebRTC Demo Web端

## 1. 前言

该项目实现的是简单的p2p局域网内点对点视频传输，旨在搞清楚webrtc信息交换的过程。后续会考虑加入stun、turn服务器，支持不同网段适配传输。

以下是相关的介绍webrtc文章:

1. https://juejin.cn/post/7323087699479838730

2. https://juejin.cn/post/7267892210304565303

## 2.准备工作

确保已经启动WebRtc Demo Server端，复制运行地址到`.env.development`文件的`VITE_WS_BASEURL`变量,例如：

```SHELL
VITE_WS_BASEURL = '192.168.31.215:8080'
```

这是用于交换sdp和icecandidate的信令服务器。


## 3.启动

安装依赖

```
npm i 
```

运行

```
npm run dev
```

登录

使用内置的账号登录
```JS
18800000000/123456 // 用户1
18800000001/123456 // 用户2
```

## 4.功能说明

目前实现了简单web端点对点的屏幕共享功能，由于作何设备上没有摄像头，所以是暂时只支持屏幕流。

虽然功能简单，但是其中有很多可以深究的知识点，例如：

- websocket 如何进行鉴权？
- sdp 和 icecandidate 协议格式说明
- peerconnection的流和track具体用法


后续可以计划加入一下功能

- [ ] 房间创建
- [ ] 一对一，一对多 
- [ ] 弹幕
- [ ] 流加工


## 5.最后

欢迎各位提出宝贵的意见和讨论，如果有帮助到你非常开心~





