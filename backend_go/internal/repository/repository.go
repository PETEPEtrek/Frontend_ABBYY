package repository

import (
	"backend_go/internal/entity"
	repo_sqlite "backend_go/internal/repository/sqlite"

	"gorm.io/gorm"
)

type Repository struct {
	User       entity.UserRepository
	Comments   entity.CommentRepository
	Games      entity.GameRepository
	Characters entity.CharacterRepository
	People     entity.PeopleRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:       repo_sqlite.NewUserSQLite(db),
		Comments:   repo_sqlite.NewCommentSQLite(db),
		Games:      repo_sqlite.NewGameSQLite(db),
		Characters: repo_sqlite.NewCharacterSQLite(db),
		People:     repo_sqlite.NewPeopleSQLite(db),
	}
}
