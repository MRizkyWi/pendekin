package service

import (
	"pendekin/config"
	"pendekin/model"
	"pendekin/repository"
	"time"

	"golang.org/x/exp/rand"
)

type UrlService struct {
	urlRepository repository.Url
}

func NewUrlService(urlRepository repository.Url) *UrlService {
	return &UrlService{
		urlRepository: urlRepository,
	}
}

func (s *UrlService) GetByShortUrl(shortUrl string) (*model.Url, error) {
	return s.urlRepository.GetByShortUrl(shortUrl)
}

func (s *UrlService) UpdateStatus(req model.UpdateStatusRequest) error {
	return s.urlRepository.UpdateStatus(req)
}

func (s *UrlService) Save(req model.NewUrlRequest) (model.NewUrlResponse, error) {
	shortKey := generateShortKey()

	url := model.Url{
		ShortUrl:  shortKey,
		ActualUrl: req.ActualUrl,
	}

	err := s.urlRepository.Save(url)
	if err != nil {
		return model.NewUrlResponse{}, err
	}

	urlResponse := model.NewUrlResponse{
		ShortUrl: config.GetEnv("BASE_URL", "") + shortKey,
	}

	return urlResponse, nil
}

// credit to https://dev.to/envitab/how-to-build-a-url-shortener-with-go-5hn5
func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	rand.Seed(uint64(time.Now().UnixNano()))
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
