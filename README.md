# go-demo

## 环境

- [IDEA15搭建Go开发环境](http://studygolang.com/articles/6251)


## Go Web框架

[go有哪些快速开发的web框架？](https://www.zhihu.com/question/27370112)文章告诉我们目前在Github上比较火的web框架有[martini](https://github.com/go-martini/martini),[beego](https://github.com/astaxie/beego),[revel](https://github.com/revel/revel),[gin](https://github.com/gin-gonic/gin)等.


[超全的Go Http路由框架性能比较](http://colobu.com/2016/03/23/Go-HTTP-request-router-and-web-framework-benchmark/)的在路由处理性能方面的测试报告说明[gin](https://github.com/gin-gonic/gin)和[iris](https://github.com/kataras/iris)在方面性能非常好,[beego](https://github.com/astaxie/beego)相对于[revel](https://github.com/revel/revel)和[martini](https://github.com/go-martini/martini)性能较好.


[iris 真的是最快的Golang 路由框架吗?](http://colobu.com/2016/04/01/Is-iris-the-fastest-golang-router-library/)支出了[超全的Go Http路由框架性能比较](http://colobu.com/2016/03/23/Go-HTTP-request-router-and-web-framework-benchmark/)中测试方法的不足已经性能测试数据的一些不准确性.

[谁是最快的Go Web框架](http://colobu.com/2016/04/06/the-fastest-golang-web-framework/)说明[fasthttp](https://github.com/valyala/fasthttp)的性能不错,并且仍然说明基于[fasthttp](https://github.com/valyala/fasthttp)的[iris](https://github.com/kataras/iris)性能不错.

如果想要知道更多关于web frameworks的性能对比请查看(Web Framework Benchmarks)[https://www.techempower.com/benchmarks/#section=data-r12&hw=peak&test=query]


总结: [gin](https://github.com/gin-gonic/gin)和[iris](https://github.com/kataras/iris)的性能都非常好,但是[iris](https://github.com/kataras/iris)是基于[fasthttp](https://github.com/valyala/fasthttp)的,迁移项目比较麻烦,[beego](https://github.com/astaxie/beego)和[revel](https://github.com/revel/revel)在国内比较普及.

