package svc

import (
	"github.com/go-redis/redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"rj97807_work_serve/api/internal/config"
	"rj97807_work_serve/api/internal/middleware"
	"rj97807_work_serve/api/models"
)

type ServiceContext struct {
	Config      config.Config
	EngineWeb   *gorm.DB
	EngineColly *gorm.DB
	RDB         *redis.Client
	Auth        rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		EngineWeb:   models.InitGorm(c.Mysql.DataSourceWeb),
		EngineColly: models.InitGorm(c.Mysql.DataSourceColly),
		RDB:         models.InitRedis(c),
		Auth:        middleware.NewAuthMiddleware().Handle,
	}
}
