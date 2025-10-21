package kit

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/night-sword/kratos-kit/errors"
	klog "github.com/night-sword/kratos-kit/log"
	"github.com/night-sword/redis-locker"

	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/data/repo"
)

type Locker struct {
	*locker.Locker
}

func NewLocker(cfg *conf.Bootstrap, redis *repo.Redis) (inst *Locker, err error) {
	if redis == nil || redis.Client() == nil {
		err = errors.InternalServer(errors.RsnInternal, "redis not init")
		return
	}

	rl := locker.NewLocker(redis.Client())

	opts := locker.GetDefaultOptions()
	opts.Logger = log.NewHelper(klog.GetLogger())
	opts.Prefix = cfg.GetBusiness().GetName()
	locker.SetDefaultOptions(opts)

	inst = &Locker{
		Locker: rl,
	}
	return
}
