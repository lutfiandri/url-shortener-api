package config

import (
	"url-shortener-api/exception"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("url.db"), &gorm.Config{})
	exception.PanicIfNeeded(err)
	return db
}
