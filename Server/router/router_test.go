package router

import (
	"dungeons_helper/internal/account"
	"dungeons_helper/internal/alignment"
	"dungeons_helper/internal/character"
	"dungeons_helper/internal/class"
	"dungeons_helper/internal/lobby"
	"dungeons_helper/internal/races"
	"dungeons_helper/internal/skills"
	"dungeons_helper/internal/stats"
	"dungeons_helper/internal/subraces"
	"dungeons_helper/internal/websocket"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitRouter(t *testing.T) {
	accountHandler := &account.Handler{}
	alignmentHandler := &alignment.Handler{}
	classHandler := &class.Handler{}
	racesHandler := &races.Handler{}
	subracesHandler := &subraces.Handler{}
	statsHandler := &stats.Handler{}
	skillHandler := &skills.Handler{}
	characterHandler := &character.Handler{}
	lobbyHandler := &lobby.Handler{}
	wsHandler := &websocket.Handler{}

	r := InitRouter(
		AccountRouter(accountHandler),
		AlignmentRouter(alignmentHandler),
		ClassRouter(classHandler),
		RacesRouter(racesHandler),
		SubracesRouter(subracesHandler),
		StatsRouter(statsHandler),
		SkillsRouter(skillHandler),
		CharacterRouter(characterHandler),
		LobbyRouter(lobbyHandler),
		WebsocketRouter(wsHandler),
	)

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
