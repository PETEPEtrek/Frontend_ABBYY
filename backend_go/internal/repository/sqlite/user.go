package repo_sqlite

import (
	"backend_go/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type UserSQLite struct {
	db *gorm.DB
}

func NewUserSQLite(db *gorm.DB) *UserSQLite {
	return &UserSQLite{db: db}
}

func (r *UserSQLite) GetAll() (*[]entity.User, error) {
	var users []entity.User

	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	} else {
		return &users, nil
	}
}

func (r *UserSQLite) Get(id uint) (*entity.User, error) {
	var user entity.User

	if result := r.db.Where("id = ?", id).First(&user); result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user, entity.ErrUserNotFound
	} else {
		return &user, result.Error
	}
}

func (r *UserSQLite) GetByEmail(email string) (*entity.User, error) {
	var user entity.User

	if result := r.db.Where("email = ?", email).First(&user); result.Error == nil {
		return &user, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user, entity.ErrUserNotFound
	} else {
		return &user, result.Error
	}
}

func (r *UserSQLite) Create(user *entity.User) error {
	if result := r.db.Create(user); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *UserSQLite) Update(user *entity.User) error {
	result := r.db.Model(user).Updates(user)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *UserSQLite) Delete(id uint) error {
	result := r.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
