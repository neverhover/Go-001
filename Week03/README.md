学习笔记

## 作业核心要点

- main要能控制住所有的子服务的启动和停止。
- 子服务，既要能接受到上级context的结束，也要能recovery住自己的panic
- 子服务一旦工作完成，父级waitGroup计数减一
- main的正常停止服务,应当通知到下层，下层有时间能够处理完任务，在此期间，main应该等待workers完成任务。
- main需要做到，强制退出而不等待所有子服务释放，这里暂时不做假设服务shutdown需要timeout的情况，由调用者来决定强制退出时机。
- 任何一个子服务的异常退出，不能影响到其他子服务...

以下为正常退出，等待所有任务完成的打印
```shell
Main: all services started
Service httpServer-01: running...
Service httpServer-02: running...

^CMain: Try to exit by signal SIGINT
httpServer-01: get a finish event by parent context , should shutdown service next
httpServer-02: get a finish event by parent context , should shutdown service next
Service httpServer-01: shutdown start... need work more 7 seconds to finish some tasks
Service httpServer-02: shutdown start... need work more 1 seconds to finish some tasks

Service httpServer-02: shutdown done
Service httpServer-01: shutdown done
Main: work done

```

以下为强制退出，不等待子任务完成的打印信息
```shell
Main: all services started
Service httpServer-01: running...
Service httpServer-02: running...
^CMain: Try to exit by signal SIGINT
httpServer-01: get a finish event by parent context , should shutdown service next
httpServer-02: get a finish event by parent context , should shutdown service next
Service httpServer-02: shutdown start... need work more 4 seconds to finish some tasks
Service httpServer-01: shutdown start... need work more 6 seconds to finish some tasks

^CMain: Exit without waiting workers done

```

## error_example 说明

目录内包含了一个错误的例子，使用errgroup创建的context传递到其他service

当一个service因为异常返回err时，errgroup会发送cancel，这样导致其他所有的service都收到退出的消息，某些场景下，这并不是我们所想要的结果

因此，尽量避免使用errgroup返回的context做为子服务的context

以下为测试打印信息：
``` shell
Main: waiting all services work done
Main: all services started
Service httpServer-01: running...
Service httpServer-02: running...

Service httpServer-02: Mock panic got -> panic httpServer-02
httpServer-02 goroutine error, quit service
httpServer-01 work done
Service httpServer-01: Mock panic got -> panic httpServer-01

```