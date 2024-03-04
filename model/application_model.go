package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(gereric interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&gereric)

	return db, nil
}
