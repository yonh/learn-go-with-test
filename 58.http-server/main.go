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
	player:= r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
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

type InMemoryPlayerStore struct {}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}
func (i *InMemoryPlayerStore) RecordWin(name string) {

}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	// ListenAndServe 会在 Handler 上监听一个端口。如果端口已被占用，它会返回一个 error，所以我们在一个 if 语句中捕获出错的场景并记录下来。
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
