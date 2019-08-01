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
// [new] 我们可以通过使用新的 RecordWin 方法扩展 StubPlayerStore 然后监视它的调用来实现这一点。
func TestGETPlayers(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Bob":    10,
		},
		nil,
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

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
	}

	server := &PlayerServer{&store}

	t.Run("it returns win on POST", func(t *testing.T) {
		player:="Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, http.StatusAccepted, response.Code)

		if len(store.winCalls)!=1 {
			t.Errorf("expect %d calls, actual calls %d", 1, len(store.winCalls))
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner, expect '%s' actual '%s'", player, store.winCalls[0])
		}
	})
}

// 我们正在尝试集成两个组件：InMemoryPlayerStore 和 PlayerServer。
// 然后我们发起 3 个请求，为玩家记录 3 次获胜。我们并不太关心测试中的返回状态码，因为和集成得好不好无关。
// 我们真正关心的是下一个响应（所以我们用变量存储 response），因为我们要尝试并获得 player 的得分。

// 构建并运行代码，然后使用 curl 来测试它。
// 运行几次这条命令 curl -X POST http://localhost:5000/players/Pepper，你换成别的玩家名称也可以
// 用 curl http://localhost:5000/players/Pepper 获取玩家得分
func TestRecordingWinsAndRetrivingThem(t *testing.T) {
	store := InMemoryPlayerStore{map[string]int{}}
	server := PlayerServer{&store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, http.StatusOK, response.Code)

	assertResponseBody(t, response.Body.String(), "3")
}

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server:= &PlayerServer{&store}

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, http.StatusOK, response.Code)
	})
}

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]

	return score
}
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func newGetScoreRequest(name string) *http.Request {
	// 我们用 http.NewRequest 来创建一个请求。第一个参数是请求方法，第二个是请求路径。nil 是请求实体，不过在这个场景中不用发送请求实体。
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)

	return req
}
func newPostWinRequest(name string) *http.Request {
	req, _:= http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)

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
