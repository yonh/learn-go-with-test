package main

import (
	"math"
)

// 假设我们需要编程计算一个给定高和宽的长方形的周长。我们可以写一个函数如下：

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

// 声明方法(method)和函数(function)比较类似，唯一不同的是方法 func 后面接的是方法调用者类型
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
// Circle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
// string 没有这种方法，所以它不满足这个接口
// 在 Go 语言中 interface resolution 是隐式的。如果传入的类型匹配接口需要的，则编译正确。
type Shape interface {
	Area() float64
}
