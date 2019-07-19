package main

import (
	"testing"
)

//func TestArea(t *testing.T) {
//
//	CheckArea := func(t *testing.T, shape Shape, expect float64) {
//		t.Helper()
//
//		actual := shape.Area()
//
//		if expect!=actual {
//			t.Errorf("got %.2f want %.2f", expect, actual)
//		}
//	}
//
//	t.Run("",func(t * testing.T) {
//		// 计算长方形面积
//		rect := Rectangle{3.0, 4.0}
//		expect := 12.0
//		CheckArea(t, rect, expect)
//	})
//
//	t.Run("", func(t *testing.T) {
//		circle := Circle{10}
//		expect := 314.1592653589793
//
//		CheckArea(t, circle, expect)
//	})
//
//}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape  Shape
		expect float64
	}{
		// 同时我们也可以用 MyStruct{key1: val1, key2: val2} 的方式定义结构体
		{name: "Rectangle", shape: Rectangle{3.0, 4.0}, expect: 12.0},
		{name: "Circle", shape: Circle{10.0}, expect: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, expect: 36.0},
		{"Rectangle",Rectangle{3.0, 4.0}, 12.0},
		{"Circle",Circle{10.0}, 314.1592653589793},
		{"Triangle",Triangle{12.0, 6.0}, 36.0},
	}

	for _, tt := range areaTests {
		// 如果包含几十个测试用例的系统里出现类似的错误我们怎么能快速从中找到错误的位置呢？
		// 我们可以使用 t.Run() 实现错误位置的输出，让人可以一眼知道错误发生的具体位置
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Area()
			if tt.expect != actual {
				t.Errorf("%#v expect %.2f actual %.2f",tt.shape, tt.expect, actual)
			}
		})


	}
}
