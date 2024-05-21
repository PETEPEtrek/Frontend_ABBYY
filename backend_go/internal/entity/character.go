package entity

import (
	"time"
)

type Character struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Story      string
	Image      string
	BirthDate  time.Time
	Tags       string
	Score      float64
	VoteNumber uint
}

type CharacterRepository interface {
	Create(*Character) error
	GetAll() (*[]Character, error)
	Get(id uint) (*Character, error)
	GetByTags(Tags string) (*[]Character, error)
	Update(*Character) error
	Delete(id uint) error
}

type CharacterService interface {
	Get(id uint) (*Character, error)
	GetAll() (*[]Character, error)
	Update(character *Character) error
	Delete(character *Character) error

	ChangeScore(id uint, score float64) error
}
