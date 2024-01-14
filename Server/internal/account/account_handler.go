package account

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account CreateAccountReq
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	ctx := r.Context()
	err = h.Service.CreateAccount(ctx, &account)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		ok := errors.As(err, &mysqlErr)
		if ok && mysqlErr.Number == 1062 {
			http.Error(w, err.Error(), http.StatusConflict)
		} else if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "Wrong password or email", http.StatusRequestTimeout)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
