package data

import (
	"github.com/redis/go-redis/v9"

	"github.com/night-sword/kratos-layout/internal/conf"
)

type Cache struct {
	redis  *redis.Client
	prefix string
}

func NewCache(cfg *conf.Business, data *Data) *Cache {
	return &Cache{
		redis:  data.redis,
		prefix: cfg.GetName(),
	}
}

func (inst *Cache) Client() (client *redis.Client) {
	return inst.redis
}

func (inst *Cache) Key(key string) string {
	return inst.prefix + ":" + key
}
