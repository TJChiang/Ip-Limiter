package config

import (
	"strconv"
	"time"
)

type LimiterConfig struct {
	Ttl        time.Duration
	MaxAttempt int64
}

func NewLimiterConfig() (*LimiterConfig, error) {
	ttl, err := strconv.ParseInt(GetEnvValue("LIMITER_TTL", "3600"), 10, 64)
	if err != nil {
		panic(err)
	}

	maxAttempt, err := strconv.ParseInt(GetEnvValue("LIMITER_MAX_ATTEMPT", "1000"), 10, 64)
	if err != nil {
		panic(err)
	}

	return &LimiterConfig{
		Ttl:        time.Duration(ttl) * time.Second,
		MaxAttempt: maxAttempt,
	}, nil
}
