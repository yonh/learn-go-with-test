package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// 你被要求编写一个叫做 WebsiteRacer 的函数，用来对比请求两个 URL 来「比赛」，并返回先响应的 URL。
// 如果两个 URL 在 10 秒内都未返回结果，那么应该返回一个 error。

func TestRacer(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	slowURL := slowServer.URL
	fastURL := fastServer.URL




	expect := fastURL
	actual := Racer(slowURL, fastURL)

	if expect != actual {
		t.Errorf("expect '%s', actual '%s'", expect, actual)
	}

	slowServer.Close()
	fastServer.Close()
}