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

##### 
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/concurrency
##### 
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/select
##### 
https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/reflection
##### 
https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/app-intro
##### 
https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/http-server
##### 
https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/json
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