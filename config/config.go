package config

import "os"

type AppConfig struct {
	Name    string
	Port    string
	GinMode string
}

func GetEnvValue(key string, initial string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return initial
}

func NewAppConfig() (*AppConfig, error) {

	return &AppConfig{
		Name:    GetEnvValue("APP_NAME", "IpLimiter"),
		Port:    GetEnvValue("APP_PORT", "8088"),
		GinMode: GetEnvValue("GIN_MODE", "debug"),
	}, nil
}
