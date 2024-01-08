package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	UrlStore
}

type ReddisRepository struct {
	client *redis.Client
}

var ctx = context.Background()

func New() (Store, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	store := &ReddisRepository{
		client: rdb,
	}

	return store, nil

}
