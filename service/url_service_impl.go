package service

import (
	"url-shortener-api/entity"
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

func (s *urlService) FindAll() ([]entity.Url, error) {
	urls, err := s.UrlRepository.FindAll()
	return urls, err
}

func (s *urlService) FindById(id string) (entity.Url, error) {
	url, err := s.UrlRepository.FindById(id)
	return url, err
}

func (s *urlService) Create(url entity.Url) error {
	err := s.UrlRepository.Create(url)
	return err
}

func (s *urlService) DeleteById(id string) error {
	err := s.UrlRepository.DeleteById(id)
	return err
}
