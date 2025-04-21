package svc

import (
	"github.com/qmcloud/game/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := c.DatabaseConf.NewGORM()
	if err != nil {
		logx.Errorf("DatabaseError", logx.Field("detail", err.Error()))
		return nil
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  c.RedisConf.MustNewUniversalRedis(),
	}
}
