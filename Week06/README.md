## 原有问题

计数器.在固定时间周期内,统计是否超过阀值，未超过，则允许通过，超过则限制；但是在一个周期的结尾和开始都刚好有等于阀值的计数，则导致在这1-2s内，产生了2倍的流量进入系统，不符合预期。

## 优化方法

- 滑动窗口。相比计数实现，滑动窗口实现会更加平滑，能自动消除毛刺。
- 漏桶
- 令牌桶

## 个人理解

优化方法一就是使用滑动窗口：
1.依然是一个周期，放多少流量
2.将周期理解为一个窗口，窗口可以分割为多个更小的时间片，随着时间的推移，窗口会向右边滑动。
3.举例，一个接口一分钟限制调用1000次，1分钟就可以理解为一个窗口，可以把1分钟分割为10个单元格，每个单元格就是6秒。
4.单元格=bucket,窗口=window
5.因此窗口的计数=bucket_counter * N
6.失败率达到阀值，则将熔断状态变更


pass为每100ms采样window中成功请求的数量
rt为一个window中采样成功率最高的bucket平均响应时间。

因此使用滑动窗口：
- 意义，统计在指定时间周期（一个窗口）内的指标
- 本质，属于metrics统计，输出的指标统计，作为其他（如熔断器）指标值（输入值）
- 输入，一系列point
- 输出，某一个window的指标统计，包括，请求成功数量，成功率最高的平均响应时间

## 参考

[Hystrix Go实现](https://github.com/afex/hystrix-go)
[Hystrix指标窗口实现原理](https://www.jianshu.com/p/249e4f22fb84?from=singlemessage)

## 当前代码说明

使用了`kratos-v2`中`/metrics/rolling/iterator.go`的代码,


了解javaHystrix版本中关于滑动窗口的实现以及后续根据失败比例来做为熔断器的状态变更依据
在对比了[Hystrix Go中rolling](https://github.com/afex/hystrix-go/blob/fa1af6a1f4f56e0e50d427fe901cd604d8c6fb8a/hystrix/rolling/rolling.go)
以及`kratos-v2`中的实现，认为`kratos`的计数器更为纯粹（抽象的也更好）

- `bucket`数组使用了ring buffer的方式来实现。
- 窗口`window`本身并不于任何的指标耦合
- 所有的指标均在`rolling_counter`中，以接口的方式来进行实现，也方便后续扩展更多的指标
- 核心思路为在创建后`rolling_counter`后，每次有新的val进入，则计算该val落入的bucket，计算时间跨度（timespan)，进而可以得到该时刻的所有指标
- 任何时刻计算出来的指标可以做为后续熔断器的熔断依据

