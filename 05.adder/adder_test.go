package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

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

	actual := Add(2,2)
	expect := 4

	assertCorrectMessage(t, actual, expect)
}

// 作为包的测试套件的一部分，示例会被编译（并可选择性地执行）。
// 与典型的测试一样，示例是存在于一个包的 _test.go 文件中的函数。
// 请注意，如果删除注释 「//Output: 6」，示例函数将不会执行。虽然函数会被编译，但是它不会执行。
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}