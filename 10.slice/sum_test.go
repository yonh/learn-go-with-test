package main

import (
	"reflect"
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

	// 质疑测试的价值是非常重要的。测试并不是越多越好，而是尽可能的使你的代码更加健壮。太多的测试会增加维护成本，因为 维护每个测试都是需要成本的。
	// 在本例中，针对该函数写两个测试其实是多余的，因为切片尺寸并不影响函数的运行。
	// Go 有内置的计算测试 覆盖率的工具，它能帮助你发现没有被测试过的区域。我们不需要追求 100% 的测试覆盖率，它只是一个供你获取测试覆盖率的方式。只要你严格遵循 TDD 规范，那你的测试覆盖率就会很接近 100%。
	// go test -cover
}


// 这回我们需要一个 SumAll 函数，它接受多个切片，并返回由每个切片元素的总和组成的新切片。
// SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}
// 或
// SumAll([]int{1,1,1}) would return []int{3}

func TestSumAll(t *testing.T) {
	expect := []int{3,9}
	actual := SumAll([]int{1,2}, []int{0,9})

	// 在 Go 中不能对切片使用等号运算符。你可以写一个函数迭代每个元素来检查它们的值。但是一种比较简单的办法是使用 reflect.DeepEqual，它在判断两个变量是否相等时十分有用。
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expected '%v' but actual '%v'", expect, actual)
	}
}


// 接下来的工作是把 SumAll 变成 SumAllTails。它会把每个切片的尾部元素想加（尾部的意思就是出去第一个元素以外的其他元素）。
func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, expect, actual []int) {
		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("got %v want %v", expect, actual)
		}
	}



	t.Run("make the sums of some slices", func(t *testing.T) {
		expect := []int{2,9}
		actual := SumAllTails([]int{1,2}, []int{0,9})

		checkSums(t, expect, actual)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		expect := []int{0,9}
		actual := SumAllTails([]int{}, []int{3, 4, 5})

		checkSums(t, expect, actual)
	})
}