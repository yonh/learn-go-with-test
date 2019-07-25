package main

import (
	"bytes"
	"testing"
)

func TestGountdown(t *testing.T) {

	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	actual := buffer.String()
	// 反引号语法是创建 string 的另一种方式，但是允许你放置东西例如放到新的一行，对我们的测试来说是完美的。
	expect := `3
2
1
Go!`

	if expect != actual {
		t.Errorf("expect '%s', actual '%s'", expect, actual)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}


