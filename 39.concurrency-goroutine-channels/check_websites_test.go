package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {

	return url != "http://have.fun"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://baidu.com",
		"http://google.com",
		"http://have.fun",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	expectLen := len(websites)
	actualLen := len(actualResults)

	if expectLen != actualLen {
		t.Fatalf("expect %v, actual %v", expectLen, actualLen)
	}

	expectResults := map[string]bool{
		"http://baidu.com": true,
		"http://google.com": true,
		"http://have.fun": false,
	}

	if !reflect.DeepEqual(expectResults, actualResults) {
		t.Fatalf("expect %v, actual %v", expectResults, actualResults)
	}
}


//这里基准测试使用一百个网址的 slice 对 CheckWebsites 进行测试，
// 并使用 WebsiteChecker 的伪造实现 slowStubWebsiteChecker 故意放慢速度来模拟请求耗时。
// 它使用 time.Sleep 明确等待 20 毫秒，然后返回 true。
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i:=0; i<len(urls); i++ {
		urls[i] = "test url"
	}

	for i := 0; i<b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	// slowStubWebsiteChecker 使用 time.Sleep 明确等待 10 毫秒，然后返回 true。
	time.Sleep(20 * time.Millisecond)

	return true
}
