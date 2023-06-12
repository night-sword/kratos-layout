//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/night-sword/kratos-layout/cmd"
	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/server"
	"github.com/night-sword/kratos-layout/internal/service"
)

func wireApp(
	cmd.Name, cmd.Version,
	log.Logger, config.Config,
	*conf.Bootstrap, *conf.Server, *conf.Data, *conf.Business,
) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			// data.ProviderSet, biz.ProviderSet,
			service.ProviderSet, server.ProviderSet,
			NewServers, cmd.NewKratos,
		),
	)
}
