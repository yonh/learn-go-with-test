package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(url1 string, url2 string) (winner string, err error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableRacer(url1 string, url2 string, timeout time.Duration) (winner string, err error) {
	// select 则允许你同时在 多个 channel 等待。第一个发送值的 channel「胜出」，case 中的代码会被执行。
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	// 使用 select 时，time.After 会在你定义的时间过后发送一个信号给 channel 并返回一个 chan 类型（就像 ping 那样）。
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping (url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
