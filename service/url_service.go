package service

import (
	"url-shortener-api/entity"
)

type UrlService interface {
	FindAll() ([]entity.Url, error)
	FindById(id string) (entity.Url, error)
	Create(url entity.Url) error
	DeleteById(id string) error
}
