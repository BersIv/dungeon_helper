package lobby

import (
	"context"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) GetAllLobby(c context.Context) ([]GetLobbyRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	res, err := s.Repository.GetAllLobby(ctx)
	if err != nil {
		return nil, err
	}

	return res, err
}
