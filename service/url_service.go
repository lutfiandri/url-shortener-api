package service

import (
	"url-shortener-api/entity"
)

type UrlService interface {
	Create(url entity.Url) error
	FindById(id string) (entity.Url, error)
}
