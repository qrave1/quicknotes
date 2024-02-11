package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Server      Server
	CacheConfig CacheConfig
	DbConfig    DbConfig
}

type Server struct {
	Port int `env:"SERVER_PORT" env-default:"8080"`
}

type CacheConfig struct {
	DSN string `env:"REDIS_DSN"`
}

type DbConfig struct {
	DSN string `env:"POSTGRES_DSN"`
}

// Load Config from env
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
