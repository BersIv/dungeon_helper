package router

import (
	"dungeons_helper/internal/account"
	"net/http"

	"github.com/gorilla/mux"
)

type responseWriterWithStatus struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseWriterWithStatus) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

type Option func(router *mux.Router)

func AccountRoutes(accountHandler *account.Handler) Option {
	return func(r *mux.Router) {
		r.HandleFunc("/auth/registration", accountHandler.CreateAccount).Methods("POST")
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
