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
	"dungeons_helper/util"
	"net/http"

	"github.com/gorilla/mux"
)

type Option func(router *mux.Router)

func AccountRouter(accountHandler *account.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/auth/registration", accountHandler.CreateAccount).Methods("POST")
		r.HandleFunc("/auth/byEmail", accountHandler.Login).Methods("POST")
		r.HandleFunc("/auth/logout", accountHandler.Logout).Methods("POST")
		r.HandleFunc("/auth/google/login", accountHandler.LoginGoogle).Methods("POST")
		r.HandleFunc("/auth/restorePassword", accountHandler.RestorePassword).Methods("POST")
		r.HandleFunc("/account/change/nickname", accountHandler.UpdateNickname).Methods("PATCH")
		r.HandleFunc("/account/change/password", accountHandler.UpdatePassword).Methods("PATCH")
		r.HandleFunc("/account/change/avatar", accountHandler.UpdateAvatar).Methods("PATCH")
	}
}

func AlignmentRouter(alignmentHandler *alignment.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/alignments/getAlignments", alignmentHandler.GetAllAlignments).Methods("GET")
	}
}

func ClassRouter(classHandler *class.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/class/getClasses", classHandler.GetAllClasses).Methods("GET")
	}
}

func RacesRouter(racesHandler *races.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/race/getRaces", racesHandler.GetAllRaces).Methods("GET")
	}
}

func SubracesRouter(subracesHandler *subraces.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/subrace/getSubraces", subracesHandler.GetAllSubraces).Methods("POST")
	}
}

func StatsRouter(statsHandler *stats.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/stats/getStatsById", statsHandler.GetStatsById).Methods("POST")
	}
}

func SkillsRouter(skillHandler *skills.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/skills/getSkills", skillHandler.GetAllSkills).Methods("GET")
	}
}

func CharacterRouter(characterHandler *character.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/character/getAllCharactersByAccId", characterHandler.GetAllCharactersByAccId).Methods("GET")
		r.HandleFunc("/character/getCharacterById", characterHandler.GetCharacterById).Methods("POST")
		r.HandleFunc("/character/createCharacter", characterHandler.CreateCharacter).Methods("POST")
		r.HandleFunc("/character/setActiveCharacter", characterHandler.SetActiveCharacterById).Methods("POST")
	}
}

func LobbyRouter(lobbyHandler *lobby.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/lobby/getAllLobby", lobbyHandler.GetAllLobby).Methods("GET")
	}
}

func WebsocketRouter(wsHandler *websocket.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/lobby/join", wsHandler.JoinLobby)
		r.HandleFunc("/lobby/create", wsHandler.CreateLobby)
	}
}

func InitRouter(options ...Option) *mux.Router {
	r := mux.NewRouter()
	r.Use(util.LoggingMiddleware)

	for _, option := range options {
		option(r)
	}
	return r
}

func Start(addr string, r *mux.Router) error {
	return http.ListenAndServe(addr, r)
}
