# tcp server

## 1.目标
用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出

## 2.功能实现说明

实现了基本的tcp server，保存会话信息，client可以输入指定格式的字符串，服务端接收后，并返回执行结果。

### 实现

- 可以多个client同时连接
- 每一个client发送消息后，可以立即发送下一条消息
- client可发送指令，server的logic模块模拟了真实了logic处理。例如执行server端上的某个command

### 未实现

- tcp读写超时处理
- 单独清除某一个会话
- 退出main时，给所有会话信息发送bye消息


## 3.使用

### Server端

```shell
cd cmd/comet/
go build
./comet
```

### Client端
输入内容格式为`act content`

例如`auth username`, `exec ps`, `exec date`

```shell
telnet 127.0.0.1 10000
# 连接成功后，可输入命令
# auth 用户，正确的用户名为cool，3秒后返回
auth username
auth cool
exec date
exec ps
```
