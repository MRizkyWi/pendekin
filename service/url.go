package service

import (
	"context"
	"fmt"
	"pendekin/cache"
	"pendekin/config"
	"pendekin/model"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	"golang.org/x/exp/rand"
)

func (s *UrlService) GetByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error) {
	redisValue, err := s.redisClient.Get(ctx, shortUrl).Result()
	if err != nil && err != redis.Nil {
		log.Error(fmt.Sprintf("error on redisClient get %s", err.Error()))

		return nil, err
	}

	if redisValue != "" {
		return &model.Url{
			ShortUrl:  shortUrl,
			ActualUrl: redisValue,
		}, nil
	}

	url, err := s.urlRepository.GetByShortUrl(shortUrl)
	if err != nil {
		log.Error(fmt.Sprintf("error on repo get by short url %s", err.Error()))

		return nil, err
	}

	err = s.redisClient.Set(ctx, shortUrl, url.ActualUrl, cache.GetDefaultExpiryTime()).Err()
	if err != nil {
		log.Error(fmt.Sprintf("error on redis client set %s", err.Error()))
	}

	return url, nil
}

func (s *UrlService) UpdateStatus(ctx context.Context, req model.UpdateStatusRequest) error {
	err := s.urlRepository.UpdateStatus(req)
	if err != nil {
		log.Error(fmt.Sprintf("error on repo update status %s", err.Error()))

		return err
	}

	if req.Status != false {
		url, err := s.urlRepository.GetByID(req.ID)
		if err != nil {
			log.Error(fmt.Sprintf("error on repo update status %s", err.Error()))

			return nil
		}

		err = s.redisClient.Del(ctx, url.ShortUrl).Err()
		if err != nil {
			log.Error(fmt.Sprintf("error on redis client set %s", err.Error()))

			return nil
		}
	}

	return nil
}

func (s *UrlService) Save(ctx context.Context, req model.NewUrlRequest) (model.NewUrlResponse, error) {
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
		ShortUrl: generateShortURL(shortKey),
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

func generateShortURL(shortKey string) string {
	return fmt.Sprintf("%s/%s", config.GetEnv("BASE_URL", ""), shortKey)
}