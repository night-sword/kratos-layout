// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/night-sword/kratos-layout/cmd/internal"
	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/server"
	"github.com/night-sword/kratos-layout/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(arg []kratos.Option, bootstrap *conf.Bootstrap, confServer *conf.Server, data *conf.Data, business *conf.Business) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(confServer)
	healthService := service.NewHealthService()
	httpServer := server.NewHTTPServer(confServer, healthService)
	servers := newServers(grpcServer, httpServer)
	app := internal.NewKratos(arg, servers, bootstrap)
	return app, func() {
	}, nil
}
