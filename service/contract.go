package service

import (
	"context"
	"pendekin/model"
	"pendekin/repository"

	"github.com/go-redis/redis/v8"
)

type UrlService struct {
	urlRepository repository.Url
	redisClient   *redis.Client
}

type UrlServiceInterface interface {
	GetByShortUrl(ctx context.Context, shortUrl string) (*model.Url, error)
	UpdateStatus(ctx context.Context, req model.UpdateStatusRequest) (*model.UpdateStatusRequest, error)
	Save(ctx context.Context, req model.NewUrlRequest) (*model.NewUrlResponse, error)
}

func NewUrlService(urlRepository repository.Url, redisClient *redis.Client) *UrlService {
	return &UrlService{
		urlRepository: urlRepository,
		redisClient:   redisClient,
	}
}
