package service

import (
	"backend_go/internal/entity"
	"errors"
)

type GameService struct {
	gameRepo entity.GameRepository
}

func NewGameService(gameRepo entity.GameRepository) *GameService {
	return &GameService{
		gameRepo: gameRepo,
	}
}

func (s *GameService) Get(id uint) (*entity.Game, error) {
	gameDB, err := s.gameRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		return gameDB, nil
	}
}

func (s *GameService) GetAll() (*[]entity.Game, error) {
	gameDB, err := s.gameRepo.GetAll()
	if err != nil {
		return nil, err
	} else {
		return gameDB, nil
	}
}

func (s *GameService) Update(game *entity.Game) error {
	_, err := s.gameRepo.Get(game.ID)
	if err != nil {
		return err
	}
	err = s.gameRepo.Update(game)
	return err
}

func (s *GameService) Delete(game *entity.Game) error {
	err := s.gameRepo.Delete(game.ID)
	return err
}

func (s *GameService) ChangeScore(id uint, score float64) error {
	if score < 1 || score > 10 {
		return errors.New("bad score")
	}

	game, err := s.gameRepo.Get(id)
	if err != nil {
		return err
	}
	game.Score = (game.Score*float64(game.VoteNumber) + score) / (float64(game.VoteNumber) + 1)
	game.VoteNumber += 1
	err = s.gameRepo.Update(game)
	if err != nil {
		return err
	}

	return nil

}
