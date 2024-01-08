package database

import (
	"fmt"

	"github.com/schoolboybru/urlShortener/internal/models"
)

type UrlStore interface {
	AddUrl(url string, shortenedUrl string) (*models.Url, error)
	GetUrl(url string) (string, error)
}

func (u *ReddisRepository) GetUrl(url string) (string, error) {
	var result = ""

	c := u.client

	err := c.Get(ctx, url).Scan(&result)

	fmt.Println(result)

	if err != nil {
		return "", err
	}
	return result, nil
}

func (u *ReddisRepository) AddUrl(longUrl string, shortenedUrl string) (*models.Url, error) {
	c := u.client
	url := models.Url{Short: shortenedUrl, Long: longUrl}

	err := c.Set(ctx, shortenedUrl, longUrl, 0).Err()

	if err != nil {
		return nil, err
	}

	return &url, err
}
