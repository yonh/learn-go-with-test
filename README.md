##### 01-04 hello world
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/hello-world
##### 05 integers
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/integers
##### 06 for
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/iteration
##### 07-10 arrays and slices
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/arrays-and-slices
##### 11-13 structs,methods and interfaces
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/structs-methods-and-interfaces
##### 14-20 pointers and errors
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/pointers-and-errors
##### 21-27 map
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/maps
##### 28-29 dependency injection
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/dependency-injection
##### 30 mocking
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/mocking  
* 没有对代码中重要的区域进行 `mock` 将会导致难以测试。在我们的例子中，我们不能测试我们的代码在每个打印之间暂停，但是还有无数其他的例子。
    * 调用一个`可能失败的服务`？
    * 想要在一个`特定的状态`测试您的系统？
    * 在不使用 mocking 的情况下测试这些场景是非常困难的。
* 如果没有 `mock`，你可能需要`设置数据库`和`其他第三方的东西`来测试简单的业务规则。你可能会`进行缓慢的测试`，从而导致 `缓慢的反馈循环`。
* 当不得不启用一个`数据库`或者 `webservice` 去测试某个功能时，由于这种服务的不可靠性，你将会得到的是一个 `脆弱的测试`。

##### 35-39 concurrency
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/concurrency  

在这一章节，我们对`CheckWebsites`进行了重构，使用`goroutine`使得程序的检查结果部分得以并行运行，这大大加速了程序的执行  
同时我们也要了解在使用goroutine并行写入的时候会导致写入竞争，我们通过`channels`来组织和控制不同进程之间的交流，使我们能够避免 race condition（竞争条件） 的问题。

##### 40-43 selete
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/select

select  
可帮助你同时在多个 channel 上等待。  
有时你想在你的某个「案例」中使用 time.After 来防止你的系统被永久阻塞。  
httptest  
一种方便地创建测试服务器的方法，这样你就可以进行可靠且可控的测试。  
使用和 net/http 相同的接口作为「真实的」服务器会和真实环境保持一致，并且只需更少的学习。

##### 44-54 reflection
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/reflection

介绍了 reflect 包中的一些概念。  
使用递归遍历任意数据结构。  
这只是 reflection 的一个小方面。Go 博客上有[一篇精彩的文章](https://blog.golang.org/laws-of-reflection)介绍了更多细节。  
现在你已经了解了反射，请尽量避免使用它。

##### 55-59 http server 
https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/http-server

**http.Handler**  
通过实现这个接口来创建 web 服务器  
用 http.HandlerFunc 把普通函数转化为 `http.Handler`  
把 httptest.NewRecorder 作为一个 `ResponseWriter` 传进去，这样让你可以监视 `handler` 发送了什么响应
使用 `http.NewRequest` 构建对服务器的请求

#####  
https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/json

Go 有一个内置的路由机制叫做 [ServeMux](https://golang.org/pkg/net/http/#ServeMux)（`request multiplexer`，多路请求复用器），它允许你将 http.Handler 附加到特定的请求路径。


##### 
https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/io
##### 
https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/command-line


#### 命令
```bash
# 测试覆盖率
go test -cover
# 基准测试
go test -bench="." # windows
go test -bench=.
```
