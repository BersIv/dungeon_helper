package races

import (
	"dungeons_helper/util"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service
	util.TokenGetter
}

func NewHandler(s Service, tg util.TokenGetter) *Handler {
	return &Handler{
		Service:     s,
		TokenGetter: tg,
	}
}

func (h *Handler) GetAllRaces(w http.ResponseWriter, r *http.Request) {
	_, err := h.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	ctx := r.Context()
	res, err := h.Service.GetAllRaces(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}
