package config

import (
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/orm/gorm"
	"github.com/suyuan32/simple-admin-common/plugins/mq/nats"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	//DatabaseConf config.DatabaseConf
	DatabaseConf gorm.Conf
	RedisConf    config.RedisConf
	CoreRpc      zrpc.RpcClientConf
	NatsConf     nats.Conf
}
