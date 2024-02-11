package main

import (
	"github.com/qrave1/quicknotes/internal/config"
	"github.com/qrave1/quicknotes/internal/storage/postgres"
	redis2 "github.com/qrave1/quicknotes/internal/storage/redis"
	"github.com/qrave1/quicknotes/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.NewLogger()

	post := postgres.MustLoad(cfg)
	redis := redis2.MustLoad(cfg)

	// todo graceful shutdown

}
