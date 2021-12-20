package service

import "url-shortener-api/entity"

type UrlService interface {
	Create(url entity.Url) string
	GetLongUrl(id string) string
}
