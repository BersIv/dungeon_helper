package subraces

import (
	"context"
	"database/sql"
	"dungeons_helper/db"
)

type repository struct {
	db db.DatabaseTX
}

func NewRepository(db db.DatabaseTX) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllSubraces(ctx context.Context, req GetSubracesReq) ([]Subraces, error) {
	var subraces []Subraces

	query := `SELECT r.id, r.subraceName, s.strength, s.dexterity, s.constitution, 
				s.intelligence, s.wisdom, s.charisma FROM subrace r
				LEFT JOIN stats s ON s.id = r.idStats
				WHERE idRace = ?`
	rows, err := r.db.QueryContext(ctx, query, req.IdRace)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	for rows.Next() {
		var subrace Subraces
		err := rows.Scan(&subrace.Id, &subrace.SubraceName, &subrace.Stats.Strength,
			&subrace.Stats.Dexterity, &subrace.Stats.Constitution, &subrace.Stats.Intelligence,
			&subrace.Stats.Wisdom, &subrace.Stats.Charisma)
		if err != nil {
			return nil, err
		}
		subraces = append(subraces, subrace)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return subraces, nil
}
