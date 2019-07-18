package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	// 我们将断言重构为函数。这减少了重复，提高了测试的可读性。在 Go 中，你可以在其他函数中声明函数并将它们分配给变量。你可以像普通函数一样调用它们。
	// 我们需要传入 t *testing.T，这样我们就可以在需要的时候令测试代码失败。
	assertCorrectMessage := func(t *testing.T, expect, actual int) {
		// t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。
		// 通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
		// 这将帮助其他开发人员更容易地跟踪问题。如果你仍然不理解，请注释掉它，使测试失败并观察测试输出。
		t.Helper()
		if expect != actual {
			t.Errorf("expected '%d' but actual '%d'", expect, actual)
		}
	}


	t.Run("collection of 5 numbers", func(t *testing.T) {
		// 我们会使用 切片类型，它可以接收不同大小的切片集合。语法上和数组非常相似，只是在声明的时候不指定长度：
		numbers := []int{1, 2, 3, 4, 5}

		expect := 15
		actual := Sum(numbers)

		assertCorrectMessage(t, expect, actual)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1,2,3}

		expect := 6
		actual := Sum(numbers)

		assertCorrectMessage(t, expect, actual)
	})

	// 质疑测试的价值是非常重要的。测试并不是越多越好，而是尽可能的使你的代码更加健壮。太多的测试会增加维护成本，因为 维护每个测试都是需要成本的。
	// 在本例中，针对该函数写两个测试其实是多余的，因为切片尺寸并不影响函数的运行。
	// Go 有内置的计算测试 覆盖率的工具，它能帮助你发现没有被测试过的区域。我们不需要追求 100% 的测试覆盖率，它只是一个供你获取测试覆盖率的方式。只要你严格遵循 TDD 规范，那你的测试覆盖率就会很接近 100%。
	// go test -cover
}
