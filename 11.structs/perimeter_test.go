package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	// 我们将断言重构为函数。这减少了重复，提高了测试的可读性。在 Go 中，你可以在其他函数中声明函数并将它们分配给变量。你可以像普通函数一样调用它们。
	// 我们需要传入 t *testing.T，这样我们就可以在需要的时候令测试代码失败。
	assertCorrectMessage := func(t *testing.T, expect, actual float64) {
		// t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。
		// 通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
		// 这将帮助其他开发人员更容易地跟踪问题。如果你仍然不理解，请注释掉它，使测试失败并观察测试输出。
		t.Helper()
		if expect != actual {
			t.Errorf("expected %.2f but actual %.2f", expect, actual)
		}
	}

	// 假设我们需要编程计算一个给定高和宽的长方形的周长。我们可以写一个函数如下：
	rectangle := Rectangle{10.0, 10.0}
	expect := 40.0
	actual := Perimeter(rectangle)

	assertCorrectMessage(t, expect, actual)
}

// 计算长方形面积
func TestArea(t *testing.T) {
	rect := Rectangle{3.0, 4.0}

	expect := 12.0
	actual := Area(rect)

	if expect != actual {
		t.Errorf("got %.2f want %.2f", actual, actual)
	}
}