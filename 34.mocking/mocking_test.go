package main

import (
	"bytes"
	"reflect"
	"testing"
)

// 测试执行步骤值
const write = "write"
const sleep = "sleep"

// 监视器（spies）是一种 mock，它可以记录依赖关系是怎样被使用的。它们可以记录被传入来的参数，多少次等等。
// 在我们的例子中，我们跟踪记录了 Sleep() 被调用了多少次，这样我们就可以在测试中检查它。
// CountdownOperationsSpy 同时实现了 io.writer 和 Sleeper，使得可以把每一次 write 和 sleep 调用记录到 slice。
type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

// 由于在返回值中定义了变量，所以，在函数退出时，可以不用显示的在return后边指定函数返回值
// 疑问：这里 实现了 io.Write 接口，但是定义了的 (n int, err error) 返回值并没有处理过, 会不会有什么问题？
// 答: 这里 n=0, 貌似不会对这里产生什么影响
func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func TestGountdown(t *testing.T) {

	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &CountdownOperationsSpy{})

		actual := buffer.String()
		// 反引号语法是创建 string 的另一种方式，但是允许你放置东西例如放到新的一行，对我们的测试来说是完美的。
		expect := `3
2
1
Go!`

		if expect != actual {
			t.Errorf("expect '%s', actual '%s'", expect, actual)
		}
	})


	// 还有一个重要的特性，我们还没有测试过。
	// Countdown 应该在第一个打印之前 sleep，然后是直到最后一个前的每一个，例如：
	// Sleep
	// Print N
	// Sleep
	// Print N-1
	// Sleep
	// etc
	// 我们最新的修改只断言它已经 sleep 了 4 次，但是那些 sleeps 可能没按顺序发生。
	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		expect := []string {
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expect, spySleepPrinter.Calls) {
			t.Errorf("expect calls %v, actual calls %v", expect, spySleepPrinter.Calls)
		}
	})
}

