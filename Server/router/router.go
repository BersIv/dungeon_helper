package router

import (
	"dungeons_helper/internal/account"
	"dungeons_helper/internal/alignment"
	"dungeons_helper/internal/class"
	"dungeons_helper/internal/races"
	"dungeons_helper/internal/stats"
	"dungeons_helper/internal/subraces"
	"net/http"

	"github.com/gorilla/mux"
)

type Option func(router *mux.Router)

func AccountRouter(accountHandler *account.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/auth/registration", accountHandler.CreateAccount).Methods("POST")
		r.HandleFunc("/auth/byEmail", accountHandler.Login).Methods("POST")
		r.HandleFunc("/logout", accountHandler.Logout).Methods("POST")
		r.HandleFunc("/auth/google/login", accountHandler.LoginGoogle).Methods("POST")
		r.HandleFunc("/auth/restorePassword", accountHandler.RestorePassword).Methods("POST")
		r.HandleFunc("/account/change/nickname", accountHandler.UpdateNickname).Methods("PATCH")
		r.HandleFunc("/account/change/password", accountHandler.UpdatePassword).Methods("PATCH")
		r.HandleFunc("/account/change/avatar", accountHandler.UpdateAvatar).Methods("PATCH")
	}
}

func AlignmentRouter(alignmentHandler *alignment.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/getAlignments", alignmentHandler.GetAllAlignments).Methods("GET")
	}
}

func ClassRouter(classHandler *class.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/getClasses", classHandler.GetAllClasses).Methods("GET")
	}
}

func RacesRouter(racesHandler *races.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/race/getRaces", racesHandler.GetAllRaces).Methods("GET")
	}
}

func SubracesRouter(subracesHandler *subraces.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/subrace/getSubraces", subracesHandler.GetAllSubraces).Methods("GET")
	}
}

func StatsRouter(statsHandler *stats.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/getStatsById", statsHandler.GetStatsById).Methods("GET")
	}
}

func InitRouter(options ...Option) *mux.Router {
	r := mux.NewRouter()

	for _, option := range options {
		option(r)
	}
	return r
}

func Start(addr string, r *mux.Router) error {
	return http.ListenAndServe(addr, r)
}
