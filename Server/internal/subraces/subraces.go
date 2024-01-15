package subraces

import (
	"context"
	"dungeons_helper/internal/stats"
)

type Subraces struct {
	Id          int64       `json:"id"`
	SubraceName string      `json:"raceName"`
	Stats       stats.Stats `json:"idStats"`
}

type RaceId struct {
	RaceId int64 `json:"raceId"`
}

type Repository interface {
	GetAllSubraces(ctx context.Context, idRace int64) ([]Subraces, error)
}

type Service interface {
	GetAllSubraces(c context.Context, idRace int64) ([]Subraces, error)
}
