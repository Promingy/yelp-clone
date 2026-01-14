package config

import (
	"os"
	"time"
)



type JWTConfig struct {
	AccessSecret     string
	RefreshSecret    string
	AccessExpiresIn  time.Duration
	RefreshExpiresIn time.Duration
}

func LoadJWTConfig() *JWTConfig {
	return &JWTConfig{
		AccessSecret:     os.Getenv("JWT_ACCESS_SECRET"),
		RefreshSecret:    os.Getenv("JWT_REFRESH_SECRET"),
		AccessExpiresIn:  15 * time.Minute,
		RefreshExpiresIn: 7 * 24 * time.Hour,
	}
}
