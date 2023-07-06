package config

import (
	"fmt"
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

	addr := GetEnvValue("REDIS_HOST", "localhost") + ":" + GetEnvValue("REDIS_PORT", "6379")
	fmt.Println(addr)
	return &RedisConfig{
		Addr:       addr,
		Password:   GetEnvValue("REDIS_PASSWORD", ""),
		DB:         db,
		MaxRetries: retires,
	}, nil
}
