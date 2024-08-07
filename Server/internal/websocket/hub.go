package websocket

import (
	"context"
	"dungeons_helper/db"
	"dungeons_helper/internal/character"
	"dungeons_helper/internal/lobby"
	"fmt"

	"github.com/gorilla/websocket"
)

type AccLobby struct {
	Id      int64 `json:"id"`
	IdAcc   int64 `json:"idAcc"`
	IdLobby int64 `json:"idLobby"`
}

type Hub struct {
	Lobbies      map[int64]*lobby.Lobby
	LobbyMembers map[int64][]*Client
	JoinRoom     chan *Client
	LeaveRoom    chan *Client
	Broadcast    chan *Command
	Db           db.DatabaseTX
	Character    character.Repository
}

func NewHub(db db.DatabaseTX) *Hub {
	return &Hub{
		Lobbies:      make(map[int64]*lobby.Lobby),
		LobbyMembers: make(map[int64][]*Client),
		JoinRoom:     make(chan *Client),
		LeaveRoom:    make(chan *Client),
		Broadcast:    make(chan *Command, 5),
		Db:           db,
	}
}

type ErrorMessage struct {
	ErrorMsg string `json:"error"`
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.JoinRoom:
			h.addToLobby(cl)
			ctx := cl.Context
			tx, err := h.Db.BeginTx(ctx, nil)
			if err != nil {
				fmt.Println(err)
				sendErrorMessage(cl.Conn, "Failed to start transaction")
				continue
			}
			query := `INSERT INTO accLobby(idAcc, idLobby) VALUES (?, ?)`
			_, err = tx.ExecContext(ctx, query, cl.Id, cl.IdLobby)
			if err != nil {
				fmt.Println(err)
				sendErrorMessage(cl.Conn, err.Error())
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					fmt.Println("Failed to rollback transaction:", rollbackErr)
				}
				continue
			}
			if err := tx.Commit(); err != nil {
				fmt.Println(err)
				sendErrorMessage(cl.Conn, err.Error())
				continue
			}

		case cl := <-h.LeaveRoom:
			ctx := context.Background()
			tx, err := h.Db.BeginTx(ctx, nil)
			if err != nil {
				fmt.Println(err)
				sendErrorMessage(cl.Conn, "Failed to start transaction")
				continue
			}
			var masterId int64
			query := `SELECT lobbyMasterId FROM lobby WHERE id = ?`
			err = tx.QueryRowContext(ctx, query, cl.IdLobby).Scan(&masterId)
			if err != nil {
				fmt.Println(err)
				sendErrorMessage(cl.Conn, err.Error())
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					fmt.Println("Failed to rollback transaction:", rollbackErr)
				}
				continue
			}
			if masterId == cl.Id {
				if clients, ok := h.LobbyMembers[cl.IdLobby]; ok {
					for _, client := range clients {
						// leaveCmd := &Command{
						// 	Type:    "leave",
						// 	Payload: cl.IdLobby,
						// }
						// err := client.Conn.WriteJSON(leaveCmd)
						// if err != nil {
						// 	sendErrorMessage(client.Conn, err.Error())
						// 	continue
						// }

						client.Conn.Close()
					}
				}

				// query = `DELETE FROM accLobby WHERE idLobby = ?`
				// _, err = tx.ExecContext(ctx, query, cl.IdLobby)
				// if err != nil {
				// 	fmt.Println(err)
				// 	sendErrorMessage(cl.Conn, err.Error())
				// 	if rollbackErr := tx.Rollback(); rollbackErr != nil {
				// 		fmt.Println("Failed to rollback transaction:", rollbackErr)
				// 	}
				// 	continue
				// }
				// query = `DELETE FROM lobby WHERE id = ?`
				// _, err = tx.ExecContext(ctx, query, cl.IdLobby)
				// if err != nil {
				// 	fmt.Println(err)
				// 	sendErrorMessage(cl.Conn, err.Error())
				// 	if rollbackErr := tx.Rollback(); rollbackErr != nil {
				// 		fmt.Println("Failed to rollback transaction:", rollbackErr)
				// 	}
				// 	continue
				// }
			} else {
				// query = `DELETE FROM accLobby WHERE idAcc = ? AND idLobby = ?`
				// _, err = tx.ExecContext(ctx, query, cl.Id, cl.IdLobby)
				// if err != nil {
				// 	fmt.Println(err)
				// 	sendErrorMessage(cl.Conn, err.Error())
				// 	if rollbackErr := tx.Rollback(); rollbackErr != nil {
				// 		fmt.Println("Failed to rollback transaction:", rollbackErr)
				// 	}
				// 	continue
				// }
			}
			h.removeFromLobby(cl)

			if err := tx.Commit(); err != nil {
				fmt.Println(err)
				sendErrorMessage(cl.Conn, err.Error())
				continue
			}
		case cmd := <-h.Broadcast:
			switch cmd.Type {
			case "join":
				for _, cl := range h.LobbyMembers[cmd.Payload.(int64)] {
					err := cl.Conn.WriteJSON(cmd)
					if err != nil {
						sendErrorMessage(cl.Conn, err.Error())
						continue
					}
				}
			case "updateChar":
				for _, cl := range h.LobbyMembers[cmd.Payload.(int64)] {
					err := cl.Conn.WriteJSON(cmd)
					if err != nil {
						sendErrorMessage(cl.Conn, err.Error())
						continue
					}
				}
			}
		}
	}
}

func sendErrorMessage(conn *websocket.Conn, errorMsg string) {
	errMsg := ErrorMessage{ErrorMsg: errorMsg}
	if err := conn.WriteJSON(errMsg); err != nil {
		fmt.Println("Error sending error message:", err)
	}
}

func (h *Hub) addToLobby(cl *Client) {
	lobbyID := cl.IdLobby
	if _, ok := h.LobbyMembers[lobbyID]; !ok {
		h.LobbyMembers[lobbyID] = []*Client{}
	}

	h.LobbyMembers[lobbyID] = append(h.LobbyMembers[lobbyID], cl)
}

func (h *Hub) removeFromLobby(cl *Client) {
	lobbyID := cl.IdLobby
	if clients, ok := h.LobbyMembers[lobbyID]; ok {
		for i, client := range clients {
			if client.Id == cl.Id {
				h.LobbyMembers[lobbyID] = append(clients[:i], clients[i+1:]...)
				ctx := context.Background()
				tx, err := h.Db.BeginTx(ctx, nil)
				if err != nil {
					fmt.Println(err)
					sendErrorMessage(cl.Conn, "Failed to start transaction")
					continue
				}
				query := `DELETE FROM accLobby WHERE idLobby = ?`
				_, err = tx.ExecContext(ctx, query, cl.IdLobby)
				if err != nil {
					fmt.Println(err)
					sendErrorMessage(cl.Conn, err.Error())
					if rollbackErr := tx.Rollback(); rollbackErr != nil {
						fmt.Println("Failed to rollback transaction:", rollbackErr)
					}
					continue
				}
				if len(h.LobbyMembers[lobbyID]) == 0 {
					delete(h.LobbyMembers, lobbyID)
					query = `DELETE FROM lobby WHERE id = ?`
					_, err = tx.ExecContext(ctx, query, cl.IdLobby)
					if err != nil {
						fmt.Println(err)
						sendErrorMessage(cl.Conn, err.Error())
						if rollbackErr := tx.Rollback(); rollbackErr != nil {
							fmt.Println("Failed to rollback transaction:", rollbackErr)
						}
						continue
					}
				}
				if err := tx.Commit(); err != nil {
					fmt.Println(err)
					sendErrorMessage(cl.Conn, err.Error())
					continue
				}
				return
			}
		}
	}
}
