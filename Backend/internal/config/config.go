package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Server      server
	CacheConfig cacheConfig
	DbConfig    postgresConfig
}

type server struct {
	Port   int    `env:"SERVER_PORT" env-default:"8080"`
	Secret string `env:"SERVER_SECRET"`
}

type cacheConfig struct {
	DSN string `env:"REDIS_DSN"`
}

type postgresConfig struct {
	DSN string `env:"POSTGRES_DSN"`
}

// Load config from env
func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var c Config
	err = cleanenv.ReadEnv(&c)
	if err != nil {
		panic(err)
	}

	return &c
}
