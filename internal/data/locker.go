package data

import (
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	klog "github.com/night-sword/kratos-kit/log"
	"github.com/night-sword/redis-locker"

	"github.com/night-sword/kratos-layout/internal/conf"
)

type Locker struct {
	*locker.Locker
}

func NewLocker(cfg *conf.Business, data *Data) *Locker {
	rl := locker.NewLocker(data.redis)

	opts := locker.GetDefaultOptions()
	opts.Logger = log.NewHelper(klog.GetLogger())
	opts.Prefix = cfg.GetName()
	locker.SetDefaultOptions(opts)

	return &Locker{
		Locker: rl,
	}
}
