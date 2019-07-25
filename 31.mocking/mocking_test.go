package main

import (
	"bytes"
	"testing"
)

func TestGountdown(t *testing.T) {

	buffer := &bytes.Buffer{}

	Countdown(buffer)

	actual := buffer.String()
	// 反引号语法是创建 string 的另一种方式，但是允许你放置东西例如放到新的一行，对我们的测试来说是完美的。
	expect := `3
2
1
Go!`

	if expect != actual {
		t.Errorf("expect '%s', actual '%s'", expect, actual)
	}

}
