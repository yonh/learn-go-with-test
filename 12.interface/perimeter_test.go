package main

import (
	"testing"
)

func TestArea(t *testing.T) {

	CheckArea := func(t *testing.T, shape Shape, expect float64) {
		t.Helper()

		actual := shape.Area()

		if expect!=actual {
			t.Errorf("got %.2f want %.2f", expect, actual)
		}
	}

	t.Run("",func(t * testing.T) {
		// 计算长方形面积
		rect := Rectangle{3.0, 4.0}
		expect := 12.0
		CheckArea(t, rect, expect)
	})

	t.Run("", func(t *testing.T) {
		circle := Circle{10}
		expect := 314.1592653589793

		CheckArea(t, circle, expect)
	})

}
