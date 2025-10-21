package kit

import (
	"github.com/night-sword/kratos-kit/errors"
	"github.com/redis/go-redis/v9"

	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/data/repo"
)

type Cache struct {
	*redis.Client
	prefix string
}

func NewCache(cfg *conf.Bootstrap, redis *repo.Redis) (inst *Cache, err error) {
	if redis == nil || redis.Client() == nil {
		err = errors.InternalServer(errors.RsnInternal, "redis not init")
		return
	}

	inst = &Cache{
		Client: redis.Client(),
		prefix: cfg.GetBusiness().GetName(),
	}
	return
}

func (inst *Cache) Key(key string) string {
	return inst.prefix + ":" + key
}
