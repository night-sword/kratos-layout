package internal

import (
	"flag"
	"time"

	"github.com/go-kratos/kratos/v2"

	"github.com/night-sword/kratos-layout/internal/conf"
)

type WireApp func(arg []kratos.Option, bootstrap *conf.Bootstrap, confServer *conf.Server, confData *conf.Data) (*kratos.App, func(), error)

func Run(name *string, version string, wireApp WireApp) {
	flag.Parse()
	bootstrap, cancel := Bootstrap()
	defer cancel()
	if bootstrap.GetBusiness().GetName() != "" {
		*name = bootstrap.GetBusiness().GetName()
	}

	fixTimeZone(bootstrap.GetData().GetTimezone())

	logger := Logger(version, bootstrap.GetData().GetLog().GetLevel())
	opts := []kratos.Option{
		kratos.Name(*name),
		kratos.Version(version),
		kratos.Logger(logger),
	}

	app, cleanup, err := wireApp(opts, bootstrap, bootstrap.GetServer(), bootstrap.GetData())
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if e := app.Run(); e != nil {
		panic(e)
	}
	return
}

func fixTimeZone(cfg *conf.Data_TimeZone) {
	if cfg.GetLocation() == "" || cfg.GetOffset() <= 0 {
		return
	}

	location, err := time.LoadLocation(cfg.GetLocation())
	if err != nil {
		location = time.FixedZone("CST", int(cfg.GetOffset()))
	}

	time.Local = location
}
