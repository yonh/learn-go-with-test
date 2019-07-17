package main

import (
	 "testing"
)
/**
编写测试和写函数很类似，其中有一些规则
它需要在一个名为 xxx_test.go 的文件中编写
测试函数的命名必须从单词 Test 开始
测试函数只接受一个参数 t *testing.T

现在这些信息足以让我们明白，类型为 *testing.T 的变量 t 是你在测试框架中的 "hook"（钩子），所以你可以在想要失败时执行 t.Fail() 之类的操作。

 */
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}