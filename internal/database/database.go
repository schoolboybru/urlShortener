package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type UrlStore struct {
	client *redis.Client
}

type UrlCache interface {
	AddUrl(url string, shortenedUrl string) (string, error)
	GetUrl(url string) (string, error)
}

var ctx = context.Background()

func New() (UrlCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	store := &UrlStore{
		client: rdb,
	}

	return store, nil

}

func (u *UrlStore) GetUrl(url string) (string, error) {
	var result = ""

	c := u.client

	err := c.Get(ctx, url).Scan(&result)

	fmt.Println(result)

	if err != nil {
		return "", err
	}
	return result, nil
}

func (u *UrlStore) AddUrl(url string, shortenedUrl string) (string, error) {
	c := u.client

	err := c.Set(ctx, shortenedUrl, url, 0).Err()

	if err != nil {
		return "", err
	}

	return shortenedUrl, err
}
