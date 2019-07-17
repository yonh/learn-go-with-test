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

	// 我们将断言重构为函数。这减少了重复，提高了测试的可读性。在 Go 中，你可以在其他函数中声明函数并将它们分配给变量。你可以像普通函数一样调用它们。
	// 我们需要传入 t *testing.T，这样我们就可以在需要的时候令测试代码失败。
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。
		// 通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
		// 这将帮助其他开发人员更容易地跟踪问题。如果你仍然不理解，请注释掉它，使测试失败并观察测试输出。
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Carrie")
		want := "Hello, Carrie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})


}