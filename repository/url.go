package repository

import (
	"database/sql"
	"log"
	errors "pendekin/error"
	"pendekin/model"
)

type Url interface {
	Save(url model.Url) error
	GetByID(id int) (*model.Url, error)
	GetByShortUrl(shortUrl string) (*model.Url, error)
	UpdateStatus(req model.UpdateStatusRequest) error
}

type UrlRepository struct {
	DB *sql.DB
}

func NewURLRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{
		DB: db,
	}
}

func (u *UrlRepository) Save(url model.Url) (err error) {
	_, err = u.DB.Exec("INSERT INTO url (short_url, actual_url, is_active) VALUES (?, ?, ?)", url.ShortUrl, url.ActualUrl, url.IsActive)
	if err != nil {
		log.Printf("Failed to save url: %v", err)

		return err
	}

	return nil
}

func (u *UrlRepository) GetByID(id int) (*model.Url, error) {
	var url model.Url

	err := u.DB.QueryRow("SELECT id, short_url, actual_url, is_active FROM url WHERE id = ?", id).
		Scan(&url.ID, &url.ShortUrl, &url.ActualUrl, &url.IsActive)
	if err == sql.ErrNoRows {
		return nil, errors.NewCustomError(errors.NotFound, "url not found")
	} else if err != nil {
		log.Printf("Failed to fetch url by ID: %v", err)
		return nil, err
	}

	return &url, nil
}

func (u *UrlRepository) GetByShortUrl(shortUrl string) (*model.Url, error) {
	var url model.Url

	err := u.DB.QueryRow("SELECT id, short_url, actual_url, is_active FROM url WHERE short_url = ? and is_active = true", shortUrl).
		Scan(&url.ID, &url.ShortUrl, &url.ActualUrl, &url.IsActive)
	if err == sql.ErrNoRows {
		return nil, errors.NewCustomError(errors.NotFound, "url not found")
	} else if err != nil {
		log.Printf("Failed to fetch url by short URL: %v", err)
		return nil, err
	}

	return &url, nil
}

func (u *UrlRepository) UpdateStatus(req model.UpdateStatusRequest) error {
	_, err := u.DB.Exec("UPDATE url SET is_active = ? WHERE id = ?", req.Status, req.ID)
	if err != nil {
		log.Printf("Failed to update url status: %v", err)
		return err
	}

	return nil
}
