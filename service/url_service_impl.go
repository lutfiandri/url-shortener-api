package service

import (
	"url-shortener-api/entity"
	"url-shortener-api/exception"
	"url-shortener-api/repository"
)

type urlService struct {
	UrlRepository repository.UrlRepository
}

func NewUrlService(urlRepository repository.UrlRepository) UrlService {
	return &urlService{
		UrlRepository: urlRepository,
	}
}

func (s *urlService) Create(url entity.Url) string {
	err := s.UrlRepository.Create(url)
	exception.PanicIfNeeded(err)
	shortUrl := "http://localhost:8080/" + url.Id
	return shortUrl
}

func (s *urlService) FindById(id string) string {
	url, err := s.UrlRepository.FindById(id)
	exception.PanicIfNeeded(err)
	return url.LongUrl
}
