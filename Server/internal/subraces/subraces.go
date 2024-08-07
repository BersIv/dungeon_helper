package subraces

import (
	"context"
	"dungeons_helper/internal/stats"
)

type Subraces struct {
	Id          int64             `json:"id"`
	SubraceName string            `json:"raceName"`
	Stats       stats.GetStatsRes `json:"stats"`
}

type GetSubracesReq struct {
	IdRace int64 `json:"idRace"`
}

type CreateCharReq struct {
	Id          int64             `json:"id"`
	SubraceName string            `json:"raceName"`
	Stats       stats.GetStatsRes `json:"Stats"`
}

type Repository interface {
	GetAllSubraces(ctx context.Context, race GetSubracesReq) ([]Subraces, error)
}

type Service interface {
	GetAllSubraces(c context.Context, race GetSubracesReq) ([]Subraces, error)
}
