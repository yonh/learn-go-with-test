package main

// 假设我们需要编程计算一个给定高和宽的长方形的周长。我们可以写一个函数如下：
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}


type Rectangle struct {
	Width float64
	Height float64
}
