学习笔记

## 目录结构

cmd目录内不要直接放main.go,最好是再加一层项目的名称的目录

可重用的目录,放入到pkg目录,按照功能划分

内部的则放到internal目录

### Kit project layout

每个公司都应当为不同的微服务建立一个统一的 kit 工具包项目(基础库/框架) 和 app 项目。

特点:
- 统一
- 标准库方式布局
- 高度抽象
- 支持插件

### Service application project layout

/api
把api相关的放到其中

/configs
推荐使用yaml文件

/test

一个公司应该有一个service tree,把所有服务管理起来,每个微服务都有一个唯一命名,推荐“业务.服务.子服务”,如“account.service.vip”

将管理员权限和普通对外业务的目录分开,以免出现权限问题.
![72587bffc64e42f1d534a12214d234e5.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p724)


推荐: 三层架构+贫血模型
![69f81f87bdbc027d4bb214f4248aeb10.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p725)

service层,实现了 api 定义的服务层,做deep copy.关注了GRPC类的实现

data层做数据库的操作

推荐使用wire工具来完成依赖注入

## API设计

pb文件放在api目录中,并制作api仓库,统一管理
![55bb4b5466c261cd434bca6771612834.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p726)

命令特别注意
![d68a82ef6747467965e7f2b35c15abc5.png](evernotecid://C6256292-0189-4229-A8DF-6DB4F0728096/appyinxiangcom/14034229/ENResource/p727)

输入,输出都定义特定的对象,方便扩展

## 错误处理

最好是翻译多次  service error -> gprc error -> service error

## References
[Facebook ent框架](https://github.com/facebook/ent)
[ent文档](https://entgo.io/)

[Google API Desgin Guide](https://www.bookstack.cn/read/API-design-guide/API-design-guide-04-%E6%A0%87%E5%87%86%E6%96%B9%E6%B3%95.md)

字段掩码