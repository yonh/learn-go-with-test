package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 你被要求创建一个 Web 服务器，用户可以在其中跟踪玩家赢了多少场游戏。
// GET /players/{name} 应该返回一个表示获胜总数的数字
// POST /players/{name} 应该为玩家赢得游戏记录一次得分，并随着每次 POST 递增
func TestGETPlayers(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Bob":    10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, http.StatusOK, response.Code)
		assertResponseBody(t, "20", response.Body.String())
	})

	t.Run("returns Bob's score", func(t *testing.T) {
		request := newGetScoreRequest("Bob")
		// net/http/httptest 自带一个名为 ResponseRecorder 的监听器，所以我们可以用这个。它有很多有用的方法可以检查应答被写入了什么。
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, http.StatusOK, response.Code)
		assertResponseBody(t, "10", response.Body.String())
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Hello")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, http.StatusNotFound, response.Code)
	})
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}

func newGetScoreRequest(name string) *http.Request {
	// 我们用 http.NewRequest 来创建一个请求。第一个参数是请求方法，第二个是请求路径。nil 是请求实体，不过在这个场景中不用发送请求实体。
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return req
}

func assertStatus(t *testing.T, expect, actual int) {
	t.Helper()

	if expect != actual {
		t.Errorf("expect status '%d', actual status '%d'", expect, actual)
	}
}

func assertResponseBody(t *testing.T, expect, actual string) {
	t.Helper()

	if expect != actual {
		t.Errorf("response body is wrong, expect '%s', actual '%s'", expect, actual)
	}
}
