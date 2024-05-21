package service

import (
	"backend_go/internal/entity"
)

type PeopleService struct {
	peopleRepo entity.PeopleRepository
}

func NewPeopleService(peopleRepo entity.PeopleRepository) *PeopleService {
	return &PeopleService{
		peopleRepo: peopleRepo,
	}
}

func (s *PeopleService) Get(id uint) (*entity.People, error) {
	peopleDB, err := s.peopleRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		return peopleDB, nil
	}
}

func (s *PeopleService) GetAll() (*[]entity.People, error) {
	peopleDB, err := s.peopleRepo.GetAll()
	if err != nil {
		return nil, err
	} else {
		return peopleDB, nil
	}
}

func (s *PeopleService) Update(people *entity.People) error {
	_, err := s.peopleRepo.Get(people.ID)
	if err != nil {
		return err
	}
	err = s.peopleRepo.Update(people)
	return err
}

func (s *PeopleService) Delete(people *entity.People) error {
	err := s.peopleRepo.Delete(people.ID)
	return err
}

func (s *PeopleService) ChangeScore(id uint, score float64) error {
	people, err := s.peopleRepo.Get(id)
	if err != nil {
		return err
	}
	people.Score = (people.Score*float64(people.VoteNumber) + score) / (float64(people.VoteNumber) + 1)
	people.VoteNumber += 1
	err = s.peopleRepo.Update(people)
	if err != nil {
		return err
	}

	return nil

}
