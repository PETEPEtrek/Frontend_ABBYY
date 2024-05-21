package service

import (
	"backend_go/internal/entity"
)

type CharacterService struct {
	characterRepo entity.CharacterRepository
}

func NewCharacterService(characterRepo entity.CharacterRepository) *CharacterService {
	return &CharacterService{
		characterRepo: characterRepo,
	}
}

func (s *CharacterService) Get(id uint) (*entity.Character, error) {
	characterDB, err := s.characterRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		return characterDB, nil
	}
}

func (s *CharacterService) GetAll() (*[]entity.Character, error) {
	characterDB, err := s.characterRepo.GetAll()
	if err != nil {
		return nil, err
	} else {
		return characterDB, nil
	}
}

func (s *CharacterService) Update(character *entity.Character) error {
	_, err := s.characterRepo.Get(character.ID)
	if err != nil {
		return err
	}
	err = s.characterRepo.Update(character)
	return err
}

func (s *CharacterService) Delete(character *entity.Character) error {
	err := s.characterRepo.Delete(character.ID)
	return err
}

func (s *CharacterService) ChangeScore(id uint, score float64) error {
	character, err := s.characterRepo.Get(id)
	if err != nil {
		return err
	}
	character.Score = (character.Score*float64(character.VoteNumber) + score) / (float64(character.VoteNumber) + 1)
	character.VoteNumber += 1
	err = s.characterRepo.Update(character)
	if err != nil {
		return err
	}

	return nil

}
