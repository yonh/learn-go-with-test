package main

import (
	"fmt"
	"log"
	"net/http"
)

func GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}

	if player == "Bob" {
		return "10"
	}

	return ""
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 当请求开始时，我们创建了一个路由，然后我们告诉它 x 路径使用 y handler。
	// 那么对于我们的新端点 /league 被请求时，我们用 http.HandlerFunc 和一个匿名函数来响应 w.WriteHeader(http.StatusOK) 使测试通过。
	// 对于 /players/ 路由我们只需剪贴代码并粘贴到另一个 http.HandlerFunc。
	// 最终，我们通过调用新路由的 ServeHTTP 方法处理到来的请求（注意到 ServeMux 也是一个 http.Handler 了吗？）


	// Go 有一个内置的路由机制叫做 ServeMux（request multiplexer，多路请求复用器），它允许你将 http.Handler 附加到特定的请求路径。
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		player:= r.URL.Path[len("/players/"):]

		switch r.Method {
		case http.MethodPost:
			p.processWin(w, player)
		case http.MethodGet:
			p.showScore(w, player)
		}
	}))

	router.ServeHTTP(w, r)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}


func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	// ListenAndServe 会在 Handler 上监听一个端口。如果端口已被占用，它会返回一个 error，所以我们在一个 if 语句中捕获出错的场景并记录下来。
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
