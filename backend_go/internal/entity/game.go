package entity

import (
	"time"
)

type Game struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"unique"`
	Synopsis   string
	Image      string
	Date       time.Time
	Tags       string
	Score      float64
	VoteNumber uint
}

type GameRepository interface {
	Create(*Game) error
	GetAll() (*[]Game, error)
	Get(id uint) (*Game, error)
	GetByTags(tags string) (*[]Game, error)
	Update(*Game) error
	Delete(id uint) error
}

type GameService interface {
	Get(id uint) (*Game, error)
	GetAll() (*[]Game, error)
	Update(game *Game) error
	Delete(game *Game) error

	ChangeScore(id uint, score float64) error
}
