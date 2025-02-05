package model

import "time"

type Url struct {
	ID        int       `json:"id"`
	ShortUrl  string    `json:"short_url"`
	ActualUrl string    `json:"actual_url"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
