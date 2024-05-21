package repo_sqlite

import (
	"backend_go/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type GameSQLite struct {
	db *gorm.DB
}

func NewGameSQLite(db *gorm.DB) *GameSQLite {
	return &GameSQLite{db: db}
}

func (r *GameSQLite) GetAll() (*[]entity.Game, error) {
	var games []entity.Game

	if result := r.db.Find(&games); result.Error != nil {
		return nil, result.Error
	} else {
		return &games, nil
	}
}

func (r *GameSQLite) Get(id uint) (*entity.Game, error) {
	var game entity.Game

	if result := r.db.Where("id = ?", id).First(&game); result.Error == nil {
		return &game, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &game, entity.ErrGameNotFound
	} else {
		return &game, result.Error
	}
}

func (r *GameSQLite) GetByTags(tags string) (*[]entity.Game, error) {
	var games []entity.Game

	if result := r.db.Where("Tags LIKE ?", tags).Find(&games); result.Error == nil {
		return &games, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else {
		return &games, result.Error
	}
}

func (r *GameSQLite) Create(game *entity.Game) error {
	if result := r.db.Create(game); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *GameSQLite) Update(game *entity.Game) error {
	result := r.db.Model(game).Updates(game)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *GameSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.Game{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
