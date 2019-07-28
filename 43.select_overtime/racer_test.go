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



	t.Run("test fast", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL


		expect := fastURL
		actual, _ := Racer(slowURL, fastURL)

		if expect != actual {
			t.Errorf("expect '%s', actual '%s'", expect, actual)
		}

	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server1 := makeDelayedServer(11 * time.Second)
		server2 := makeDelayedServer(12 * time.Second)
		defer server1.Close()
		defer server2.Close()

		_, err := Racer(server1.URL, server2.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
