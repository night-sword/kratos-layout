package repo

import (
	"crypto/tls"

	"github.com/redis/go-redis/v9"

	"github.com/night-sword/kratos-layout/internal/conf"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(cfg *conf.Bootstrap) (inst *Redis, cleanup func()) {
	inst = &Redis{}

	client := inst.newClient(cfg.GetData().GetRedis())
	cleanup = func() { closeResource(client, "redis") }

	inst = &Redis{
		client: client,
	}
	return
}

func (inst *Redis) Client() *redis.Client {
	return inst.client
}

func (inst *Redis) newClient(cfg *conf.Data_Redis) (client *redis.Client) {
	opts := &redis.Options{
		Addr:     cfg.GetAddr(),
		Password: cfg.GetPwd(),
	}

	if cfg.GetTls() {
		opts.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	return redis.NewClient(opts)
}
