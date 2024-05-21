package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	UserRegister
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type UserLogin struct {
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null" json:"Password,omitempty""`
}

type UserRegister struct {
	UserLogin

	UserName string
}

type UserRepository interface {
	Create(*User) error
	GetAll() (*[]User, error)
	Get(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(*User) error
	Delete(id uint) error
}

type UserService interface {
	Get(id uint) (*User, error)
	GetAll() (*[]User, error)
	Update(user *User) error
	Delete(user *User) error

	Register(userReg *UserRegister) error
	Login(userLogin *UserLogin) (uint, error)
}

func (u *User) OmitPassword() {
	u.Password = ""
}
