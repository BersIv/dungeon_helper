package stats

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
func (h *Handler) GetStatsById(w http.ResponseWriter, r *http.Request) {
	_, err := h.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req GetStatsReq
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Id == 0 {
		http.Error(w, "RaceId cannot be zero", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	res, err := h.Service.GetStatsById(ctx, req.Id)
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
