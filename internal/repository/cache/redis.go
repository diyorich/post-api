package cache

import (
	"context"
	"fmt"
	"github.com/diyorich/post-api/internal/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	*redis.Client
}

func Dial(ctx context.Context, cfg config.Cache) (*Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis %v", err)
	}

	return &Cache{rdb}, nil
}
