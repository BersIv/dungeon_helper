package lobby

import (
	"context"
	"dungeons_helper/db"
)

type repository struct {
	db db.DatabaseTX
}

func NewRepository(db db.DatabaseTX) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllLobby(ctx context.Context) ([]GetLobbyRes, error) {
	var lobbyList []GetLobbyRes

	query := `SELECT l.id, l.lobbyName, count(ac.idAcc) FROM lobby l 
    LEFT JOIN accLobby ac on l.id = ac.idLobby 
    LEFT JOIN account a on ac.idAcc = a.id
	GROUP BY l.id, l.lobbyName, a.id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var lobby GetLobbyRes
		err := rows.Scan(&lobby.Id, &lobby.LobbyName, &lobby.PlayersInLobby)
		if err != nil {
			return nil, err
		}
		lobbyList = append(lobbyList, lobby)
	}

	return lobbyList, nil
}
