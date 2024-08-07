package websocket

import (
	"dungeons_helper/db"
	"dungeons_helper/internal/character"
	"dungeons_helper/util"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
)

type Handler struct {
	db          db.DatabaseTX
	hub         *Hub
	tokenGetter util.TokenGetter
}

func NewHandler(db db.DatabaseTX, h *Hub, tg util.TokenGetter) *Handler {
	return &Handler{
		db:          db,
		hub:         h,
		tokenGetter: tg,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) CreateLobby(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	accountId, err := h.tokenGetter.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	lobbyName := r.URL.Query().Get("lobbyName")
	lobbyPassword := r.URL.Query().Get("lobbyPassword")
	amount := r.URL.Query().Get("amount")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := `INSERT INTO lobby(lobbyMasterId, lobbyName,
		                  lobbyPassword, amount) VALUES (?, ?, ?, ?)`
	res, err := tx.ExecContext(ctx, query, accountId, lobbyName, lobbyPassword, amount)
	if err != nil {
		var dbErr *mysql.MySQLError
		ok := errors.As(err, &dbErr)
		if ok && dbErr.Number == 1062 {
			http.Error(w, "Duplicate entry", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
	idLobby, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Res: ", idLobby)
	nickname, err := h.tokenGetter.GetNickNameFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	h.Reg(w, r, accountId, idLobby, nickname, true)
}

func (h *Handler) JoinLobby(w http.ResponseWriter, r *http.Request) {
	accountId, err := h.tokenGetter.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	nickname, err := h.tokenGetter.GetNickNameFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	idLobbyStr := r.URL.Query().Get("idLobby")
	idLobby, err := strconv.Atoi(idLobbyStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	h.Reg(w, r, accountId, int64(idLobby), nickname, false)
}

func (h *Handler) Reg(w http.ResponseWriter, r *http.Request, accountId int64, idLobby int64, nickname string, master bool) {
	conn, err := upgrader.Upgrade(&util.HijackableResponseWriter{ResponseWriter: w}, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()
	ctx := r.Context()
	cl := &Client{
		Conn:     conn,
		Id:       accountId,
		IdLobby:  idLobby,
		Nickname: nickname,
		Context:  ctx,
	}
	charRepo := character.NewRepository(h.db)

	if !master {
		var idChar int64
		query := `SELECT idChar FROM accChar WHERE idAccount = ? AND act = 1`
		_ = h.db.QueryRowContext(ctx, query, accountId).Scan(&idChar)
		if idChar == 0 {
			conn.Close()
			log.Println("No char")
		}
		char, _ := charRepo.GetCharacterById(ctx, idChar)

		cl.Character = char

		joinMessage := &Command{
			Type:      "join",
			Payload:   cl.IdLobby,
			Character: cl.Character,
		}
		cl.Conn.WriteJSON(joinMessage)

		h.hub.Broadcast <- joinMessage
	}

	h.hub.JoinRoom <- cl
	cl.readCommand(h.hub, charRepo)
	log.Println("Client connected")
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Println("Connection closed:", err)
				return
			}
			log.Println("Error reading message:", err)
			return
		}
	}

}
