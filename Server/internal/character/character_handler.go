package character

import (
	"dungeons_helper/util"
	"encoding/base64"
	"encoding/json"
	"io"
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

func (h *Handler) GetAllCharactersByAccId(w http.ResponseWriter, r *http.Request) {
	idAcc, err := h.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	ctx := r.Context()
	res, err := h.Service.GetAllCharactersByAccId(ctx, int64(idAcc))
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

func (h *Handler) GetCharacterById(w http.ResponseWriter, r *http.Request) {
	_, err := h.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var id GetCharacterReq
	err = json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	res, err := h.Service.GetCharacterById(ctx, id.Id)
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

func (h *Handler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var char CreateCharacterReq
	err := json.NewDecoder(r.Body).Decode(&char)
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

	char.IdAcc, err = h.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	imageData, err := base64.StdEncoding.DecodeString(char.Avatar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	char.Avatar = string(imageData)

	err = h.Service.CreateCharacter(ctx, &char)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) SetActiveCharacterById(w http.ResponseWriter, r *http.Request) {
	var req SetActiveCharReq
	err := json.NewDecoder(r.Body).Decode(&req)
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

	req.IdAcc, err = h.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	ctx := r.Context()

	err = h.Service.SetActiveCharacterById(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
