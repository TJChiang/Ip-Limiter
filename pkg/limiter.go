package pkg

import (
	"IpLimiter/config"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Limiter struct {
	rdb        *redis.Client
	ctx        context.Context
	expiration time.Duration
}

func NewLimiter() (*Limiter, error) {
	redisConfig, err := config.NewRedisConfig()
	if err != nil {
		panic("Invalid redis config.")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:       redisConfig.Addr,
		Password:   redisConfig.Password,
		DB:         redisConfig.DB,
		MaxRetries: redisConfig.MaxRetries,
	})

	ctx := context.Background()

	_, err = rdb.Ping(ctx).Result()
	if err == redis.Nil || err != nil {
		return nil, err
	}

	return &Limiter{
		rdb,
		ctx,
		redisConfig.CacheExpiration,
	}, nil
}
