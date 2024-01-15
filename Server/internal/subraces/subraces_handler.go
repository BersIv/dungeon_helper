package subraces

import (
	"dungeons_helper/util"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) GetAllSubraces(w http.ResponseWriter, r *http.Request) {
	_, err := util.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var raceId RaceId
	err = json.NewDecoder(r.Body).Decode(&raceId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if raceId.RaceId == 0 {
		http.Error(w, "RaceId cannot be zero", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	res, err := h.Service.GetAllSubraces(ctx, raceId.RaceId)
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
