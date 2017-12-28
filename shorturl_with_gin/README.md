## My first web project demo of Go


## 用到的第三方代码
- gin
    - 据说适合做API，拿来做基本框架处理路由什么的了。

- beego
    - 感觉功能挺全，文档不错（中文，详细）。可以学习下各种功能代码的实现。


## Step by step
1. 第一天
- 目标：短网址服务
- 内容：借助beego的短网址生成算法和memory-cache，和gin的基本路由，实现了长网址缩短。
- 接口：
    - /ping
    - /api/v1/shorten
- 评价：
    - 基本没有写什么代码，主要是组合gin+beego，熟悉go的package机制。
    - 直接从beego的源码拷贝cache部分的实现，一开始只复制了一个cache.go，调用的实现用了memory的adapter，然而并没有memory-cache的具体实现代码，所以直接内存错误。随后看了下beego的cache源码结构，发现支持多种adapter，如redis，file，memory等。
    - 了解了下go的几种json包，然而易用性跟python相比实在是差远了（python就import一次，两种方法基本通吃……），go的几种json库，简直不知所云。。。