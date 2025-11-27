package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewClient(ctx context.Context, port string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", port),
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	return rdb
}

func Close(rdb *redis.Client) error {
	err := rdb.Close()
	if err != nil {
		return err
	}

	return nil
}
