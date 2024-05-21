package repo_sqlite

import (
	"backend_go/internal/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB(db_uri string) (*gorm.DB, error) {
	db, err := gorm.Open(
		sqlite.Open(db_uri),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Game{},
		&entity.Character{},
		&entity.People{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
