package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	"github.com/night-sword/kratos-layout/cmd"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
)

func main() {
	flag.Parse()
	configs := cmd.Bootstrap()
	name, version := cmd.Name(Name), cmd.Version(Version)
	logger := cmd.Logger(version)
	log.SetLogger(logger) // set default logger

	app, cleanup, err := wireApp(
		name, version, logger,
		configs.Server, configs.Data, configs.Business,
	)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func NewServers(gs *grpc.Server, hs *http.Server) *cmd.Servers {
	ss := []transport.Server{gs, hs}
	return cmd.NewServers(ss...)
}
