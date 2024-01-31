//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	cmd "github.com/night-sword/kratos-layout/cmd/internal"
	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/server"
	"github.com/night-sword/kratos-layout/internal/service"
)

func wireApp([]kratos.Option, *conf.Bootstrap, *conf.Server, *conf.Data) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			// data.ProviderSet, biz.ProviderSet,
			service.ProviderSet, server.ProviderSet,
			newServers, cmd.NewKratos,
		),
	)
}
