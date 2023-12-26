package data

import (
	"github.com/redis/go-redis/v9"

	"github.com/night-sword/kratos-layout/internal/conf"
)

type Cache struct {
	*redis.Client
	prefix string
}

func NewCache(cfg *conf.Business, data *Data) *Cache {
	return &Cache{
		Client: data.redis,
		prefix: cfg.GetName(),
	}
}

func (inst *Cache) Key(key string) string {
	return inst.prefix + ":" + key
}
