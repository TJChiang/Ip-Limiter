package pkg

import (
	"IpLimiter/config"
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"time"
)

type Limiter struct {
	rdb    *redis.Client
	ctx    context.Context
	prefix string
	config *config.LimiterConfig
}

func NewLimiter(prefix string) (*Limiter, error) {
	redisConfig, err := config.NewRedisConfig()
	if err != nil {
		panic("Invalid redis config.")
	}

	limiterConfig, err := config.NewLimiterConfig()
	if err != nil {
		panic(err)
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
		prefix,
		limiterConfig,
	}, nil
}

func (l *Limiter) Limit(ipaddr string) (int64, time.Duration, error) {
	key, err := l.getKey(ipaddr)
	if err != nil {
		return 0, 0, err
	}

	value, err := l.rdb.Get(l.ctx, key).Result()
	if err != nil {
		return 0, 0, err
	}

	ttl, err := l.rdb.TTL(l.ctx, key).Result()
	if err != nil {
		return 0, 0, err
	}

	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return val, ttl, nil
}

func (l *Limiter) Hit(ipaddr string) (int64, error) {
	key, err := l.getKey(ipaddr)
	if err != nil {
		return 0, err
	}
	if _, err := l.rdb.SetNX(l.ctx, key, 0, l.config.Ttl).Result(); err != nil {
		return 0, err
	}

	value, err := l.rdb.Incr(l.ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (l *Limiter) Clean(ipaddr string) error {
	key, _ := l.getKey(ipaddr)
	if _, err := l.rdb.Del(l.ctx, key).Result(); err != nil {
		return err
	}

	return nil
}

func (l *Limiter) getKey(key string) (string, error) {
	var buffer strings.Builder
	if _, err := buffer.WriteString(l.prefix); err != nil {
		return "", err
	}
	if _, err := buffer.WriteString(key); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
