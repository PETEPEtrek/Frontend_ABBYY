package service

import (
	"backend_go/internal/entity"
	"backend_go/internal/repository"
)

type Service struct {
	User      entity.UserService
	Comment   entity.CommentService
	Game      entity.GameService
	Character entity.CharacterService
	People    entity.PeopleService
}

func NewService(repo *repository.Repository) *Service {
	userService := NewUserService(repo.User)
	return &Service{
		User:      userService,
		Comment:   NewCommentService(repo.Comments),
		Game:      NewGameService(repo.Games),
		Character: NewCharacterService(repo.Characters),
		People:    NewPeopleService(repo.People),
	}
}
