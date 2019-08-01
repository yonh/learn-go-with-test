package main

import (
	"encoding/json"
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
	//router *http.ServeMux
	// 我们更改了 PlayerServer 的第二个属性，删除了命名属性 router http.ServeMux，并用 http.Handler 替换了它；这被称为 嵌入。
	// 这意味着我们的 PlayerServer 现在已经有了 http.Handler 所有的方法，也就是 ServeHTTP。
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	leagueTable := []Player {
		{"Carrie", 18},
	}
	json.NewEncoder(w).Encode(leagueTable)

	w.WriteHeader(http.StatusOK)
}
func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player:= r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

//func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//
//	// Go 有一个内置的路由机制叫做 ServeMux（request multiplexer，多路请求复用器），它允许你将 http.Handler 附加到特定的请求路径。
//	router := http.NewServeMux()
//
//	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//	}))
//
//	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		player:= r.URL.Path[len("/players/"):]
//
//		switch r.Method {
//		case http.MethodPost:
//			p.processWin(w, player)
//		case http.MethodGet:
//			p.showScore(w, player)
//		}
//	}))
//
//	router.ServeHTTP(w, r)
//}

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
	server := &PlayerServer{NewInMemoryPlayerStore(),&http.ServeMux{}}

	// ListenAndServe 会在 Handler 上监听一个端口。如果端口已被占用，它会返回一个 error，所以我们在一个 if 语句中捕获出错的场景并记录下来。
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
