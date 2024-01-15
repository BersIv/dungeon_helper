package lobby

import (
	"context"
)

type Lobby struct {
	Id            int64  `json:"id"`
	LobbyMasterID int64  `json:"lobbyMasterID"`
	LobbyName     string `json:"lobbyName"`
	LobbyPassword string `json:"lobbyPassword"`
	Amount        int64  `json:"amount"`
}

type GetLobbyRes struct {
	Id             int64  `json:"id"`
	LobbyName      string `json:"lobbyName"`
	Amount         int64  `json:"amount"`
	PlayersInLobby int64  `json:"playersInLobby"`
}

type Repository interface {
	GetAllLobby(ctx context.Context) ([]GetLobbyRes, error)
}

type Service interface {
	GetAllLobby(c context.Context) ([]GetLobbyRes, error)
}
