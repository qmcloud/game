package svc

import (
	"github.com/nats-io/nats.go/jetstream"
	"github.com/qmcloud/game/ent"
	"github.com/qmcloud/game/internal/config"
	"github.com/redis/go-redis/v9"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
	Redis  redis.UniversalClient
	Nats   jetstream.JetStream
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	nats, err := c.NatsConf.NewJetStream()
	logx.Must(err)
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  c.RedisConf.MustNewUniversalRedis(),
		Nats:   nats,
	}
}
