package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld  = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWorld)
}

func main() {
	Countdown(os.Stdout)
}
