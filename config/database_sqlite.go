package config

import (
	"url-shortener-api/exception"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDatabase(location string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	exception.PanicIfNeeded(err)
	return db
}
