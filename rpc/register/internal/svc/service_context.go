package svc

import (
	"github.com/go-redis/redis/v9"
	"rj97807_work_serve/rpc/register/internal/config"
	"rj97807_work_serve/rpc/register/models"
)

type ServiceContext struct {
	Config config.Config
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RDB:    models.InitRedis(c),
	}
}
