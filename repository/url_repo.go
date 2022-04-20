package repository

import "url-shortener-api/entity"

type UrlRepository interface {
	FindAll() ([]entity.Url, error)
	FindById(id string) (entity.Url, error)
	FindByUserId(userId string) ([]entity.Url, error)
	Create(url entity.Url) error
	DeleteById(id string) error
}
