package repository

import (
	"fmt"

	"url-shortener-api/entity"

	"gorm.io/gorm"
)

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(database *gorm.DB) UrlRepository {
	database.AutoMigrate(&entity.Url{})
	return &urlRepository{
		db: database,
	}
}

func (r *urlRepository) FindAll() ([]entity.Url, error) {
	var urls []entity.Url
	err := r.db.Find(&urls).Error
	return urls, err
}

func (r *urlRepository) FindById(id string) (entity.Url, error) {
	var url entity.Url
	err := r.db.First(&url, "id = ?", id).Error
	return url, err
}

func (r *urlRepository) Create(url entity.Url) error {
	err := r.db.Create(&url).Error
	return err
}

func (r *urlRepository) DeleteById(id string) error {
	res := r.db.Delete(entity.Url{}, "id = ?", id)

	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected < 1 {
		return fmt.Errorf("rl with id %v does not exist", id)
	}
	return nil
}
