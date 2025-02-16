package database

import (
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"web/model"
)

func ProvideDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return db, nil
}

var DatabaseModule = fx.Module("database", fx.Provide(ProvideDatabase))
