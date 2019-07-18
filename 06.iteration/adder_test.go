package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	// 我们将断言重构为函数。这减少了重复，提高了测试的可读性。在 Go 中，你可以在其他函数中声明函数并将它们分配给变量。你可以像普通函数一样调用它们。
	// 我们需要传入 t *testing.T，这样我们就可以在需要的时候令测试代码失败。
	assertCorrectMessage := func(t *testing.T, expect, actual string) {
		// t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。
		// 通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
		// 这将帮助其他开发人员更容易地跟踪问题。如果你仍然不理解，请注释掉它，使测试失败并观察测试输出。
		t.Helper()
		if expect != actual {
			t.Errorf("expected '%s' but actual '%s'", expect, actual)
		}
	}

	actual := Repeat("a", 5)
	expect := "aaaaa"

	assertCorrectMessage(t, actual, expect)
}

func ExampleRepeat() {
	s := Repeat("*", 10)
	fmt.Println(s)
	// Output: **********
}

// 在 Go 中编写基准测试（benchmarks）是该语言的另一个一级特性，它与编写测试非常相似。
// testing.B 可使你访问隐性命名（cryptically named）b.N。
// 基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
// 代码运行的次数应该不影响你，框架将决定什么是「好」的值，以便让你获得一些得体的结果。
// 用 go test -bench=. 来运行基准测试。 (如果在 Windows Powershell 环境下使用 go test -bench=".")
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
