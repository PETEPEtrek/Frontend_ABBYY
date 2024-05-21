package repo_sqlite

import (
	"backend_go/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type CharacterSQLite struct {
	db *gorm.DB
}

func NewCharacterSQLite(db *gorm.DB) *CharacterSQLite {
	return &CharacterSQLite{db: db}
}

func (r *CharacterSQLite) GetAll() (*[]entity.Character, error) {
	var characters []entity.Character

	if result := r.db.Find(&characters); result.Error != nil {
		return nil, result.Error
	} else {
		return &characters, nil
	}
}

func (r *CharacterSQLite) Get(id uint) (*entity.Character, error) {
	var character entity.Character

	if result := r.db.Where("id = ?", id).First(&character); result.Error == nil {
		return &character, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &character, entity.ErrCharacterNotFound
	} else {
		return &character, result.Error
	}
}

func (r *CharacterSQLite) GetByTags(tags string) (*[]entity.Character, error) {
	var characters []entity.Character

	if result := r.db.Where("Tags LIKE ?", tags).Find(&characters); result.Error == nil {
		return &characters, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else {
		return &characters, result.Error
	}
}

func (r *CharacterSQLite) Create(character *entity.Character) error {
	if result := r.db.Create(character); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *CharacterSQLite) Update(character *entity.Character) error {
	result := r.db.Model(character).Updates(character)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *CharacterSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.Character{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
