package router

import (
	"dungeons_helper/internal/account"
	"dungeons_helper/internal/alignment"
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
