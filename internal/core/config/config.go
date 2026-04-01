package core_config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	TimeZone *time.Location
}

func NewConfig() (*Config, error) {
	tz := os.Getenv("TIME_ZONE")
	if tz == "" {
		tz = "UTC"
	}

	timeZone, err := time.LoadLocation(tz)
	if err != nil {
		return nil, fmt.Errorf("Load location: %s: %w", tz, err)
	}

	return &Config{
		TimeZone: timeZone,
	}, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		panic("Get core config")
	}

	return config
}
