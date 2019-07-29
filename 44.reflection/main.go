package main

import "reflect"

// 这段代码 非常不安全，也非常幼稚，但请记住，当我们处于「红色」状态（测试失败）时，我们的目标是编写尽可能少的代码。
// 然后我们编写更多的测试来解决我们的问题。
func walk(x interface{}, fn func(input string)) {

	// 反射包有一个函数 ValueOf，该函数值返回一个给定变量的 Value。这为我们提供了检查值的方法，包括我们在下一行中使用的字段。
	// 然后我们对传入的值做了一些非常乐观的假设：
	// 我们只看第一个也是唯一的字段，可能根本就没有字段会引起 panic
	// 然后我们调用 String()，它以字符串的形式返回底层值，但是我们知道，如果这个字段不是字符串，程序就会出错。
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}