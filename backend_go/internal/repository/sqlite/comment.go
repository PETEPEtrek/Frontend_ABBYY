package repo_sqlite

import (
	"backend_go/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type CommentSQLite struct {
	db *gorm.DB
}

func NewCommentSQLite(db *gorm.DB) *CommentSQLite {
	return &CommentSQLite{db: db}
}

func (r *CommentSQLite) GetAll() (*[]entity.Comment, error) {
	var comments []entity.Comment

	if result := r.db.Find(&comments); result.Error != nil {
		return nil, result.Error
	} else {
		return &comments, nil
	}
}

func (r *CommentSQLite) Get(id uint) (*entity.Comment, error) {
	var comment entity.Comment

	if result := r.db.Where("id = ?", id).First(&comment); result.Error == nil {
		return &comment, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &comment, entity.ErrCommentNotFound
	} else {
		return &comment, result.Error
	}
}

func (r *CommentSQLite) GetByPostID(postID uint) (*[]entity.Comment, error) {
	var comments []entity.Comment

	if result := r.db.Where("userid = ?", postID).Find(&comments); result.Error == nil {
		return &comments, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, entity.ErrCommentNotFound
	} else {
		return &comments, result.Error
	}
}

func (r *CommentSQLite) Create(comment *entity.Comment) error {
	if result := r.db.Create(comment); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *CommentSQLite) Update(comment *entity.Comment) error {
	result := r.db.Model(comment).Updates(comment)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *CommentSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.Comment{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
