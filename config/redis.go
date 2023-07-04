package config

import (
	"strconv"
)

type RedisConfig struct {
	Addr       string
	Password   string
	DB         int
	MaxRetries int
}

func NewRedisConfig() (*RedisConfig, error) {
	db, err := strconv.Atoi(GetEnvValue("REDIS_DB", "0"))
	if err != nil {
		return nil, err
	}

	retires, err := strconv.Atoi(GetEnvValue("REDIS_RETRIES", "3"))
	if err != nil {
		return nil, err
	}

	return &RedisConfig{
		Addr:       GetEnvValue("REDIS_ADDR", "localhost:6379"),
		Password:   GetEnvValue("REDIS_PASSWORD", ""),
		DB:         db,
		MaxRetries: retires,
	}, nil
}
