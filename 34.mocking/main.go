package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld  = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWorld)
}

// 我做了一个设计的决定，我们的 Countdown 函数将不会负责 sleep 的时间长度。
// 这至少简化了我们的代码，也就是说，我们函数的使用者可以根据喜好配置休眠的时长。
type Sleeper interface {
	Sleep()
}


type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
