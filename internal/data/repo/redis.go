package repo

import (
	"context"
	"crypto/tls"

	"github.com/night-sword/kratos-kit/errors"
	"github.com/redis/go-redis/v9"

	"github.com/night-sword/kratos-layout/internal/conf"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(cfg *conf.Bootstrap) (inst *Redis, cleanup func(), err error) {
	redisCfg := cfg.GetData().GetRedis()
	if redisCfg == nil || redisCfg.GetAddr() == "" {
		err = errors.InternalServer(errors.RsnInternal, "redis addr is empty")
		return
	}

	client := newRedisClient(redisCfg)

	if err = client.Ping(context.Background()).Err(); err != nil {
		_ = client.Close()
		return
	}

	cleanup = func() { closeResource(client, "redis") }
	inst = &Redis{client: client}
	return
}

func (inst *Redis) Client() *redis.Client {
	return inst.client
}

func newRedisClient(cfg *conf.Data_Redis) (client *redis.Client) {
	opts := &redis.Options{
		Addr:     cfg.GetAddr(),
		Password: cfg.GetPwd(),
	}

	if cfg.GetNetwork() != "" {
		opts.Network = cfg.GetNetwork()
	}
	if cfg.GetReadTimeout() != nil {
		opts.ReadTimeout = cfg.GetReadTimeout().AsDuration()
	}
	if cfg.GetWriteTimeout() != nil {
		opts.WriteTimeout = cfg.GetWriteTimeout().AsDuration()
	}

	if cfg.GetTls() {
		opts.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	return redis.NewClient(opts)
}
