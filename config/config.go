package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env        string `env:"PLANNER_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"80"`
	DBHost     string `env:"PLANNER_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"PLANNER_DB_PORT" envDefault:"36306"`
	DBUser     string `env:"PLANNER_DB_USER" envDefault:"planner"`
	DBPassword string `env:"PLANNER_DB_PASSWORD" envDefault:"planner"`
	DBName     string `env:"PLANNER_DB_NAME" envDefault:"planner"`
	RedisHost  string `env:"PLANNER_REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort  int    `env:"PLANNER_REDIS_PORT" envDefault:"36379"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
