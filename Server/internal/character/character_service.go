package character

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

func (s *service) GetAllCharactersByAccId(c context.Context, idAcc int64) ([]Character, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	chars, err := s.Repository.GetAllCharactersByAccId(ctx, idAcc)
	if err != nil {
		return nil, err
	}

	return chars, nil
}

func (s *service) GetCharacterById(c context.Context, id int64) (*Character, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	char, err := s.Repository.GetCharacterById(ctx, id)
	if err != nil {
		return nil, nil
	}

	return char, nil
}

func (s *service) CreateCharacter(c context.Context, character *CreateCharacterReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	character.Stats.Strength = character.Stats.Strength + character.Subrace.Stats.Strength
	character.Stats.Dexterity = character.Stats.Dexterity + character.Subrace.Stats.Dexterity
	character.Stats.Constitution = character.Stats.Constitution + character.Subrace.Stats.Constitution
	character.Stats.Intelligence = character.Stats.Intelligence + character.Subrace.Stats.Intelligence
	character.Stats.Wisdom = character.Stats.Wisdom + character.Subrace.Stats.Wisdom
	character.Stats.Charisma = character.Stats.Charisma + character.Subrace.Stats.Charisma

	err := s.Repository.CreateCharacter(ctx, character)
	if err != nil {
		return err
	}

	return nil
}
