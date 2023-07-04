package config

import "os"

type AppConfig struct {
	Name string
}

func GetEnvValue(key string, initial string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return initial
}

func NewAppConfig() (*AppConfig, error) {
	appName := GetEnvValue("APP_NAME", "IpLimiter")

	return &AppConfig{
		Name: appName,
	}, nil
}
