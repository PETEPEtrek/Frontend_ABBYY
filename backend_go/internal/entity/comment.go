package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	Text        string
	IsGame      bool
	IsCharacter bool
	IsPeople    bool
	PostID      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type CommentRepository interface {
	Create(*Comment) error
	GetAll() (*[]Comment, error)
	Get(id uint) (*Comment, error)
	GetByPostID(postID uint) (*[]Comment, error)
	Update(*Comment) error
	Delete(id uint) error
}

type CommentService interface {
	Get(id uint) (*Comment, error)
	Update(comment *Comment) error
	Delete(comment *Comment) error
}
