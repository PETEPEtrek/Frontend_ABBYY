package entity

import (
	"time"
)

type People struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Story      string
	Image      string
	BirthDate  time.Time
	Tags       string
	Score      float64
	VoteNumber uint
}

type PeopleRepository interface {
	Create(*People) error
	GetAll() (*[]People, error)
	Get(id uint) (*People, error)
	GetByTags(tags string) (*[]People, error)
	Update(*People) error
	Delete(id uint) error
}

type PeopleService interface {
	Get(id uint) (*People, error)
	GetAll() (*[]People, error)
	Update(people *People) error
	Delete(people *People) error

	ChangeScore(id uint, score float64) error
}
