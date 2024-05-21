package repo_sqlite

import (
	"backend_go/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type PeopleSQLite struct {
	db *gorm.DB
}

func NewPeopleSQLite(db *gorm.DB) *PeopleSQLite {
	return &PeopleSQLite{db: db}
}

func (r *PeopleSQLite) GetAll() (*[]entity.People, error) {
	var peoples []entity.People

	if result := r.db.Find(&peoples); result.Error != nil {
		return nil, result.Error
	} else {
		return &peoples, nil
	}
}

func (r *PeopleSQLite) Get(id uint) (*entity.People, error) {
	var people entity.People

	if result := r.db.Where("id = ?", id).First(&people); result.Error == nil {
		return &people, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &people, entity.ErrPeopleNotFound
	} else {
		return &people, result.Error
	}
}

func (r *PeopleSQLite) GetByTags(tags string) (*[]entity.People, error) {
	var peoples []entity.People

	if result := r.db.Where("Tags LIKE ?", tags).Find(&peoples); result.Error == nil {
		return &peoples, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else {
		return &peoples, result.Error
	}
}

func (r *PeopleSQLite) Create(people *entity.People) error {
	if result := r.db.Create(people); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *PeopleSQLite) Update(people *entity.People) error {
	result := r.db.Model(people).Updates(people)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *PeopleSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.People{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
