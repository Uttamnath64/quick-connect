package storage

import (
	"github.com/Uttamnath64/logger"
	"github.com/Uttamnath64/quick-connect/internal/app/config"
)

type Container struct {
	Config *config.Config
	Logger *logger.Logger
	Redis  *RedisClient
	Env    *config.AppEnv
}

// NewContainer initializes the DI container
func NewContainer(cfg *config.Config, log *logger.Logger, redis *RedisClient, env *config.AppEnv) *Container {
	return &Container{
		Config: cfg,
		Logger: log,
		Redis:  redis,
		Env:    env,
	}
}
