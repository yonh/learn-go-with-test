package main

import (
	"reflect"
	"testing"
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
