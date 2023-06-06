// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/night-sword/kratos-layout/cmd"
	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/server"
	"github.com/night-sword/kratos-layout/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(name cmd.Name, version cmd.Version, logger log.Logger, confServer *conf.Server, data *conf.Data, business *conf.Business) (*kratos.App, func(), error) {
	healthService := service.NewHealthService()
	grpcServer := server.NewGRPCServer(confServer, healthService, logger)
	httpServer := server.NewHTTPServer(confServer, healthService, logger)
	servers := NewServers(grpcServer, httpServer)
	app := cmd.NewKratos(name, version, logger, servers)
	return app, func() {
	}, nil
}
